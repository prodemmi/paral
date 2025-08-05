// Code generated from antlr/ParalParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ParalParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ParalParser struct {
	*antlr.BaseParser
}

var ParalParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func paralparserParserInit() {
	staticData := &ParalParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'->'", "'@'", "", "'if'", "'endif'", "'task'", "'='", "", "", "",
		"", "", "", "", "", "':'", "'::'", "", "", "", "'('", "", "'{'", "'}'",
		"", "", "", "", "", "'@stash['", "", "", "", "", "", "", "", "", "",
		"", "", "'@endif'",
	}
	staticData.SymbolicNames = []string{
		"", "CMD_ARROW", "AT", "FUNCTION_START", "IF", "ENDIF", "TASK", "ASSIGN",
		"STRING", "SINGLE_QUOTE_STRING", "FLOAT", "NUMBER", "BOOLEAN", "ZERO_ONE",
		"DURATION", "URL", "COLON", "COLONCOLON", "COMMA", "LBRACK", "RBRACK",
		"LRBRACK", "RRBRACK", "LBRACE", "RBRACE", "DOUBLE_BACK_ARROW", "IDENTIFIER",
		"NEWLINE", "WS", "PIPELINE_NEWLINE", "PIPELINE_STASH", "PIPELINE_IF",
		"PIPELINE_LOOP_KEY", "PIPELINE_LOOP_VALUE", "FUNCTION_CALL_START", "PIPELINE_WS",
		"COMMAND_RAW_TEXT", "STASH_WS", "NESTED_FUNCTION_START", "FUNCTION_END",
		"FUNCTION_WS", "NESTED_IF", "IF_END", "IF_WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "variable_assignment", "task_definition", "task_directive",
		"pipeline_block", "pipeline_content", "pipeline_item", "directive",
		"stash", "condition", "function", "nested_function", "if_condition",
		"argument_list", "expression", "number_expr", "string_expr", "boolean_expr",
		"duration_expr", "matrix_expr", "list_expr", "loop_variable",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 43, 240, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 5, 0, 48, 8, 0, 10, 0, 12, 0, 51, 9, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 58, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 5, 3, 68, 8, 3, 10, 3, 12, 3, 71, 9, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 5, 3, 77, 8, 3, 10, 3, 12, 3, 80, 9, 3, 1, 3, 5, 3, 83, 8,
		3, 10, 3, 12, 3, 86, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 6, 5, 6, 97, 8, 6, 10, 6, 12, 6, 100, 9, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 106, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 111, 8, 8, 10, 8, 12, 8, 114,
		9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 3,
		11, 126, 8, 11, 1, 11, 1, 11, 1, 11, 3, 11, 131, 8, 11, 1, 11, 3, 11, 134,
		8, 11, 1, 12, 1, 12, 3, 12, 138, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 5, 13, 148, 8, 13, 10, 13, 12, 13, 151, 9, 13,
		1, 13, 1, 13, 5, 13, 155, 8, 13, 10, 13, 12, 13, 158, 9, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 5, 14, 165, 8, 14, 10, 14, 12, 14, 168, 9, 14,
		1, 14, 1, 14, 5, 14, 172, 8, 14, 10, 14, 12, 14, 175, 9, 14, 1, 14, 5,
		14, 178, 8, 14, 10, 14, 12, 14, 181, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 194, 8, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 3, 19, 204, 8, 19,
		1, 20, 1, 20, 1, 20, 4, 20, 209, 8, 20, 11, 20, 12, 20, 210, 1, 21, 1,
		21, 1, 21, 5, 21, 216, 8, 21, 10, 21, 12, 21, 219, 9, 21, 1, 21, 1, 21,
		5, 21, 223, 8, 21, 10, 21, 12, 21, 226, 9, 21, 1, 21, 5, 21, 229, 8, 21,
		10, 21, 12, 21, 232, 9, 21, 3, 21, 234, 8, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 0, 0, 23, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 0, 4, 1, 0, 10, 11, 1, 0, 8, 9, 1,
		0, 12, 13, 1, 0, 32, 33, 252, 0, 49, 1, 0, 0, 0, 2, 57, 1, 0, 0, 0, 4,
		59, 1, 0, 0, 0, 6, 69, 1, 0, 0, 0, 8, 89, 1, 0, 0, 0, 10, 91, 1, 0, 0,
		0, 12, 98, 1, 0, 0, 0, 14, 105, 1, 0, 0, 0, 16, 107, 1, 0, 0, 0, 18, 115,
		1, 0, 0, 0, 20, 121, 1, 0, 0, 0, 22, 133, 1, 0, 0, 0, 24, 135, 1, 0, 0,
		0, 26, 141, 1, 0, 0, 0, 28, 162, 1, 0, 0, 0, 30, 193, 1, 0, 0, 0, 32, 195,
		1, 0, 0, 0, 34, 197, 1, 0, 0, 0, 36, 199, 1, 0, 0, 0, 38, 203, 1, 0, 0,
		0, 40, 205, 1, 0, 0, 0, 42, 212, 1, 0, 0, 0, 44, 237, 1, 0, 0, 0, 46, 48,
		3, 2, 1, 0, 47, 46, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0,
		49, 50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 53, 5,
		0, 0, 1, 53, 1, 1, 0, 0, 0, 54, 58, 3, 4, 2, 0, 55, 58, 3, 6, 3, 0, 56,
		58, 5, 27, 0, 0, 57, 54, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 56, 1, 0,
		0, 0, 58, 3, 1, 0, 0, 0, 59, 60, 5, 26, 0, 0, 60, 61, 5, 7, 0, 0, 61, 62,
		3, 30, 15, 0, 62, 63, 5, 27, 0, 0, 63, 5, 1, 0, 0, 0, 64, 65, 3, 8, 4,
		0, 65, 66, 5, 27, 0, 0, 66, 68, 1, 0, 0, 0, 67, 64, 1, 0, 0, 0, 68, 71,
		1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72, 1, 0, 0, 0,
		71, 69, 1, 0, 0, 0, 72, 73, 5, 6, 0, 0, 73, 74, 3, 34, 17, 0, 74, 78, 5,
		23, 0, 0, 75, 77, 5, 27, 0, 0, 76, 75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0,
		78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 84, 1, 0, 0, 0, 80, 78, 1,
		0, 0, 0, 81, 83, 3, 10, 5, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84,
		82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 84, 1, 0, 0,
		0, 87, 88, 5, 24, 0, 0, 88, 7, 1, 0, 0, 0, 89, 90, 3, 16, 8, 0, 90, 9,
		1, 0, 0, 0, 91, 92, 5, 1, 0, 0, 92, 93, 3, 12, 6, 0, 93, 94, 5, 29, 0,
		0, 94, 11, 1, 0, 0, 0, 95, 97, 3, 14, 7, 0, 96, 95, 1, 0, 0, 0, 97, 100,
		1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 13, 1, 0, 0, 0,
		100, 98, 1, 0, 0, 0, 101, 106, 3, 18, 9, 0, 102, 106, 3, 20, 10, 0, 103,
		106, 3, 22, 11, 0, 104, 106, 5, 36, 0, 0, 105, 101, 1, 0, 0, 0, 105, 102,
		1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106, 15, 1, 0,
		0, 0, 107, 108, 5, 2, 0, 0, 108, 112, 5, 26, 0, 0, 109, 111, 3, 30, 15,
		0, 110, 109, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112,
		113, 1, 0, 0, 0, 113, 17, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 5,
		30, 0, 0, 116, 117, 3, 34, 17, 0, 117, 118, 5, 20, 0, 0, 118, 119, 5, 25,
		0, 0, 119, 120, 3, 12, 6, 0, 120, 19, 1, 0, 0, 0, 121, 122, 3, 26, 13,
		0, 122, 21, 1, 0, 0, 0, 123, 125, 5, 34, 0, 0, 124, 126, 3, 28, 14, 0,
		125, 124, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127,
		134, 5, 39, 0, 0, 128, 130, 5, 3, 0, 0, 129, 131, 3, 28, 14, 0, 130, 129,
		1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 5, 39,
		0, 0, 133, 123, 1, 0, 0, 0, 133, 128, 1, 0, 0, 0, 134, 23, 1, 0, 0, 0,
		135, 137, 5, 38, 0, 0, 136, 138, 3, 28, 14, 0, 137, 136, 1, 0, 0, 0, 137,
		138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 5, 39, 0, 0, 140, 25,
		1, 0, 0, 0, 141, 142, 5, 2, 0, 0, 142, 143, 5, 4, 0, 0, 143, 144, 5, 21,
		0, 0, 144, 145, 3, 30, 15, 0, 145, 149, 5, 22, 0, 0, 146, 148, 5, 27, 0,
		0, 147, 146, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149,
		150, 1, 0, 0, 0, 150, 152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 156,
		3, 12, 6, 0, 153, 155, 5, 27, 0, 0, 154, 153, 1, 0, 0, 0, 155, 158, 1,
		0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 159, 1, 0, 0,
		0, 158, 156, 1, 0, 0, 0, 159, 160, 5, 2, 0, 0, 160, 161, 5, 5, 0, 0, 161,
		27, 1, 0, 0, 0, 162, 179, 3, 30, 15, 0, 163, 165, 5, 27, 0, 0, 164, 163,
		1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0,
		0, 0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 173, 5, 18, 0, 0,
		170, 172, 5, 27, 0, 0, 171, 170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173,
		171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 173,
		1, 0, 0, 0, 176, 178, 3, 30, 15, 0, 177, 166, 1, 0, 0, 0, 178, 181, 1,
		0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 29, 1, 0, 0,
		0, 181, 179, 1, 0, 0, 0, 182, 194, 3, 44, 22, 0, 183, 194, 3, 22, 11, 0,
		184, 194, 3, 24, 12, 0, 185, 194, 5, 15, 0, 0, 186, 194, 3, 32, 16, 0,
		187, 194, 3, 34, 17, 0, 188, 194, 3, 36, 18, 0, 189, 194, 3, 38, 19, 0,
		190, 194, 3, 40, 20, 0, 191, 194, 3, 42, 21, 0, 192, 194, 5, 26, 0, 0,
		193, 182, 1, 0, 0, 0, 193, 183, 1, 0, 0, 0, 193, 184, 1, 0, 0, 0, 193,
		185, 1, 0, 0, 0, 193, 186, 1, 0, 0, 0, 193, 187, 1, 0, 0, 0, 193, 188,
		1, 0, 0, 0, 193, 189, 1, 0, 0, 0, 193, 190, 1, 0, 0, 0, 193, 191, 1, 0,
		0, 0, 193, 192, 1, 0, 0, 0, 194, 31, 1, 0, 0, 0, 195, 196, 7, 0, 0, 0,
		196, 33, 1, 0, 0, 0, 197, 198, 7, 1, 0, 0, 198, 35, 1, 0, 0, 0, 199, 200,
		7, 2, 0, 0, 200, 37, 1, 0, 0, 0, 201, 204, 5, 14, 0, 0, 202, 204, 3, 32,
		16, 0, 203, 201, 1, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204, 39, 1, 0, 0, 0,
		205, 208, 3, 42, 21, 0, 206, 207, 5, 17, 0, 0, 207, 209, 3, 42, 21, 0,
		208, 206, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210,
		211, 1, 0, 0, 0, 211, 41, 1, 0, 0, 0, 212, 233, 5, 19, 0, 0, 213, 230,
		3, 30, 15, 0, 214, 216, 5, 27, 0, 0, 215, 214, 1, 0, 0, 0, 216, 219, 1,
		0, 0, 0, 217, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 220, 1, 0, 0,
		0, 219, 217, 1, 0, 0, 0, 220, 224, 5, 18, 0, 0, 221, 223, 5, 27, 0, 0,
		222, 221, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224,
		225, 1, 0, 0, 0, 225, 227, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 229,
		3, 30, 15, 0, 228, 217, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1,
		0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0, 232, 230, 1, 0, 0,
		0, 233, 213, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235,
		236, 5, 20, 0, 0, 236, 43, 1, 0, 0, 0, 237, 238, 7, 3, 0, 0, 238, 45, 1,
		0, 0, 0, 24, 49, 57, 69, 78, 84, 98, 105, 112, 125, 130, 133, 137, 149,
		156, 166, 173, 179, 193, 203, 210, 217, 224, 230, 233,
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

// ParalParserInit initializes any static state used to implement ParalParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewParalParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParalParserInit() {
	staticData := &ParalParserParserStaticData
	staticData.once.Do(paralparserParserInit)
}

// NewParalParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParalParser(input antlr.TokenStream) *ParalParser {
	ParalParserInit()
	this := new(ParalParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ParalParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ParalParser.g4"

	return this
}

// ParalParser tokens.
const (
	ParalParserEOF                   = antlr.TokenEOF
	ParalParserCMD_ARROW             = 1
	ParalParserAT                    = 2
	ParalParserFUNCTION_START        = 3
	ParalParserIF                    = 4
	ParalParserENDIF                 = 5
	ParalParserTASK                  = 6
	ParalParserASSIGN                = 7
	ParalParserSTRING                = 8
	ParalParserSINGLE_QUOTE_STRING   = 9
	ParalParserFLOAT                 = 10
	ParalParserNUMBER                = 11
	ParalParserBOOLEAN               = 12
	ParalParserZERO_ONE              = 13
	ParalParserDURATION              = 14
	ParalParserURL                   = 15
	ParalParserCOLON                 = 16
	ParalParserCOLONCOLON            = 17
	ParalParserCOMMA                 = 18
	ParalParserLBRACK                = 19
	ParalParserRBRACK                = 20
	ParalParserLRBRACK               = 21
	ParalParserRRBRACK               = 22
	ParalParserLBRACE                = 23
	ParalParserRBRACE                = 24
	ParalParserDOUBLE_BACK_ARROW     = 25
	ParalParserIDENTIFIER            = 26
	ParalParserNEWLINE               = 27
	ParalParserWS                    = 28
	ParalParserPIPELINE_NEWLINE      = 29
	ParalParserPIPELINE_STASH        = 30
	ParalParserPIPELINE_IF           = 31
	ParalParserPIPELINE_LOOP_KEY     = 32
	ParalParserPIPELINE_LOOP_VALUE   = 33
	ParalParserFUNCTION_CALL_START   = 34
	ParalParserPIPELINE_WS           = 35
	ParalParserCOMMAND_RAW_TEXT      = 36
	ParalParserSTASH_WS              = 37
	ParalParserNESTED_FUNCTION_START = 38
	ParalParserFUNCTION_END          = 39
	ParalParserFUNCTION_WS           = 40
	ParalParserNESTED_IF             = 41
	ParalParserIF_END                = 42
	ParalParserIF_WS                 = 43
)

