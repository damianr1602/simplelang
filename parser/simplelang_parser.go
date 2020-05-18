// Code generated from simplelang.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // simplelang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 126,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 5, 3, 24, 10, 3,
	3, 3, 7, 3, 27, 10, 3, 12, 3, 14, 3, 30, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 42, 10, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 50, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 70, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	5, 5, 81, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	5, 6, 92, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 110, 10, 7, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 7, 10, 121, 10, 10, 12, 10,
	14, 10, 124, 11, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2,
	2, 2, 139, 2, 20, 3, 2, 2, 2, 4, 28, 3, 2, 2, 2, 6, 69, 3, 2, 2, 2, 8,
	80, 3, 2, 2, 2, 10, 91, 3, 2, 2, 2, 12, 109, 3, 2, 2, 2, 14, 111, 3, 2,
	2, 2, 16, 113, 3, 2, 2, 2, 18, 117, 3, 2, 2, 2, 20, 21, 5, 4, 3, 2, 21,
	3, 3, 2, 2, 2, 22, 24, 5, 6, 4, 2, 23, 22, 3, 2, 2, 2, 23, 24, 3, 2, 2,
	2, 24, 25, 3, 2, 2, 2, 25, 27, 7, 32, 2, 2, 26, 23, 3, 2, 2, 2, 27, 30,
	3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 5, 3, 2, 2, 2,
	30, 28, 3, 2, 2, 2, 31, 32, 7, 16, 2, 2, 32, 70, 7, 21, 2, 2, 33, 34, 7,
	21, 2, 2, 34, 35, 7, 27, 2, 2, 35, 70, 5, 8, 5, 2, 36, 37, 7, 18, 2, 2,
	37, 70, 7, 21, 2, 2, 38, 39, 7, 21, 2, 2, 39, 41, 7, 3, 2, 2, 40, 42, 7,
	25, 2, 2, 41, 40, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43,
	49, 7, 4, 2, 2, 44, 45, 7, 27, 2, 2, 45, 46, 7, 5, 2, 2, 46, 47, 5, 18,
	10, 2, 47, 48, 7, 6, 2, 2, 48, 50, 3, 2, 2, 2, 49, 44, 3, 2, 2, 2, 49,
	50, 3, 2, 2, 2, 50, 70, 3, 2, 2, 2, 51, 52, 7, 19, 2, 2, 52, 70, 7, 21,
	2, 2, 53, 54, 7, 17, 2, 2, 54, 55, 7, 21, 2, 2, 55, 56, 7, 3, 2, 2, 56,
	57, 7, 25, 2, 2, 57, 70, 7, 4, 2, 2, 58, 59, 7, 13, 2, 2, 59, 60, 5, 16,
	9, 2, 60, 61, 7, 14, 2, 2, 61, 62, 5, 4, 3, 2, 62, 63, 7, 15, 2, 2, 63,
	70, 3, 2, 2, 2, 64, 65, 7, 11, 2, 2, 65, 66, 5, 14, 8, 2, 66, 67, 5, 4,
	3, 2, 67, 68, 7, 12, 2, 2, 68, 70, 3, 2, 2, 2, 69, 31, 3, 2, 2, 2, 69,
	33, 3, 2, 2, 2, 69, 36, 3, 2, 2, 2, 69, 38, 3, 2, 2, 2, 69, 51, 3, 2, 2,
	2, 69, 53, 3, 2, 2, 2, 69, 58, 3, 2, 2, 2, 69, 64, 3, 2, 2, 2, 70, 7, 3,
	2, 2, 2, 71, 81, 5, 10, 6, 2, 72, 73, 5, 10, 6, 2, 73, 74, 7, 28, 2, 2,
	74, 75, 5, 10, 6, 2, 75, 81, 3, 2, 2, 2, 76, 77, 5, 10, 6, 2, 77, 78, 7,
	30, 2, 2, 78, 79, 5, 10, 6, 2, 79, 81, 3, 2, 2, 2, 80, 71, 3, 2, 2, 2,
	80, 72, 3, 2, 2, 2, 80, 76, 3, 2, 2, 2, 81, 9, 3, 2, 2, 2, 82, 92, 5, 12,
	7, 2, 83, 84, 5, 12, 7, 2, 84, 85, 7, 29, 2, 2, 85, 86, 5, 12, 7, 2, 86,
	92, 3, 2, 2, 2, 87, 88, 5, 12, 7, 2, 88, 89, 7, 31, 2, 2, 89, 90, 5, 12,
	7, 2, 90, 92, 3, 2, 2, 2, 91, 82, 3, 2, 2, 2, 91, 83, 3, 2, 2, 2, 91, 87,
	3, 2, 2, 2, 92, 11, 3, 2, 2, 2, 93, 110, 7, 25, 2, 2, 94, 110, 7, 21, 2,
	2, 95, 96, 7, 21, 2, 2, 96, 97, 7, 3, 2, 2, 97, 98, 7, 25, 2, 2, 98, 110,
	7, 4, 2, 2, 99, 110, 7, 26, 2, 2, 100, 110, 7, 22, 2, 2, 101, 102, 7, 23,
	2, 2, 102, 110, 5, 12, 7, 2, 103, 104, 7, 24, 2, 2, 104, 110, 5, 12, 7,
	2, 105, 106, 7, 7, 2, 2, 106, 107, 5, 8, 5, 2, 107, 108, 7, 8, 2, 2, 108,
	110, 3, 2, 2, 2, 109, 93, 3, 2, 2, 2, 109, 94, 3, 2, 2, 2, 109, 95, 3,
	2, 2, 2, 109, 99, 3, 2, 2, 2, 109, 100, 3, 2, 2, 2, 109, 101, 3, 2, 2,
	2, 109, 103, 3, 2, 2, 2, 109, 105, 3, 2, 2, 2, 110, 13, 3, 2, 2, 2, 111,
	112, 5, 12, 7, 2, 112, 15, 3, 2, 2, 2, 113, 114, 7, 21, 2, 2, 114, 115,
	7, 9, 2, 2, 115, 116, 7, 25, 2, 2, 116, 17, 3, 2, 2, 2, 117, 122, 7, 25,
	2, 2, 118, 119, 7, 10, 2, 2, 119, 121, 7, 25, 2, 2, 120, 118, 3, 2, 2,
	2, 121, 124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123,
	19, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 11, 23, 28, 41, 49, 69, 80, 91, 109,
	122,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'['", "']'", "'{'", "'}'", "'('", "')'", "'=='", "','", "'repeat'",
	"'endrepeat'", "'if'", "'then'", "'endif'", "'show'", "'show_array_elem'",
	"'read_int'", "'read_double'", "'read_array'", "", "", "'(int)'", "'(real)'",
	"", "", "'='", "'+'", "'*'", "'-'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "REPEAT", "ENDREPEAT", "IF", "THEN",
	"ENDIF", "SHOW", "SHOWARRAYELEM", "READINT", "READDOUBLE", "READARRAY",
	"ID", "STRING", "TOINT", "TOREAL", "INT", "REAL", "ASSIGN", "ADD", "MUL",
	"SUB", "DIV", "NEWLINE", "WS",
}

