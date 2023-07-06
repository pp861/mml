// Code generated from Mml.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mml

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MmlParser struct {
	*antlr.BaseParser
}

var mmlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mmlParserInit() {
	staticData := &mmlParserStaticData
	staticData.literalNames = []string{
		"", "':'", "';'", "','", "'='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "ACT", "BLK", "CHECK", "CLB", "DEL", "DLD", "DSP",
		"HIS", "HISTORY", "LST", "MOD", "PING", "POWEROP", "QRY", "RBK", "RMV",
		"RST", "SCAN", "SET", "SHW", "SND", "STOP", "SWITCH", "SYNC", "UNACT",
		"UNBLK", "UPLD", "ADD", "Attribute", "Identifier", "Letter", "Value",
		"Digit", "NumberInt", "IntegerLiteral", "Number", "QuotedString", "ESC",
		"Pattern", "WS",
	}
	staticData.ruleNames = []string{
		"prog", "ruleLine", "opType", "mmlList", "condition",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 36, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 4, 0, 12, 8, 0, 11, 0, 12, 0, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 27, 8, 3, 10, 3, 12, 3, 30, 9,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 0, 0, 5, 0, 2, 4, 6, 8, 0, 1, 1, 0, 5,
		32, 32, 0, 11, 1, 0, 0, 0, 2, 15, 1, 0, 0, 0, 4, 21, 1, 0, 0, 0, 6, 23,
		1, 0, 0, 0, 8, 31, 1, 0, 0, 0, 10, 12, 3, 2, 1, 0, 11, 10, 1, 0, 0, 0,
		12, 13, 1, 0, 0, 0, 13, 11, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 1, 1, 0,
		0, 0, 15, 16, 3, 4, 2, 0, 16, 17, 5, 33, 0, 0, 17, 18, 5, 1, 0, 0, 18,
		19, 3, 6, 3, 0, 19, 20, 5, 2, 0, 0, 20, 3, 1, 0, 0, 0, 21, 22, 7, 0, 0,
		0, 22, 5, 1, 0, 0, 0, 23, 28, 3, 8, 4, 0, 24, 25, 5, 3, 0, 0, 25, 27, 3,
		8, 4, 0, 26, 24, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28,
		29, 1, 0, 0, 0, 29, 7, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 32, 5, 33, 0,
		0, 32, 33, 5, 4, 0, 0, 33, 34, 5, 36, 0, 0, 34, 9, 1, 0, 0, 0, 2, 13, 28,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MmlParserInit initializes any static state used to implement MmlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMmlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MmlParserInit() {
	staticData := &mmlParserStaticData
	staticData.once.Do(mmlParserInit)
}

// NewMmlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMmlParser(input antlr.TokenStream) *MmlParser {
	MmlParserInit()
	this := new(MmlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &mmlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Mml.g4"

	return this
}

// MmlParser tokens.
const (
	MmlParserEOF            = antlr.TokenEOF
	MmlParserT__0           = 1
	MmlParserT__1           = 2
	MmlParserT__2           = 3
	MmlParserT__3           = 4
	MmlParserACT            = 5
	MmlParserBLK            = 6
	MmlParserCHECK          = 7
	MmlParserCLB            = 8
	MmlParserDEL            = 9
	MmlParserDLD            = 10
	MmlParserDSP            = 11
	MmlParserHIS            = 12
	MmlParserHISTORY        = 13
	MmlParserLST            = 14
	MmlParserMOD            = 15
	MmlParserPING           = 16
	MmlParserPOWEROP        = 17
	MmlParserQRY            = 18
	MmlParserRBK            = 19
	MmlParserRMV            = 20
	MmlParserRST            = 21
	MmlParserSCAN           = 22
	MmlParserSET            = 23
	MmlParserSHW            = 24
	MmlParserSND            = 25
	MmlParserSTOP           = 26
	MmlParserSWITCH         = 27
	MmlParserSYNC           = 28
	MmlParserUNACT          = 29
	MmlParserUNBLK          = 30
	MmlParserUPLD           = 31
	MmlParserADD            = 32
	MmlParserAttribute      = 33
	MmlParserIdentifier     = 34
	MmlParserLetter         = 35
	MmlParserValue          = 36
	MmlParserDigit          = 37
	MmlParserNumberInt      = 38
	MmlParserIntegerLiteral = 39
	MmlParserNumber         = 40
	MmlParserQuotedString   = 41
	MmlParserESC            = 42
	MmlParserPattern        = 43
	MmlParserWS             = 44
)

// MmlParser rules.
const (
	MmlParserRULE_prog      = 0
	MmlParserRULE_ruleLine  = 1
	MmlParserRULE_opType    = 2
	MmlParserRULE_mmlList   = 3
	MmlParserRULE_condition = 4
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllRuleLine() []IRuleLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleLineContext); ok {
			len++
		}
	}

	tst := make([]IRuleLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleLineContext); ok {
			tst[i] = t.(IRuleLineContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) RuleLine(i int) IRuleLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleLineContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *MmlParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MmlParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(MmlParserACT-5))|(1<<(MmlParserBLK-5))|(1<<(MmlParserCHECK-5))|(1<<(MmlParserCLB-5))|(1<<(MmlParserDEL-5))|(1<<(MmlParserDLD-5))|(1<<(MmlParserDSP-5))|(1<<(MmlParserHIS-5))|(1<<(MmlParserHISTORY-5))|(1<<(MmlParserLST-5))|(1<<(MmlParserMOD-5))|(1<<(MmlParserPING-5))|(1<<(MmlParserPOWEROP-5))|(1<<(MmlParserQRY-5))|(1<<(MmlParserRBK-5))|(1<<(MmlParserRMV-5))|(1<<(MmlParserRST-5))|(1<<(MmlParserSCAN-5))|(1<<(MmlParserSET-5))|(1<<(MmlParserSHW-5))|(1<<(MmlParserSND-5))|(1<<(MmlParserSTOP-5))|(1<<(MmlParserSWITCH-5))|(1<<(MmlParserSYNC-5))|(1<<(MmlParserUNACT-5))|(1<<(MmlParserUNBLK-5))|(1<<(MmlParserUPLD-5))|(1<<(MmlParserADD-5)))) != 0) {
		{
			p.SetState(10)
			p.RuleLine()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRuleLineContext is an interface to support dynamic dispatch.
type IRuleLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleLineContext differentiates from other interfaces.
	IsRuleLineContext()
}

type RuleLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleLineContext() *RuleLineContext {
	var p = new(RuleLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_ruleLine
	return p
}

func (*RuleLineContext) IsRuleLineContext() {}

func NewRuleLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleLineContext {
	var p = new(RuleLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_ruleLine

	return p
}

func (s *RuleLineContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleLineContext) OpType() IOpTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpTypeContext)
}

func (s *RuleLineContext) Attribute() antlr.TerminalNode {
	return s.GetToken(MmlParserAttribute, 0)
}

func (s *RuleLineContext) MmlList() IMmlListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMmlListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMmlListContext)
}