// ParalParser rules.
const (
	ParalParserRULE_program             = 0
	ParalParserRULE_statement           = 1
	ParalParserRULE_variable_assignment = 2
	ParalParserRULE_task_definition     = 3
	ParalParserRULE_task_directive      = 4
	ParalParserRULE_pipeline_block      = 5
	ParalParserRULE_pipeline_content    = 6
	ParalParserRULE_pipeline_item       = 7
	ParalParserRULE_directive           = 8
	ParalParserRULE_stash               = 9
	ParalParserRULE_condition           = 10
	ParalParserRULE_function            = 11
	ParalParserRULE_nested_function     = 12
	ParalParserRULE_if_condition        = 13
	ParalParserRULE_argument_list       = 14
	ParalParserRULE_expression          = 15
	ParalParserRULE_number_expr         = 16
	ParalParserRULE_string_expr         = 17
	ParalParserRULE_boolean_expr        = 18
	ParalParserRULE_duration_expr       = 19
	ParalParserRULE_matrix_expr         = 20
	ParalParserRULE_list_expr           = 21
	ParalParserRULE_loop_variable       = 22
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParalParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *ParalParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParalParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&201326660) != 0 {
		{
			p.SetState(46)
			p.Statement()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
		p.Match(ParalParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable_assignment() IVariable_assignmentContext
	Task_definition() ITask_definitionContext
	NEWLINE() antlr.TerminalNode

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Variable_assignment() IVariable_assignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_assignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_assignmentContext)
}

func (s *StatementContext) Task_definition() ITask_definitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITask_definitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITask_definitionContext)
}

func (s *StatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *ParalParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParalParserRULE_statement)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Variable_assignment()
		}

	case ParalParserAT, ParalParserTASK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Task_definition()
		}

	case ParalParserNEWLINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariable_assignmentContext is an interface to support dynamic dispatch.
type IVariable_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	NEWLINE() antlr.TerminalNode

	// IsVariable_assignmentContext differentiates from other interfaces.
	IsVariable_assignmentContext()
}

type Variable_assignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariable_assignmentContext() *Variable_assignmentContext {
	var p = new(Variable_assignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_variable_assignment
	return p
}

func InitEmptyVariable_assignmentContext(p *Variable_assignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_variable_assignment
}

func (*Variable_assignmentContext) IsVariable_assignmentContext() {}

func NewVariable_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_assignmentContext {
	var p = new(Variable_assignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_variable_assignment

	return p
}

func (s *Variable_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_assignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *Variable_assignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ParalParserASSIGN, 0)
}

func (s *Variable_assignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Variable_assignmentContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, 0)
}

