// Code generated from simplelang.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 33, 240,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 6, 20, 170, 10, 20, 13, 20, 14,
	20, 171, 3, 21, 3, 21, 3, 21, 7, 21, 177, 10, 21, 12, 21, 14, 21, 180,
	11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 188, 10, 22, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 25, 6, 25, 204, 10, 25, 13, 25, 14, 25, 205, 3, 26, 6,
	26, 209, 10, 26, 13, 26, 14, 26, 210, 3, 26, 3, 26, 6, 26, 215, 10, 26,
	13, 26, 14, 26, 216, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 32, 5, 32, 230, 10, 32, 3, 32, 3, 32, 3, 33, 6, 33,
	235, 10, 33, 13, 33, 14, 33, 236, 3, 33, 3, 33, 2, 2, 34, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 2, 45,
	23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63,
	32, 65, 33, 3, 2, 5, 4, 2, 67, 92, 99, 124, 6, 2, 12, 12, 15, 15, 36, 36,
	94, 94, 4, 2, 11, 11, 34, 34, 2, 247, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2,
	2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2,
	2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2,
	2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 3, 67, 3, 2, 2, 2, 5, 69, 3,
	2, 2, 2, 7, 71, 3, 2, 2, 2, 9, 73, 3, 2, 2, 2, 11, 75, 3, 2, 2, 2, 13,
	77, 3, 2, 2, 2, 15, 79, 3, 2, 2, 2, 17, 82, 3, 2, 2, 2, 19, 84, 3, 2, 2,
	2, 21, 91, 3, 2, 2, 2, 23, 101, 3, 2, 2, 2, 25, 104, 3, 2, 2, 2, 27, 109,
	3, 2, 2, 2, 29, 115, 3, 2, 2, 2, 31, 120, 3, 2, 2, 2, 33, 136, 3, 2, 2,
	2, 35, 145, 3, 2, 2, 2, 37, 157, 3, 2, 2, 2, 39, 169, 3, 2, 2, 2, 41, 173,
	3, 2, 2, 2, 43, 187, 3, 2, 2, 2, 45, 189, 3, 2, 2, 2, 47, 195, 3, 2, 2,
	2, 49, 203, 3, 2, 2, 2, 51, 208, 3, 2, 2, 2, 53, 218, 3, 2, 2, 2, 55, 220,
	3, 2, 2, 2, 57, 222, 3, 2, 2, 2, 59, 224, 3, 2, 2, 2, 61, 226, 3, 2, 2,
	2, 63, 229, 3, 2, 2, 2, 65, 234, 3, 2, 2, 2, 67, 68, 7, 93, 2, 2, 68, 4,
	3, 2, 2, 2, 69, 70, 7, 95, 2, 2, 70, 6, 3, 2, 2, 2, 71, 72, 7, 125, 2,
	2, 72, 8, 3, 2, 2, 2, 73, 74, 7, 127, 2, 2, 74, 10, 3, 2, 2, 2, 75, 76,
	7, 42, 2, 2, 76, 12, 3, 2, 2, 2, 77, 78, 7, 43, 2, 2, 78, 14, 3, 2, 2,
	2, 79, 80, 7, 63, 2, 2, 80, 81, 7, 63, 2, 2, 81, 16, 3, 2, 2, 2, 82, 83,
	7, 46, 2, 2, 83, 18, 3, 2, 2, 2, 84, 85, 7, 116, 2, 2, 85, 86, 7, 103,
	2, 2, 86, 87, 7, 114, 2, 2, 87, 88, 7, 103, 2, 2, 88, 89, 7, 99, 2, 2,
	89, 90, 7, 118, 2, 2, 90, 20, 3, 2, 2, 2, 91, 92, 7, 103, 2, 2, 92, 93,
	7, 112, 2, 2, 93, 94, 7, 102, 2, 2, 94, 95, 7, 116, 2, 2, 95, 96, 7, 103,
	2, 2, 96, 97, 7, 114, 2, 2, 97, 98, 7, 103, 2, 2, 98, 99, 7, 99, 2, 2,
	99, 100, 7, 118, 2, 2, 100, 22, 3, 2, 2, 2, 101, 102, 7, 107, 2, 2, 102,
	103, 7, 104, 2, 2, 103, 24, 3, 2, 2, 2, 104, 105, 7, 118, 2, 2, 105, 106,
	7, 106, 2, 2, 106, 107, 7, 103, 2, 2, 107, 108, 7, 112, 2, 2, 108, 26,
	3, 2, 2, 2, 109, 110, 7, 103, 2, 2, 110, 111, 7, 112, 2, 2, 111, 112, 7,
	102, 2, 2, 112, 113, 7, 107, 2, 2, 113, 114, 7, 104, 2, 2, 114, 28, 3,
	2, 2, 2, 115, 116, 7, 117, 2, 2, 116, 117, 7, 106, 2, 2, 117, 118, 7, 113,
	2, 2, 118, 119, 7, 121, 2, 2, 119, 30, 3, 2, 2, 2, 120, 121, 7, 117, 2,
	2, 121, 122, 7, 106, 2, 2, 122, 123, 7, 113, 2, 2, 123, 124, 7, 121, 2,
	2, 124, 125, 7, 97, 2, 2, 125, 126, 7, 99, 2, 2, 126, 127, 7, 116, 2, 2,
	127, 128, 7, 116, 2, 2, 128, 129, 7, 99, 2, 2, 129, 130, 7, 123, 2, 2,
	130, 131, 7, 97, 2, 2, 131, 132, 7, 103, 2, 2, 132, 133, 7, 110, 2, 2,
	133, 134, 7, 103, 2, 2, 134, 135, 7, 111, 2, 2, 135, 32, 3, 2, 2, 2, 136,
	137, 7, 116, 2, 2, 137, 138, 7, 103, 2, 2, 138, 139, 7, 99, 2, 2, 139,
	140, 7, 102, 2, 2, 140, 141, 7, 97, 2, 2, 141, 142, 7, 107, 2, 2, 142,
	143, 7, 112, 2, 2, 143, 144, 7, 118, 2, 2, 144, 34, 3, 2, 2, 2, 145, 146,
	7, 116, 2, 2, 146, 147, 7, 103, 2, 2, 147, 148, 7, 99, 2, 2, 148, 149,
	7, 102, 2, 2, 149, 150, 7, 97, 2, 2, 150, 151, 7, 102, 2, 2, 151, 152,
	7, 113, 2, 2, 152, 153, 7, 119, 2, 2, 153, 154, 7, 100, 2, 2, 154, 155,
	7, 110, 2, 2, 155, 156, 7, 103, 2, 2, 156, 36, 3, 2, 2, 2, 157, 158, 7,
	116, 2, 2, 158, 159, 7, 103, 2, 2, 159, 160, 7, 99, 2, 2, 160, 161, 7,
	102, 2, 2, 161, 162, 7, 97, 2, 2, 162, 163, 7, 99, 2, 2, 163, 164, 7, 116,
	2, 2, 164, 165, 7, 116, 2, 2, 165, 166, 7, 99, 2, 2, 166, 167, 7, 123,
	2, 2, 167, 38, 3, 2, 2, 2, 168, 170, 9, 2, 2, 2, 169, 168, 3, 2, 2, 2,
	170, 171, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172,
	40, 3, 2, 2, 2, 173, 178, 7, 36, 2, 2, 174, 177, 5, 43, 22, 2, 175, 177,
	10, 3, 2, 2, 176, 174, 3, 2, 2, 2, 176, 175, 3, 2, 2, 2, 177, 180, 3, 2,
	2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 3, 2, 2, 2,
	180, 178, 3, 2, 2, 2, 181, 182, 7, 36, 2, 2, 182, 42, 3, 2, 2, 2, 183,
	184, 7, 94, 2, 2, 184, 188, 7, 36, 2, 2, 185, 186, 7, 94, 2, 2, 186, 188,
	7, 94, 2, 2, 187, 183, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 188, 44, 3, 2,
	2, 2, 189, 190, 7, 42, 2, 2, 190, 191, 7, 107, 2, 2, 191, 192, 7, 112,
	2, 2, 192, 193, 7, 118, 2, 2, 193, 194, 7, 43, 2, 2, 194, 46, 3, 2, 2,
	2, 195, 196, 7, 42, 2, 2, 196, 197, 7, 116, 2, 2, 197, 198, 7, 103, 2,
	2, 198, 199, 7, 99, 2, 2, 199, 200, 7, 110, 2, 2, 200, 201, 7, 43, 2, 2,
	201, 48, 3, 2, 2, 2, 202, 204, 4, 50, 59, 2, 203, 202, 3, 2, 2, 2, 204,
	205, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 50, 3,
	2, 2, 2, 207, 209, 4, 50, 59, 2, 208, 207, 3, 2, 2, 2, 209, 210, 3, 2,
	2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2,
	212, 214, 7, 48, 2, 2, 213, 215, 4, 50, 59, 2, 214, 213, 3, 2, 2, 2, 215,
	216, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 52, 3,
	2, 2, 2, 218, 219, 7, 63, 2, 2, 219, 54, 3, 2, 2, 2, 220, 221, 7, 45, 2,
	2, 221, 56, 3, 2, 2, 2, 222, 223, 7, 44, 2, 2, 223, 58, 3, 2, 2, 2, 224,
	225, 7, 47, 2, 2, 225, 60, 3, 2, 2, 2, 226, 227, 7, 49, 2, 2, 227, 62,
	3, 2, 2, 2, 228, 230, 7, 15, 2, 2, 229, 228, 3, 2, 2, 2, 229, 230, 3, 2,
	2, 2, 230, 231, 3, 2, 2, 2, 231, 232, 7, 12, 2, 2, 232, 64, 3, 2, 2, 2,
	233, 235, 9, 4, 2, 2, 234, 233, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236,
	234, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 239,
	8, 33, 2, 2, 239, 66, 3, 2, 2, 2, 12, 2, 171, 176, 178, 187, 205, 210,
	216, 229, 236, 3, 8, 2, 2,
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
	"", "'['", "']'", "'{'", "'}'", "'('", "')'", "'=='", "','", "'repeat'",
	"'endrepeat'", "'if'", "'then'", "'endif'", "'show'", "'show_array_elem'",
	"'read_int'", "'read_double'", "'read_array'", "", "", "'(int)'", "'(real)'",
	"", "", "'='", "'+'", "'*'", "'-'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "REPEAT", "ENDREPEAT", "IF", "THEN",
	"ENDIF", "SHOW", "SHOWARRAYELEM", "READINT", "READDOUBLE", "READARRAY",
	"ID", "STRING", "TOINT", "TOREAL", "INT", "REAL", "ASSIGN", "ADD", "MUL",
	"SUB", "DIV", "NEWLINE", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "REPEAT",
	"ENDREPEAT", "IF", "THEN", "ENDIF", "SHOW", "SHOWARRAYELEM", "READINT",
	"READDOUBLE", "READARRAY", "ID", "STRING", "ESC", "TOINT", "TOREAL", "INT",
	"REAL", "ASSIGN", "ADD", "MUL", "SUB", "DIV", "NEWLINE", "WS",
}

