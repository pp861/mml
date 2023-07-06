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
		"", "':'", "';'", "'ACT'", "'BLK'", "'CHECK'", "'CLB'", "'DEL'", "'DLD'",
		"'DSP'", "'HIS'", "'HISTORY'", "'LST'", "'MOD'", "'PING'", "'POWEROP'",
		"'QRY'", "'RBK'", "'RMV'", "'RST'", "'SCAN'", "'SET'", "'SHW'", "'SND'",
		"'STOP'", "'SWITCH'", "'SYNC'", "'UNACT'", "'UNBLK'", "'UPLD'", "'ADD'",
		"'_'", "','", "'='", "'.'", "'\"'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "Letter", "Digit", "IntegerLiteral", "ESC", "WS",
	}
	staticData.ruleNames = []string{
		"prog", "ruleLine", "opType", "attribute", "identifier", "underScore",
		"mmlList", "condition", "value", "numberInt", "number", "quotedString",
		"pattern",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 119, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 4, 0, 28, 8, 0, 11, 0, 12, 0, 29,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 49, 8, 4, 10, 4, 12, 4, 52, 9, 4, 1,
		4, 4, 4, 55, 8, 4, 11, 4, 12, 4, 56, 3, 4, 59, 8, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 5, 6, 66, 8, 6, 10, 6, 12, 6, 69, 9, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 3, 8, 77, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 83, 8, 9,
		1, 10, 4, 10, 86, 8, 10, 11, 10, 12, 10, 87, 1, 10, 5, 10, 91, 8, 10, 10,
		10, 12, 10, 94, 9, 10, 1, 10, 1, 10, 5, 10, 98, 8, 10, 10, 10, 12, 10,
		101, 9, 10, 3, 10, 103, 8, 10, 1, 11, 1, 11, 1, 11, 5, 11, 108, 8, 11,
		10, 11, 12, 11, 111, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 109, 0, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 2,
		1, 0, 3, 30, 1, 0, 36, 37, 123, 0, 27, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4,
		37, 1, 0, 0, 0, 6, 39, 1, 0, 0, 0, 8, 58, 1, 0, 0, 0, 10, 60, 1, 0, 0,
		0, 12, 62, 1, 0, 0, 0, 14, 70, 1, 0, 0, 0, 16, 76, 1, 0, 0, 0, 18, 82,
		1, 0, 0, 0, 20, 102, 1, 0, 0, 0, 22, 104, 1, 0, 0, 0, 24, 114, 1, 0, 0,
		0, 26, 28, 3, 2, 1, 0, 27, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 27,
		1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 1, 1, 0, 0, 0, 31, 32, 3, 4, 2, 0,
		32, 33, 3, 6, 3, 0, 33, 34, 5, 1, 0, 0, 34, 35, 3, 12, 6, 0, 35, 36, 5,
		2, 0, 0, 36, 3, 1, 0, 0, 0, 37, 38, 7, 0, 0, 0, 38, 5, 1, 0, 0, 0, 39,
		40, 3, 8, 4, 0, 40, 7, 1, 0, 0, 0, 41, 59, 5, 36, 0, 0, 42, 43, 5, 36,
		0, 0, 43, 59, 5, 37, 0, 0, 44, 50, 5, 36, 0, 0, 45, 49, 5, 36, 0, 0, 46,
		49, 5, 37, 0, 0, 47, 49, 3, 10, 5, 0, 48, 45, 1, 0, 0, 0, 48, 46, 1, 0,
		0, 0, 48, 47, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51,
		1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 55, 7, 1, 0, 0,
		54, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1,
		0, 0, 0, 57, 59, 1, 0, 0, 0, 58, 41, 1, 0, 0, 0, 58, 42, 1, 0, 0, 0, 58,
		44, 1, 0, 0, 0, 59, 9, 1, 0, 0, 0, 60, 61, 5, 31, 0, 0, 61, 11, 1, 0, 0,
		0, 62, 67, 3, 14, 7, 0, 63, 64, 5, 32, 0, 0, 64, 66, 3, 14, 7, 0, 65, 63,
		1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0,
		68, 13, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 71, 3, 6, 3, 0, 71, 72, 5,
		33, 0, 0, 72, 73, 3, 16, 8, 0, 73, 15, 1, 0, 0, 0, 74, 77, 3, 18, 9, 0,
		75, 77, 3, 24, 12, 0, 76, 74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 17, 1,
		0, 0, 0, 78, 83, 5, 38, 0, 0, 79, 83, 5, 37, 0, 0, 80, 83, 3, 20, 10, 0,
		81, 83, 3, 22, 11, 0, 82, 78, 1, 0, 0, 0, 82, 79, 1, 0, 0, 0, 82, 80, 1,
		0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 19, 1, 0, 0, 0, 84, 86, 5, 37, 0, 0, 85,
		84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0,
		0, 88, 103, 1, 0, 0, 0, 89, 91, 5, 37, 0, 0, 90, 89, 1, 0, 0, 0, 91, 94,
		1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0,
		94, 92, 1, 0, 0, 0, 95, 99, 5, 34, 0, 0, 96, 98, 5, 37, 0, 0, 97, 96, 1,
		0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0,
		100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 85, 1, 0, 0, 0, 102, 92,
		1, 0, 0, 0, 103, 21, 1, 0, 0, 0, 104, 109, 5, 35, 0, 0, 105, 108, 5, 39,
		0, 0, 106, 108, 9, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 106, 1, 0, 0, 0,
		108, 111, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110,
		112, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 113, 5, 35, 0, 0, 113, 23,
		1, 0, 0, 0, 114, 115, 5, 35, 0, 0, 115, 116, 3, 8, 4, 0, 116, 117, 5, 35,
		0, 0, 117, 25, 1, 0, 0, 0, 14, 29, 48, 50, 56, 58, 67, 76, 82, 87, 92,
		99, 102, 107, 109,
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
	MmlParserT__4           = 5
	MmlParserT__5           = 6
	MmlParserT__6           = 7
	MmlParserT__7           = 8
	MmlParserT__8           = 9
	MmlParserT__9           = 10
	MmlParserT__10          = 11
	MmlParserT__11          = 12
	MmlParserT__12          = 13
	MmlParserT__13          = 14
	MmlParserT__14          = 15
	MmlParserT__15          = 16
	MmlParserT__16          = 17
	MmlParserT__17          = 18
	MmlParserT__18          = 19
	MmlParserT__19          = 20
	MmlParserT__20          = 21
	MmlParserT__21          = 22
	MmlParserT__22          = 23
	MmlParserT__23          = 24
	MmlParserT__24          = 25
	MmlParserT__25          = 26
	MmlParserT__26          = 27
	MmlParserT__27          = 28
	MmlParserT__28          = 29
	MmlParserT__29          = 30
	MmlParserT__30          = 31
	MmlParserT__31          = 32
	MmlParserT__32          = 33
	MmlParserT__33          = 34
	MmlParserT__34          = 35
	MmlParserLetter         = 36
	MmlParserDigit          = 37
	MmlParserIntegerLiteral = 38
	MmlParserESC            = 39
	MmlParserWS             = 40
)