func (s *Variable_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterVariable_assignment(s)
	}
}

func (s *Variable_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitVariable_assignment(s)
	}
}

func (p *ParalParser) Variable_assignment() (localctx IVariable_assignmentContext) {
	localctx = NewVariable_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParalParserRULE_variable_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(ParalParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Expression()
	}
	{
		p.SetState(62)
		p.Match(ParalParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITask_definitionContext is an interface to support dynamic dispatch.
type ITask_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TASK() antlr.TerminalNode
	String_expr() IString_exprContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllTask_directive() []ITask_directiveContext
	Task_directive(i int) ITask_directiveContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllPipeline_block() []IPipeline_blockContext
	Pipeline_block(i int) IPipeline_blockContext

	// IsTask_definitionContext differentiates from other interfaces.
	IsTask_definitionContext()
}

type Task_definitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTask_definitionContext() *Task_definitionContext {
	var p = new(Task_definitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_task_definition
	return p
}

func InitEmptyTask_definitionContext(p *Task_definitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_task_definition
}

func (*Task_definitionContext) IsTask_definitionContext() {}

func NewTask_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Task_definitionContext {
	var p = new(Task_definitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_task_definition

	return p
}

func (s *Task_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Task_definitionContext) TASK() antlr.TerminalNode {
	return s.GetToken(ParalParserTASK, 0)
}

func (s *Task_definitionContext) String_expr() IString_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_exprContext)
}

func (s *Task_definitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ParalParserLBRACE, 0)
}

func (s *Task_definitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ParalParserRBRACE, 0)
}

func (s *Task_definitionContext) AllTask_directive() []ITask_directiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITask_directiveContext); ok {
			len++
		}
	}

	tst := make([]ITask_directiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITask_directiveContext); ok {
			tst[i] = t.(ITask_directiveContext)
			i++
		}
	}

	return tst
}

func (s *Task_definitionContext) Task_directive(i int) ITask_directiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITask_directiveContext); ok {
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

	return t.(ITask_directiveContext)
}

func (s *Task_definitionContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNEWLINE)
}

func (s *Task_definitionContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, i)
}

func (s *Task_definitionContext) AllPipeline_block() []IPipeline_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPipeline_blockContext); ok {
			len++
		}
	}

	tst := make([]IPipeline_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPipeline_blockContext); ok {
			tst[i] = t.(IPipeline_blockContext)
			i++
		}
	}

	return tst
}

func (s *Task_definitionContext) Pipeline_block(i int) IPipeline_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeline_blockContext); ok {
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

	return t.(IPipeline_blockContext)
}

func (s *Task_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Task_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Task_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterTask_definition(s)
	}
}

func (s *Task_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitTask_definition(s)
	}
}

func (p *ParalParser) Task_definition() (localctx ITask_definitionContext) {
	localctx = NewTask_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParalParserRULE_task_definition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserAT {
		{
			p.SetState(64)
			p.Task_directive()
		}
		{
			p.SetState(65)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(ParalParserTASK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.String_expr()
	}
	{
		p.SetState(74)
		p.Match(ParalParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(75)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserCMD_ARROW {
		{
			p.SetState(81)
			p.Pipeline_block()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(ParalParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITask_directiveContext is an interface to support dynamic dispatch.
type ITask_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Directive() IDirectiveContext

	// IsTask_directiveContext differentiates from other interfaces.
	IsTask_directiveContext()
}

type Task_directiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTask_directiveContext() *Task_directiveContext {
	var p = new(Task_directiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_task_directive
	return p
}

func InitEmptyTask_directiveContext(p *Task_directiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_task_directive
}

func (*Task_directiveContext) IsTask_directiveContext() {}

func NewTask_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Task_directiveContext {
	var p = new(Task_directiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_task_directive

	return p
}

func (s *Task_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Task_directiveContext) Directive() IDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *Task_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Task_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Task_directiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterTask_directive(s)
	}
}

func (s *Task_directiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitTask_directive(s)
	}
}

func (p *ParalParser) Task_directive() (localctx ITask_directiveContext) {
	localctx = NewTask_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ParalParserRULE_task_directive)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Directive()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPipeline_blockContext is an interface to support dynamic dispatch.
type IPipeline_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CMD_ARROW() antlr.TerminalNode
	Pipeline_content() IPipeline_contentContext
	PIPELINE_NEWLINE() antlr.TerminalNode

	// IsPipeline_blockContext differentiates from other interfaces.
	IsPipeline_blockContext()
}

type Pipeline_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeline_blockContext() *Pipeline_blockContext {
	var p = new(Pipeline_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_pipeline_block
	return p
}

func InitEmptyPipeline_blockContext(p *Pipeline_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_pipeline_block
}

func (*Pipeline_blockContext) IsPipeline_blockContext() {}

func NewPipeline_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pipeline_blockContext {
	var p = new(Pipeline_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_pipeline_block

	return p
}

func (s *Pipeline_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Pipeline_blockContext) CMD_ARROW() antlr.TerminalNode {
	return s.GetToken(ParalParserCMD_ARROW, 0)
}

func (s *Pipeline_blockContext) Pipeline_content() IPipeline_contentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeline_contentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeline_contentContext)
}

func (s *Pipeline_blockContext) PIPELINE_NEWLINE() antlr.TerminalNode {
	return s.GetToken(ParalParserPIPELINE_NEWLINE, 0)
}

func (s *Pipeline_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pipeline_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pipeline_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterPipeline_block(s)
	}
}

func (s *Pipeline_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitPipeline_block(s)
	}
}

func (p *ParalParser) Pipeline_block() (localctx IPipeline_blockContext) {
	localctx = NewPipeline_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ParalParserRULE_pipeline_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(ParalParserCMD_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Pipeline_content()
	}
	{
		p.SetState(93)
		p.Match(ParalParserPIPELINE_NEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPipeline_contentContext is an interface to support dynamic dispatch.
type IPipeline_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPipeline_item() []IPipeline_itemContext
	Pipeline_item(i int) IPipeline_itemContext

	// IsPipeline_contentContext differentiates from other interfaces.
	IsPipeline_contentContext()
}

type Pipeline_contentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeline_contentContext() *Pipeline_contentContext {
	var p = new(Pipeline_contentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_pipeline_content
	return p
}

func InitEmptyPipeline_contentContext(p *Pipeline_contentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_pipeline_content
}

func (*Pipeline_contentContext) IsPipeline_contentContext() {}

func NewPipeline_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pipeline_contentContext {
	var p = new(Pipeline_contentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_pipeline_content

	return p
}

func (s *Pipeline_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Pipeline_contentContext) AllPipeline_item() []IPipeline_itemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPipeline_itemContext); ok {
			len++
		}
	}

	tst := make([]IPipeline_itemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPipeline_itemContext); ok {
			tst[i] = t.(IPipeline_itemContext)
			i++
		}
	}

	return tst
}

func (s *Pipeline_contentContext) Pipeline_item(i int) IPipeline_itemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeline_itemContext); ok {
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

	return t.(IPipeline_itemContext)
}

func (s *Pipeline_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pipeline_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pipeline_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterPipeline_content(s)
	}
}

func (s *Pipeline_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitPipeline_content(s)
	}
}