type simplelangLexer struct {
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

func NewsimplelangLexer(input antlr.CharStream) *simplelangLexer {

	l := new(simplelangLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "simplelang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// simplelangLexer tokens.
const (
	simplelangLexerT__0          = 1
	simplelangLexerT__1          = 2
	simplelangLexerT__2          = 3
	simplelangLexerT__3          = 4
	simplelangLexerT__4          = 5
	simplelangLexerT__5          = 6
	simplelangLexerT__6          = 7
	simplelangLexerT__7          = 8
	simplelangLexerREPEAT        = 9
	simplelangLexerENDREPEAT     = 10
	simplelangLexerIF            = 11
	simplelangLexerTHEN          = 12
	simplelangLexerENDIF         = 13
	simplelangLexerSHOW          = 14
	simplelangLexerSHOWARRAYELEM = 15
	simplelangLexerREADINT       = 16
	simplelangLexerREADDOUBLE    = 17
	simplelangLexerREADARRAY     = 18
	simplelangLexerID            = 19
	simplelangLexerSTRING        = 20
	simplelangLexerTOINT         = 21
	simplelangLexerTOREAL        = 22
	simplelangLexerINT           = 23
	simplelangLexerREAL          = 24
	simplelangLexerASSIGN        = 25
	simplelangLexerADD           = 26
	simplelangLexerMUL           = 27
	simplelangLexerSUB           = 28
	simplelangLexerDIV           = 29
	simplelangLexerNEWLINE       = 30
	simplelangLexerWS            = 31
)