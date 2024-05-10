package visitors

import (
	"bytes"
	"fmt"
	"github.com/EngineeringSoftware/gocgo/pkg/parsing"
	"github.com/antlr4-go/antlr/v4"
	"testing"
)

type SaveVisitor struct {
	*BaseVisitorImpl
	buff bytes.Buffer
}

func (v *SaveVisitor) VisitTerminal(n antlr.TerminalNode) error {
	s := n.GetSymbol().GetText()
	if s == "<EOF>" {
		return nil
	}
	v.buff.WriteString(s)
	v.buff.WriteString(" ")
	return nil
}

func TestAllVisited(t *testing.T) {
	v := &SaveVisitor{}
	code := "int main ( void ) { for ( int i = 0 ; i < 10 ; i ++ ) { int j = i ++ ; } return 0 ; } "
	err := VisitCode(v, code)

	if err != nil {
		t.Errorf("should not have errors during this visit")
	}
	if v.buff.String() != code {
		t.Errorf("incorrect code output *%v*", v.buff.String())
	}
}

type ErrorVisitor struct {
	*BaseVisitorImpl
}

func (v *ErrorVisitor) VisitStatement(ctx *parsing.StatementContext) (bool, error) {
	return true, fmt.Errorf("this should propagate")
}

func TestErrorPropagation(t *testing.T) {
	v := &ErrorVisitor{}
	code := "int main() { int i; i = 5; }"
	err := VisitCode(v, code)
	if err == nil {
		t.Errorf("should have an error during this visit")
	}
}

type SkipVisitor struct {
	*BaseVisitorImpl
	end   bool
	block bool
}

func (v *SkipVisitor) VisitFunctionDefinition(ctx *parsing.FunctionDefinitionContext) (bool, error) {
	return false, nil
}

func (v *SkipVisitor) VisitFunctionDefinitionEnd(ctx *parsing.FunctionDefinitionContext) error {
	v.end = true
	return nil
}

func (v *SkipVisitor) VisitBlockItemList(ctx *parsing.BlockItemListContext) (bool, error) {
	v.block = true
	return true, nil
}

func TestSkipChildren(t *testing.T) {
	v := &SkipVisitor{end: false, block: false}
	code := "int main() { int i; i = 5; }"

	err := VisitCode(v, code)
	if err != nil {
		t.Errorf("should not have errors during this visit")
	}
	if !v.end {
		t.Errorf("end visit function was not invoked (but should be)")
	}
	if v.block {
		t.Errorf("block visit function is invoked (but shold not be)")
	}
}

func TestParseError(t *testing.T) {
	codes := []string{"int main() {", "int main() { return 0;"}

	for _, code := range codes {
		v := &SaveVisitor{}
		err := VisitCode(v, code)
		if err == nil {
			t.Errorf("there should be a parser error")
		}
	}
}

func TestLexError(t *testing.T) {
	v := &SaveVisitor{}
	code := "@"

	err := VisitCode(v, code)
	if err == nil {
		t.Errorf("there should be a lexer error")
	}
}
