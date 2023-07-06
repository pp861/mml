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

// EnterMmlList is called when production mmlList is entered.
func (s *BaseMmlListener) EnterMmlList(ctx *MmlListContext) {}

// ExitMmlList is called when production mmlList is exited.
func (s *BaseMmlListener) ExitMmlList(ctx *MmlListContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseMmlListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseMmlListener) ExitCondition(ctx *ConditionContext) {}
