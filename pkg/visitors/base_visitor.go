package visitors

import (
	"github.com/EngineeringSoftware/gocgo/pkg/parsing"
	"github.com/antlr4-go/antlr/v4"
)

type BaseVisitor interface {
	VisitTerminal(n antlr.TerminalNode) error

	VisitPrimaryExpression(ctx *parsing.PrimaryExpressionContext) (bool, error)

	VisitPrimaryExpressionEnd(ctx *parsing.PrimaryExpressionContext) error

	VisitGenericSelection(ctx *parsing.GenericSelectionContext) (bool, error)

	VisitGenericSelectionEnd(ctx *parsing.GenericSelectionContext) error

	VisitGenericAssocList(ctx *parsing.GenericAssocListContext) (bool, error)

	VisitGenericAssocListEnd(ctx *parsing.GenericAssocListContext) error

	VisitGenericAssociation(ctx *parsing.GenericAssociationContext) (bool, error)

	VisitGenericAssociationEnd(ctx *parsing.GenericAssociationContext) error

	VisitPostfixExpression(ctx *parsing.PostfixExpressionContext) (bool, error)

	VisitPostfixExpressionEnd(ctx *parsing.PostfixExpressionContext) error

	VisitArgumentExpressionList(ctx *parsing.ArgumentExpressionListContext) (bool, error)

	VisitArgumentExpressionListEnd(ctx *parsing.ArgumentExpressionListContext) error

	VisitUnaryExpression(ctx *parsing.UnaryExpressionContext) (bool, error)

	VisitUnaryExpressionEnd(ctx *parsing.UnaryExpressionContext) error

	VisitUnaryOperator(ctx *parsing.UnaryOperatorContext) (bool, error)

	VisitUnaryOperatorEnd(ctx *parsing.UnaryOperatorContext) error

	VisitCastExpression(ctx *parsing.CastExpressionContext) (bool, error)

	VisitCastExpressionEnd(ctx *parsing.CastExpressionContext) error

	VisitMultiplicativeExpression(ctx *parsing.MultiplicativeExpressionContext) (bool, error)

	VisitMultiplicativeExpressionEnd(ctx *parsing.MultiplicativeExpressionContext) error

	VisitAdditiveExpression(ctx *parsing.AdditiveExpressionContext) (bool, error)

	VisitAdditiveExpressionEnd(ctx *parsing.AdditiveExpressionContext) error

	VisitShiftExpression(ctx *parsing.ShiftExpressionContext) (bool, error)

	VisitShiftExpressionEnd(ctx *parsing.ShiftExpressionContext) error

	VisitRelationalExpression(ctx *parsing.RelationalExpressionContext) (bool, error)

	VisitRelationalExpressionEnd(ctx *parsing.RelationalExpressionContext) error

	VisitEqualityExpression(ctx *parsing.EqualityExpressionContext) (bool, error)

	VisitEqualityExpressionEnd(ctx *parsing.EqualityExpressionContext) error

	VisitAndExpression(ctx *parsing.AndExpressionContext) (bool, error)

	VisitAndExpressionEnd(ctx *parsing.AndExpressionContext) error

	VisitExclusiveOrExpression(ctx *parsing.ExclusiveOrExpressionContext) (bool, error)

	VisitExclusiveOrExpressionEnd(ctx *parsing.ExclusiveOrExpressionContext) error

	VisitInclusiveOrExpression(ctx *parsing.InclusiveOrExpressionContext) (bool, error)

	VisitInclusiveOrExpressionEnd(ctx *parsing.InclusiveOrExpressionContext) error

	VisitLogicalAndExpression(ctx *parsing.LogicalAndExpressionContext) (bool, error)

	VisitLogicalAndExpressionEnd(ctx *parsing.LogicalAndExpressionContext) error

	VisitLogicalOrExpression(ctx *parsing.LogicalOrExpressionContext) (bool, error)

	VisitLogicalOrExpressionEnd(ctx *parsing.LogicalOrExpressionContext) error

	VisitConditionalExpression(ctx *parsing.ConditionalExpressionContext) (bool, error)

	VisitConditionalExpressionEnd(ctx *parsing.ConditionalExpressionContext) error

	VisitAssignmentExpression(ctx *parsing.AssignmentExpressionContext) (bool, error)

	VisitAssignmentExpressionEnd(ctx *parsing.AssignmentExpressionContext) error

	VisitAssignmentOperator(ctx *parsing.AssignmentOperatorContext) (bool, error)

	VisitAssignmentOperatorEnd(ctx *parsing.AssignmentOperatorContext) error

	VisitExpression(ctx *parsing.ExpressionContext) (bool, error)

	VisitExpressionEnd(ctx *parsing.ExpressionContext) error

	VisitConstantExpression(ctx *parsing.ConstantExpressionContext) (bool, error)

	VisitConstantExpressionEnd(ctx *parsing.ConstantExpressionContext) error

	VisitDeclaration(ctx *parsing.DeclarationContext) (bool, error)

	VisitDeclarationEnd(ctx *parsing.DeclarationContext) error

	VisitDeclarationSpecifiers(ctx *parsing.DeclarationSpecifiersContext) (bool, error)

	VisitDeclarationSpecifiersEnd(ctx *parsing.DeclarationSpecifiersContext) error

	VisitDeclarationSpecifiers2(ctx *parsing.DeclarationSpecifiers2Context) (bool, error)

	VisitDeclarationSpecifiers2End(ctx *parsing.DeclarationSpecifiers2Context) error

	VisitDeclarationSpecifier(ctx *parsing.DeclarationSpecifierContext) (bool, error)

	VisitDeclarationSpecifierEnd(ctx *parsing.DeclarationSpecifierContext) error

	VisitInitDeclaratorList(ctx *parsing.InitDeclaratorListContext) (bool, error)

	VisitInitDeclaratorListEnd(ctx *parsing.InitDeclaratorListContext) error

	VisitInitDeclarator(ctx *parsing.InitDeclaratorContext) (bool, error)

	VisitInitDeclaratorEnd(ctx *parsing.InitDeclaratorContext) error

	VisitStorageClassSpecifier(ctx *parsing.StorageClassSpecifierContext) (bool, error)

	VisitStorageClassSpecifierEnd(ctx *parsing.StorageClassSpecifierContext) error

	VisitTypeSpecifier(ctx *parsing.TypeSpecifierContext) (bool, error)

	VisitTypeSpecifierEnd(ctx *parsing.TypeSpecifierContext) error

	VisitStructOrUnionSpecifier(ctx *parsing.StructOrUnionSpecifierContext) (bool, error)

	VisitStructOrUnionSpecifierEnd(ctx *parsing.StructOrUnionSpecifierContext) error

	VisitStructOrUnion(ctx *parsing.StructOrUnionContext) (bool, error)

	VisitStructOrUnionEnd(ctx *parsing.StructOrUnionContext) error

	VisitStructDeclarationList(ctx *parsing.StructDeclarationListContext) (bool, error)

	VisitStructDeclarationListEnd(ctx *parsing.StructDeclarationListContext) error

	VisitStructDeclaration(ctx *parsing.StructDeclarationContext) (bool, error)

	VisitStructDeclarationEnd(ctx *parsing.StructDeclarationContext) error

	VisitSpecifierQualifierList(ctx *parsing.SpecifierQualifierListContext) (bool, error)

	VisitSpecifierQualifierListEnd(ctx *parsing.SpecifierQualifierListContext) error

	VisitStructDeclaratorList(ctx *parsing.StructDeclaratorListContext) (bool, error)

	VisitStructDeclaratorListEnd(ctx *parsing.StructDeclaratorListContext) error

	VisitStructDeclarator(ctx *parsing.StructDeclaratorContext) (bool, error)

	VisitStructDeclaratorEnd(ctx *parsing.StructDeclaratorContext) error

	VisitEnumSpecifier(ctx *parsing.EnumSpecifierContext) (bool, error)

	VisitEnumSpecifierEnd(ctx *parsing.EnumSpecifierContext) error

	VisitEnumeratorList(ctx *parsing.EnumeratorListContext) (bool, error)

	VisitEnumeratorListEnd(ctx *parsing.EnumeratorListContext) error

	VisitEnumerator(ctx *parsing.EnumeratorContext) (bool, error)

	VisitEnumeratorEnd(ctx *parsing.EnumeratorContext) error

	VisitEnumerationConstant(ctx *parsing.EnumerationConstantContext) (bool, error)

	VisitEnumerationConstantEnd(ctx *parsing.EnumerationConstantContext) error

	VisitAtomicTypeSpecifier(ctx *parsing.AtomicTypeSpecifierContext) (bool, error)

	VisitAtomicTypeSpecifierEnd(ctx *parsing.AtomicTypeSpecifierContext) error

	VisitTypeQualifier(ctx *parsing.TypeQualifierContext) (bool, error)

	VisitTypeQualifierEnd(ctx *parsing.TypeQualifierContext) error

	VisitFunctionSpecifier(ctx *parsing.FunctionSpecifierContext) (bool, error)

	VisitFunctionSpecifierEnd(ctx *parsing.FunctionSpecifierContext) error

	VisitAlignmentSpecifier(ctx *parsing.AlignmentSpecifierContext) (bool, error)

	VisitAlignmentSpecifierEnd(ctx *parsing.AlignmentSpecifierContext) error

	VisitDeclarator(ctx *parsing.DeclaratorContext) (bool, error)

	VisitDeclaratorEnd(ctx *parsing.DeclaratorContext) error

	VisitDirectDeclarator(ctx *parsing.DirectDeclaratorContext) (bool, error)

	VisitDirectDeclaratorEnd(ctx *parsing.DirectDeclaratorContext) error

	VisitVcSpecificModifer(ctx *parsing.VcSpecificModiferContext) (bool, error)

	VisitVcSpecificModiferEnd(ctx *parsing.VcSpecificModiferContext) error

	VisitGccDeclaratorExtension(ctx *parsing.GccDeclaratorExtensionContext) (bool, error)

	VisitGccDeclaratorExtensionEnd(ctx *parsing.GccDeclaratorExtensionContext) error

	VisitGccAttributeSpecifier(ctx *parsing.GccAttributeSpecifierContext) (bool, error)

	VisitGccAttributeSpecifierEnd(ctx *parsing.GccAttributeSpecifierContext) error

	VisitGccAttributeList(ctx *parsing.GccAttributeListContext) (bool, error)

	VisitGccAttributeListEnd(ctx *parsing.GccAttributeListContext) error

	VisitGccAttribute(ctx *parsing.GccAttributeContext) (bool, error)

	VisitGccAttributeEnd(ctx *parsing.GccAttributeContext) error

	VisitNestedParenthesesBlock(ctx *parsing.NestedParenthesesBlockContext) (bool, error)

	VisitNestedParenthesesBlockEnd(ctx *parsing.NestedParenthesesBlockContext) error

	VisitPointer(ctx *parsing.PointerContext) (bool, error)

	VisitPointerEnd(ctx *parsing.PointerContext) error

	VisitTypeQualifierList(ctx *parsing.TypeQualifierListContext) (bool, error)

	VisitTypeQualifierListEnd(ctx *parsing.TypeQualifierListContext) error

	VisitParameterTypeList(ctx *parsing.ParameterTypeListContext) (bool, error)

	VisitParameterTypeListEnd(ctx *parsing.ParameterTypeListContext) error

	VisitParameterList(ctx *parsing.ParameterListContext) (bool, error)

	VisitParameterListEnd(ctx *parsing.ParameterListContext) error

	VisitParameterDeclaration(ctx *parsing.ParameterDeclarationContext) (bool, error)

	VisitParameterDeclarationEnd(ctx *parsing.ParameterDeclarationContext) error

	VisitIdentifierList(ctx *parsing.IdentifierListContext) (bool, error)

	VisitIdentifierListEnd(ctx *parsing.IdentifierListContext) error

	VisitTypeName(ctx *parsing.TypeNameContext) (bool, error)

	VisitTypeNameEnd(ctx *parsing.TypeNameContext) error

	VisitAbstractDeclarator(ctx *parsing.AbstractDeclaratorContext) (bool, error)

	VisitAbstractDeclaratorEnd(ctx *parsing.AbstractDeclaratorContext) error

	VisitDirectAbstractDeclarator(ctx *parsing.DirectAbstractDeclaratorContext) (bool, error)

	VisitDirectAbstractDeclaratorEnd(ctx *parsing.DirectAbstractDeclaratorContext) error

	VisitTypedefName(ctx *parsing.TypedefNameContext) (bool, error)

	VisitTypedefNameEnd(ctx *parsing.TypedefNameContext) error

	VisitInitializer(ctx *parsing.InitializerContext) (bool, error)

	VisitInitializerEnd(ctx *parsing.InitializerContext) error

	VisitInitializerList(ctx *parsing.InitializerListContext) (bool, error)

	VisitInitializerListEnd(ctx *parsing.InitializerListContext) error

	VisitDesignation(ctx *parsing.DesignationContext) (bool, error)

	VisitDesignationEnd(ctx *parsing.DesignationContext) error

	VisitDesignatorList(ctx *parsing.DesignatorListContext) (bool, error)

	VisitDesignatorListEnd(ctx *parsing.DesignatorListContext) error

	VisitDesignator(ctx *parsing.DesignatorContext) (bool, error)

	VisitDesignatorEnd(ctx *parsing.DesignatorContext) error

	VisitStaticAssertDeclaration(ctx *parsing.StaticAssertDeclarationContext) (bool, error)

	VisitStaticAssertDeclarationEnd(ctx *parsing.StaticAssertDeclarationContext) error

	VisitStatement(ctx *parsing.StatementContext) (bool, error)

	VisitStatementEnd(ctx *parsing.StatementContext) error

	VisitLabeledStatement(ctx *parsing.LabeledStatementContext) (bool, error)

	VisitLabeledStatementEnd(ctx *parsing.LabeledStatementContext) error

	VisitCompoundStatement(ctx *parsing.CompoundStatementContext) (bool, error)

	VisitCompoundStatementEnd(ctx *parsing.CompoundStatementContext) error

	VisitBlockItemList(ctx *parsing.BlockItemListContext) (bool, error)

	VisitBlockItemListEnd(ctx *parsing.BlockItemListContext) error

	VisitBlockItem(ctx *parsing.BlockItemContext) (bool, error)

	VisitBlockItemEnd(ctx *parsing.BlockItemContext) error

	VisitExpressionStatement(ctx *parsing.ExpressionStatementContext) (bool, error)

	VisitExpressionStatementEnd(ctx *parsing.ExpressionStatementContext) error

	VisitSelectionStatement(ctx *parsing.SelectionStatementContext) (bool, error)

	VisitSelectionStatementEnd(ctx *parsing.SelectionStatementContext) error

	VisitIterationStatement(ctx *parsing.IterationStatementContext) (bool, error)

	VisitIterationStatementEnd(ctx *parsing.IterationStatementContext) error

	VisitForCondition(ctx *parsing.ForConditionContext) (bool, error)

	VisitForConditionEnd(ctx *parsing.ForConditionContext) error

	VisitForDeclaration(ctx *parsing.ForDeclarationContext) (bool, error)

	VisitForDeclarationEnd(ctx *parsing.ForDeclarationContext) error

	VisitForExpression(ctx *parsing.ForExpressionContext) (bool, error)

	VisitForExpressionEnd(ctx *parsing.ForExpressionContext) error

	VisitJumpStatement(ctx *parsing.JumpStatementContext) (bool, error)

	VisitJumpStatementEnd(ctx *parsing.JumpStatementContext) error

	VisitCompilationUnit(ctx *parsing.CompilationUnitContext) (bool, error)

	VisitCompilationUnitEnd(ctx *parsing.CompilationUnitContext) error

	VisitTranslationUnit(ctx *parsing.TranslationUnitContext) (bool, error)

	VisitTranslationUnitEnd(ctx *parsing.TranslationUnitContext) error

	VisitExternalDeclaration(ctx *parsing.ExternalDeclarationContext) (bool, error)

	VisitExternalDeclarationEnd(ctx *parsing.ExternalDeclarationContext) error

	VisitFunctionDefinition(ctx *parsing.FunctionDefinitionContext) (bool, error)

	VisitFunctionDefinitionEnd(ctx *parsing.FunctionDefinitionContext) error

	VisitDeclarationList(ctx *parsing.DeclarationListContext) (bool, error)

	VisitDeclarationListEnd(ctx *parsing.DeclarationListContext) error
}

