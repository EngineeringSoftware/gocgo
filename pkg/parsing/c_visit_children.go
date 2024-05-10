package parsing

import (
	"github.com/antlr4-go/antlr/v4"
)

// VisitChildren is missing in generated code:
// https://github.com/antlr/antlr4/issues/4398

func VisitChildren(v CVisitor, ctx antlr.RuleNode) interface{} {
	res := make([]interface{}, 0)
	for i := 0; i < ctx.GetChildCount(); i++ {
		c := ctx.GetChild(i)
		switch n := c.(type) {
		case antlr.RuleNode:
			out := n.Accept(v)
			if out != nil {
				if err, ok := out.(error); ok {
					return err
				}
				res = append(res, out)
			}
		case antlr.TerminalNode:
			out := v.VisitTerminal(n)
			if out != nil {
				if err, ok := out.(error); ok {
					return err
				}
				res = append(res, out)
			}
		default:
			// error
		}
	}
	if len(res) == 0 {
		return nil
	}
	return res
}
