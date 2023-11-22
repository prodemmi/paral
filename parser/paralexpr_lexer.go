// Code generated from ParalExpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ParalExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ParalExprLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func paralexprlexerLexerInit() {
	staticData := &ParalExprLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.SymbolicNames = []string{
		"", "IDENT", "VALUE", "NEWLINE", "WS", "VARIABLE",
	}
	staticData.RuleNames = []string{
		"SPACE", "VAR", "IDENT", "VALUE", "NEWLINE", "WS", "VARIABLE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 47, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 24, 8, 1, 1, 2, 4, 2, 27, 8, 2, 11, 2, 12, 2, 28, 1, 3,
		4, 3, 32, 8, 3, 11, 3, 12, 3, 33, 1, 4, 4, 4, 37, 8, 4, 11, 4, 12, 4, 38,
		1, 5, 4, 5, 42, 8, 5, 11, 5, 12, 5, 43, 1, 6, 1, 6, 0, 0, 7, 1, 0, 3, 0,
		5, 1, 7, 2, 9, 3, 11, 4, 13, 5, 1, 0, 4, 3, 0, 9, 10, 12, 13, 32, 32, 3,
		0, 65, 90, 95, 95, 97, 122, 3, 0, 48, 57, 65, 90, 97, 122, 2, 0, 10, 10,
		13, 13, 49, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 1, 15, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5,
		26, 1, 0, 0, 0, 7, 31, 1, 0, 0, 0, 9, 36, 1, 0, 0, 0, 11, 41, 1, 0, 0,
		0, 13, 45, 1, 0, 0, 0, 15, 16, 7, 0, 0, 0, 16, 2, 1, 0, 0, 0, 17, 18, 5,
		118, 0, 0, 18, 19, 5, 97, 0, 0, 19, 24, 5, 114, 0, 0, 20, 21, 5, 86, 0,
		0, 21, 22, 5, 65, 0, 0, 22, 24, 5, 82, 0, 0, 23, 17, 1, 0, 0, 0, 23, 20,
		1, 0, 0, 0, 24, 4, 1, 0, 0, 0, 25, 27, 7, 1, 0, 0, 26, 25, 1, 0, 0, 0,
		27, 28, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 6, 1, 0,
		0, 0, 30, 32, 7, 2, 0, 0, 31, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 31,
		1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 8, 1, 0, 0, 0, 35, 37, 7, 3, 0, 0,
		36, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1,
		0, 0, 0, 39, 10, 1, 0, 0, 0, 40, 42, 3, 1, 0, 0, 41, 40, 1, 0, 0, 0, 42,
		43, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 12, 1, 0, 0,
		0, 45, 46, 3, 3, 1, 0, 46, 14, 1, 0, 0, 0, 7, 0, 23, 28, 31, 33, 38, 43,
		0,
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

// ParalExprLexerInit initializes any static state used to implement ParalExprLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewParalExprLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParalExprLexerInit() {
	staticData := &ParalExprLexerLexerStaticData
	staticData.once.Do(paralexprlexerLexerInit)
}

// NewParalExprLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewParalExprLexer(input antlr.CharStream) *ParalExprLexer {
	ParalExprLexerInit()
	l := new(ParalExprLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ParalExprLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ParalExpr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ParalExprLexer tokens.
const (
	ParalExprLexerIDENT    = 1
	ParalExprLexerVALUE    = 2
	ParalExprLexerNEWLINE  = 3
	ParalExprLexerWS       = 4
	ParalExprLexerVARIABLE = 5
)
