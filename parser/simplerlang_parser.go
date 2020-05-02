// Code generated from simplerlang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // simplerlang

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 67, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 5, 2, 14,
	10, 2, 3, 2, 7, 2, 17, 10, 2, 12, 2, 14, 2, 20, 11, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 31, 10, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 42, 10, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 53, 10, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 65, 10, 6, 3, 6, 2, 2,
	7, 2, 4, 6, 8, 10, 2, 2, 2, 74, 2, 18, 3, 2, 2, 2, 4, 30, 3, 2, 2, 2, 6,
	41, 3, 2, 2, 2, 8, 52, 3, 2, 2, 2, 10, 64, 3, 2, 2, 2, 12, 14, 5, 4, 3,
	2, 13, 12, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2, 14, 15, 3, 2, 2, 2, 15, 17,
	7, 18, 2, 2, 16, 13, 3, 2, 2, 2, 17, 20, 3, 2, 2, 2, 18, 16, 3, 2, 2, 2,
	18, 19, 3, 2, 2, 2, 19, 3, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 21, 22, 7, 6,
	2, 2, 22, 31, 7, 9, 2, 2, 23, 24, 7, 9, 2, 2, 24, 25, 7, 3, 2, 2, 25, 31,
	5, 6, 4, 2, 26, 27, 7, 7, 2, 2, 27, 31, 7, 9, 2, 2, 28, 29, 7, 8, 2, 2,
	29, 31, 7, 9, 2, 2, 30, 21, 3, 2, 2, 2, 30, 23, 3, 2, 2, 2, 30, 26, 3,
	2, 2, 2, 30, 28, 3, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32, 42, 5, 8, 5, 2, 33,
	34, 5, 8, 5, 2, 34, 35, 7, 14, 2, 2, 35, 36, 5, 8, 5, 2, 36, 42, 3, 2,
	2, 2, 37, 38, 5, 8, 5, 2, 38, 39, 7, 16, 2, 2, 39, 40, 5, 8, 5, 2, 40,
	42, 3, 2, 2, 2, 41, 32, 3, 2, 2, 2, 41, 33, 3, 2, 2, 2, 41, 37, 3, 2, 2,
	2, 42, 7, 3, 2, 2, 2, 43, 53, 5, 10, 6, 2, 44, 45, 5, 10, 6, 2, 45, 46,
	7, 15, 2, 2, 46, 47, 5, 10, 6, 2, 47, 53, 3, 2, 2, 2, 48, 49, 5, 10, 6,
	2, 49, 50, 7, 17, 2, 2, 50, 51, 5, 10, 6, 2, 51, 53, 3, 2, 2, 2, 52, 43,
	3, 2, 2, 2, 52, 44, 3, 2, 2, 2, 52, 48, 3, 2, 2, 2, 53, 9, 3, 2, 2, 2,
	54, 65, 7, 12, 2, 2, 55, 65, 7, 13, 2, 2, 56, 57, 7, 10, 2, 2, 57, 65,
	5, 10, 6, 2, 58, 59, 7, 11, 2, 2, 59, 65, 5, 10, 6, 2, 60, 61, 7, 4, 2,
	2, 61, 62, 5, 6, 4, 2, 62, 63, 7, 5, 2, 2, 63, 65, 3, 2, 2, 2, 64, 54,
	3, 2, 2, 2, 64, 55, 3, 2, 2, 2, 64, 56, 3, 2, 2, 2, 64, 58, 3, 2, 2, 2,
	64, 60, 3, 2, 2, 2, 65, 11, 3, 2, 2, 2, 8, 13, 18, 30, 41, 52, 64,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'('", "')'", "'show'", "'read_int'", "'read_double'", "", "'(int)'",
	"'(real)'", "", "", "'+'", "'*'", "'-'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "SHOW", "READINT", "READDOUBLE", "ID", "TOINT", "TOREAL",
	"INT", "REAL", "ADD", "MUL", "SUB", "DIV", "NEWLINE", "WS",
}