var ruleNames = []string{
	"prog", "block", "stat", "expr0", "expr1", "expr2", "repetitions", "equal",
	"array_items",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type simplelangParser struct {
	*antlr.BaseParser
}

func NewsimplelangParser(input antlr.TokenStream) *simplelangParser {
	this := new(simplelangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "simplelang.g4"

	return this
}

// simplelangParser tokens.
const (
	simplelangParserEOF           = antlr.TokenEOF
	simplelangParserT__0          = 1
	simplelangParserT__1          = 2
	simplelangParserT__2          = 3
	simplelangParserT__3          = 4
	simplelangParserT__4          = 5
	simplelangParserT__5          = 6
	simplelangParserT__6          = 7
	simplelangParserT__7          = 8
	simplelangParserREPEAT        = 9
	simplelangParserENDREPEAT     = 10
	simplelangParserIF            = 11
	simplelangParserTHEN          = 12
	simplelangParserENDIF         = 13
	simplelangParserSHOW          = 14
	simplelangParserSHOWARRAYELEM = 15
	simplelangParserREADINT       = 16
	simplelangParserREADDOUBLE    = 17
	simplelangParserREADARRAY     = 18
	simplelangParserID            = 19
	simplelangParserSTRING        = 20
	simplelangParserTOINT         = 21
	simplelangParserTOREAL        = 22
	simplelangParserINT           = 23
	simplelangParserREAL          = 24
	simplelangParserASSIGN        = 25
	simplelangParserADD           = 26
	simplelangParserMUL           = 27
	simplelangParserSUB           = 28
	simplelangParserDIV           = 29
	simplelangParserNEWLINE       = 30
	simplelangParserWS            = 31
)

// simplelangParser rules.
const (
	simplelangParserRULE_prog        = 0
	simplelangParserRULE_block       = 1
	simplelangParserRULE_stat        = 2
	simplelangParserRULE_expr0       = 3
	simplelangParserRULE_expr1       = 4
	simplelangParserRULE_expr2       = 5
	simplelangParserRULE_repetitions = 6
	simplelangParserRULE_equal       = 7
	simplelangParserRULE_array_items = 8
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
	p.RuleIndex = simplelangParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *simplelangParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, simplelangParserRULE_prog)

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
		p.SetState(18)
		p.Block()
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(simplelangParserNEWLINE)
}

func (s *BlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(simplelangParserNEWLINE, i)
}

func (s *BlockContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *simplelangParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, simplelangParserRULE_block)
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
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<simplelangParserREPEAT)|(1<<simplelangParserIF)|(1<<simplelangParserSHOW)|(1<<simplelangParserSHOWARRAYELEM)|(1<<simplelangParserREADINT)|(1<<simplelangParserREADDOUBLE)|(1<<simplelangParserID)|(1<<simplelangParserNEWLINE))) != 0 {
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<simplelangParserREPEAT)|(1<<simplelangParserIF)|(1<<simplelangParserSHOW)|(1<<simplelangParserSHOWARRAYELEM)|(1<<simplelangParserREADINT)|(1<<simplelangParserREADDOUBLE)|(1<<simplelangParserID))) != 0 {
			{
				p.SetState(20)
				p.Stat()
			}

		}
		{
			p.SetState(23)
			p.Match(simplelangParserNEWLINE)
		}

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyFrom(ctx *StatContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RepeatContext struct {
	*StatContext
}

func NewRepeatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RepeatContext {
	var p = new(RepeatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *RepeatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatContext) REPEAT() antlr.TerminalNode {
	return s.GetToken(simplelangParserREPEAT, 0)
}

func (s *RepeatContext) Repetitions() IRepetitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepetitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRepetitionsContext)
}