func (p *ParalParser) Pipeline_content() (localctx IPipeline_contentContext) {
	localctx = NewPipeline_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ParalParserRULE_pipeline_content)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(95)
				p.Pipeline_item()
			}

		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPipeline_itemContext is an interface to support dynamic dispatch.
type IPipeline_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Stash() IStashContext
	Condition() IConditionContext
	Function() IFunctionContext
	COMMAND_RAW_TEXT() antlr.TerminalNode

	// IsPipeline_itemContext differentiates from other interfaces.
	IsPipeline_itemContext()
}

type Pipeline_itemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeline_itemContext() *Pipeline_itemContext {
	var p = new(Pipeline_itemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_pipeline_item
	return p
}

func InitEmptyPipeline_itemContext(p *Pipeline_itemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_pipeline_item
}

func (*Pipeline_itemContext) IsPipeline_itemContext() {}

func NewPipeline_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pipeline_itemContext {
	var p = new(Pipeline_itemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_pipeline_item

	return p
}

func (s *Pipeline_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Pipeline_itemContext) Stash() IStashContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStashContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStashContext)
}

func (s *Pipeline_itemContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Pipeline_itemContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *Pipeline_itemContext) COMMAND_RAW_TEXT() antlr.TerminalNode {
	return s.GetToken(ParalParserCOMMAND_RAW_TEXT, 0)
}

func (s *Pipeline_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pipeline_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pipeline_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterPipeline_item(s)
	}
}

func (s *Pipeline_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitPipeline_item(s)
	}
}

func (p *ParalParser) Pipeline_item() (localctx IPipeline_itemContext) {
	localctx = NewPipeline_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParalParserRULE_pipeline_item)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserPIPELINE_STASH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Stash()
		}

	case ParalParserAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Condition()
		}

	case ParalParserFUNCTION_START, ParalParserFUNCTION_CALL_START:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.Function()
		}

	case ParalParserCOMMAND_RAW_TEXT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(104)
			p.Match(ParalParserCOMMAND_RAW_TEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_directive
	return p
}

func InitEmptyDirectiveContext(p *DirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_directive
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) AT() antlr.TerminalNode {
	return s.GetToken(ParalParserAT, 0)
}

func (s *DirectiveContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *DirectiveContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DirectiveContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (p *ParalParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParalParserRULE_directive)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(ParalParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&305010376456) != 0 {
		{
			p.SetState(109)
			p.Expression()
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStashContext is an interface to support dynamic dispatch.
type IStashContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPELINE_STASH() antlr.TerminalNode
	String_expr() IString_exprContext
	RBRACK() antlr.TerminalNode
	DOUBLE_BACK_ARROW() antlr.TerminalNode
	Pipeline_content() IPipeline_contentContext

	// IsStashContext differentiates from other interfaces.
	IsStashContext()
}

type StashContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStashContext() *StashContext {
	var p = new(StashContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_stash
	return p
}

func InitEmptyStashContext(p *StashContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_stash
}

func (*StashContext) IsStashContext() {}

func NewStashContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StashContext {
	var p = new(StashContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_stash

	return p
}

func (s *StashContext) GetParser() antlr.Parser { return s.parser }

func (s *StashContext) PIPELINE_STASH() antlr.TerminalNode {
	return s.GetToken(ParalParserPIPELINE_STASH, 0)
}

func (s *StashContext) String_expr() IString_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_exprContext)
}

func (s *StashContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRBRACK, 0)
}

func (s *StashContext) DOUBLE_BACK_ARROW() antlr.TerminalNode {
	return s.GetToken(ParalParserDOUBLE_BACK_ARROW, 0)
}

func (s *StashContext) Pipeline_content() IPipeline_contentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeline_contentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeline_contentContext)
}

func (s *StashContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StashContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StashContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterStash(s)
	}
}

func (s *StashContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitStash(s)
	}
}

func (p *ParalParser) Stash() (localctx IStashContext) {
	localctx = NewStashContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ParalParserRULE_stash)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(ParalParserPIPELINE_STASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.String_expr()
	}
	{
		p.SetState(117)
		p.Match(ParalParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(ParalParserDOUBLE_BACK_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Pipeline_content()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_condition() IIf_conditionContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) If_condition() IIf_conditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_conditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_conditionContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *ParalParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParalParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.If_condition()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION_CALL_START() antlr.TerminalNode
	FUNCTION_END() antlr.TerminalNode
	Argument_list() IArgument_listContext
	FUNCTION_START() antlr.TerminalNode

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FUNCTION_CALL_START() antlr.TerminalNode {
	return s.GetToken(ParalParserFUNCTION_CALL_START, 0)
}

func (s *FunctionContext) FUNCTION_END() antlr.TerminalNode {
	return s.GetToken(ParalParserFUNCTION_END, 0)
}

func (s *FunctionContext) Argument_list() IArgument_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgument_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgument_listContext)
}

