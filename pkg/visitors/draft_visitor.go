package visitors

import (
	"github.com/EngineeringSoftware/gocgo/pkg/parsing"
)

type DraftVisitor struct {
	parsing.CVisitor
}

func (v *DraftVisitor) VisitPrimaryExpression(ctx *parsing.PrimaryExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitGenericSelection(ctx *parsing.GenericSelectionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitGenericAssocList(ctx *parsing.GenericAssocListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitGenericAssociation(ctx *parsing.GenericAssociationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitPostfixExpression(ctx *parsing.PostfixExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitArgumentExpressionList(ctx *parsing.ArgumentExpressionListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitUnaryExpression(ctx *parsing.UnaryExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitUnaryOperator(ctx *parsing.UnaryOperatorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitCastExpression(ctx *parsing.CastExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitMultiplicativeExpression(ctx *parsing.MultiplicativeExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitAdditiveExpression(ctx *parsing.AdditiveExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitShiftExpression(ctx *parsing.ShiftExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitRelationalExpression(ctx *parsing.RelationalExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitEqualityExpression(ctx *parsing.EqualityExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitAndExpression(ctx *parsing.AndExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitExclusiveOrExpression(ctx *parsing.ExclusiveOrExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitInclusiveOrExpression(ctx *parsing.InclusiveOrExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitLogicalAndExpression(ctx *parsing.LogicalAndExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitLogicalOrExpression(ctx *parsing.LogicalOrExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitConditionalExpression(ctx *parsing.ConditionalExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitAssignmentExpression(ctx *parsing.AssignmentExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitAssignmentOperator(ctx *parsing.AssignmentOperatorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitExpression(ctx *parsing.ExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitConstantExpression(ctx *parsing.ConstantExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitDeclaration(ctx *parsing.DeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitDeclarationSpecifiers(ctx *parsing.DeclarationSpecifiersContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitDeclarationSpecifiers2(ctx *parsing.DeclarationSpecifiers2Context) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitDeclarationSpecifier(ctx *parsing.DeclarationSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitInitDeclaratorList(ctx *parsing.InitDeclaratorListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitInitDeclarator(ctx *parsing.InitDeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitStorageClassSpecifier(ctx *parsing.StorageClassSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitTypeSpecifier(ctx *parsing.TypeSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitStructOrUnionSpecifier(ctx *parsing.StructOrUnionSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitStructOrUnion(ctx *parsing.StructOrUnionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitStructDeclarationList(ctx *parsing.StructDeclarationListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitStructDeclaration(ctx *parsing.StructDeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitSpecifierQualifierList(ctx *parsing.SpecifierQualifierListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitStructDeclaratorList(ctx *parsing.StructDeclaratorListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitStructDeclarator(ctx *parsing.StructDeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitEnumSpecifier(ctx *parsing.EnumSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitEnumeratorList(ctx *parsing.EnumeratorListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitEnumerator(ctx *parsing.EnumeratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitEnumerationConstant(ctx *parsing.EnumerationConstantContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitAtomicTypeSpecifier(ctx *parsing.AtomicTypeSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitTypeQualifier(ctx *parsing.TypeQualifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitFunctionSpecifier(ctx *parsing.FunctionSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitAlignmentSpecifier(ctx *parsing.AlignmentSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitDeclarator(ctx *parsing.DeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitDirectDeclarator(ctx *parsing.DirectDeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitVcSpecificModifer(ctx *parsing.VcSpecificModiferContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitGccDeclaratorExtension(ctx *parsing.GccDeclaratorExtensionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitGccAttributeSpecifier(ctx *parsing.GccAttributeSpecifierContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitGccAttributeList(ctx *parsing.GccAttributeListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitGccAttribute(ctx *parsing.GccAttributeContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitNestedParenthesesBlock(ctx *parsing.NestedParenthesesBlockContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitPointer(ctx *parsing.PointerContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitTypeQualifierList(ctx *parsing.TypeQualifierListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitParameterTypeList(ctx *parsing.ParameterTypeListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitParameterList(ctx *parsing.ParameterListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitParameterDeclaration(ctx *parsing.ParameterDeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitIdentifierList(ctx *parsing.IdentifierListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitTypeName(ctx *parsing.TypeNameContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitAbstractDeclarator(ctx *parsing.AbstractDeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitDirectAbstractDeclarator(ctx *parsing.DirectAbstractDeclaratorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitTypedefName(ctx *parsing.TypedefNameContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitInitializer(ctx *parsing.InitializerContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitInitializerList(ctx *parsing.InitializerListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitDesignation(ctx *parsing.DesignationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitDesignatorList(ctx *parsing.DesignatorListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitDesignator(ctx *parsing.DesignatorContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitStaticAssertDeclaration(ctx *parsing.StaticAssertDeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitStatement(ctx *parsing.StatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitLabeledStatement(ctx *parsing.LabeledStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitCompoundStatement(ctx *parsing.CompoundStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitBlockItemList(ctx *parsing.BlockItemListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitBlockItem(ctx *parsing.BlockItemContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitExpressionStatement(ctx *parsing.ExpressionStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitSelectionStatement(ctx *parsing.SelectionStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitIterationStatement(ctx *parsing.IterationStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitForCondition(ctx *parsing.ForConditionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitForDeclaration(ctx *parsing.ForDeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitForExpression(ctx *parsing.ForExpressionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitJumpStatement(ctx *parsing.JumpStatementContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitCompilationUnit(ctx *parsing.CompilationUnitContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitTranslationUnit(ctx *parsing.TranslationUnitContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitExternalDeclaration(ctx *parsing.ExternalDeclarationContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitFunctionDefinition(ctx *parsing.FunctionDefinitionContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}

func (v *DraftVisitor) VisitDeclarationList(ctx *parsing.DeclarationListContext) interface{} {
	return parsing.VisitChildren(v, ctx)
}