func (s *RepeatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RepeatContext) ENDREPEAT() antlr.TerminalNode {
	return s.GetToken(simplelangParserENDREPEAT, 0)
}

func (s *RepeatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterRepeat(s)
	}
}

func (s *RepeatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitRepeat(s)
	}
}

type ShowContext struct {
	*StatContext
}

func NewShowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShowContext {
	var p = new(ShowContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *ShowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowContext) SHOW() antlr.TerminalNode {
	return s.GetToken(simplelangParserSHOW, 0)
}

func (s *ShowContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *ShowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterShow(s)
	}
}

func (s *ShowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitShow(s)
	}
}

type LetContext struct {
	*StatContext
}

func NewLetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetContext {
	var p = new(LetContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *LetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *LetContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(simplelangParserASSIGN, 0)
}

func (s *LetContext) Expr0() IExpr0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr0Context)
}

func (s *LetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterLet(s)
	}
}

func (s *LetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitLet(s)
	}
}

type ReaddoubleContext struct {
	*StatContext
}

func NewReaddoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReaddoubleContext {
	var p = new(ReaddoubleContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *ReaddoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReaddoubleContext) READDOUBLE() antlr.TerminalNode {
	return s.GetToken(simplelangParserREADDOUBLE, 0)
}

func (s *ReaddoubleContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *ReaddoubleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterReaddouble(s)
	}
}

func (s *ReaddoubleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitReaddouble(s)
	}
}

type ShowArrayElemContext struct {
	*StatContext
}

func NewShowArrayElemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShowArrayElemContext {
	var p = new(ShowArrayElemContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *ShowArrayElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowArrayElemContext) SHOWARRAYELEM() antlr.TerminalNode {
	return s.GetToken(simplelangParserSHOWARRAYELEM, 0)
}

func (s *ShowArrayElemContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *ShowArrayElemContext) INT() antlr.TerminalNode {
	return s.GetToken(simplelangParserINT, 0)
}

func (s *ShowArrayElemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterShowArrayElem(s)
	}
}

func (s *ShowArrayElemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitShowArrayElem(s)
	}
}

type ReadintContext struct {
	*StatContext
}

func NewReadintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReadintContext {
	var p = new(ReadintContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *ReadintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadintContext) READINT() antlr.TerminalNode {
	return s.GetToken(simplelangParserREADINT, 0)
}

func (s *ReadintContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *ReadintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterReadint(s)
	}
}

func (s *ReadintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitReadint(s)
	}
}

type IntarrayContext struct {
	*StatContext
}

func NewIntarrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntarrayContext {
	var p = new(IntarrayContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *IntarrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntarrayContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *IntarrayContext) INT() antlr.TerminalNode {
	return s.GetToken(simplelangParserINT, 0)
}

func (s *IntarrayContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(simplelangParserASSIGN, 0)
}

func (s *IntarrayContext) Array_items() IArray_itemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_itemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_itemsContext)
}

func (s *IntarrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterIntarray(s)
	}
}

func (s *IntarrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitIntarray(s)
	}
}

type IfContext struct {
	*StatContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(simplelangParserIF, 0)
}

func (s *IfContext) Equal() IEqualContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualContext)
}

func (s *IfContext) THEN() antlr.TerminalNode {
	return s.GetToken(simplelangParserTHEN, 0)
}

func (s *IfContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfContext) ENDIF() antlr.TerminalNode {
	return s.GetToken(simplelangParserENDIF, 0)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitIf(s)
	}
}

func (p *simplelangParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, simplelangParserRULE_stat)
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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewShowContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(simplelangParserSHOW)
		}
		{
			p.SetState(30)
			p.Match(simplelangParserID)
		}

	case 2:
		localctx = NewLetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Match(simplelangParserID)
		}
		{
			p.SetState(32)
			p.Match(simplelangParserASSIGN)
		}
		{
			p.SetState(33)
			p.Expr0()
		}

	case 3:
		localctx = NewReadintContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.Match(simplelangParserREADINT)
		}
		{
			p.SetState(35)
			p.Match(simplelangParserID)
		}

	case 4:
		localctx = NewIntarrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(36)
			p.Match(simplelangParserID)
		}
		{
			p.SetState(37)
			p.Match(simplelangParserT__0)
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == simplelangParserINT {
			{
				p.SetState(38)
				p.Match(simplelangParserINT)
			}

		}
		{
			p.SetState(41)
			p.Match(simplelangParserT__1)
		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == simplelangParserASSIGN {
			{
				p.SetState(42)
				p.Match(simplelangParserASSIGN)
			}
			{
				p.SetState(43)
				p.Match(simplelangParserT__2)
			}
			{
				p.SetState(44)
				p.Array_items()
			}
			{
				p.SetState(45)
				p.Match(simplelangParserT__3)
			}

		}

	case 5:
		localctx = NewReaddoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(49)
			p.Match(simplelangParserREADDOUBLE)
		}
		{
			p.SetState(50)
			p.Match(simplelangParserID)
		}

	case 6:
		localctx = NewShowArrayElemContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(51)
			p.Match(simplelangParserSHOWARRAYELEM)
		}
		{
			p.SetState(52)
			p.Match(simplelangParserID)
		}
		{
			p.SetState(53)
			p.Match(simplelangParserT__0)
		}

		{
			p.SetState(54)
			p.Match(simplelangParserINT)
		}

		{
			p.SetState(55)
			p.Match(simplelangParserT__1)
		}

	case 7:
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(56)
			p.Match(simplelangParserIF)
		}
		{
			p.SetState(57)
			p.Equal()
		}
		{
			p.SetState(58)
			p.Match(simplelangParserTHEN)
		}
		{
			p.SetState(59)
			p.Block()
		}
		{
			p.SetState(60)
			p.Match(simplelangParserENDIF)
		}

	case 8:
		localctx = NewRepeatContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(62)
			p.Match(simplelangParserREPEAT)
		}
		{
			p.SetState(63)
			p.Repetitions()
		}
		{
			p.SetState(64)
			p.Block()
		}
		{
			p.SetState(65)
			p.Match(simplelangParserENDREPEAT)
		}

	}

	return localctx
}

// IExpr0Context is an interface to support dynamic dispatch.
type IExpr0Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr0Context differentiates from other interfaces.
	IsExpr0Context()
}