// MmlParser rules.
const (
	MmlParserRULE_prog         = 0
	MmlParserRULE_ruleLine     = 1
	MmlParserRULE_opType       = 2
	MmlParserRULE_attribute    = 3
	MmlParserRULE_identifier   = 4
	MmlParserRULE_underScore   = 5
	MmlParserRULE_mmlList      = 6
	MmlParserRULE_condition    = 7
	MmlParserRULE_value        = 8
	MmlParserRULE_numberInt    = 9
	MmlParserRULE_number       = 10
	MmlParserRULE_quotedString = 11
	MmlParserRULE_pattern      = 12
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MmlParserT__2)|(1<<MmlParserT__3)|(1<<MmlParserT__4)|(1<<MmlParserT__5)|(1<<MmlParserT__6)|(1<<MmlParserT__7)|(1<<MmlParserT__8)|(1<<MmlParserT__9)|(1<<MmlParserT__10)|(1<<MmlParserT__11)|(1<<MmlParserT__12)|(1<<MmlParserT__13)|(1<<MmlParserT__14)|(1<<MmlParserT__15)|(1<<MmlParserT__16)|(1<<MmlParserT__17)|(1<<MmlParserT__18)|(1<<MmlParserT__19)|(1<<MmlParserT__20)|(1<<MmlParserT__21)|(1<<MmlParserT__22)|(1<<MmlParserT__23)|(1<<MmlParserT__24)|(1<<MmlParserT__25)|(1<<MmlParserT__26)|(1<<MmlParserT__27)|(1<<MmlParserT__28)|(1<<MmlParserT__29))) != 0) {
		{
			p.SetState(26)
			p.RuleLine()
		}

		p.SetState(29)
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

func (s *RuleLineContext) Attribute() IAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
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
		p.SetState(31)
		p.OpType()
	}
	{
		p.SetState(32)
		p.Attribute()
	}
	{
		p.SetState(33)
		p.Match(MmlParserT__0)
	}
	{
		p.SetState(34)
		p.MmlList()
	}
	{
		p.SetState(35)
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
		p.SetState(37)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MmlParserT__2)|(1<<MmlParserT__3)|(1<<MmlParserT__4)|(1<<MmlParserT__5)|(1<<MmlParserT__6)|(1<<MmlParserT__7)|(1<<MmlParserT__8)|(1<<MmlParserT__9)|(1<<MmlParserT__10)|(1<<MmlParserT__11)|(1<<MmlParserT__12)|(1<<MmlParserT__13)|(1<<MmlParserT__14)|(1<<MmlParserT__15)|(1<<MmlParserT__16)|(1<<MmlParserT__17)|(1<<MmlParserT__18)|(1<<MmlParserT__19)|(1<<MmlParserT__20)|(1<<MmlParserT__21)|(1<<MmlParserT__22)|(1<<MmlParserT__23)|(1<<MmlParserT__24)|(1<<MmlParserT__25)|(1<<MmlParserT__26)|(1<<MmlParserT__27)|(1<<MmlParserT__28)|(1<<MmlParserT__29))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (p *MmlParser) Attribute() (localctx IAttributeContext) {
	this := p
	_ = this

	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MmlParserRULE_attribute)

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
		p.SetState(39)
		p.Identifier()
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) AllLetter() []antlr.TerminalNode {
	return s.GetTokens(MmlParserLetter)
}

