// Code generated from Mml.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mml

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MmlListener is a complete listener for a parse tree produced by MmlParser.
type MmlListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterRuleLine is called when entering the ruleLine production.
	EnterRuleLine(c *RuleLineContext)

	// EnterOpType is called when entering the opType production.
	EnterOpType(c *OpTypeContext)

	// EnterMmlList is called when entering the mmlList production.
	EnterMmlList(c *MmlListContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitRuleLine is called when exiting the ruleLine production.
	ExitRuleLine(c *RuleLineContext)

	// ExitOpType is called when exiting the opType production.
	ExitOpType(c *OpTypeContext)

	// ExitMmlList is called when exiting the mmlList production.
	ExitMmlList(c *MmlListContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)
}