var ruleNames = []string{
	"program", "stat", "expr0", "expr1", "expr2",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type simplerlangParser struct {
	*antlr.BaseParser
}

func NewsimplerlangParser(input antlr.TokenStream) *simplerlangParser {
	this := new(simplerlangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "simplerlang.g4"

	return this
}

// simplerlangParser tokens.
const (
	simplerlangParserEOF        = antlr.TokenEOF
	simplerlangParserT__0       = 1
	simplerlangParserT__1       = 2
	simplerlangParserT__2       = 3
	simplerlangParserSHOW       = 4
	simplerlangParserREADINT    = 5
	simplerlangParserREADDOUBLE = 6
	simplerlangParserID         = 7
	simplerlangParserTOINT      = 8
	simplerlangParserTOREAL     = 9
	simplerlangParserINT        = 10
	simplerlangParserREAL       = 11
	simplerlangParserADD        = 12
	simplerlangParserMUL        = 13
	simplerlangParserSUB        = 14
	simplerlangParserDIV        = 15
	simplerlangParserNEWLINE    = 16
	simplerlangParserWS         = 17
)

// simplerlangParser rules.
const (
	simplerlangParserRULE_program = 0
	simplerlangParserRULE_stat    = 1
	simplerlangParserRULE_expr0   = 2
	simplerlangParserRULE_expr1   = 3
	simplerlangParserRULE_expr2   = 4
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplerlangParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplerlangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(simplerlangParserNEWLINE)
}

func (s *ProgramContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(simplerlangParserNEWLINE, i)
}

func (s *ProgramContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *ProgramContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *simplerlangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, simplerlangParserRULE_program)
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
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<simplerlangParserSHOW)|(1<<simplerlangParserREADINT)|(1<<simplerlangParserREADDOUBLE)|(1<<simplerlangParserID)|(1<<simplerlangParserNEWLINE))) != 0 {
		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<simplerlangParserSHOW)|(1<<simplerlangParserREADINT)|(1<<simplerlangParserREADDOUBLE)|(1<<simplerlangParserID))) != 0 {
			{
				p.SetState(10)
				p.Stat()
			}

		}
		{
			p.SetState(13)
			p.Match(simplerlangParserNEWLINE)
		}

		p.SetState(18)
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
	p.RuleIndex = simplerlangParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplerlangParserRULE_stat

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
	return s.GetToken(simplerlangParserSHOW, 0)
}

func (s *ShowContext) ID() antlr.TerminalNode {
	return s.GetToken(simplerlangParserID, 0)
}

func (s *ShowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterShow(s)
	}
}

func (s *ShowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
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
	return s.GetToken(simplerlangParserID, 0)
}

func (s *LetContext) Expr0() IExpr0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr0Context)
}

func (s *LetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterLet(s)
	}
}

func (s *LetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
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
	return s.GetToken(simplerlangParserREADDOUBLE, 0)
}

func (s *ReaddoubleContext) ID() antlr.TerminalNode {
	return s.GetToken(simplerlangParserID, 0)
}

func (s *ReaddoubleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterReaddouble(s)
	}
}

func (s *ReaddoubleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.ExitReaddouble(s)
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
	return s.GetToken(simplerlangParserREADINT, 0)
}

func (s *ReadintContext) ID() antlr.TerminalNode {
	return s.GetToken(simplerlangParserID, 0)
}

func (s *ReadintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterReadint(s)
	}
}

func (s *ReadintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.ExitReadint(s)
	}
}

func (p *simplerlangParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, simplerlangParserRULE_stat)

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

	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case simplerlangParserSHOW:
		localctx = NewShowContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
			p.Match(simplerlangParserSHOW)
		}
		{
			p.SetState(20)
			p.Match(simplerlangParserID)
		}

	case simplerlangParserID:
		localctx = NewLetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.Match(simplerlangParserID)
		}
		{
			p.SetState(22)
			p.Match(simplerlangParserT__0)
		}
		{
			p.SetState(23)
			p.Expr0()
		}

	case simplerlangParserREADINT:
		localctx = NewReadintContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(24)
			p.Match(simplerlangParserREADINT)
		}
		{
			p.SetState(25)
			p.Match(simplerlangParserID)
		}

	case simplerlangParserREADDOUBLE:
		localctx = NewReaddoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(26)
			p.Match(simplerlangParserREADDOUBLE)
		}
		{
			p.SetState(27)
			p.Match(simplerlangParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = simplerlangParserRULE_expr0
	return p
}

func (*Expr0Context) IsExpr0Context() {}

func NewExpr0Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr0Context {
	var p = new(Expr0Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplerlangParserRULE_expr0

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
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterSingle0(s)
	}
}

func (s *Single0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
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
	return s.GetToken(simplerlangParserADD, 0)
}

func (s *AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterAdd(s)
	}
}