func (s *IdentifierContext) Letter(i int) antlr.TerminalNode {
	return s.GetToken(MmlParserLetter, i)
}

func (s *IdentifierContext) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(MmlParserDigit)
}

func (s *IdentifierContext) Digit(i int) antlr.TerminalNode {
	return s.GetToken(MmlParserDigit, i)
}

func (s *IdentifierContext) AllUnderScore() []IUnderScoreContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnderScoreContext); ok {
			len++
		}
	}

	tst := make([]IUnderScoreContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnderScoreContext); ok {
			tst[i] = t.(IUnderScoreContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierContext) UnderScore(i int) IUnderScoreContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnderScoreContext); ok {
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

	return t.(IUnderScoreContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *MmlParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MmlParserRULE_identifier)
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

	var _alt int

	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(MmlParserLetter)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Match(MmlParserLetter)
		}
		{
			p.SetState(43)
			p.Match(MmlParserDigit)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.Match(MmlParserLetter)
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(48)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case MmlParserLetter:
					{
						p.SetState(45)
						p.Match(MmlParserLetter)
					}

				case MmlParserDigit:
					{
						p.SetState(46)
						p.Match(MmlParserDigit)
					}

				case MmlParserT__30:
					{
						p.SetState(47)
						p.UnderScore()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			}
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == MmlParserLetter || _la == MmlParserDigit {
			{
				p.SetState(53)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MmlParserLetter || _la == MmlParserDigit) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IUnderScoreContext is an interface to support dynamic dispatch.
type IUnderScoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnderScoreContext differentiates from other interfaces.
	IsUnderScoreContext()
}

type UnderScoreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnderScoreContext() *UnderScoreContext {
	var p = new(UnderScoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_underScore
	return p
}

func (*UnderScoreContext) IsUnderScoreContext() {}

func NewUnderScoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnderScoreContext {
	var p = new(UnderScoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_underScore

	return p
}

func (s *UnderScoreContext) GetParser() antlr.Parser { return s.parser }
func (s *UnderScoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnderScoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnderScoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterUnderScore(s)
	}
}

func (s *UnderScoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitUnderScore(s)
	}
}

func (p *MmlParser) UnderScore() (localctx IUnderScoreContext) {
	this := p
	_ = this

	localctx = NewUnderScoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MmlParserRULE_underScore)

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
		p.SetState(60)
		p.Match(MmlParserT__30)
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
	p.EnterRule(localctx, 12, MmlParserRULE_mmlList)
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
		p.SetState(62)
		p.Condition()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MmlParserT__31 {
		{
			p.SetState(63)
			p.Match(MmlParserT__31)
		}
		{
			p.SetState(64)
			p.Condition()
		}

		p.SetState(69)
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

func (s *ConditionContext) Attribute() IAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *ConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
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
	p.EnterRule(localctx, 14, MmlParserRULE_condition)

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
		p.SetState(70)
		p.Attribute()
	}
	{
		p.SetState(71)
		p.Match(MmlParserT__32)
	}
	{
		p.SetState(72)
		p.Value()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) NumberInt() INumberIntContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberIntContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberIntContext)
}