type Expr0Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr0Context() *Expr0Context {
	var p = new(Expr0Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_expr0
	return p
}

func (*Expr0Context) IsExpr0Context() {}

func NewExpr0Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr0Context {
	var p = new(Expr0Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_expr0

	return p
}

func (s *Expr0Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr0Context) CopyFrom(ctx *Expr0Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Expr0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr0Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Single0Context struct {
	*Expr0Context
}

func NewSingle0Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Single0Context {
	var p = new(Single0Context)

	p.Expr0Context = NewEmptyExpr0Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr0Context))

	return p
}

func (s *Single0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single0Context) Expr1() IExpr1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr1Context)
}

func (s *Single0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterSingle0(s)
	}
}

func (s *Single0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitSingle0(s)
	}
}

type AddContext struct {
	*Expr0Context
}

func NewAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddContext {
	var p = new(AddContext)

	p.Expr0Context = NewEmptyExpr0Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr0Context))

	return p
}

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) AllExpr1() []IExpr1Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr1Context)(nil)).Elem())
	var tst = make([]IExpr1Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr1Context)
		}
	}

	return tst
}

func (s *AddContext) Expr1(i int) IExpr1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr1Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr1Context)
}

func (s *AddContext) ADD() antlr.TerminalNode {
	return s.GetToken(simplelangParserADD, 0)
}

func (s *AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterAdd(s)
	}
}

func (s *AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitAdd(s)
	}
}

type SubContext struct {
	*Expr0Context
}

func NewSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubContext {
	var p = new(SubContext)

	p.Expr0Context = NewEmptyExpr0Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr0Context))

	return p
}

func (s *SubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubContext) AllExpr1() []IExpr1Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr1Context)(nil)).Elem())
	var tst = make([]IExpr1Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr1Context)
		}
	}

	return tst
}

func (s *SubContext) Expr1(i int) IExpr1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr1Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr1Context)
}

func (s *SubContext) SUB() antlr.TerminalNode {
	return s.GetToken(simplelangParserSUB, 0)
}

func (s *SubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterSub(s)
	}
}

func (s *SubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitSub(s)
	}
}

func (p *simplelangParser) Expr0() (localctx IExpr0Context) {
	localctx = NewExpr0Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, simplelangParserRULE_expr0)

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

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingle0Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Expr1()
		}

	case 2:
		localctx = NewAddContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Expr1()
		}
		{
			p.SetState(71)
			p.Match(simplelangParserADD)
		}
		{
			p.SetState(72)
			p.Expr1()
		}

	case 3:
		localctx = NewSubContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)
			p.Expr1()
		}
		{
			p.SetState(75)
			p.Match(simplelangParserSUB)
		}
		{
			p.SetState(76)
			p.Expr1()
		}

	}

	return localctx
}

// IExpr1Context is an interface to support dynamic dispatch.
type IExpr1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr1Context differentiates from other interfaces.
	IsExpr1Context()
}

type Expr1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr1Context() *Expr1Context {
	var p = new(Expr1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_expr1
	return p
}

func (*Expr1Context) IsExpr1Context() {}

func NewExpr1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr1Context {
	var p = new(Expr1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_expr1

	return p
}

func (s *Expr1Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr1Context) CopyFrom(ctx *Expr1Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Expr1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DivContext struct {
	*Expr1Context
}

func NewDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivContext {
	var p = new(DivContext)

	p.Expr1Context = NewEmptyExpr1Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr1Context))

	return p
}

func (s *DivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivContext) AllExpr2() []IExpr2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr2Context)(nil)).Elem())
	var tst = make([]IExpr2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr2Context)
		}
	}

	return tst
}

func (s *DivContext) Expr2(i int) IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
}

func (s *DivContext) DIV() antlr.TerminalNode {
	return s.GetToken(simplelangParserDIV, 0)
}

func (s *DivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterDiv(s)
	}
}

func (s *DivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitDiv(s)
	}
}

type Single1Context struct {
	*Expr1Context
}

func NewSingle1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Single1Context {
	var p = new(Single1Context)

	p.Expr1Context = NewEmptyExpr1Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr1Context))

	return p
}

func (s *Single1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single1Context) Expr2() IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
}

func (s *Single1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterSingle1(s)
	}
}

func (s *Single1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitSingle1(s)
	}
}

type MulContext struct {
	*Expr1Context
}

func NewMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulContext {
	var p = new(MulContext)

	p.Expr1Context = NewEmptyExpr1Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr1Context))

	return p
}

func (s *MulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulContext) AllExpr2() []IExpr2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr2Context)(nil)).Elem())
	var tst = make([]IExpr2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr2Context)
		}
	}

	return tst
}

func (s *MulContext) Expr2(i int) IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
}

