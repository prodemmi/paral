// Code generated from /home/emad/Documents/paral/ParalExpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
	staticData.LiteralNames = []string{
		"", "'var'", "' '", "'exec'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "IDENT", "VALUE", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "IDENT", "VALUE", "NEWLINE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 6, 39, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 4, 3, 26, 8, 3, 11, 3, 12, 3, 27, 1, 4, 4, 4, 31,
		8, 4, 11, 4, 12, 4, 32, 1, 5, 4, 5, 36, 8, 5, 11, 5, 12, 5, 37, 0, 0, 6,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 1, 0, 3, 3, 0, 65, 90, 95, 95, 97,
		122, 5, 0, 46, 46, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13,
		41, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0,
		0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 1, 13, 1, 0, 0, 0, 3, 17, 1, 0,
		0, 0, 5, 19, 1, 0, 0, 0, 7, 25, 1, 0, 0, 0, 9, 30, 1, 0, 0, 0, 11, 35,
		1, 0, 0, 0, 13, 14, 5, 118, 0, 0, 14, 15, 5, 97, 0, 0, 15, 16, 5, 114,
		0, 0, 16, 2, 1, 0, 0, 0, 17, 18, 5, 32, 0, 0, 18, 4, 1, 0, 0, 0, 19, 20,
		5, 101, 0, 0, 20, 21, 5, 120, 0, 0, 21, 22, 5, 101, 0, 0, 22, 23, 5, 99,
		0, 0, 23, 6, 1, 0, 0, 0, 24, 26, 7, 0, 0, 0, 25, 24, 1, 0, 0, 0, 26, 27,
		1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 8, 1, 0, 0, 0,
		29, 31, 7, 1, 0, 0, 30, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 30, 1,
		0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 10, 1, 0, 0, 0, 34, 36, 7, 2, 0, 0, 35,
		34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0,
		0, 38, 12, 1, 0, 0, 0, 4, 0, 27, 32, 37, 0,
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
	ParalExprLexerT__0    = 1
	ParalExprLexerT__1    = 2
	ParalExprLexerT__2    = 3
	ParalExprLexerIDENT   = 4
	ParalExprLexerVALUE   = 5
	ParalExprLexerNEWLINE = 6
)
