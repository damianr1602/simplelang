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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 31, 120,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 5, 3, 24, 10, 3,
	3, 3, 7, 3, 27, 10, 3, 12, 3, 14, 3, 30, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 48, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 56, 10, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 65, 10, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 76, 10, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 87, 10, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 5, 7, 104, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 7, 10, 115, 10, 10, 12, 10, 14, 10, 118, 11, 10, 3, 10, 2, 2, 11,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 2, 2, 131, 2, 20, 3, 2, 2, 2, 4, 28,
	3, 2, 2, 2, 6, 64, 3, 2, 2, 2, 8, 75, 3, 2, 2, 2, 10, 86, 3, 2, 2, 2, 12,
	103, 3, 2, 2, 2, 14, 105, 3, 2, 2, 2, 16, 107, 3, 2, 2, 2, 18, 111, 3,
	2, 2, 2, 20, 21, 5, 4, 3, 2, 21, 3, 3, 2, 2, 2, 22, 24, 5, 6, 4, 2, 23,
	22, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 27, 7, 30,
	2, 2, 26, 23, 3, 2, 2, 2, 27, 30, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29,
	3, 2, 2, 2, 29, 5, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 31, 32, 7, 14, 2, 2,
	32, 65, 7, 19, 2, 2, 33, 34, 7, 19, 2, 2, 34, 35, 7, 25, 2, 2, 35, 65,
	5, 8, 5, 2, 36, 37, 7, 11, 2, 2, 37, 38, 5, 16, 9, 2, 38, 39, 7, 12, 2,
	2, 39, 40, 5, 14, 8, 2, 40, 41, 7, 13, 2, 2, 41, 65, 3, 2, 2, 2, 42, 43,
	7, 16, 2, 2, 43, 65, 7, 19, 2, 2, 44, 45, 7, 19, 2, 2, 45, 47, 7, 3, 2,
	2, 46, 48, 7, 23, 2, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 49,
	3, 2, 2, 2, 49, 55, 7, 4, 2, 2, 50, 51, 7, 25, 2, 2, 51, 52, 7, 5, 2, 2,
	52, 53, 5, 18, 10, 2, 53, 54, 7, 6, 2, 2, 54, 56, 3, 2, 2, 2, 55, 50, 3,
	2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 65, 3, 2, 2, 2, 57, 58, 7, 17, 2, 2, 58,
	65, 7, 19, 2, 2, 59, 60, 7, 15, 2, 2, 60, 61, 7, 19, 2, 2, 61, 62, 7, 3,
	2, 2, 62, 63, 7, 23, 2, 2, 63, 65, 7, 4, 2, 2, 64, 31, 3, 2, 2, 2, 64,
	33, 3, 2, 2, 2, 64, 36, 3, 2, 2, 2, 64, 42, 3, 2, 2, 2, 64, 44, 3, 2, 2,
	2, 64, 57, 3, 2, 2, 2, 64, 59, 3, 2, 2, 2, 65, 7, 3, 2, 2, 2, 66, 76, 5,
	10, 6, 2, 67, 68, 5, 10, 6, 2, 68, 69, 7, 26, 2, 2, 69, 70, 5, 10, 6, 2,
	70, 76, 3, 2, 2, 2, 71, 72, 5, 10, 6, 2, 72, 73, 7, 28, 2, 2, 73, 74, 5,
	10, 6, 2, 74, 76, 3, 2, 2, 2, 75, 66, 3, 2, 2, 2, 75, 67, 3, 2, 2, 2, 75,
	71, 3, 2, 2, 2, 76, 9, 3, 2, 2, 2, 77, 87, 5, 12, 7, 2, 78, 79, 5, 12,
	7, 2, 79, 80, 7, 27, 2, 2, 80, 81, 5, 12, 7, 2, 81, 87, 3, 2, 2, 2, 82,
	83, 5, 12, 7, 2, 83, 84, 7, 29, 2, 2, 84, 85, 5, 12, 7, 2, 85, 87, 3, 2,
	2, 2, 86, 77, 3, 2, 2, 2, 86, 78, 3, 2, 2, 2, 86, 82, 3, 2, 2, 2, 87, 11,
	3, 2, 2, 2, 88, 104, 7, 23, 2, 2, 89, 90, 7, 19, 2, 2, 90, 91, 7, 3, 2,
	2, 91, 92, 7, 23, 2, 2, 92, 104, 7, 4, 2, 2, 93, 104, 7, 24, 2, 2, 94,
	104, 7, 20, 2, 2, 95, 96, 7, 21, 2, 2, 96, 104, 5, 12, 7, 2, 97, 98, 7,
	22, 2, 2, 98, 104, 5, 12, 7, 2, 99, 100, 7, 7, 2, 2, 100, 101, 5, 8, 5,
	2, 101, 102, 7, 8, 2, 2, 102, 104, 3, 2, 2, 2, 103, 88, 3, 2, 2, 2, 103,
	89, 3, 2, 2, 2, 103, 93, 3, 2, 2, 2, 103, 94, 3, 2, 2, 2, 103, 95, 3, 2,
	2, 2, 103, 97, 3, 2, 2, 2, 103, 99, 3, 2, 2, 2, 104, 13, 3, 2, 2, 2, 105,
	106, 5, 4, 3, 2, 106, 15, 3, 2, 2, 2, 107, 108, 7, 19, 2, 2, 108, 109,
	7, 9, 2, 2, 109, 110, 7, 23, 2, 2, 110, 17, 3, 2, 2, 2, 111, 116, 7, 23,
	2, 2, 112, 113, 7, 10, 2, 2, 113, 115, 7, 23, 2, 2, 114, 112, 3, 2, 2,
	2, 115, 118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117,
	19, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 11, 23, 28, 47, 55, 64, 75, 86, 103,
	116,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'['", "']'", "'{'", "'}'", "'('", "')'", "'=='", "','", "'if'", "'then'",
	"'endif'", "'show'", "'show_array_elem'", "'read_int'", "'read_double'",
	"'read_array'", "", "", "'(int)'", "'(real)'", "", "", "'='", "'+'", "'*'",
	"'-'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "IF", "THEN", "ENDIF", "SHOW", "SHOWARRAYELEM",
	"READINT", "READDOUBLE", "READARRAY", "ID", "STRING", "TOINT", "TOREAL",
	"INT", "REAL", "ASSIGN", "ADD", "MUL", "SUB", "DIV", "NEWLINE", "WS",
}

