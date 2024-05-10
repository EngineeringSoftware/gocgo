package visitors

import (
	"bytes"
	"github.com/EngineeringSoftware/gocgo/pkg/parsing"
	"github.com/antlr4-go/antlr/v4"
)

type FuncDeleteVisitor struct {
	*BaseVisitorImpl
	in   bool
	Buff bytes.Buffer
}

func NewFuncDeleteVisitor() *FuncDeleteVisitor {
	return &FuncDeleteVisitor{in: false}
}

func (v *FuncDeleteVisitor) VisitTerminal(n antlr.TerminalNode) error {
	v.Buff.WriteString(n.GetSymbol().GetText())
	v.Buff.WriteString(" ")
	return nil
}

func (v *FuncDeleteVisitor) VisitFunctionDefinition(ctx *parsing.FunctionDefinitionContext) (bool, error) {
	v.in = true
	return true, nil
}

func (v *FuncDeleteVisitor) VisitFunctionDefinitionEnd(ctx *parsing.FunctionDefinitionContext) error {
	v.in = false
	return nil
}

func (v *FuncDeleteVisitor) VisitBlockItemList(ctx *parsing.BlockItemListContext) (bool, error) {
	if v.in {
		return false, nil
	} else {
		return true, nil
	}
}