func (s *MulContext) MUL() antlr.TerminalNode {
	return s.GetToken(simplelangParserMUL, 0)
}

func (s *MulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterMul(s)
	}
}

func (s *MulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitMul(s)
	}
}

func (p *simplelangParser) Expr1() (localctx IExpr1Context) {
	localctx = NewExpr1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, simplelangParserRULE_expr1)

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

	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingle1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Expr2()
		}

	case 2:
		localctx = NewMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Expr2()
		}
		{
			p.SetState(82)
			p.Match(simplelangParserMUL)
		}
		{
			p.SetState(83)
			p.Expr2()
		}

	case 3:
		localctx = NewDivContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.Expr2()
		}
		{
			p.SetState(86)
			p.Match(simplelangParserDIV)
		}
		{
			p.SetState(87)
			p.Expr2()
		}

	}

	return localctx
}

// IExpr2Context is an interface to support dynamic dispatch.
type IExpr2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr2Context differentiates from other interfaces.
	IsExpr2Context()
}

type Expr2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr2Context() *Expr2Context {
	var p = new(Expr2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_expr2
	return p
}

func (*Expr2Context) IsExpr2Context() {}

func NewExpr2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr2Context {
	var p = new(Expr2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_expr2

	return p
}

func (s *Expr2Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr2Context) CopyFrom(ctx *Expr2Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Expr2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParContext struct {
	*Expr2Context
}

func NewParContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParContext {
	var p = new(ParContext)

	p.Expr2Context = NewEmptyExpr2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr2Context))

	return p
}

func (s *ParContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParContext) Expr0() IExpr0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr0Context)
}

func (s *ParContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterPar(s)
	}
}

func (s *ParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitPar(s)
	}
}

type TointContext struct {
	*Expr2Context
}

func NewTointContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TointContext {
	var p = new(TointContext)

	p.Expr2Context = NewEmptyExpr2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr2Context))

	return p
}

func (s *TointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TointContext) TOINT() antlr.TerminalNode {
	return s.GetToken(simplelangParserTOINT, 0)
}

func (s *TointContext) Expr2() IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
}

func (s *TointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterToint(s)
	}
}

func (s *TointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitToint(s)
	}
}

type StringContext struct {
	*Expr2Context
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.Expr2Context = NewEmptyExpr2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr2Context))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(simplelangParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitString(s)
	}
}

type AssignArrayElemContext struct {
	*Expr2Context
}

func NewAssignArrayElemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignArrayElemContext {
	var p = new(AssignArrayElemContext)

	p.Expr2Context = NewEmptyExpr2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr2Context))

	return p
}

func (s *AssignArrayElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignArrayElemContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *AssignArrayElemContext) INT() antlr.TerminalNode {
	return s.GetToken(simplelangParserINT, 0)
}

func (s *AssignArrayElemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterAssignArrayElem(s)
	}
}

func (s *AssignArrayElemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitAssignArrayElem(s)
	}
}

type TorealContext struct {
	*Expr2Context
}

func NewTorealContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TorealContext {
	var p = new(TorealContext)

	p.Expr2Context = NewEmptyExpr2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr2Context))

	return p
}

func (s *TorealContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TorealContext) TOREAL() antlr.TerminalNode {
	return s.GetToken(simplelangParserTOREAL, 0)
}

func (s *TorealContext) Expr2() IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
}

func (s *TorealContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterToreal(s)
	}
}

func (s *TorealContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitToreal(s)
	}
}

type IdContext struct {
	*Expr2Context
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	p.Expr2Context = NewEmptyExpr2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr2Context))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitId(s)
	}
}

type RealContext struct {
	*Expr2Context
}

func NewRealContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RealContext {
	var p = new(RealContext)

	p.Expr2Context = NewEmptyExpr2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr2Context))

	return p
}

func (s *RealContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealContext) REAL() antlr.TerminalNode {
	return s.GetToken(simplelangParserREAL, 0)
}

func (s *RealContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterReal(s)
	}
}

func (s *RealContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitReal(s)
	}
}

type IntContext struct {
	*Expr2Context
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.Expr2Context = NewEmptyExpr2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr2Context))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(simplelangParserINT, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitInt(s)
	}
}

