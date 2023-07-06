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

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterUnderScore is called when entering the underScore production.
	EnterUnderScore(c *UnderScoreContext)

	// EnterMmlList is called when entering the mmlList production.
	EnterMmlList(c *MmlListContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterNumberInt is called when entering the numberInt production.
	EnterNumberInt(c *NumberIntContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterQuotedString is called when entering the quotedString production.
	EnterQuotedString(c *QuotedStringContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitRuleLine is called when exiting the ruleLine production.
	ExitRuleLine(c *RuleLineContext)

	// ExitOpType is called when exiting the opType production.
	ExitOpType(c *OpTypeContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitUnderScore is called when exiting the underScore production.
	ExitUnderScore(c *UnderScoreContext)

	// ExitMmlList is called when exiting the mmlList production.
	ExitMmlList(c *MmlListContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitNumberInt is called when exiting the numberInt production.
	ExitNumberInt(c *NumberIntContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitQuotedString is called when exiting the quotedString production.
	ExitQuotedString(c *QuotedStringContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)
}