var ruleNames = []string{
	"prog", "block", "stat", "expr0", "expr1", "expr2", "blockif", "equal",
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
	simplelangParserIF            = 9
	simplelangParserTHEN          = 10
	simplelangParserENDIF         = 11
	simplelangParserSHOW          = 12
	simplelangParserSHOWARRAYELEM = 13
	simplelangParserREADINT       = 14
	simplelangParserREADDOUBLE    = 15
	simplelangParserREADARRAY     = 16
	simplelangParserID            = 17
	simplelangParserSTRING        = 18
	simplelangParserTOINT         = 19
	simplelangParserTOREAL        = 20
	simplelangParserINT           = 21
	simplelangParserREAL          = 22
	simplelangParserASSIGN        = 23
	simplelangParserADD           = 24
	simplelangParserMUL           = 25
	simplelangParserSUB           = 26
	simplelangParserDIV           = 27
	simplelangParserNEWLINE       = 28
	simplelangParserWS            = 29
)

// simplelangParser rules.
const (
	simplelangParserRULE_prog        = 0
	simplelangParserRULE_block       = 1
	simplelangParserRULE_stat        = 2
	simplelangParserRULE_expr0       = 3
	simplelangParserRULE_expr1       = 4
	simplelangParserRULE_expr2       = 5
	simplelangParserRULE_blockif     = 6
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

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<simplelangParserIF)|(1<<simplelangParserSHOW)|(1<<simplelangParserSHOWARRAYELEM)|(1<<simplelangParserREADINT)|(1<<simplelangParserREADDOUBLE)|(1<<simplelangParserID)|(1<<simplelangParserNEWLINE))) != 0 {
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<simplelangParserIF)|(1<<simplelangParserSHOW)|(1<<simplelangParserSHOWARRAYELEM)|(1<<simplelangParserREADINT)|(1<<simplelangParserREADDOUBLE)|(1<<simplelangParserID))) != 0 {
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

func (s *IfContext) Blockif() IBlockifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockifContext)
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

	p.SetState(62)
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
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.Match(simplelangParserIF)
		}
		{
			p.SetState(35)
			p.Equal()
		}
		{
			p.SetState(36)
			p.Match(simplelangParserTHEN)
		}
		{
			p.SetState(37)
			p.Blockif()
		}
		{
			p.SetState(38)
			p.Match(simplelangParserENDIF)
		}

	case 4:
		localctx = NewReadintContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.Match(simplelangParserREADINT)
		}
		{
			p.SetState(41)
			p.Match(simplelangParserID)
		}

	case 5:
		localctx = NewIntarrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(42)
			p.Match(simplelangParserID)
		}
		{
			p.SetState(43)
			p.Match(simplelangParserT__0)
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == simplelangParserINT {
			{
				p.SetState(44)
				p.Match(simplelangParserINT)
			}

		}
		{
			p.SetState(47)
			p.Match(simplelangParserT__1)
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == simplelangParserASSIGN {
			{
				p.SetState(48)
				p.Match(simplelangParserASSIGN)
			}
			{
				p.SetState(49)
				p.Match(simplelangParserT__2)
			}
			{
				p.SetState(50)
				p.Array_items()
			}
			{
				p.SetState(51)
				p.Match(simplelangParserT__3)
			}

		}

	case 6:
		localctx = NewReaddoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(55)
			p.Match(simplelangParserREADDOUBLE)
		}
		{
			p.SetState(56)
			p.Match(simplelangParserID)
		}

	case 7:
		localctx = NewShowArrayElemContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(57)
			p.Match(simplelangParserSHOWARRAYELEM)
		}
		{
			p.SetState(58)
			p.Match(simplelangParserID)
		}
		{
			p.SetState(59)
			p.Match(simplelangParserT__0)
		}

		{
			p.SetState(60)
			p.Match(simplelangParserINT)
		}

		{
			p.SetState(61)
			p.Match(simplelangParserT__1)
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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingle0Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Expr1()
		}

	case 2:
		localctx = NewAddContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Expr1()
		}
		{
			p.SetState(66)
			p.Match(simplelangParserADD)
		}
		{
			p.SetState(67)
			p.Expr1()
		}

	case 3:
		localctx = NewSubContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.Expr1()
		}
		{
			p.SetState(70)
			p.Match(simplelangParserSUB)
		}
		{
			p.SetState(71)
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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingle1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Expr2()
		}

	case 2:
		localctx = NewMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Expr2()
		}
		{
			p.SetState(77)
			p.Match(simplelangParserMUL)
		}
		{
			p.SetState(78)
			p.Expr2()
		}

	case 3:
		localctx = NewDivContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.Expr2()
		}
		{
			p.SetState(81)
			p.Match(simplelangParserDIV)
		}
		{
			p.SetState(82)
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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case simplelangParserINT:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(simplelangParserINT)
		}

	case simplelangParserID:
		localctx = NewAssignArrayElemContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Match(simplelangParserID)
		}
		{
			p.SetState(88)
			p.Match(simplelangParserT__0)
		}

		{
			p.SetState(89)
			p.Match(simplelangParserINT)
		}

		{
			p.SetState(90)
			p.Match(simplelangParserT__1)
		}

	case simplelangParserREAL:
		localctx = NewRealContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.Match(simplelangParserREAL)
		}

	case simplelangParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(92)
			p.Match(simplelangParserSTRING)
		}

	case simplelangParserTOINT:
		localctx = NewTointContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(93)
			p.Match(simplelangParserTOINT)
		}
		{
			p.SetState(94)
			p.Expr2()
		}

	case simplelangParserTOREAL:
		localctx = NewTorealContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(95)
			p.Match(simplelangParserTOREAL)
		}
		{
			p.SetState(96)
			p.Expr2()
		}

	case simplelangParserT__4:
		localctx = NewParContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(97)
			p.Match(simplelangParserT__4)
		}
		{
			p.SetState(98)
			p.Expr0()
		}
		{
			p.SetState(99)
			p.Match(simplelangParserT__5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlockifContext is an interface to support dynamic dispatch.
type IBlockifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockifContext differentiates from other interfaces.
	IsBlockifContext()
}

type BlockifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockifContext() *BlockifContext {
	var p = new(BlockifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_blockif
	return p
}

func (*BlockifContext) IsBlockifContext() {}

func NewBlockifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockifContext {
	var p = new(BlockifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_blockif

	return p
}

func (s *BlockifContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockifContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterBlockif(s)
	}
}

func (s *BlockifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitBlockif(s)
	}
}

func (p *simplelangParser) Blockif() (localctx IBlockifContext) {
	localctx = NewBlockifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, simplelangParserRULE_blockif)

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
		p.SetState(103)
		p.Block()
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
		p.SetState(105)
		p.Match(simplelangParserID)
	}
	{
		p.SetState(106)
		p.Match(simplelangParserT__6)
	}
	{
		p.SetState(107)
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
		p.SetState(109)
		p.Match(simplelangParserINT)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == simplelangParserT__7 {
		{
			p.SetState(110)
			p.Match(simplelangParserT__7)
		}
		{
			p.SetState(111)
			p.Match(simplelangParserINT)
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
