// Code generated from simplerlang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 123,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 6, 8,
	71, 10, 8, 13, 8, 14, 8, 72, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 6, 11, 89, 10, 11, 13,
	11, 14, 11, 90, 3, 12, 6, 12, 94, 10, 12, 13, 12, 14, 12, 95, 3, 12, 3,
	12, 6, 12, 100, 10, 12, 13, 12, 14, 12, 101, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 5, 17, 113, 10, 17, 3, 17, 3, 17, 3,
	18, 6, 18, 118, 10, 18, 13, 18, 14, 18, 119, 3, 18, 3, 18, 2, 2, 19, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 3, 2, 4, 4, 2, 67, 92,
	99, 124, 4, 2, 11, 12, 34, 34, 2, 128, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 3,
	37, 3, 2, 2, 2, 5, 39, 3, 2, 2, 2, 7, 41, 3, 2, 2, 2, 9, 43, 3, 2, 2, 2,
	11, 48, 3, 2, 2, 2, 13, 57, 3, 2, 2, 2, 15, 70, 3, 2, 2, 2, 17, 74, 3,
	2, 2, 2, 19, 80, 3, 2, 2, 2, 21, 88, 3, 2, 2, 2, 23, 93, 3, 2, 2, 2, 25,
	103, 3, 2, 2, 2, 27, 105, 3, 2, 2, 2, 29, 107, 3, 2, 2, 2, 31, 109, 3,
	2, 2, 2, 33, 112, 3, 2, 2, 2, 35, 117, 3, 2, 2, 2, 37, 38, 7, 63, 2, 2,
	38, 4, 3, 2, 2, 2, 39, 40, 7, 42, 2, 2, 40, 6, 3, 2, 2, 2, 41, 42, 7, 43,
	2, 2, 42, 8, 3, 2, 2, 2, 43, 44, 7, 117, 2, 2, 44, 45, 7, 106, 2, 2, 45,
	46, 7, 113, 2, 2, 46, 47, 7, 121, 2, 2, 47, 10, 3, 2, 2, 2, 48, 49, 7,
	116, 2, 2, 49, 50, 7, 103, 2, 2, 50, 51, 7, 99, 2, 2, 51, 52, 7, 102, 2,
	2, 52, 53, 7, 97, 2, 2, 53, 54, 7, 107, 2, 2, 54, 55, 7, 112, 2, 2, 55,
	56, 7, 118, 2, 2, 56, 12, 3, 2, 2, 2, 57, 58, 7, 116, 2, 2, 58, 59, 7,
	103, 2, 2, 59, 60, 7, 99, 2, 2, 60, 61, 7, 102, 2, 2, 61, 62, 7, 97, 2,
	2, 62, 63, 7, 102, 2, 2, 63, 64, 7, 113, 2, 2, 64, 65, 7, 119, 2, 2, 65,
	66, 7, 100, 2, 2, 66, 67, 7, 110, 2, 2, 67, 68, 7, 103, 2, 2, 68, 14, 3,
	2, 2, 2, 69, 71, 9, 2, 2, 2, 70, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72,
	70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 16, 3, 2, 2, 2, 74, 75, 7, 42,
	2, 2, 75, 76, 7, 107, 2, 2, 76, 77, 7, 112, 2, 2, 77, 78, 7, 118, 2, 2,
	78, 79, 7, 43, 2, 2, 79, 18, 3, 2, 2, 2, 80, 81, 7, 42, 2, 2, 81, 82, 7,
	116, 2, 2, 82, 83, 7, 103, 2, 2, 83, 84, 7, 99, 2, 2, 84, 85, 7, 110, 2,
	2, 85, 86, 7, 43, 2, 2, 86, 20, 3, 2, 2, 2, 87, 89, 4, 50, 59, 2, 88, 87,
	3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2,
	91, 22, 3, 2, 2, 2, 92, 94, 4, 50, 59, 2, 93, 92, 3, 2, 2, 2, 94, 95, 3,
	2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97,
	99, 7, 48, 2, 2, 98, 100, 4, 50, 59, 2, 99, 98, 3, 2, 2, 2, 100, 101, 3,
	2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 24, 3, 2, 2, 2,
	103, 104, 7, 45, 2, 2, 104, 26, 3, 2, 2, 2, 105, 106, 7, 44, 2, 2, 106,
	28, 3, 2, 2, 2, 107, 108, 7, 47, 2, 2, 108, 30, 3, 2, 2, 2, 109, 110, 7,
	49, 2, 2, 110, 32, 3, 2, 2, 2, 111, 113, 7, 15, 2, 2, 112, 111, 3, 2, 2,
	2, 112, 113, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115, 7, 12, 2, 2, 115,
	34, 3, 2, 2, 2, 116, 118, 9, 3, 2, 2, 117, 116, 3, 2, 2, 2, 118, 119, 3,
	2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 121, 3, 2, 2,
	2, 121, 122, 8, 18, 2, 2, 122, 36, 3, 2, 2, 2, 9, 2, 72, 90, 95, 101, 112,
	119, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'('", "')'", "'show'", "'read_int'", "'read_double'", "", "'(int)'",
	"'(real)'", "", "", "'+'", "'*'", "'-'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "SHOW", "READINT", "READDOUBLE", "ID", "TOINT", "TOREAL",
	"INT", "REAL", "ADD", "MUL", "SUB", "DIV", "NEWLINE", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "SHOW", "READINT", "READDOUBLE", "ID", "TOINT",
	"TOREAL", "INT", "REAL", "ADD", "MUL", "SUB", "DIV", "NEWLINE", "WS",
}

type simplerlangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewsimplerlangLexer(input antlr.CharStream) *simplerlangLexer {

	l := new(simplerlangLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "simplerlang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// simplerlangLexer tokens.
const (
	simplerlangLexerT__0       = 1
	simplerlangLexerT__1       = 2
	simplerlangLexerT__2       = 3
	simplerlangLexerSHOW       = 4
	simplerlangLexerREADINT    = 5
	simplerlangLexerREADDOUBLE = 6
	simplerlangLexerID         = 7
	simplerlangLexerTOINT      = 8
	simplerlangLexerTOREAL     = 9
	simplerlangLexerINT        = 10
	simplerlangLexerREAL       = 11
	simplerlangLexerADD        = 12
	simplerlangLexerMUL        = 13
	simplerlangLexerSUB        = 14
	simplerlangLexerDIV        = 15
	simplerlangLexerNEWLINE    = 16
	simplerlangLexerWS         = 17
)