func (s *ValueContext) Pattern() IPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *MmlParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MmlParserRULE_value)

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

	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.NumberInt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Pattern()
		}

	}

	return localctx
}

// INumberIntContext is an interface to support dynamic dispatch.
type INumberIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberIntContext differentiates from other interfaces.
	IsNumberIntContext()
}

type NumberIntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberIntContext() *NumberIntContext {
	var p = new(NumberIntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_numberInt
	return p
}

func (*NumberIntContext) IsNumberIntContext() {}

func NewNumberIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberIntContext {
	var p = new(NumberIntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_numberInt

	return p
}

func (s *NumberIntContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberIntContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(MmlParserIntegerLiteral, 0)
}

func (s *NumberIntContext) Digit() antlr.TerminalNode {
	return s.GetToken(MmlParserDigit, 0)
}

func (s *NumberIntContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *NumberIntContext) QuotedString() IQuotedStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuotedStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuotedStringContext)
}

func (s *NumberIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberIntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterNumberInt(s)
	}
}

func (s *NumberIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitNumberInt(s)
	}
}

func (p *MmlParser) NumberInt() (localctx INumberIntContext) {
	this := p
	_ = this

	localctx = NewNumberIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MmlParserRULE_numberInt)

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

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Match(MmlParserIntegerLiteral)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Match(MmlParserDigit)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.Number()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(81)
			p.QuotedString()
		}

	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(MmlParserDigit)
}

func (s *NumberContext) Digit(i int) antlr.TerminalNode {
	return s.GetToken(MmlParserDigit, i)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *MmlParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MmlParserRULE_number)
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

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == MmlParserDigit {
			{
				p.SetState(84)
				p.Match(MmlParserDigit)
			}

			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MmlParserDigit {
			{
				p.SetState(89)
				p.Match(MmlParserDigit)
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(95)
			p.Match(MmlParserT__33)
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MmlParserDigit {
			{
				p.SetState(96)
				p.Match(MmlParserDigit)
			}

			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IQuotedStringContext is an interface to support dynamic dispatch.
type IQuotedStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedStringContext differentiates from other interfaces.
	IsQuotedStringContext()
}

type QuotedStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedStringContext() *QuotedStringContext {
	var p = new(QuotedStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_quotedString
	return p
}

func (*QuotedStringContext) IsQuotedStringContext() {}

func NewQuotedStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedStringContext {
	var p = new(QuotedStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_quotedString

	return p
}

func (s *QuotedStringContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedStringContext) AllESC() []antlr.TerminalNode {
	return s.GetTokens(MmlParserESC)
}

func (s *QuotedStringContext) ESC(i int) antlr.TerminalNode {
	return s.GetToken(MmlParserESC, i)
}

func (s *QuotedStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterQuotedString(s)
	}
}

func (s *QuotedStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitQuotedString(s)
	}
}

func (p *MmlParser) QuotedString() (localctx IQuotedStringContext) {
	this := p
	_ = this

	localctx = NewQuotedStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MmlParserRULE_quotedString)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(MmlParserT__34)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(105)
					p.Match(MmlParserESC)
				}

			case 2:
				p.SetState(106)
				p.MatchWildcard()

			}

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	{
		p.SetState(112)
		p.Match(MmlParserT__34)
	}

	return localctx
}

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MmlParserRULE_pattern
	return p
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MmlParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MmlListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (p *MmlParser) Pattern() (localctx IPatternContext) {
	this := p
	_ = this

	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MmlParserRULE_pattern)

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
		p.SetState(114)
		p.Match(MmlParserT__34)
	}
	{
		p.SetState(115)
		p.Identifier()
	}
	{
		p.SetState(116)
		p.Match(MmlParserT__34)
	}

	return localctx
}