type BaseVisitorImpl struct {
}

func (v *BaseVisitorImpl) VisitTerminal(n antlr.TerminalNode) error {
	return nil
}

func (v *BaseVisitorImpl) VisitPrimaryExpression(ctx *parsing.PrimaryExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitPrimaryExpressionEnd(ctx *parsing.PrimaryExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitGenericSelection(ctx *parsing.GenericSelectionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitGenericSelectionEnd(ctx *parsing.GenericSelectionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitGenericAssocList(ctx *parsing.GenericAssocListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitGenericAssocListEnd(ctx *parsing.GenericAssocListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitGenericAssociation(ctx *parsing.GenericAssociationContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitGenericAssociationEnd(ctx *parsing.GenericAssociationContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitPostfixExpression(ctx *parsing.PostfixExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitPostfixExpressionEnd(ctx *parsing.PostfixExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitArgumentExpressionList(ctx *parsing.ArgumentExpressionListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitArgumentExpressionListEnd(ctx *parsing.ArgumentExpressionListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitUnaryExpression(ctx *parsing.UnaryExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitUnaryExpressionEnd(ctx *parsing.UnaryExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitUnaryOperator(ctx *parsing.UnaryOperatorContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitUnaryOperatorEnd(ctx *parsing.UnaryOperatorContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitCastExpression(ctx *parsing.CastExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitCastExpressionEnd(ctx *parsing.CastExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitMultiplicativeExpression(ctx *parsing.MultiplicativeExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitMultiplicativeExpressionEnd(ctx *parsing.MultiplicativeExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitAdditiveExpression(ctx *parsing.AdditiveExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitAdditiveExpressionEnd(ctx *parsing.AdditiveExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitShiftExpression(ctx *parsing.ShiftExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitShiftExpressionEnd(ctx *parsing.ShiftExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitRelationalExpression(ctx *parsing.RelationalExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitRelationalExpressionEnd(ctx *parsing.RelationalExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitEqualityExpression(ctx *parsing.EqualityExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitEqualityExpressionEnd(ctx *parsing.EqualityExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitAndExpression(ctx *parsing.AndExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitAndExpressionEnd(ctx *parsing.AndExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitExclusiveOrExpression(ctx *parsing.ExclusiveOrExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitExclusiveOrExpressionEnd(ctx *parsing.ExclusiveOrExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitInclusiveOrExpression(ctx *parsing.InclusiveOrExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitInclusiveOrExpressionEnd(ctx *parsing.InclusiveOrExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitLogicalAndExpression(ctx *parsing.LogicalAndExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitLogicalAndExpressionEnd(ctx *parsing.LogicalAndExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitLogicalOrExpression(ctx *parsing.LogicalOrExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitLogicalOrExpressionEnd(ctx *parsing.LogicalOrExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitConditionalExpression(ctx *parsing.ConditionalExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitConditionalExpressionEnd(ctx *parsing.ConditionalExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitAssignmentExpression(ctx *parsing.AssignmentExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitAssignmentExpressionEnd(ctx *parsing.AssignmentExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitAssignmentOperator(ctx *parsing.AssignmentOperatorContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitAssignmentOperatorEnd(ctx *parsing.AssignmentOperatorContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitExpression(ctx *parsing.ExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitExpressionEnd(ctx *parsing.ExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitConstantExpression(ctx *parsing.ConstantExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitConstantExpressionEnd(ctx *parsing.ConstantExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitDeclaration(ctx *parsing.DeclarationContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitDeclarationEnd(ctx *parsing.DeclarationContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitDeclarationSpecifiers(ctx *parsing.DeclarationSpecifiersContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitDeclarationSpecifiersEnd(ctx *parsing.DeclarationSpecifiersContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitDeclarationSpecifiers2(ctx *parsing.DeclarationSpecifiers2Context) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitDeclarationSpecifiers2End(ctx *parsing.DeclarationSpecifiers2Context) error {
	return nil
}

func (v *BaseVisitorImpl) VisitDeclarationSpecifier(ctx *parsing.DeclarationSpecifierContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitDeclarationSpecifierEnd(ctx *parsing.DeclarationSpecifierContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitInitDeclaratorList(ctx *parsing.InitDeclaratorListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitInitDeclaratorListEnd(ctx *parsing.InitDeclaratorListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitInitDeclarator(ctx *parsing.InitDeclaratorContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitInitDeclaratorEnd(ctx *parsing.InitDeclaratorContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitStorageClassSpecifier(ctx *parsing.StorageClassSpecifierContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitStorageClassSpecifierEnd(ctx *parsing.StorageClassSpecifierContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitTypeSpecifier(ctx *parsing.TypeSpecifierContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitTypeSpecifierEnd(ctx *parsing.TypeSpecifierContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitStructOrUnionSpecifier(ctx *parsing.StructOrUnionSpecifierContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitStructOrUnionSpecifierEnd(ctx *parsing.StructOrUnionSpecifierContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitStructOrUnion(ctx *parsing.StructOrUnionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitStructOrUnionEnd(ctx *parsing.StructOrUnionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitStructDeclarationList(ctx *parsing.StructDeclarationListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitStructDeclarationListEnd(ctx *parsing.StructDeclarationListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitStructDeclaration(ctx *parsing.StructDeclarationContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitStructDeclarationEnd(ctx *parsing.StructDeclarationContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitSpecifierQualifierList(ctx *parsing.SpecifierQualifierListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitSpecifierQualifierListEnd(ctx *parsing.SpecifierQualifierListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitStructDeclaratorList(ctx *parsing.StructDeclaratorListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitStructDeclaratorListEnd(ctx *parsing.StructDeclaratorListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitStructDeclarator(ctx *parsing.StructDeclaratorContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitStructDeclaratorEnd(ctx *parsing.StructDeclaratorContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitEnumSpecifier(ctx *parsing.EnumSpecifierContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitEnumSpecifierEnd(ctx *parsing.EnumSpecifierContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitEnumeratorList(ctx *parsing.EnumeratorListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitEnumeratorListEnd(ctx *parsing.EnumeratorListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitEnumerator(ctx *parsing.EnumeratorContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitEnumeratorEnd(ctx *parsing.EnumeratorContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitEnumerationConstant(ctx *parsing.EnumerationConstantContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitEnumerationConstantEnd(ctx *parsing.EnumerationConstantContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitAtomicTypeSpecifier(ctx *parsing.AtomicTypeSpecifierContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitAtomicTypeSpecifierEnd(ctx *parsing.AtomicTypeSpecifierContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitTypeQualifier(ctx *parsing.TypeQualifierContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitTypeQualifierEnd(ctx *parsing.TypeQualifierContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitFunctionSpecifier(ctx *parsing.FunctionSpecifierContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitFunctionSpecifierEnd(ctx *parsing.FunctionSpecifierContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitAlignmentSpecifier(ctx *parsing.AlignmentSpecifierContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitAlignmentSpecifierEnd(ctx *parsing.AlignmentSpecifierContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitDeclarator(ctx *parsing.DeclaratorContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitDeclaratorEnd(ctx *parsing.DeclaratorContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitDirectDeclarator(ctx *parsing.DirectDeclaratorContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitDirectDeclaratorEnd(ctx *parsing.DirectDeclaratorContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitVcSpecificModifer(ctx *parsing.VcSpecificModiferContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitVcSpecificModiferEnd(ctx *parsing.VcSpecificModiferContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitGccDeclaratorExtension(ctx *parsing.GccDeclaratorExtensionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitGccDeclaratorExtensionEnd(ctx *parsing.GccDeclaratorExtensionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitGccAttributeSpecifier(ctx *parsing.GccAttributeSpecifierContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitGccAttributeSpecifierEnd(ctx *parsing.GccAttributeSpecifierContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitGccAttributeList(ctx *parsing.GccAttributeListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitGccAttributeListEnd(ctx *parsing.GccAttributeListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitGccAttribute(ctx *parsing.GccAttributeContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitGccAttributeEnd(ctx *parsing.GccAttributeContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitNestedParenthesesBlock(ctx *parsing.NestedParenthesesBlockContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitNestedParenthesesBlockEnd(ctx *parsing.NestedParenthesesBlockContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitPointer(ctx *parsing.PointerContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitPointerEnd(ctx *parsing.PointerContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitTypeQualifierList(ctx *parsing.TypeQualifierListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitTypeQualifierListEnd(ctx *parsing.TypeQualifierListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitParameterTypeList(ctx *parsing.ParameterTypeListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitParameterTypeListEnd(ctx *parsing.ParameterTypeListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitParameterList(ctx *parsing.ParameterListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitParameterListEnd(ctx *parsing.ParameterListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitParameterDeclaration(ctx *parsing.ParameterDeclarationContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitParameterDeclarationEnd(ctx *parsing.ParameterDeclarationContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitIdentifierList(ctx *parsing.IdentifierListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitIdentifierListEnd(ctx *parsing.IdentifierListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitTypeName(ctx *parsing.TypeNameContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitTypeNameEnd(ctx *parsing.TypeNameContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitAbstractDeclarator(ctx *parsing.AbstractDeclaratorContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitAbstractDeclaratorEnd(ctx *parsing.AbstractDeclaratorContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitDirectAbstractDeclarator(ctx *parsing.DirectAbstractDeclaratorContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitDirectAbstractDeclaratorEnd(ctx *parsing.DirectAbstractDeclaratorContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitTypedefName(ctx *parsing.TypedefNameContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitTypedefNameEnd(ctx *parsing.TypedefNameContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitInitializer(ctx *parsing.InitializerContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitInitializerEnd(ctx *parsing.InitializerContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitInitializerList(ctx *parsing.InitializerListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitInitializerListEnd(ctx *parsing.InitializerListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitDesignation(ctx *parsing.DesignationContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitDesignationEnd(ctx *parsing.DesignationContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitDesignatorList(ctx *parsing.DesignatorListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitDesignatorListEnd(ctx *parsing.DesignatorListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitDesignator(ctx *parsing.DesignatorContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitDesignatorEnd(ctx *parsing.DesignatorContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitStaticAssertDeclaration(ctx *parsing.StaticAssertDeclarationContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitStaticAssertDeclarationEnd(ctx *parsing.StaticAssertDeclarationContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitStatement(ctx *parsing.StatementContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitStatementEnd(ctx *parsing.StatementContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitLabeledStatement(ctx *parsing.LabeledStatementContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitLabeledStatementEnd(ctx *parsing.LabeledStatementContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitCompoundStatement(ctx *parsing.CompoundStatementContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitCompoundStatementEnd(ctx *parsing.CompoundStatementContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitBlockItemList(ctx *parsing.BlockItemListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitBlockItemListEnd(ctx *parsing.BlockItemListContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitBlockItem(ctx *parsing.BlockItemContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitBlockItemEnd(ctx *parsing.BlockItemContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitExpressionStatement(ctx *parsing.ExpressionStatementContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitExpressionStatementEnd(ctx *parsing.ExpressionStatementContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitSelectionStatement(ctx *parsing.SelectionStatementContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitSelectionStatementEnd(ctx *parsing.SelectionStatementContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitIterationStatement(ctx *parsing.IterationStatementContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitIterationStatementEnd(ctx *parsing.IterationStatementContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitForCondition(ctx *parsing.ForConditionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitForConditionEnd(ctx *parsing.ForConditionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitForDeclaration(ctx *parsing.ForDeclarationContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitForDeclarationEnd(ctx *parsing.ForDeclarationContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitForExpression(ctx *parsing.ForExpressionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitForExpressionEnd(ctx *parsing.ForExpressionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitJumpStatement(ctx *parsing.JumpStatementContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitJumpStatementEnd(ctx *parsing.JumpStatementContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitCompilationUnit(ctx *parsing.CompilationUnitContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitCompilationUnitEnd(ctx *parsing.CompilationUnitContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitTranslationUnit(ctx *parsing.TranslationUnitContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitTranslationUnitEnd(ctx *parsing.TranslationUnitContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitExternalDeclaration(ctx *parsing.ExternalDeclarationContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitExternalDeclarationEnd(ctx *parsing.ExternalDeclarationContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitFunctionDefinition(ctx *parsing.FunctionDefinitionContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitFunctionDefinitionEnd(ctx *parsing.FunctionDefinitionContext) error {
	return nil
}

func (v *BaseVisitorImpl) VisitDeclarationList(ctx *parsing.DeclarationListContext) (bool, error) {
	return true, nil
}

func (v *BaseVisitorImpl) VisitDeclarationListEnd(ctx *parsing.DeclarationListContext) error {
	return nil
}
