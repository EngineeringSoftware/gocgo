package visitors

import (
	"bytes"
	"github.com/EngineeringSoftware/gocgo/pkg/parsing"
	"github.com/antlr4-go/antlr/v4"
)

type PrintingVisitor struct {
	parsing.CVisitor
	Buff bytes.Buffer
}

func NewPrintingVisitor() *PrintingVisitor {
	return &PrintingVisitor{}
}

func (v *PrintingVisitor) VisitTerminal(n antlr.TerminalNode) interface{} {
	v.Buff.WriteString(n.GetSymbol().GetText())
	v.Buff.WriteString(" ")
	// We could also collect into a result: n.GetSymbol().GetText()
	return nil
}

func (v *PrintingVisitor) VisitPrimaryExpression(ctx *parsing.PrimaryExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitGenericSelection(ctx *parsing.GenericSelectionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitGenericAssocList(ctx *parsing.GenericAssocListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitGenericAssociation(ctx *parsing.GenericAssociationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitPostfixExpression(ctx *parsing.PostfixExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitArgumentExpressionList(ctx *parsing.ArgumentExpressionListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitUnaryExpression(ctx *parsing.UnaryExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitUnaryOperator(ctx *parsing.UnaryOperatorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitCastExpression(ctx *parsing.CastExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitMultiplicativeExpression(ctx *parsing.MultiplicativeExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitAdditiveExpression(ctx *parsing.AdditiveExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitShiftExpression(ctx *parsing.ShiftExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitRelationalExpression(ctx *parsing.RelationalExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitEqualityExpression(ctx *parsing.EqualityExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitAndExpression(ctx *parsing.AndExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitExclusiveOrExpression(ctx *parsing.ExclusiveOrExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitInclusiveOrExpression(ctx *parsing.InclusiveOrExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitLogicalAndExpression(ctx *parsing.LogicalAndExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitLogicalOrExpression(ctx *parsing.LogicalOrExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitConditionalExpression(ctx *parsing.ConditionalExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitAssignmentExpression(ctx *parsing.AssignmentExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitAssignmentOperator(ctx *parsing.AssignmentOperatorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitExpression(ctx *parsing.ExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitConstantExpression(ctx *parsing.ConstantExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitDeclaration(ctx *parsing.DeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitDeclarationSpecifiers(ctx *parsing.DeclarationSpecifiersContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitDeclarationSpecifiers2(ctx *parsing.DeclarationSpecifiers2Context) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitDeclarationSpecifier(ctx *parsing.DeclarationSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitInitDeclaratorList(ctx *parsing.InitDeclaratorListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitInitDeclarator(ctx *parsing.InitDeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitStorageClassSpecifier(ctx *parsing.StorageClassSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitTypeSpecifier(ctx *parsing.TypeSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitStructOrUnionSpecifier(ctx *parsing.StructOrUnionSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitStructOrUnion(ctx *parsing.StructOrUnionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitStructDeclarationList(ctx *parsing.StructDeclarationListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitStructDeclaration(ctx *parsing.StructDeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitSpecifierQualifierList(ctx *parsing.SpecifierQualifierListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitStructDeclaratorList(ctx *parsing.StructDeclaratorListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitStructDeclarator(ctx *parsing.StructDeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitEnumSpecifier(ctx *parsing.EnumSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitEnumeratorList(ctx *parsing.EnumeratorListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitEnumerator(ctx *parsing.EnumeratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitEnumerationConstant(ctx *parsing.EnumerationConstantContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitAtomicTypeSpecifier(ctx *parsing.AtomicTypeSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitTypeQualifier(ctx *parsing.TypeQualifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitFunctionSpecifier(ctx *parsing.FunctionSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitAlignmentSpecifier(ctx *parsing.AlignmentSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitDeclarator(ctx *parsing.DeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitDirectDeclarator(ctx *parsing.DirectDeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitVcSpecificModifer(ctx *parsing.VcSpecificModiferContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitGccDeclaratorExtension(ctx *parsing.GccDeclaratorExtensionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitGccAttributeSpecifier(ctx *parsing.GccAttributeSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitGccAttributeList(ctx *parsing.GccAttributeListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitGccAttribute(ctx *parsing.GccAttributeContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitNestedParenthesesBlock(ctx *parsing.NestedParenthesesBlockContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitPointer(ctx *parsing.PointerContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitTypeQualifierList(ctx *parsing.TypeQualifierListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitParameterTypeList(ctx *parsing.ParameterTypeListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitParameterList(ctx *parsing.ParameterListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitParameterDeclaration(ctx *parsing.ParameterDeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitIdentifierList(ctx *parsing.IdentifierListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitTypeName(ctx *parsing.TypeNameContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitAbstractDeclarator(ctx *parsing.AbstractDeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitDirectAbstractDeclarator(ctx *parsing.DirectAbstractDeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitTypedefName(ctx *parsing.TypedefNameContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitInitializer(ctx *parsing.InitializerContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitInitializerList(ctx *parsing.InitializerListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitDesignation(ctx *parsing.DesignationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitDesignatorList(ctx *parsing.DesignatorListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitDesignator(ctx *parsing.DesignatorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitStaticAssertDeclaration(ctx *parsing.StaticAssertDeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitStatement(ctx *parsing.StatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitLabeledStatement(ctx *parsing.LabeledStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitCompoundStatement(ctx *parsing.CompoundStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitBlockItemList(ctx *parsing.BlockItemListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitBlockItem(ctx *parsing.BlockItemContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitExpressionStatement(ctx *parsing.ExpressionStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitSelectionStatement(ctx *parsing.SelectionStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitIterationStatement(ctx *parsing.IterationStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitForCondition(ctx *parsing.ForConditionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitForDeclaration(ctx *parsing.ForDeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitForExpression(ctx *parsing.ForExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitJumpStatement(ctx *parsing.JumpStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitCompilationUnit(ctx *parsing.CompilationUnitContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitTranslationUnit(ctx *parsing.TranslationUnitContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitExternalDeclaration(ctx *parsing.ExternalDeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitFunctionDefinition(ctx *parsing.FunctionDefinitionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *PrintingVisitor) VisitDeclarationList(ctx *parsing.DeclarationListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}