func (s *AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
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
	return s.GetToken(simplerlangParserSUB, 0)
}

func (s *SubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterSub(s)
	}
}

func (s *SubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.ExitSub(s)
	}
}

func (p *simplerlangParser) Expr0() (localctx IExpr0Context) {
	localctx = NewExpr0Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, simplerlangParserRULE_expr0)

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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingle0Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Expr1()
		}

	case 2:
		localctx = NewAddContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Expr1()
		}
		{
			p.SetState(32)
			p.Match(simplerlangParserADD)
		}
		{
			p.SetState(33)
			p.Expr1()
		}

	case 3:
		localctx = NewSubContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.Expr1()
		}
		{
			p.SetState(36)
			p.Match(simplerlangParserSUB)
		}
		{
			p.SetState(37)
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
	p.RuleIndex = simplerlangParserRULE_expr1
	return p
}

func (*Expr1Context) IsExpr1Context() {}

func NewExpr1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr1Context {
	var p = new(Expr1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplerlangParserRULE_expr1

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
	return s.GetToken(simplerlangParserDIV, 0)
}

func (s *DivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterDiv(s)
	}
}

func (s *DivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
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
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterSingle1(s)
	}
}

func (s *Single1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
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
	return s.GetToken(simplerlangParserMUL, 0)
}

func (s *MulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterMul(s)
	}
}

func (s *MulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.ExitMul(s)
	}
}

func (p *simplerlangParser) Expr1() (localctx IExpr1Context) {
	localctx = NewExpr1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, simplerlangParserRULE_expr1)

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

	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingle1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Expr2()
		}

	case 2:
		localctx = NewMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Expr2()
		}
		{
			p.SetState(43)
			p.Match(simplerlangParserMUL)
		}
		{
			p.SetState(44)
			p.Expr2()
		}

	case 3:
		localctx = NewDivContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(46)
			p.Expr2()
		}
		{
			p.SetState(47)
			p.Match(simplerlangParserDIV)
		}
		{
			p.SetState(48)
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
	p.RuleIndex = simplerlangParserRULE_expr2
	return p
}

func (*Expr2Context) IsExpr2Context() {}

func NewExpr2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr2Context {
	var p = new(Expr2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplerlangParserRULE_expr2

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
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterPar(s)
	}
}

func (s *ParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
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
	return s.GetToken(simplerlangParserTOINT, 0)
}

func (s *TointContext) Expr2() IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
}

func (s *TointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterToint(s)
	}
}

func (s *TointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.ExitToint(s)
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
	return s.GetToken(simplerlangParserTOREAL, 0)
}

func (s *TorealContext) Expr2() IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
}

func (s *TorealContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterToreal(s)
	}
}

func (s *TorealContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
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
	return s.GetToken(simplerlangParserREAL, 0)
}

func (s *RealContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterReal(s)
	}
}

func (s *RealContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
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
	return s.GetToken(simplerlangParserINT, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplerlangListener); ok {
		listenerT.ExitInt(s)
	}
}

func (p *simplerlangParser) Expr2() (localctx IExpr2Context) {
	localctx = NewExpr2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, simplerlangParserRULE_expr2)

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

	switch p.GetTokenStream().LA(1) {
	case simplerlangParserINT:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(simplerlangParserINT)
		}

	case simplerlangParserREAL:
		localctx = NewRealContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Match(simplerlangParserREAL)
		}

	case simplerlangParserTOINT:
		localctx = NewTointContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Match(simplerlangParserTOINT)
		}
		{
			p.SetState(55)
			p.Expr2()
		}

	case simplerlangParserTOREAL:
		localctx = NewTorealContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Match(simplerlangParserTOREAL)
		}
		{
			p.SetState(57)
			p.Expr2()
		}

	case simplerlangParserT__1:
		localctx = NewParContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(58)
			p.Match(simplerlangParserT__1)
		}
		{
			p.SetState(59)
			p.Expr0()
		}
		{
			p.SetState(60)
			p.Match(simplerlangParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