func (s *FunctionContext) FUNCTION_START() antlr.TerminalNode {
	return s.GetToken(ParalParserFUNCTION_START, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *ParalParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ParalParserRULE_function)
	var _la int

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserFUNCTION_CALL_START:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(ParalParserFUNCTION_CALL_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&305010376456) != 0 {
			{
				p.SetState(124)
				p.Argument_list()
			}

		}
		{
			p.SetState(127)
			p.Match(ParalParserFUNCTION_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserFUNCTION_START:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Match(ParalParserFUNCTION_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&305010376456) != 0 {
			{
				p.SetState(129)
				p.Argument_list()
			}

		}
		{
			p.SetState(132)
			p.Match(ParalParserFUNCTION_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INested_functionContext is an interface to support dynamic dispatch.
type INested_functionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NESTED_FUNCTION_START() antlr.TerminalNode
	FUNCTION_END() antlr.TerminalNode
	Argument_list() IArgument_listContext

	// IsNested_functionContext differentiates from other interfaces.
	IsNested_functionContext()
}

type Nested_functionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNested_functionContext() *Nested_functionContext {
	var p = new(Nested_functionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_nested_function
	return p
}

func InitEmptyNested_functionContext(p *Nested_functionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_nested_function
}

func (*Nested_functionContext) IsNested_functionContext() {}

func NewNested_functionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nested_functionContext {
	var p = new(Nested_functionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_nested_function

	return p
}

func (s *Nested_functionContext) GetParser() antlr.Parser { return s.parser }

func (s *Nested_functionContext) NESTED_FUNCTION_START() antlr.TerminalNode {
	return s.GetToken(ParalParserNESTED_FUNCTION_START, 0)
}

func (s *Nested_functionContext) FUNCTION_END() antlr.TerminalNode {
	return s.GetToken(ParalParserFUNCTION_END, 0)
}

func (s *Nested_functionContext) Argument_list() IArgument_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgument_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgument_listContext)
}

func (s *Nested_functionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nested_functionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nested_functionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterNested_function(s)
	}
}

func (s *Nested_functionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitNested_function(s)
	}
}

func (p *ParalParser) Nested_function() (localctx INested_functionContext) {
	localctx = NewNested_functionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ParalParserRULE_nested_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(ParalParserNESTED_FUNCTION_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&305010376456) != 0 {
		{
			p.SetState(136)
			p.Argument_list()
		}

	}
	{
		p.SetState(139)
		p.Match(ParalParserFUNCTION_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_conditionContext is an interface to support dynamic dispatch.
type IIf_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAT() []antlr.TerminalNode
	AT(i int) antlr.TerminalNode
	IF() antlr.TerminalNode
	LRBRACK() antlr.TerminalNode
	Expression() IExpressionContext
	RRBRACK() antlr.TerminalNode
	Pipeline_content() IPipeline_contentContext
	ENDIF() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsIf_conditionContext differentiates from other interfaces.
	IsIf_conditionContext()
}

type If_conditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_conditionContext() *If_conditionContext {
	var p = new(If_conditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_if_condition
	return p
}

func InitEmptyIf_conditionContext(p *If_conditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_if_condition
}

func (*If_conditionContext) IsIf_conditionContext() {}

func NewIf_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_conditionContext {
	var p = new(If_conditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_if_condition

	return p
}

func (s *If_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *If_conditionContext) AllAT() []antlr.TerminalNode {
	return s.GetTokens(ParalParserAT)
}

func (s *If_conditionContext) AT(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserAT, i)
}

func (s *If_conditionContext) IF() antlr.TerminalNode {
	return s.GetToken(ParalParserIF, 0)
}

func (s *If_conditionContext) LRBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserLRBRACK, 0)
}

func (s *If_conditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_conditionContext) RRBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRRBRACK, 0)
}

func (s *If_conditionContext) Pipeline_content() IPipeline_contentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeline_contentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeline_contentContext)
}

func (s *If_conditionContext) ENDIF() antlr.TerminalNode {
	return s.GetToken(ParalParserENDIF, 0)
}

func (s *If_conditionContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNEWLINE)
}

func (s *If_conditionContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, i)
}

func (s *If_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterIf_condition(s)
	}
}

func (s *If_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitIf_condition(s)
	}
}

func (p *ParalParser) If_condition() (localctx IIf_conditionContext) {
	localctx = NewIf_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ParalParserRULE_if_condition)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(ParalParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(ParalParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(ParalParserLRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Expression()
	}
	{
		p.SetState(145)
		p.Match(ParalParserRRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(146)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Pipeline_content()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(153)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(159)
		p.Match(ParalParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(ParalParserENDIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgument_listContext is an interface to support dynamic dispatch.
type IArgument_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsArgument_listContext differentiates from other interfaces.
	IsArgument_listContext()
}

type Argument_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgument_listContext() *Argument_listContext {
	var p = new(Argument_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_argument_list
	return p
}

func InitEmptyArgument_listContext(p *Argument_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_argument_list
}

func (*Argument_listContext) IsArgument_listContext() {}

func NewArgument_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Argument_listContext {
	var p = new(Argument_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_argument_list

	return p
}

func (s *Argument_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Argument_listContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Argument_listContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *Argument_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ParalParserCOMMA)
}

func (s *Argument_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserCOMMA, i)
}

func (s *Argument_listContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNEWLINE)
}

func (s *Argument_listContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, i)
}

func (s *Argument_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Argument_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Argument_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterArgument_list(s)
	}
}

func (s *Argument_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitArgument_list(s)
	}
}

func (p *ParalParser) Argument_list() (localctx IArgument_listContext) {
	localctx = NewArgument_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ParalParserRULE_argument_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Expression()
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserCOMMA || _la == ParalParserNEWLINE {
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(163)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(168)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(169)
			p.Match(ParalParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(170)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(176)
			p.Expression()
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Loop_variable() ILoop_variableContext
	Function() IFunctionContext
	Nested_function() INested_functionContext
	URL() antlr.TerminalNode
	Number_expr() INumber_exprContext
	String_expr() IString_exprContext
	Boolean_expr() IBoolean_exprContext
	Duration_expr() IDuration_exprContext
	Matrix_expr() IMatrix_exprContext
	List_expr() IList_exprContext
	IDENTIFIER() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Loop_variable() ILoop_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoop_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoop_variableContext)
}

func (s *ExpressionContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *ExpressionContext) Nested_function() INested_functionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INested_functionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INested_functionContext)
}

