package visitors

import (
	"github.com/EngineeringSoftware/gocgo/pkg/parsing"
	"github.com/antlr4-go/antlr/v4"
)

type AdapterVisitor struct {
	parsing.CVisitor
	w BaseVisitor
}

func NewAdapterVisitor(w BaseVisitor) *AdapterVisitor {
	return &AdapterVisitor{w: w}
}

func (v *AdapterVisitor) visit(s func() (bool, error), c func() interface{}, e func() error) error {
	if pro, err := s(); err != nil {
		return err
	} else if pro {
		res := c()
		if res != nil {
			if err, ok := res.(error); ok {
				return err
			}
		}
	}
	if err := e(); err != nil {
		return err
	}
	return nil
}

func (v *AdapterVisitor) VisitTerminal(n antlr.TerminalNode) interface{} {
	return v.w.VisitTerminal(n)
}

func (v *AdapterVisitor) VisitPrimaryExpression(ctx *parsing.PrimaryExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitPrimaryExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitPrimaryExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitGenericSelection(ctx *parsing.GenericSelectionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitGenericSelection(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitGenericSelectionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitGenericAssocList(ctx *parsing.GenericAssocListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitGenericAssocList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitGenericAssocListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitGenericAssociation(ctx *parsing.GenericAssociationContext) interface{} {
	s := func() (bool, error) { return v.w.VisitGenericAssociation(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitGenericAssociationEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitPostfixExpression(ctx *parsing.PostfixExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitPostfixExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitPostfixExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitArgumentExpressionList(ctx *parsing.ArgumentExpressionListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitArgumentExpressionList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitArgumentExpressionListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitUnaryExpression(ctx *parsing.UnaryExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitUnaryExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitUnaryExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitUnaryOperator(ctx *parsing.UnaryOperatorContext) interface{} {
	s := func() (bool, error) { return v.w.VisitUnaryOperator(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitUnaryOperatorEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitCastExpression(ctx *parsing.CastExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitCastExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitCastExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitMultiplicativeExpression(ctx *parsing.MultiplicativeExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitMultiplicativeExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitMultiplicativeExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitAdditiveExpression(ctx *parsing.AdditiveExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitAdditiveExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitAdditiveExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitShiftExpression(ctx *parsing.ShiftExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitShiftExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitShiftExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitRelationalExpression(ctx *parsing.RelationalExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitRelationalExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitRelationalExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitEqualityExpression(ctx *parsing.EqualityExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitEqualityExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitEqualityExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitAndExpression(ctx *parsing.AndExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitAndExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitAndExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitExclusiveOrExpression(ctx *parsing.ExclusiveOrExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitExclusiveOrExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitExclusiveOrExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitInclusiveOrExpression(ctx *parsing.InclusiveOrExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitInclusiveOrExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitInclusiveOrExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitLogicalAndExpression(ctx *parsing.LogicalAndExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitLogicalAndExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitLogicalAndExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitLogicalOrExpression(ctx *parsing.LogicalOrExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitLogicalOrExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitLogicalOrExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitConditionalExpression(ctx *parsing.ConditionalExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitConditionalExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitConditionalExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitAssignmentExpression(ctx *parsing.AssignmentExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitAssignmentExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitAssignmentExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitAssignmentOperator(ctx *parsing.AssignmentOperatorContext) interface{} {
	s := func() (bool, error) { return v.w.VisitAssignmentOperator(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitAssignmentOperatorEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitExpression(ctx *parsing.ExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitConstantExpression(ctx *parsing.ConstantExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitConstantExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitConstantExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitDeclaration(ctx *parsing.DeclarationContext) interface{} {
	s := func() (bool, error) { return v.w.VisitDeclaration(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitDeclarationEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitDeclarationSpecifiers(ctx *parsing.DeclarationSpecifiersContext) interface{} {
	s := func() (bool, error) { return v.w.VisitDeclarationSpecifiers(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitDeclarationSpecifiersEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitDeclarationSpecifiers2(ctx *parsing.DeclarationSpecifiers2Context) interface{} {
	s := func() (bool, error) { return v.w.VisitDeclarationSpecifiers2(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitDeclarationSpecifiers2End(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitDeclarationSpecifier(ctx *parsing.DeclarationSpecifierContext) interface{} {
	s := func() (bool, error) { return v.w.VisitDeclarationSpecifier(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitDeclarationSpecifierEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitInitDeclaratorList(ctx *parsing.InitDeclaratorListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitInitDeclaratorList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitInitDeclaratorListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitInitDeclarator(ctx *parsing.InitDeclaratorContext) interface{} {
	s := func() (bool, error) { return v.w.VisitInitDeclarator(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitInitDeclaratorEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitStorageClassSpecifier(ctx *parsing.StorageClassSpecifierContext) interface{} {
	s := func() (bool, error) { return v.w.VisitStorageClassSpecifier(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitStorageClassSpecifierEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitTypeSpecifier(ctx *parsing.TypeSpecifierContext) interface{} {
	s := func() (bool, error) { return v.w.VisitTypeSpecifier(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitTypeSpecifierEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitStructOrUnionSpecifier(ctx *parsing.StructOrUnionSpecifierContext) interface{} {
	s := func() (bool, error) { return v.w.VisitStructOrUnionSpecifier(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitStructOrUnionSpecifierEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitStructOrUnion(ctx *parsing.StructOrUnionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitStructOrUnion(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitStructOrUnionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitStructDeclarationList(ctx *parsing.StructDeclarationListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitStructDeclarationList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitStructDeclarationListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitStructDeclaration(ctx *parsing.StructDeclarationContext) interface{} {
	s := func() (bool, error) { return v.w.VisitStructDeclaration(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitStructDeclarationEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitSpecifierQualifierList(ctx *parsing.SpecifierQualifierListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitSpecifierQualifierList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitSpecifierQualifierListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitStructDeclaratorList(ctx *parsing.StructDeclaratorListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitStructDeclaratorList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitStructDeclaratorListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitStructDeclarator(ctx *parsing.StructDeclaratorContext) interface{} {
	s := func() (bool, error) { return v.w.VisitStructDeclarator(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitStructDeclaratorEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitEnumSpecifier(ctx *parsing.EnumSpecifierContext) interface{} {
	s := func() (bool, error) { return v.w.VisitEnumSpecifier(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitEnumSpecifierEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitEnumeratorList(ctx *parsing.EnumeratorListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitEnumeratorList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitEnumeratorListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitEnumerator(ctx *parsing.EnumeratorContext) interface{} {
	s := func() (bool, error) { return v.w.VisitEnumerator(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitEnumeratorEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitEnumerationConstant(ctx *parsing.EnumerationConstantContext) interface{} {
	s := func() (bool, error) { return v.w.VisitEnumerationConstant(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitEnumerationConstantEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitAtomicTypeSpecifier(ctx *parsing.AtomicTypeSpecifierContext) interface{} {
	s := func() (bool, error) { return v.w.VisitAtomicTypeSpecifier(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitAtomicTypeSpecifierEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitTypeQualifier(ctx *parsing.TypeQualifierContext) interface{} {
	s := func() (bool, error) { return v.w.VisitTypeQualifier(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitTypeQualifierEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitFunctionSpecifier(ctx *parsing.FunctionSpecifierContext) interface{} {
	s := func() (bool, error) { return v.w.VisitFunctionSpecifier(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitFunctionSpecifierEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitAlignmentSpecifier(ctx *parsing.AlignmentSpecifierContext) interface{} {
	s := func() (bool, error) { return v.w.VisitAlignmentSpecifier(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitAlignmentSpecifierEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitDeclarator(ctx *parsing.DeclaratorContext) interface{} {
	s := func() (bool, error) { return v.w.VisitDeclarator(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitDeclaratorEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitDirectDeclarator(ctx *parsing.DirectDeclaratorContext) interface{} {
	s := func() (bool, error) { return v.w.VisitDirectDeclarator(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitDirectDeclaratorEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitVcSpecificModifer(ctx *parsing.VcSpecificModiferContext) interface{} {
	s := func() (bool, error) { return v.w.VisitVcSpecificModifer(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitVcSpecificModiferEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitGccDeclaratorExtension(ctx *parsing.GccDeclaratorExtensionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitGccDeclaratorExtension(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitGccDeclaratorExtensionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitGccAttributeSpecifier(ctx *parsing.GccAttributeSpecifierContext) interface{} {
	s := func() (bool, error) { return v.w.VisitGccAttributeSpecifier(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitGccAttributeSpecifierEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitGccAttributeList(ctx *parsing.GccAttributeListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitGccAttributeList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitGccAttributeListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitGccAttribute(ctx *parsing.GccAttributeContext) interface{} {
	s := func() (bool, error) { return v.w.VisitGccAttribute(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitGccAttributeEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitNestedParenthesesBlock(ctx *parsing.NestedParenthesesBlockContext) interface{} {
	s := func() (bool, error) { return v.w.VisitNestedParenthesesBlock(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitNestedParenthesesBlockEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitPointer(ctx *parsing.PointerContext) interface{} {
	s := func() (bool, error) { return v.w.VisitPointer(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitPointerEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitTypeQualifierList(ctx *parsing.TypeQualifierListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitTypeQualifierList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitTypeQualifierListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitParameterTypeList(ctx *parsing.ParameterTypeListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitParameterTypeList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitParameterTypeListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitParameterList(ctx *parsing.ParameterListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitParameterList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitParameterListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitParameterDeclaration(ctx *parsing.ParameterDeclarationContext) interface{} {
	s := func() (bool, error) { return v.w.VisitParameterDeclaration(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitParameterDeclarationEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitIdentifierList(ctx *parsing.IdentifierListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitIdentifierList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitIdentifierListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitTypeName(ctx *parsing.TypeNameContext) interface{} {
	s := func() (bool, error) { return v.w.VisitTypeName(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitTypeNameEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitAbstractDeclarator(ctx *parsing.AbstractDeclaratorContext) interface{} {
	s := func() (bool, error) { return v.w.VisitAbstractDeclarator(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitAbstractDeclaratorEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitDirectAbstractDeclarator(ctx *parsing.DirectAbstractDeclaratorContext) interface{} {
	s := func() (bool, error) { return v.w.VisitDirectAbstractDeclarator(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitDirectAbstractDeclaratorEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitTypedefName(ctx *parsing.TypedefNameContext) interface{} {
	s := func() (bool, error) { return v.w.VisitTypedefName(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitTypedefNameEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitInitializer(ctx *parsing.InitializerContext) interface{} {
	s := func() (bool, error) { return v.w.VisitInitializer(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitInitializerEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitInitializerList(ctx *parsing.InitializerListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitInitializerList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitInitializerListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitDesignation(ctx *parsing.DesignationContext) interface{} {
	s := func() (bool, error) { return v.w.VisitDesignation(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitDesignationEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitDesignatorList(ctx *parsing.DesignatorListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitDesignatorList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitDesignatorListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitDesignator(ctx *parsing.DesignatorContext) interface{} {
	s := func() (bool, error) { return v.w.VisitDesignator(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitDesignatorEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitStaticAssertDeclaration(ctx *parsing.StaticAssertDeclarationContext) interface{} {
	s := func() (bool, error) { return v.w.VisitStaticAssertDeclaration(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitStaticAssertDeclarationEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitStatement(ctx *parsing.StatementContext) interface{} {
	s := func() (bool, error) { return v.w.VisitStatement(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitStatementEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitLabeledStatement(ctx *parsing.LabeledStatementContext) interface{} {
	s := func() (bool, error) { return v.w.VisitLabeledStatement(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitLabeledStatementEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitCompoundStatement(ctx *parsing.CompoundStatementContext) interface{} {
	s := func() (bool, error) { return v.w.VisitCompoundStatement(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitCompoundStatementEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitBlockItemList(ctx *parsing.BlockItemListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitBlockItemList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitBlockItemListEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitBlockItem(ctx *parsing.BlockItemContext) interface{} {
	s := func() (bool, error) { return v.w.VisitBlockItem(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitBlockItemEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitExpressionStatement(ctx *parsing.ExpressionStatementContext) interface{} {
	s := func() (bool, error) { return v.w.VisitExpressionStatement(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitExpressionStatementEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitSelectionStatement(ctx *parsing.SelectionStatementContext) interface{} {
	s := func() (bool, error) { return v.w.VisitSelectionStatement(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitSelectionStatementEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitIterationStatement(ctx *parsing.IterationStatementContext) interface{} {
	s := func() (bool, error) { return v.w.VisitIterationStatement(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitIterationStatementEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitForCondition(ctx *parsing.ForConditionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitForCondition(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitForConditionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitForDeclaration(ctx *parsing.ForDeclarationContext) interface{} {
	s := func() (bool, error) { return v.w.VisitForDeclaration(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitForDeclarationEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitForExpression(ctx *parsing.ForExpressionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitForExpression(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitForExpressionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitJumpStatement(ctx *parsing.JumpStatementContext) interface{} {
	s := func() (bool, error) { return v.w.VisitJumpStatement(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitJumpStatementEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitCompilationUnit(ctx *parsing.CompilationUnitContext) interface{} {
	s := func() (bool, error) { return v.w.VisitCompilationUnit(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitCompilationUnitEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitTranslationUnit(ctx *parsing.TranslationUnitContext) interface{} {
	s := func() (bool, error) { return v.w.VisitTranslationUnit(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitTranslationUnitEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitExternalDeclaration(ctx *parsing.ExternalDeclarationContext) interface{} {
	s := func() (bool, error) { return v.w.VisitExternalDeclaration(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitExternalDeclarationEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitFunctionDefinition(ctx *parsing.FunctionDefinitionContext) interface{} {
	s := func() (bool, error) { return v.w.VisitFunctionDefinition(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitFunctionDefinitionEnd(ctx) }
	return v.visit(s, c, e)
}

func (v *AdapterVisitor) VisitDeclarationList(ctx *parsing.DeclarationListContext) interface{} {
	s := func() (bool, error) { return v.w.VisitDeclarationList(ctx) }
	c := func() interface{} { return parsing.VisitChildren(v, ctx) }
	e := func() error { return v.w.VisitDeclarationListEnd(ctx) }
	return v.visit(s, c, e)
}