func (p *simplelangParser) Expr2() (localctx IExpr2Context) {
	localctx = NewExpr2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, simplelangParserRULE_expr2)

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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.Match(simplelangParserINT)
		}

	case 2:
		localctx = NewIdContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Match(simplelangParserID)
		}

	case 3:
		localctx = NewAssignArrayElemContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Match(simplelangParserID)
		}
		{
			p.SetState(94)
			p.Match(simplelangParserT__0)
		}

		{
			p.SetState(95)
			p.Match(simplelangParserINT)
		}

		{
			p.SetState(96)
			p.Match(simplelangParserT__1)
		}

	case 4:
		localctx = NewRealContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(97)
			p.Match(simplelangParserREAL)
		}

	case 5:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(98)
			p.Match(simplelangParserSTRING)
		}

	case 6:
		localctx = NewTointContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(99)
			p.Match(simplelangParserTOINT)
		}
		{
			p.SetState(100)
			p.Expr2()
		}

	case 7:
		localctx = NewTorealContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(101)
			p.Match(simplelangParserTOREAL)
		}
		{
			p.SetState(102)
			p.Expr2()
		}

	case 8:
		localctx = NewParContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(103)
			p.Match(simplelangParserT__4)
		}
		{
			p.SetState(104)
			p.Expr0()
		}
		{
			p.SetState(105)
			p.Match(simplelangParserT__5)
		}

	}

	return localctx
}

// IRepetitionsContext is an interface to support dynamic dispatch.
type IRepetitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRepetitionsContext differentiates from other interfaces.
	IsRepetitionsContext()
}

type RepetitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepetitionsContext() *RepetitionsContext {
	var p = new(RepetitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_repetitions
	return p
}

func (*RepetitionsContext) IsRepetitionsContext() {}

func NewRepetitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepetitionsContext {
	var p = new(RepetitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_repetitions

	return p
}

func (s *RepetitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *RepetitionsContext) Expr2() IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
}

func (s *RepetitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepetitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepetitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterRepetitions(s)
	}
}

func (s *RepetitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitRepetitions(s)
	}
}

func (p *simplelangParser) Repetitions() (localctx IRepetitionsContext) {
	localctx = NewRepetitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, simplelangParserRULE_repetitions)

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
		p.SetState(109)
		p.Expr2()
	}

	return localctx
}

// IEqualContext is an interface to support dynamic dispatch.
type IEqualContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualContext differentiates from other interfaces.
	IsEqualContext()
}

type EqualContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualContext() *EqualContext {
	var p = new(EqualContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_equal
	return p
}

func (*EqualContext) IsEqualContext() {}

func NewEqualContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualContext {
	var p = new(EqualContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_equal

	return p
}

func (s *EqualContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *EqualContext) INT() antlr.TerminalNode {
	return s.GetToken(simplelangParserINT, 0)
}

func (s *EqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterEqual(s)
	}
}

func (s *EqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitEqual(s)
	}
}

func (p *simplelangParser) Equal() (localctx IEqualContext) {
	localctx = NewEqualContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, simplelangParserRULE_equal)

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
		p.SetState(111)
		p.Match(simplelangParserID)
	}
	{
		p.SetState(112)
		p.Match(simplelangParserT__6)
	}
	{
		p.SetState(113)
		p.Match(simplelangParserINT)
	}

	return localctx
}

// IArray_itemsContext is an interface to support dynamic dispatch.
type IArray_itemsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_itemsContext differentiates from other interfaces.
	IsArray_itemsContext()
}

type Array_itemsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_itemsContext() *Array_itemsContext {
	var p = new(Array_itemsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_array_items
	return p
}

func (*Array_itemsContext) IsArray_itemsContext() {}

func NewArray_itemsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_itemsContext {
	var p = new(Array_itemsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_array_items

	return p
}

func (s *Array_itemsContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_itemsContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(simplelangParserINT)
}

func (s *Array_itemsContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(simplelangParserINT, i)
}

func (s *Array_itemsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_itemsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_itemsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterArray_items(s)
	}
}

func (s *Array_itemsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitArray_items(s)
	}
}

func (p *simplelangParser) Array_items() (localctx IArray_itemsContext) {
	localctx = NewArray_itemsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, simplelangParserRULE_array_items)
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
		p.SetState(115)
		p.Match(simplelangParserINT)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == simplelangParserT__7 {
		{
			p.SetState(116)
			p.Match(simplelangParserT__7)
		}
		{
			p.SetState(117)
			p.Match(simplelangParserINT)
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
