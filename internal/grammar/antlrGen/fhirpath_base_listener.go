// Code generated from Fhirpath.g4 by ANTLR 4.13.0. DO NOT EDIT.

package antlrGen // Fhirpath
import "github.com/antlr4-go/antlr/v4"

// BaseFhirpathListener is a complete listener for a parse tree produced by FhirpathParser.
type BaseFhirpathListener struct{}

var _ FhirpathListener = &BaseFhirpathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFhirpathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFhirpathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFhirpathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFhirpathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterIndexerExpression is called when production indexerExpression is entered.
func (s *BaseFhirpathListener) EnterIndexerExpression(ctx *IndexerExpressionContext) {}

// ExitIndexerExpression is called when production indexerExpression is exited.
func (s *BaseFhirpathListener) ExitIndexerExpression(ctx *IndexerExpressionContext) {}

// EnterPolarityExpression is called when production polarityExpression is entered.
func (s *BaseFhirpathListener) EnterPolarityExpression(ctx *PolarityExpressionContext) {}

// ExitPolarityExpression is called when production polarityExpression is exited.
func (s *BaseFhirpathListener) ExitPolarityExpression(ctx *PolarityExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseFhirpathListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseFhirpathListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseFhirpathListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseFhirpathListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterUnionExpression is called when production unionExpression is entered.
func (s *BaseFhirpathListener) EnterUnionExpression(ctx *UnionExpressionContext) {}

// ExitUnionExpression is called when production unionExpression is exited.
func (s *BaseFhirpathListener) ExitUnionExpression(ctx *UnionExpressionContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseFhirpathListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseFhirpathListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseFhirpathListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseFhirpathListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterMembershipExpression is called when production membershipExpression is entered.
func (s *BaseFhirpathListener) EnterMembershipExpression(ctx *MembershipExpressionContext) {}

// ExitMembershipExpression is called when production membershipExpression is exited.
func (s *BaseFhirpathListener) ExitMembershipExpression(ctx *MembershipExpressionContext) {}

// EnterInequalityExpression is called when production inequalityExpression is entered.
func (s *BaseFhirpathListener) EnterInequalityExpression(ctx *InequalityExpressionContext) {}

// ExitInequalityExpression is called when production inequalityExpression is exited.
func (s *BaseFhirpathListener) ExitInequalityExpression(ctx *InequalityExpressionContext) {}

// EnterInvocationExpression is called when production invocationExpression is entered.
func (s *BaseFhirpathListener) EnterInvocationExpression(ctx *InvocationExpressionContext) {}

// ExitInvocationExpression is called when production invocationExpression is exited.
func (s *BaseFhirpathListener) ExitInvocationExpression(ctx *InvocationExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseFhirpathListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseFhirpathListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterImpliesExpression is called when production impliesExpression is entered.
func (s *BaseFhirpathListener) EnterImpliesExpression(ctx *ImpliesExpressionContext) {}

// ExitImpliesExpression is called when production impliesExpression is exited.
func (s *BaseFhirpathListener) ExitImpliesExpression(ctx *ImpliesExpressionContext) {}

// EnterTermExpression is called when production termExpression is entered.
func (s *BaseFhirpathListener) EnterTermExpression(ctx *TermExpressionContext) {}

// ExitTermExpression is called when production termExpression is exited.
func (s *BaseFhirpathListener) ExitTermExpression(ctx *TermExpressionContext) {}

// EnterTypeExpression is called when production typeExpression is entered.
func (s *BaseFhirpathListener) EnterTypeExpression(ctx *TypeExpressionContext) {}

// ExitTypeExpression is called when production typeExpression is exited.
func (s *BaseFhirpathListener) ExitTypeExpression(ctx *TypeExpressionContext) {}

// EnterInvocationTerm is called when production invocationTerm is entered.
func (s *BaseFhirpathListener) EnterInvocationTerm(ctx *InvocationTermContext) {}

// ExitInvocationTerm is called when production invocationTerm is exited.
func (s *BaseFhirpathListener) ExitInvocationTerm(ctx *InvocationTermContext) {}

// EnterLiteralTerm is called when production literalTerm is entered.
func (s *BaseFhirpathListener) EnterLiteralTerm(ctx *LiteralTermContext) {}

// ExitLiteralTerm is called when production literalTerm is exited.
func (s *BaseFhirpathListener) ExitLiteralTerm(ctx *LiteralTermContext) {}

// EnterExternalConstantTerm is called when production externalConstantTerm is entered.
func (s *BaseFhirpathListener) EnterExternalConstantTerm(ctx *ExternalConstantTermContext) {}

// ExitExternalConstantTerm is called when production externalConstantTerm is exited.
func (s *BaseFhirpathListener) ExitExternalConstantTerm(ctx *ExternalConstantTermContext) {}

// EnterParenthesizedTerm is called when production parenthesizedTerm is entered.
func (s *BaseFhirpathListener) EnterParenthesizedTerm(ctx *ParenthesizedTermContext) {}

// ExitParenthesizedTerm is called when production parenthesizedTerm is exited.
func (s *BaseFhirpathListener) ExitParenthesizedTerm(ctx *ParenthesizedTermContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseFhirpathListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseFhirpathListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseFhirpathListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseFhirpathListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseFhirpathListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseFhirpathListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseFhirpathListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseFhirpathListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterDateLiteral is called when production dateLiteral is entered.
func (s *BaseFhirpathListener) EnterDateLiteral(ctx *DateLiteralContext) {}

// ExitDateLiteral is called when production dateLiteral is exited.
func (s *BaseFhirpathListener) ExitDateLiteral(ctx *DateLiteralContext) {}

// EnterDateTimeLiteral is called when production dateTimeLiteral is entered.
func (s *BaseFhirpathListener) EnterDateTimeLiteral(ctx *DateTimeLiteralContext) {}

// ExitDateTimeLiteral is called when production dateTimeLiteral is exited.
func (s *BaseFhirpathListener) ExitDateTimeLiteral(ctx *DateTimeLiteralContext) {}

// EnterTimeLiteral is called when production timeLiteral is entered.
func (s *BaseFhirpathListener) EnterTimeLiteral(ctx *TimeLiteralContext) {}

// ExitTimeLiteral is called when production timeLiteral is exited.
func (s *BaseFhirpathListener) ExitTimeLiteral(ctx *TimeLiteralContext) {}

// EnterQuantityLiteral is called when production quantityLiteral is entered.
func (s *BaseFhirpathListener) EnterQuantityLiteral(ctx *QuantityLiteralContext) {}

// ExitQuantityLiteral is called when production quantityLiteral is exited.
func (s *BaseFhirpathListener) ExitQuantityLiteral(ctx *QuantityLiteralContext) {}

// EnterExternalConstant is called when production externalConstant is entered.
func (s *BaseFhirpathListener) EnterExternalConstant(ctx *ExternalConstantContext) {}

// ExitExternalConstant is called when production externalConstant is exited.
func (s *BaseFhirpathListener) ExitExternalConstant(ctx *ExternalConstantContext) {}

// EnterMemberInvocation is called when production memberInvocation is entered.
func (s *BaseFhirpathListener) EnterMemberInvocation(ctx *MemberInvocationContext) {}

// ExitMemberInvocation is called when production memberInvocation is exited.
func (s *BaseFhirpathListener) ExitMemberInvocation(ctx *MemberInvocationContext) {}

// EnterFunctionInvocation is called when production functionInvocation is entered.
func (s *BaseFhirpathListener) EnterFunctionInvocation(ctx *FunctionInvocationContext) {}

// ExitFunctionInvocation is called when production functionInvocation is exited.
func (s *BaseFhirpathListener) ExitFunctionInvocation(ctx *FunctionInvocationContext) {}

// EnterThisInvocation is called when production thisInvocation is entered.
func (s *BaseFhirpathListener) EnterThisInvocation(ctx *ThisInvocationContext) {}

// ExitThisInvocation is called when production thisInvocation is exited.
func (s *BaseFhirpathListener) ExitThisInvocation(ctx *ThisInvocationContext) {}

// EnterIndexInvocation is called when production indexInvocation is entered.
func (s *BaseFhirpathListener) EnterIndexInvocation(ctx *IndexInvocationContext) {}

// ExitIndexInvocation is called when production indexInvocation is exited.
func (s *BaseFhirpathListener) ExitIndexInvocation(ctx *IndexInvocationContext) {}

// EnterTotalInvocation is called when production totalInvocation is entered.
func (s *BaseFhirpathListener) EnterTotalInvocation(ctx *TotalInvocationContext) {}

// ExitTotalInvocation is called when production totalInvocation is exited.
func (s *BaseFhirpathListener) ExitTotalInvocation(ctx *TotalInvocationContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseFhirpathListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseFhirpathListener) ExitFunction(ctx *FunctionContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseFhirpathListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseFhirpathListener) ExitParamList(ctx *ParamListContext) {}

// EnterQuantity is called when production quantity is entered.
func (s *BaseFhirpathListener) EnterQuantity(ctx *QuantityContext) {}

// ExitQuantity is called when production quantity is exited.
func (s *BaseFhirpathListener) ExitQuantity(ctx *QuantityContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseFhirpathListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseFhirpathListener) ExitUnit(ctx *UnitContext) {}

// EnterDateTimePrecision is called when production dateTimePrecision is entered.
func (s *BaseFhirpathListener) EnterDateTimePrecision(ctx *DateTimePrecisionContext) {}

// ExitDateTimePrecision is called when production dateTimePrecision is exited.
func (s *BaseFhirpathListener) ExitDateTimePrecision(ctx *DateTimePrecisionContext) {}

// EnterPluralDateTimePrecision is called when production pluralDateTimePrecision is entered.
func (s *BaseFhirpathListener) EnterPluralDateTimePrecision(ctx *PluralDateTimePrecisionContext) {}

// ExitPluralDateTimePrecision is called when production pluralDateTimePrecision is exited.
func (s *BaseFhirpathListener) ExitPluralDateTimePrecision(ctx *PluralDateTimePrecisionContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BaseFhirpathListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BaseFhirpathListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterQualifiedIdentifier is called when production qualifiedIdentifier is entered.
func (s *BaseFhirpathListener) EnterQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// ExitQualifiedIdentifier is called when production qualifiedIdentifier is exited.
func (s *BaseFhirpathListener) ExitQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseFhirpathListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseFhirpathListener) ExitIdentifier(ctx *IdentifierContext) {}
