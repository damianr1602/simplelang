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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 180,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 5, 3, 40, 10, 3, 3, 3, 7, 3, 43, 10, 3, 12, 3, 14,
	3, 46, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 5, 4, 58, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 66, 10, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 87, 10, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 7, 6, 97, 10, 6, 12, 6, 14, 6, 100,
	11, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 108, 10, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 10, 5, 10, 115, 10, 10, 3, 10, 7, 10, 118, 10, 10, 12, 10,
	14, 10, 121, 11, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 135, 10, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 146, 10, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 164, 10, 14, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 7, 17, 175, 10, 17, 12, 17, 14,
	17, 178, 11, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 2, 2, 2, 193, 2, 34, 3, 2, 2, 2, 4, 44, 3, 2, 2, 2,
	6, 86, 3, 2, 2, 2, 8, 88, 3, 2, 2, 2, 10, 93, 3, 2, 2, 2, 12, 101, 3, 2,
	2, 2, 14, 103, 3, 2, 2, 2, 16, 111, 3, 2, 2, 2, 18, 119, 3, 2, 2, 2, 20,
	122, 3, 2, 2, 2, 22, 134, 3, 2, 2, 2, 24, 145, 3, 2, 2, 2, 26, 163, 3,
	2, 2, 2, 28, 165, 3, 2, 2, 2, 30, 167, 3, 2, 2, 2, 32, 171, 3, 2, 2, 2,
	34, 35, 5, 4, 3, 2, 35, 3, 3, 2, 2, 2, 36, 40, 5, 6, 4, 2, 37, 40, 5, 14,
	8, 2, 38, 40, 5, 8, 5, 2, 39, 36, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 38,
	3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 43, 7, 38, 2, 2,
	42, 39, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3,
	2, 2, 2, 45, 5, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 48, 7, 22, 2, 2, 48,
	87, 7, 27, 2, 2, 49, 50, 7, 27, 2, 2, 50, 51, 7, 33, 2, 2, 51, 87, 5, 22,
	12, 2, 52, 53, 7, 24, 2, 2, 53, 87, 7, 27, 2, 2, 54, 55, 7, 27, 2, 2, 55,
	57, 7, 3, 2, 2, 56, 58, 7, 31, 2, 2, 57, 56, 3, 2, 2, 2, 57, 58, 3, 2,
	2, 2, 58, 59, 3, 2, 2, 2, 59, 65, 7, 4, 2, 2, 60, 61, 7, 33, 2, 2, 61,
	62, 7, 5, 2, 2, 62, 63, 5, 32, 17, 2, 63, 64, 7, 6, 2, 2, 64, 66, 3, 2,
	2, 2, 65, 60, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 87, 3, 2, 2, 2, 67, 68,
	7, 25, 2, 2, 68, 87, 7, 27, 2, 2, 69, 70, 7, 23, 2, 2, 70, 71, 7, 27, 2,
	2, 71, 72, 7, 3, 2, 2, 72, 73, 7, 31, 2, 2, 73, 87, 7, 4, 2, 2, 74, 75,
	7, 19, 2, 2, 75, 76, 5, 30, 16, 2, 76, 77, 7, 20, 2, 2, 77, 78, 5, 4, 3,
	2, 78, 79, 7, 21, 2, 2, 79, 87, 3, 2, 2, 2, 80, 81, 7, 17, 2, 2, 81, 82,
	5, 28, 15, 2, 82, 83, 5, 4, 3, 2, 83, 84, 7, 18, 2, 2, 84, 87, 3, 2, 2,
	2, 85, 87, 7, 27, 2, 2, 86, 47, 3, 2, 2, 2, 86, 49, 3, 2, 2, 2, 86, 52,
	3, 2, 2, 2, 86, 54, 3, 2, 2, 2, 86, 67, 3, 2, 2, 2, 86, 69, 3, 2, 2, 2,
	86, 74, 3, 2, 2, 2, 86, 80, 3, 2, 2, 2, 86, 85, 3, 2, 2, 2, 87, 7, 3, 2,
	2, 2, 88, 89, 7, 13, 2, 2, 89, 90, 5, 12, 7, 2, 90, 91, 5, 10, 6, 2, 91,
	92, 7, 14, 2, 2, 92, 9, 3, 2, 2, 2, 93, 98, 7, 27, 2, 2, 94, 95, 7, 7,
	2, 2, 95, 97, 7, 27, 2, 2, 96, 94, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2, 98,
	96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 11, 3, 2, 2, 2, 100, 98, 3, 2,
	2, 2, 101, 102, 7, 27, 2, 2, 102, 13, 3, 2, 2, 2, 103, 104, 7, 15, 2, 2,
	104, 105, 5, 16, 9, 2, 105, 107, 5, 18, 10, 2, 106, 108, 5, 20, 11, 2,
	107, 106, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109,
	110, 7, 16, 2, 2, 110, 15, 3, 2, 2, 2, 111, 112, 7, 27, 2, 2, 112, 17,
	3, 2, 2, 2, 113, 115, 5, 6, 4, 2, 114, 113, 3, 2, 2, 2, 114, 115, 3, 2,
	2, 2, 115, 116, 3, 2, 2, 2, 116, 118, 7, 38, 2, 2, 117, 114, 3, 2, 2, 2,
	118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120,
	19, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 123, 7, 8, 2, 2, 123, 124, 7,
	27, 2, 2, 124, 21, 3, 2, 2, 2, 125, 135, 5, 24, 13, 2, 126, 127, 5, 24,
	13, 2, 127, 128, 7, 34, 2, 2, 128, 129, 5, 24, 13, 2, 129, 135, 3, 2, 2,
	2, 130, 131, 5, 24, 13, 2, 131, 132, 7, 36, 2, 2, 132, 133, 5, 24, 13,
	2, 133, 135, 3, 2, 2, 2, 134, 125, 3, 2, 2, 2, 134, 126, 3, 2, 2, 2, 134,
	130, 3, 2, 2, 2, 135, 23, 3, 2, 2, 2, 136, 146, 5, 26, 14, 2, 137, 138,
	5, 26, 14, 2, 138, 139, 7, 35, 2, 2, 139, 140, 5, 26, 14, 2, 140, 146,
	3, 2, 2, 2, 141, 142, 5, 26, 14, 2, 142, 143, 7, 37, 2, 2, 143, 144, 5,
	26, 14, 2, 144, 146, 3, 2, 2, 2, 145, 136, 3, 2, 2, 2, 145, 137, 3, 2,
	2, 2, 145, 141, 3, 2, 2, 2, 146, 25, 3, 2, 2, 2, 147, 164, 7, 31, 2, 2,
	148, 164, 7, 27, 2, 2, 149, 150, 7, 27, 2, 2, 150, 151, 7, 3, 2, 2, 151,
	152, 7, 31, 2, 2, 152, 164, 7, 4, 2, 2, 153, 164, 7, 32, 2, 2, 154, 164,
	7, 28, 2, 2, 155, 156, 7, 29, 2, 2, 156, 164, 5, 26, 14, 2, 157, 158, 7,
	30, 2, 2, 158, 164, 5, 26, 14, 2, 159, 160, 7, 9, 2, 2, 160, 161, 5, 22,
	12, 2, 161, 162, 7, 10, 2, 2, 162, 164, 3, 2, 2, 2, 163, 147, 3, 2, 2,
	2, 163, 148, 3, 2, 2, 2, 163, 149, 3, 2, 2, 2, 163, 153, 3, 2, 2, 2, 163,
	154, 3, 2, 2, 2, 163, 155, 3, 2, 2, 2, 163, 157, 3, 2, 2, 2, 163, 159,
	3, 2, 2, 2, 164, 27, 3, 2, 2, 2, 165, 166, 5, 26, 14, 2, 166, 29, 3, 2,
	2, 2, 167, 168, 7, 27, 2, 2, 168, 169, 7, 11, 2, 2, 169, 170, 7, 31, 2,
	2, 170, 31, 3, 2, 2, 2, 171, 176, 7, 31, 2, 2, 172, 173, 7, 12, 2, 2, 173,
	175, 7, 31, 2, 2, 174, 172, 3, 2, 2, 2, 175, 178, 3, 2, 2, 2, 176, 174,
	3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 33, 3, 2, 2, 2, 178, 176, 3, 2,
	2, 2, 15, 39, 44, 57, 65, 86, 98, 107, 114, 119, 134, 145, 163, 176,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'['", "']'", "'{'", "'}'", "';'", "'return'", "'('", "')'", "'=='",
	"','", "'struct'", "'endstruct'", "'func'", "'endfunc'", "'repeat'", "'endrepeat'",
	"'if'", "'then'", "'endif'", "'show'", "'show_array_elem'", "'read_int'",
	"'read_double'", "'read_array'", "", "", "'(int)'", "'(real)'", "", "",
	"'='", "'+'", "'*'", "'-'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "STRUCT", "ENDSTRUCT", "FUNC",
	"ENDFUNC", "REPEAT", "ENDREPEAT", "IF", "THEN", "ENDIF", "SHOW", "SHOWARRAYELEM",
	"READINT", "READDOUBLE", "READARRAY", "ID", "STRING", "TOINT", "TOREAL",
	"INT", "REAL", "ASSIGN", "ADD", "MUL", "SUB", "DIV", "NEWLINE", "WS",
}