func (s *RuleLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterRuleLine(s)
	}
}

func (s *RuleLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitRuleLine(s)
	}
}

func (p *MmlParser) RuleLine() (localctx IRuleLineContext) {
	this := p
	_ = this

	localctx = NewRuleLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MmlParserRULE_ruleLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(15)
		p.OpType()
	}
	{
		p.SetState(16)
		p.Match(MmlParserAttribute)
	}
	{
		p.SetState(17)
		p.Match(MmlParserT__0)
	}
	{
		p.SetState(18)
		p.MmlList()
	}
	{
		p.SetState(19)
		p.Match(MmlParserT__1)
	}

	return localctx
}

// IOpTypeContext is an interface to support dynamic dispatch.
type IOpTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpTypeContext differentiates from other interfaces.
	IsOpTypeContext()
}

type OpTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpTypeContext() *OpTypeContext {
	var p = new(OpTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_opType
	return p
}

func (*OpTypeContext) IsOpTypeContext() {}

func NewOpTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpTypeContext {
	var p = new(OpTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_opType

	return p
}

func (s *OpTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *OpTypeContext) ACT() antlr.TerminalNode {
	return s.GetToken(MmlParserACT, 0)
}

func (s *OpTypeContext) BLK() antlr.TerminalNode {
	return s.GetToken(MmlParserBLK, 0)
}

func (s *OpTypeContext) CHECK() antlr.TerminalNode {
	return s.GetToken(MmlParserCHECK, 0)
}

func (s *OpTypeContext) CLB() antlr.TerminalNode {
	return s.GetToken(MmlParserCLB, 0)
}

func (s *OpTypeContext) DEL() antlr.TerminalNode {
	return s.GetToken(MmlParserDEL, 0)
}

func (s *OpTypeContext) DLD() antlr.TerminalNode {
	return s.GetToken(MmlParserDLD, 0)
}

func (s *OpTypeContext) DSP() antlr.TerminalNode {
	return s.GetToken(MmlParserDSP, 0)
}

func (s *OpTypeContext) HIS() antlr.TerminalNode {
	return s.GetToken(MmlParserHIS, 0)
}

func (s *OpTypeContext) HISTORY() antlr.TerminalNode {
	return s.GetToken(MmlParserHISTORY, 0)
}

func (s *OpTypeContext) LST() antlr.TerminalNode {
	return s.GetToken(MmlParserLST, 0)
}

func (s *OpTypeContext) MOD() antlr.TerminalNode {
	return s.GetToken(MmlParserMOD, 0)
}

func (s *OpTypeContext) PING() antlr.TerminalNode {
	return s.GetToken(MmlParserPING, 0)
}

func (s *OpTypeContext) POWEROP() antlr.TerminalNode {
	return s.GetToken(MmlParserPOWEROP, 0)
}

func (s *OpTypeContext) QRY() antlr.TerminalNode {
	return s.GetToken(MmlParserQRY, 0)
}

func (s *OpTypeContext) RBK() antlr.TerminalNode {
	return s.GetToken(MmlParserRBK, 0)
}

func (s *OpTypeContext) RMV() antlr.TerminalNode {
	return s.GetToken(MmlParserRMV, 0)
}

func (s *OpTypeContext) RST() antlr.TerminalNode {
	return s.GetToken(MmlParserRST, 0)
}

func (s *OpTypeContext) SCAN() antlr.TerminalNode {
	return s.GetToken(MmlParserSCAN, 0)
}

func (s *OpTypeContext) SET() antlr.TerminalNode {
	return s.GetToken(MmlParserSET, 0)
}

func (s *OpTypeContext) SHW() antlr.TerminalNode {
	return s.GetToken(MmlParserSHW, 0)
}

func (s *OpTypeContext) SND() antlr.TerminalNode {
	return s.GetToken(MmlParserSND, 0)
}

func (s *OpTypeContext) STOP() antlr.TerminalNode {
	return s.GetToken(MmlParserSTOP, 0)
}

func (s *OpTypeContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(MmlParserSWITCH, 0)
}

func (s *OpTypeContext) SYNC() antlr.TerminalNode {
	return s.GetToken(MmlParserSYNC, 0)
}

func (s *OpTypeContext) UNACT() antlr.TerminalNode {
	return s.GetToken(MmlParserUNACT, 0)
}

func (s *OpTypeContext) UNBLK() antlr.TerminalNode {
	return s.GetToken(MmlParserUNBLK, 0)
}

func (s *OpTypeContext) UPLD() antlr.TerminalNode {
	return s.GetToken(MmlParserUPLD, 0)
}

func (s *OpTypeContext) ADD() antlr.TerminalNode {
	return s.GetToken(MmlParserADD, 0)
}

func (s *OpTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterOpType(s)
	}
}

func (s *OpTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitOpType(s)
	}
}

