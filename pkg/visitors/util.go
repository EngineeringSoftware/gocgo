package visitors

import (
	"fmt"
	"github.com/EngineeringSoftware/gocgo/pkg/parsing"
	"github.com/antlr4-go/antlr/v4"
	"strings"
)

type saveErrorStrategy struct {
	*antlr.DefaultErrorStrategy
	happened bool
}

func newSaveErrorStrategy() *saveErrorStrategy {
	s := new(saveErrorStrategy)
	s.DefaultErrorStrategy = antlr.NewDefaultErrorStrategy()
	return s
}

func (s *saveErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	s.happened = true
	s.DefaultErrorStrategy.Recover(recognizer, e)
}

func (s *saveErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	s.happened = true
	return s.DefaultErrorStrategy.RecoverInline(recognizer)
}

type saveErrorListener struct {
	*antlr.DefaultErrorListener
	happened bool
}

func (l *saveErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.happened = true
	l.DefaultErrorListener.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
}

func visit(visitor BaseVisitor, chars antlr.CharStream) error {
	lexer := parsing.NewCLexer(chars)
	listener := &saveErrorListener{}
	lexer.AddErrorListener(listener)

	toks := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewCParser(toks)

	strategy := newSaveErrorStrategy()
	p.SetErrorHandler(strategy)

	tree := p.CompilationUnit()
	if listener.happened {
		return fmt.Errorf("error during lexing")
	}
	if strategy.happened {
		return fmt.Errorf("error during parsing")
	}
	vadapter := NewAdapterVisitor(visitor)
	res := tree.Accept(vadapter)
	if res != nil {
		if err, ok := res.(error); ok && err != nil {
			return err
		}
	}
	return nil
}

func VisitFile(visitor BaseVisitor, fname string) error {
	input, _ := antlr.NewFileStream(fname)
	return visit(visitor, input)
}

func VisitCode(visitor BaseVisitor, code string) error {
	input := antlr.NewIoStream(strings.NewReader(code))
	return visit(visitor, input)
}