func (s *ExpressionContext) URL() antlr.TerminalNode {
	return s.GetToken(ParalParserURL, 0)
}

func (s *ExpressionContext) Number_expr() INumber_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_exprContext)
}

func (s *ExpressionContext) String_expr() IString_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_exprContext)
}

func (s *ExpressionContext) Boolean_expr() IBoolean_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_exprContext)
}

func (s *ExpressionContext) Duration_expr() IDuration_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDuration_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDuration_exprContext)
}

func (s *ExpressionContext) Matrix_expr() IMatrix_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_exprContext)
}

func (s *ExpressionContext) List_expr() IList_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_exprContext)
}

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ParalParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ParalParserRULE_expression)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(182)
			p.Loop_variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Function()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(184)
			p.Nested_function()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(185)
			p.Match(ParalParserURL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(186)
			p.Number_expr()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(187)
			p.String_expr()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(188)
			p.Boolean_expr()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(189)
			p.Duration_expr()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(190)
			p.Matrix_expr()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(191)
			p.List_expr()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(192)
			p.Match(ParalParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumber_exprContext is an interface to support dynamic dispatch.
type INumber_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsNumber_exprContext differentiates from other interfaces.
	IsNumber_exprContext()
}

type Number_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_exprContext() *Number_exprContext {
	var p = new(Number_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_number_expr
	return p
}

func InitEmptyNumber_exprContext(p *Number_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_number_expr
}

func (*Number_exprContext) IsNumber_exprContext() {}

func NewNumber_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_exprContext {
	var p = new(Number_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_number_expr

	return p
}

func (s *Number_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_exprContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ParalParserFLOAT, 0)
}

func (s *Number_exprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ParalParserNUMBER, 0)
}

func (s *Number_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterNumber_expr(s)
	}
}

func (s *Number_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitNumber_expr(s)
	}
}

func (p *ParalParser) Number_expr() (localctx INumber_exprContext) {
	localctx = NewNumber_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ParalParserRULE_number_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ParalParserFLOAT || _la == ParalParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_exprContext is an interface to support dynamic dispatch.
type IString_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	SINGLE_QUOTE_STRING() antlr.TerminalNode

	// IsString_exprContext differentiates from other interfaces.
	IsString_exprContext()
}

type String_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_exprContext() *String_exprContext {
	var p = new(String_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_string_expr
	return p
}

func InitEmptyString_exprContext(p *String_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_string_expr
}

func (*String_exprContext) IsString_exprContext() {}

func NewString_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_exprContext {
	var p = new(String_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_string_expr

	return p
}

func (s *String_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *String_exprContext) STRING() antlr.TerminalNode {
	return s.GetToken(ParalParserSTRING, 0)
}

func (s *String_exprContext) SINGLE_QUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(ParalParserSINGLE_QUOTE_STRING, 0)
}

func (s *String_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterString_expr(s)
	}
}

func (s *String_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitString_expr(s)
	}
}

func (p *ParalParser) String_expr() (localctx IString_exprContext) {
	localctx = NewString_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ParalParserRULE_string_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ParalParserSTRING || _la == ParalParserSINGLE_QUOTE_STRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolean_exprContext is an interface to support dynamic dispatch.
type IBoolean_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN() antlr.TerminalNode
	ZERO_ONE() antlr.TerminalNode

	// IsBoolean_exprContext differentiates from other interfaces.
	IsBoolean_exprContext()
}

type Boolean_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_exprContext() *Boolean_exprContext {
	var p = new(Boolean_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_boolean_expr
	return p
}

func InitEmptyBoolean_exprContext(p *Boolean_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_boolean_expr
}

func (*Boolean_exprContext) IsBoolean_exprContext() {}

func NewBoolean_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_exprContext {
	var p = new(Boolean_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_boolean_expr

	return p
}

func (s *Boolean_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_exprContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ParalParserBOOLEAN, 0)
}

func (s *Boolean_exprContext) ZERO_ONE() antlr.TerminalNode {
	return s.GetToken(ParalParserZERO_ONE, 0)
}

func (s *Boolean_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterBoolean_expr(s)
	}
}

func (s *Boolean_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitBoolean_expr(s)
	}
}

func (p *ParalParser) Boolean_expr() (localctx IBoolean_exprContext) {
	localctx = NewBoolean_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ParalParserRULE_boolean_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ParalParserBOOLEAN || _la == ParalParserZERO_ONE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDuration_exprContext is an interface to support dynamic dispatch.
type IDuration_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DURATION() antlr.TerminalNode
	Number_expr() INumber_exprContext

	// IsDuration_exprContext differentiates from other interfaces.
	IsDuration_exprContext()
}