var ruleNames = []string{
	"prog", "block", "stat", "structure", "sBlock", "structName", "function",
	"funcName", "fBlock", "result", "expr0", "expr1", "expr2", "repetitions",
	"equal", "array_items",
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
	simplelangParserT__8          = 9
	simplelangParserT__9          = 10
	simplelangParserSTRUCT        = 11
	simplelangParserENDSTRUCT     = 12
	simplelangParserFUNC          = 13
	simplelangParserENDFUNC       = 14
	simplelangParserREPEAT        = 15
	simplelangParserENDREPEAT     = 16
	simplelangParserIF            = 17
	simplelangParserTHEN          = 18
	simplelangParserENDIF         = 19
	simplelangParserSHOW          = 20
	simplelangParserSHOWARRAYELEM = 21
	simplelangParserREADINT       = 22
	simplelangParserREADDOUBLE    = 23
	simplelangParserREADARRAY     = 24
	simplelangParserID            = 25
	simplelangParserSTRING        = 26
	simplelangParserTOINT         = 27
	simplelangParserTOREAL        = 28
	simplelangParserINT           = 29
	simplelangParserREAL          = 30
	simplelangParserASSIGN        = 31
	simplelangParserADD           = 32
	simplelangParserMUL           = 33
	simplelangParserSUB           = 34
	simplelangParserDIV           = 35
	simplelangParserNEWLINE       = 36
	simplelangParserWS            = 37
)

