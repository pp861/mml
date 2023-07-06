// Code generated from Mml.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mml

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMmlListener is a complete listener for a parse tree produced by MmlParser.
type BaseMmlListener struct{}

var _ MmlListener = &BaseMmlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMmlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMmlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMmlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMmlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseMmlListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseMmlListener) ExitProg(ctx *ProgContext) {}

// EnterRuleLine is called when production ruleLine is entered.
func (s *BaseMmlListener) EnterRuleLine(ctx *RuleLineContext) {}

// ExitRuleLine is called when production ruleLine is exited.
func (s *BaseMmlListener) ExitRuleLine(ctx *RuleLineContext) {}

// EnterOpType is called when production opType is entered.
func (s *BaseMmlListener) EnterOpType(ctx *OpTypeContext) {}

// ExitOpType is called when production opType is exited.
func (s *BaseMmlListener) ExitOpType(ctx *OpTypeContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseMmlListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseMmlListener) ExitAttribute(ctx *AttributeContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseMmlListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseMmlListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterUnderScore is called when production underScore is entered.
func (s *BaseMmlListener) EnterUnderScore(ctx *UnderScoreContext) {}

// ExitUnderScore is called when production underScore is exited.
func (s *BaseMmlListener) ExitUnderScore(ctx *UnderScoreContext) {}

// EnterMmlList is called when production mmlList is entered.
func (s *BaseMmlListener) EnterMmlList(ctx *MmlListContext) {}

// ExitMmlList is called when production mmlList is exited.
func (s *BaseMmlListener) ExitMmlList(ctx *MmlListContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseMmlListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseMmlListener) ExitCondition(ctx *ConditionContext) {}

// EnterValue is called when production value is entered.
func (s *BaseMmlListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseMmlListener) ExitValue(ctx *ValueContext) {}

// EnterNumberInt is called when production numberInt is entered.
func (s *BaseMmlListener) EnterNumberInt(ctx *NumberIntContext) {}

// ExitNumberInt is called when production numberInt is exited.
func (s *BaseMmlListener) ExitNumberInt(ctx *NumberIntContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseMmlListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseMmlListener) ExitNumber(ctx *NumberContext) {}

// EnterQuotedString is called when production quotedString is entered.
func (s *BaseMmlListener) EnterQuotedString(ctx *QuotedStringContext) {}

// ExitQuotedString is called when production quotedString is exited.
func (s *BaseMmlListener) ExitQuotedString(ctx *QuotedStringContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseMmlListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseMmlListener) ExitPattern(ctx *PatternContext) {}