type Duration_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDuration_exprContext() *Duration_exprContext {
	var p = new(Duration_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_duration_expr
	return p
}

func InitEmptyDuration_exprContext(p *Duration_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_duration_expr
}

func (*Duration_exprContext) IsDuration_exprContext() {}

func NewDuration_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Duration_exprContext {
	var p = new(Duration_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_duration_expr

	return p
}

func (s *Duration_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Duration_exprContext) DURATION() antlr.TerminalNode {
	return s.GetToken(ParalParserDURATION, 0)
}

func (s *Duration_exprContext) Number_expr() INumber_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_exprContext)
}

func (s *Duration_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Duration_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Duration_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterDuration_expr(s)
	}
}

func (s *Duration_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitDuration_expr(s)
	}
}

func (p *ParalParser) Duration_expr() (localctx IDuration_exprContext) {
	localctx = NewDuration_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ParalParserRULE_duration_expr)
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserDURATION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.Match(ParalParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserFLOAT, ParalParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Number_expr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatrix_exprContext is an interface to support dynamic dispatch.
type IMatrix_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllList_expr() []IList_exprContext
	List_expr(i int) IList_exprContext
	AllCOLONCOLON() []antlr.TerminalNode
	COLONCOLON(i int) antlr.TerminalNode

	// IsMatrix_exprContext differentiates from other interfaces.
	IsMatrix_exprContext()
}

type Matrix_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrix_exprContext() *Matrix_exprContext {
	var p = new(Matrix_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_matrix_expr
	return p
}

func InitEmptyMatrix_exprContext(p *Matrix_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_matrix_expr
}

func (*Matrix_exprContext) IsMatrix_exprContext() {}

func NewMatrix_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matrix_exprContext {
	var p = new(Matrix_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_matrix_expr

	return p
}

func (s *Matrix_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Matrix_exprContext) AllList_expr() []IList_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IList_exprContext); ok {
			len++
		}
	}

	tst := make([]IList_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IList_exprContext); ok {
			tst[i] = t.(IList_exprContext)
			i++
		}
	}

	return tst
}

func (s *Matrix_exprContext) List_expr(i int) IList_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_exprContext); ok {
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

	return t.(IList_exprContext)
}

func (s *Matrix_exprContext) AllCOLONCOLON() []antlr.TerminalNode {
	return s.GetTokens(ParalParserCOLONCOLON)
}

func (s *Matrix_exprContext) COLONCOLON(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserCOLONCOLON, i)
}

func (s *Matrix_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matrix_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matrix_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterMatrix_expr(s)
	}
}

func (s *Matrix_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitMatrix_expr(s)
	}
}

func (p *ParalParser) Matrix_expr() (localctx IMatrix_exprContext) {
	localctx = NewMatrix_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ParalParserRULE_matrix_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.List_expr()
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ParalParserCOLONCOLON {
		{
			p.SetState(206)
			p.Match(ParalParserCOLONCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.List_expr()
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_exprContext is an interface to support dynamic dispatch.
type IList_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsList_exprContext differentiates from other interfaces.
	IsList_exprContext()
}

type List_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_exprContext() *List_exprContext {
	var p = new(List_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_list_expr
	return p
}

func InitEmptyList_exprContext(p *List_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_list_expr
}

func (*List_exprContext) IsList_exprContext() {}

func NewList_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_exprContext {
	var p = new(List_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_list_expr

	return p
}

func (s *List_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *List_exprContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserLBRACK, 0)
}

func (s *List_exprContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRBRACK, 0)
}

func (s *List_exprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *List_exprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *List_exprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ParalParserCOMMA)
}

func (s *List_exprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserCOMMA, i)
}

func (s *List_exprContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNEWLINE)
}

func (s *List_exprContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, i)
}

func (s *List_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterList_expr(s)
	}
}

func (s *List_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitList_expr(s)
	}
}

func (p *ParalParser) List_expr() (localctx IList_exprContext) {
	localctx = NewList_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ParalParserRULE_list_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(ParalParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&305010376456) != 0 {
		{
			p.SetState(213)
			p.Expression()
		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserCOMMA || _la == ParalParserNEWLINE {
			p.SetState(217)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParalParserNEWLINE {
				{
					p.SetState(214)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(219)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(220)
				p.Match(ParalParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(224)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParalParserNEWLINE {
				{
					p.SetState(221)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(226)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(227)
				p.Expression()
			}

			p.SetState(232)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(235)
		p.Match(ParalParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoop_variableContext is an interface to support dynamic dispatch.
type ILoop_variableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPELINE_LOOP_KEY() antlr.TerminalNode
	PIPELINE_LOOP_VALUE() antlr.TerminalNode

	// IsLoop_variableContext differentiates from other interfaces.
	IsLoop_variableContext()
}

type Loop_variableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoop_variableContext() *Loop_variableContext {
	var p = new(Loop_variableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_loop_variable
	return p
}

func InitEmptyLoop_variableContext(p *Loop_variableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_loop_variable
}

func (*Loop_variableContext) IsLoop_variableContext() {}

func NewLoop_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Loop_variableContext {
	var p = new(Loop_variableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_loop_variable

	return p
}

func (s *Loop_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Loop_variableContext) PIPELINE_LOOP_KEY() antlr.TerminalNode {
	return s.GetToken(ParalParserPIPELINE_LOOP_KEY, 0)
}

func (s *Loop_variableContext) PIPELINE_LOOP_VALUE() antlr.TerminalNode {
	return s.GetToken(ParalParserPIPELINE_LOOP_VALUE, 0)
}

func (s *Loop_variableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Loop_variableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Loop_variableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterLoop_variable(s)
	}
}

func (s *Loop_variableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitLoop_variable(s)
	}
}

func (p *ParalParser) Loop_variable() (localctx ILoop_variableContext) {
	localctx = NewLoop_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ParalParserRULE_loop_variable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ParalParserPIPELINE_LOOP_KEY || _la == ParalParserPIPELINE_LOOP_VALUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