// simplelangParser rules.
const (
	simplelangParserRULE_prog        = 0
	simplelangParserRULE_block       = 1
	simplelangParserRULE_stat        = 2
	simplelangParserRULE_structure   = 3
	simplelangParserRULE_sBlock      = 4
	simplelangParserRULE_structName  = 5
	simplelangParserRULE_function    = 6
	simplelangParserRULE_funcName    = 7
	simplelangParserRULE_fBlock      = 8
	simplelangParserRULE_result      = 9
	simplelangParserRULE_expr0       = 10
	simplelangParserRULE_expr1       = 11
	simplelangParserRULE_expr2       = 12
	simplelangParserRULE_repetitions = 13
	simplelangParserRULE_equal       = 14
	simplelangParserRULE_array_items = 15
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
		p.SetState(32)
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

func (s *BlockContext) AllFunction() []IFunctionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionContext)(nil)).Elem())
	var tst = make([]IFunctionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionContext)
		}
	}

	return tst
}

func (s *BlockContext) Function(i int) IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *BlockContext) AllStructure() []IStructureContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructureContext)(nil)).Elem())
	var tst = make([]IStructureContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructureContext)
		}
	}

	return tst
}

func (s *BlockContext) Structure(i int) IStructureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructureContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructureContext)
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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(simplelangParserSTRUCT-11))|(1<<(simplelangParserFUNC-11))|(1<<(simplelangParserREPEAT-11))|(1<<(simplelangParserIF-11))|(1<<(simplelangParserSHOW-11))|(1<<(simplelangParserSHOWARRAYELEM-11))|(1<<(simplelangParserREADINT-11))|(1<<(simplelangParserREADDOUBLE-11))|(1<<(simplelangParserID-11))|(1<<(simplelangParserNEWLINE-11)))) != 0 {
		p.SetState(37)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case simplelangParserREPEAT, simplelangParserIF, simplelangParserSHOW, simplelangParserSHOWARRAYELEM, simplelangParserREADINT, simplelangParserREADDOUBLE, simplelangParserID:
			{
				p.SetState(34)
				p.Stat()
			}

		case simplelangParserFUNC:
			{
				p.SetState(35)
				p.Function()
			}

		case simplelangParserSTRUCT:
			{
				p.SetState(36)
				p.Structure()
			}

		case simplelangParserNEWLINE:

		default:
		}
		{
			p.SetState(39)
			p.Match(simplelangParserNEWLINE)
		}

		p.SetState(44)
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

type CallContext struct {
	*StatContext
}

func NewCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallContext {
	var p = new(CallContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitCall(s)
	}
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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewShowContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(simplelangParserSHOW)
		}
		{
			p.SetState(46)
			p.Match(simplelangParserID)
		}

	case 2:
		localctx = NewLetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Match(simplelangParserID)
		}
		{
			p.SetState(48)
			p.Match(simplelangParserASSIGN)
		}
		{
			p.SetState(49)
			p.Expr0()
		}

	case 3:
		localctx = NewReadintContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Match(simplelangParserREADINT)
		}
		{
			p.SetState(51)
			p.Match(simplelangParserID)
		}

	case 4:
		localctx = NewIntarrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(52)
			p.Match(simplelangParserID)
		}
		{
			p.SetState(53)
			p.Match(simplelangParserT__0)
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == simplelangParserINT {
			{
				p.SetState(54)
				p.Match(simplelangParserINT)
			}

		}
		{
			p.SetState(57)
			p.Match(simplelangParserT__1)
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == simplelangParserASSIGN {
			{
				p.SetState(58)
				p.Match(simplelangParserASSIGN)
			}
			{
				p.SetState(59)
				p.Match(simplelangParserT__2)
			}
			{
				p.SetState(60)
				p.Array_items()
			}
			{
				p.SetState(61)
				p.Match(simplelangParserT__3)
			}

		}

	case 5:
		localctx = NewReaddoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(65)
			p.Match(simplelangParserREADDOUBLE)
		}
		{
			p.SetState(66)
			p.Match(simplelangParserID)
		}

	case 6:
		localctx = NewShowArrayElemContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(67)
			p.Match(simplelangParserSHOWARRAYELEM)
		}
		{
			p.SetState(68)
			p.Match(simplelangParserID)
		}
		{
			p.SetState(69)
			p.Match(simplelangParserT__0)
		}

		{
			p.SetState(70)
			p.Match(simplelangParserINT)
		}

		{
			p.SetState(71)
			p.Match(simplelangParserT__1)
		}

	case 7:
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(72)
			p.Match(simplelangParserIF)
		}
		{
			p.SetState(73)
			p.Equal()
		}
		{
			p.SetState(74)
			p.Match(simplelangParserTHEN)
		}
		{
			p.SetState(75)
			p.Block()
		}
		{
			p.SetState(76)
			p.Match(simplelangParserENDIF)
		}

	case 8:
		localctx = NewRepeatContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(78)
			p.Match(simplelangParserREPEAT)
		}
		{
			p.SetState(79)
			p.Repetitions()
		}
		{
			p.SetState(80)
			p.Block()
		}
		{
			p.SetState(81)
			p.Match(simplelangParserENDREPEAT)
		}

	case 9:
		localctx = NewCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(83)
			p.Match(simplelangParserID)
		}

	}

	return localctx
}

// IStructureContext is an interface to support dynamic dispatch.
type IStructureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructureContext differentiates from other interfaces.
	IsStructureContext()
}

type StructureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructureContext() *StructureContext {
	var p = new(StructureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_structure
	return p
}

func (*StructureContext) IsStructureContext() {}

func NewStructureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructureContext {
	var p = new(StructureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_structure

	return p
}

func (s *StructureContext) GetParser() antlr.Parser { return s.parser }

func (s *StructureContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(simplelangParserSTRUCT, 0)
}

func (s *StructureContext) StructName() IStructNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructNameContext)
}

func (s *StructureContext) SBlock() ISBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISBlockContext)
}

func (s *StructureContext) ENDSTRUCT() antlr.TerminalNode {
	return s.GetToken(simplelangParserENDSTRUCT, 0)
}

func (s *StructureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterStructure(s)
	}
}

func (s *StructureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitStructure(s)
	}
}

func (p *simplelangParser) Structure() (localctx IStructureContext) {
	localctx = NewStructureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, simplelangParserRULE_structure)

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
		p.SetState(86)
		p.Match(simplelangParserSTRUCT)
	}
	{
		p.SetState(87)
		p.StructName()
	}
	{
		p.SetState(88)
		p.SBlock()
	}
	{
		p.SetState(89)
		p.Match(simplelangParserENDSTRUCT)
	}

	return localctx
}

// ISBlockContext is an interface to support dynamic dispatch.
type ISBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSBlockContext differentiates from other interfaces.
	IsSBlockContext()
}

type SBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySBlockContext() *SBlockContext {
	var p = new(SBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_sBlock
	return p
}

func (*SBlockContext) IsSBlockContext() {}

func NewSBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SBlockContext {
	var p = new(SBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_sBlock

	return p
}

func (s *SBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SBlockContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(simplelangParserID)
}

func (s *SBlockContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(simplelangParserID, i)
}

func (s *SBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterSBlock(s)
	}
}

func (s *SBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitSBlock(s)
	}
}

func (p *simplelangParser) SBlock() (localctx ISBlockContext) {
	localctx = NewSBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, simplelangParserRULE_sBlock)
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
		p.SetState(91)
		p.Match(simplelangParserID)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == simplelangParserT__4 {
		{
			p.SetState(92)
			p.Match(simplelangParserT__4)
		}
		{
			p.SetState(93)
			p.Match(simplelangParserID)
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStructNameContext is an interface to support dynamic dispatch.
type IStructNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructNameContext differentiates from other interfaces.
	IsStructNameContext()
}

type StructNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructNameContext() *StructNameContext {
	var p = new(StructNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_structName
	return p
}

func (*StructNameContext) IsStructNameContext() {}

func NewStructNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructNameContext {
	var p = new(StructNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_structName

	return p
}

func (s *StructNameContext) GetParser() antlr.Parser { return s.parser }

func (s *StructNameContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *StructNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterStructName(s)
	}
}

func (s *StructNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitStructName(s)
	}
}

func (p *simplelangParser) StructName() (localctx IStructNameContext) {
	localctx = NewStructNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, simplelangParserRULE_structName)

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
		p.SetState(99)
		p.Match(simplelangParserID)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(simplelangParserFUNC, 0)
}

func (s *FunctionContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *FunctionContext) FBlock() IFBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFBlockContext)
}

func (s *FunctionContext) ENDFUNC() antlr.TerminalNode {
	return s.GetToken(simplelangParserENDFUNC, 0)
}

func (s *FunctionContext) Result() IResultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResultContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *simplelangParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, simplelangParserRULE_function)
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
		p.SetState(101)
		p.Match(simplelangParserFUNC)
	}
	{
		p.SetState(102)
		p.FuncName()
	}
	{
		p.SetState(103)
		p.FBlock()
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == simplelangParserT__5 {
		{
			p.SetState(104)
			p.Result()
		}

	}
	{
		p.SetState(107)
		p.Match(simplelangParserENDFUNC)
	}

	return localctx
}

// IFuncNameContext is an interface to support dynamic dispatch.
type IFuncNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncNameContext differentiates from other interfaces.
	IsFuncNameContext()
}

type FuncNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncNameContext() *FuncNameContext {
	var p = new(FuncNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_funcName
	return p
}

func (*FuncNameContext) IsFuncNameContext() {}

func NewFuncNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncNameContext {
	var p = new(FuncNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_funcName

	return p
}

func (s *FuncNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncNameContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *FuncNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterFuncName(s)
	}
}

func (s *FuncNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitFuncName(s)
	}
}

func (p *simplelangParser) FuncName() (localctx IFuncNameContext) {
	localctx = NewFuncNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, simplelangParserRULE_funcName)

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
		p.Match(simplelangParserID)
	}

	return localctx
}

// IFBlockContext is an interface to support dynamic dispatch.
type IFBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFBlockContext differentiates from other interfaces.
	IsFBlockContext()
}

type FBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFBlockContext() *FBlockContext {
	var p = new(FBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_fBlock
	return p
}

func (*FBlockContext) IsFBlockContext() {}

func NewFBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FBlockContext {
	var p = new(FBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_fBlock

	return p
}

func (s *FBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *FBlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(simplelangParserNEWLINE)
}

func (s *FBlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(simplelangParserNEWLINE, i)
}

func (s *FBlockContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *FBlockContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *FBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterFBlock(s)
	}
}

func (s *FBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitFBlock(s)
	}
}

func (p *simplelangParser) FBlock() (localctx IFBlockContext) {
	localctx = NewFBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, simplelangParserRULE_fBlock)
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
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-15)&-(0x1f+1)) == 0 && ((1<<uint((_la-15)))&((1<<(simplelangParserREPEAT-15))|(1<<(simplelangParserIF-15))|(1<<(simplelangParserSHOW-15))|(1<<(simplelangParserSHOWARRAYELEM-15))|(1<<(simplelangParserREADINT-15))|(1<<(simplelangParserREADDOUBLE-15))|(1<<(simplelangParserID-15))|(1<<(simplelangParserNEWLINE-15)))) != 0 {
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<simplelangParserREPEAT)|(1<<simplelangParserIF)|(1<<simplelangParserSHOW)|(1<<simplelangParserSHOWARRAYELEM)|(1<<simplelangParserREADINT)|(1<<simplelangParserREADDOUBLE)|(1<<simplelangParserID))) != 0 {
			{
				p.SetState(111)
				p.Stat()
			}

		}
		{
			p.SetState(114)
			p.Match(simplelangParserNEWLINE)
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IResultContext is an interface to support dynamic dispatch.
type IResultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResultContext differentiates from other interfaces.
	IsResultContext()
}

type ResultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultContext() *ResultContext {
	var p = new(ResultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simplelangParserRULE_result
	return p
}

func (*ResultContext) IsResultContext() {}

func NewResultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultContext {
	var p = new(ResultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simplelangParserRULE_result

	return p
}

func (s *ResultContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultContext) ID() antlr.TerminalNode {
	return s.GetToken(simplelangParserID, 0)
}

func (s *ResultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.EnterResult(s)
	}
}

func (s *ResultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simplelangListener); ok {
		listenerT.ExitResult(s)
	}
}

func (p *simplelangParser) Result() (localctx IResultContext) {
	localctx = NewResultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, simplelangParserRULE_result)

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
		p.SetState(120)
		p.Match(simplelangParserT__5)
	}
	{
		p.SetState(121)
		p.Match(simplelangParserID)
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
	p.EnterRule(localctx, 20, simplelangParserRULE_expr0)

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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingle0Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Expr1()
		}

	case 2:
		localctx = NewAddContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Expr1()
		}
		{
			p.SetState(125)
			p.Match(simplelangParserADD)
		}
		{
			p.SetState(126)
			p.Expr1()
		}

	case 3:
		localctx = NewSubContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.Expr1()
		}
		{
			p.SetState(129)
			p.Match(simplelangParserSUB)
		}
		{
			p.SetState(130)
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
	p.EnterRule(localctx, 22, simplelangParserRULE_expr1)

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

	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingle1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Expr2()
		}

	case 2:
		localctx = NewMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Expr2()
		}
		{
			p.SetState(136)
			p.Match(simplelangParserMUL)
		}
		{
			p.SetState(137)
			p.Expr2()
		}

	case 3:
		localctx = NewDivContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(139)
			p.Expr2()
		}
		{
			p.SetState(140)
			p.Match(simplelangParserDIV)
		}
		{
			p.SetState(141)
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
	p.EnterRule(localctx, 24, simplelangParserRULE_expr2)

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

	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Match(simplelangParserINT)
		}

	case 2:
		localctx = NewIdContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Match(simplelangParserID)
		}

	case 3:
		localctx = NewAssignArrayElemContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(147)
			p.Match(simplelangParserID)
		}
		{
			p.SetState(148)
			p.Match(simplelangParserT__0)
		}

		{
			p.SetState(149)
			p.Match(simplelangParserINT)
		}

		{
			p.SetState(150)
			p.Match(simplelangParserT__1)
		}

	case 4:
		localctx = NewRealContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(151)
			p.Match(simplelangParserREAL)
		}

	case 5:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(152)
			p.Match(simplelangParserSTRING)
		}

	case 6:
		localctx = NewTointContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(153)
			p.Match(simplelangParserTOINT)
		}
		{
			p.SetState(154)
			p.Expr2()
		}

	case 7:
		localctx = NewTorealContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(155)
			p.Match(simplelangParserTOREAL)
		}
		{
			p.SetState(156)
			p.Expr2()
		}

	case 8:
		localctx = NewParContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(157)
			p.Match(simplelangParserT__6)
		}
		{
			p.SetState(158)
			p.Expr0()
		}
		{
			p.SetState(159)
			p.Match(simplelangParserT__7)
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
	p.EnterRule(localctx, 26, simplelangParserRULE_repetitions)

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
		p.SetState(163)
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
	p.EnterRule(localctx, 28, simplelangParserRULE_equal)

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
		p.SetState(165)
		p.Match(simplelangParserID)
	}
	{
		p.SetState(166)
		p.Match(simplelangParserT__8)
	}
	{
		p.SetState(167)
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
	p.EnterRule(localctx, 30, simplelangParserRULE_array_items)
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
		p.SetState(169)
		p.Match(simplelangParserINT)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == simplelangParserT__9 {
		{
			p.SetState(170)
			p.Match(simplelangParserT__9)
		}
		{
			p.SetState(171)
			p.Match(simplelangParserINT)
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