func (p *MmlParser) OpType() (localctx IOpTypeContext) {
	this := p
	_ = this

	localctx = NewOpTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MmlParserRULE_opType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(21)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(MmlParserACT-5))|(1<<(MmlParserBLK-5))|(1<<(MmlParserCHECK-5))|(1<<(MmlParserCLB-5))|(1<<(MmlParserDEL-5))|(1<<(MmlParserDLD-5))|(1<<(MmlParserDSP-5))|(1<<(MmlParserHIS-5))|(1<<(MmlParserHISTORY-5))|(1<<(MmlParserLST-5))|(1<<(MmlParserMOD-5))|(1<<(MmlParserPING-5))|(1<<(MmlParserPOWEROP-5))|(1<<(MmlParserQRY-5))|(1<<(MmlParserRBK-5))|(1<<(MmlParserRMV-5))|(1<<(MmlParserRST-5))|(1<<(MmlParserSCAN-5))|(1<<(MmlParserSET-5))|(1<<(MmlParserSHW-5))|(1<<(MmlParserSND-5))|(1<<(MmlParserSTOP-5))|(1<<(MmlParserSWITCH-5))|(1<<(MmlParserSYNC-5))|(1<<(MmlParserUNACT-5))|(1<<(MmlParserUNBLK-5))|(1<<(MmlParserUPLD-5))|(1<<(MmlParserADD-5)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMmlListContext is an interface to support dynamic dispatch.
type IMmlListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMmlListContext differentiates from other interfaces.
	IsMmlListContext()
}

type MmlListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMmlListContext() *MmlListContext {
	var p = new(MmlListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_mmlList
	return p
}

func (*MmlListContext) IsMmlListContext() {}

func NewMmlListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MmlListContext {
	var p = new(MmlListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_mmlList

	return p
}

func (s *MmlListContext) GetParser() antlr.Parser { return s.parser }

func (s *MmlListContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *MmlListContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *MmlListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MmlListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MmlListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterMmlList(s)
	}
}

func (s *MmlListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitMmlList(s)
	}
}

func (p *MmlParser) MmlList() (localctx IMmlListContext) {
	this := p
	_ = this

	localctx = NewMmlListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MmlParserRULE_mmlList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Condition()
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MmlParserT__2 {
		{
			p.SetState(24)
			p.Match(MmlParserT__2)
		}
		{
			p.SetState(25)
			p.Condition()
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Attribute() antlr.TerminalNode {
	return s.GetToken(MmlParserAttribute, 0)
}

func (s *ConditionContext) Value() antlr.TerminalNode {
	return s.GetToken(MmlParserValue, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *MmlParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MmlParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(MmlParserAttribute)
	}
	{
		p.SetState(32)
		p.Match(MmlParserT__3)
	}
	{
		p.SetState(33)
		p.Match(MmlParserValue)
	}

	return localctx
}
