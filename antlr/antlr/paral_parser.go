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
		"", "'task'", "'='", "", "", "", "", "", "", "", "", "", "':'", "'::'",
		"", "", "", "'('", "", "", "", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "TASK", "ASSIGN", "STRING", "SINGLE_QUOTE_STRING", "FLOAT", "NUMBER",
		"BOOLEAN", "ZERO_ONE", "DURATION", "URL", "PATH", "COLON", "COLONCOLON",
		"COMMA", "LBRACK", "RBRACK", "LRBRACK", "RRBRACK", "BLOCK_START", "BLOCK_END",
		"AT", "FUNCTION_START", "PIPELINE_START", "LOOP_KEY", "LOOP_VALUE",
		"IDENTIFIER", "NEWLINE", "WS", "UNKNOWN_DEFAULT", "PIPELINE_BUF", "PIPELINE_STASH",
		"PIPELINE_IF_CALL_START", "PIPELINE_ELSEIF_CALL_START", "PIPELINE_ELSE_CALL",
		"PIPELINE_FUNCTION_CALL_START", "PIPELINE_WS", "UNKNOWN_TEXT", "BUF_DOUBLE_BACK_ARROW",
		"BUF_WS", "STASH_DOUBLE_BACK_ARROW", "STASH_WS", "IF_CONDITION_END",
		"EXPRESSION_FUNCTION_CALL_START", "EXPRESSION_WS", "NESTED_FUNCTION_START",
		"FUNCTION_END", "FUNCTION_WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "variable_assignment", "task_definition", "task_directive",
		"pipeline_block", "pipeline_content", "pipeline_item", "unknown_command",
		"directive", "buf", "stash", "condition", "if_condition", "elseif_condition",
		"else_condition", "function", "nested_function", "argument_list", "expression",
		"number_expr", "string_expr", "boolean_expr", "duration_expr", "matrix_expr",
		"list_expr", "loop_variable",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 47, 375, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 1, 0, 5, 0, 56, 8, 0, 10, 0, 12, 0, 59, 9, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 3, 1, 66, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 5, 3, 76, 8, 3, 10, 3, 12, 3, 79, 9, 3, 1, 3, 1, 3, 1, 3, 5, 3, 84,
		8, 3, 10, 3, 12, 3, 87, 9, 3, 1, 3, 1, 3, 5, 3, 91, 8, 3, 10, 3, 12, 3,
		94, 9, 3, 1, 3, 5, 3, 97, 8, 3, 10, 3, 12, 3, 100, 9, 3, 1, 3, 5, 3, 103,
		8, 3, 10, 3, 12, 3, 106, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5,
		114, 8, 5, 10, 5, 12, 5, 117, 9, 5, 1, 5, 1, 5, 3, 5, 121, 8, 5, 1, 6,
		5, 6, 124, 8, 6, 10, 6, 12, 6, 127, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		3, 7, 134, 8, 7, 1, 8, 4, 8, 137, 8, 8, 11, 8, 12, 8, 138, 1, 8, 5, 8,
		142, 8, 8, 10, 8, 12, 8, 145, 9, 8, 1, 9, 1, 9, 1, 9, 3, 9, 150, 8, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 170, 8, 13,
		10, 13, 12, 13, 173, 9, 13, 1, 13, 1, 13, 5, 13, 177, 8, 13, 10, 13, 12,
		13, 180, 9, 13, 1, 13, 5, 13, 183, 8, 13, 10, 13, 12, 13, 186, 9, 13, 1,
		13, 5, 13, 189, 8, 13, 10, 13, 12, 13, 192, 9, 13, 1, 13, 1, 13, 5, 13,
		196, 8, 13, 10, 13, 12, 13, 199, 9, 13, 1, 13, 5, 13, 202, 8, 13, 10, 13,
		12, 13, 205, 9, 13, 1, 13, 5, 13, 208, 8, 13, 10, 13, 12, 13, 211, 9, 13,
		1, 13, 3, 13, 214, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 220, 8, 14,
		10, 14, 12, 14, 223, 9, 14, 1, 14, 1, 14, 5, 14, 227, 8, 14, 10, 14, 12,
		14, 230, 9, 14, 1, 14, 5, 14, 233, 8, 14, 10, 14, 12, 14, 236, 9, 14, 1,
		14, 5, 14, 239, 8, 14, 10, 14, 12, 14, 242, 9, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 5, 15, 248, 8, 15, 10, 15, 12, 15, 251, 9, 15, 1, 15, 1, 15, 5,
		15, 255, 8, 15, 10, 15, 12, 15, 258, 9, 15, 1, 15, 5, 15, 261, 8, 15, 10,
		15, 12, 15, 264, 9, 15, 1, 15, 5, 15, 267, 8, 15, 10, 15, 12, 15, 270,
		9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 276, 8, 16, 1, 16, 1, 16, 1,
		16, 3, 16, 281, 8, 16, 1, 16, 1, 16, 1, 16, 3, 16, 286, 8, 16, 1, 16, 3,
		16, 289, 8, 16, 1, 17, 1, 17, 3, 17, 293, 8, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 5, 18, 299, 8, 18, 10, 18, 12, 18, 302, 9, 18, 1, 18, 1, 18, 5, 18,
		306, 8, 18, 10, 18, 12, 18, 309, 9, 18, 1, 18, 5, 18, 312, 8, 18, 10, 18,
		12, 18, 315, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 329, 8, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 3, 23, 339, 8, 23, 1, 24, 1, 24, 1,
		24, 4, 24, 344, 8, 24, 11, 24, 12, 24, 345, 1, 25, 1, 25, 1, 25, 5, 25,
		351, 8, 25, 10, 25, 12, 25, 354, 9, 25, 1, 25, 1, 25, 5, 25, 358, 8, 25,
		10, 25, 12, 25, 361, 9, 25, 1, 25, 5, 25, 364, 8, 25, 10, 25, 12, 25, 367,
		9, 25, 3, 25, 369, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 0, 0, 27,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 50, 52, 0, 6, 2, 1, 20, 20, 27, 27, 3, 0, 10, 11,
		26, 26, 37, 37, 1, 0, 5, 6, 1, 0, 3, 4, 1, 0, 7, 8, 1, 0, 24, 25, 407,
		0, 57, 1, 0, 0, 0, 2, 65, 1, 0, 0, 0, 4, 67, 1, 0, 0, 0, 6, 77, 1, 0, 0,
		0, 8, 109, 1, 0, 0, 0, 10, 111, 1, 0, 0, 0, 12, 125, 1, 0, 0, 0, 14, 133,
		1, 0, 0, 0, 16, 136, 1, 0, 0, 0, 18, 146, 1, 0, 0, 0, 20, 151, 1, 0, 0,
		0, 22, 157, 1, 0, 0, 0, 24, 163, 1, 0, 0, 0, 26, 165, 1, 0, 0, 0, 28, 215,
		1, 0, 0, 0, 30, 245, 1, 0, 0, 0, 32, 288, 1, 0, 0, 0, 34, 290, 1, 0, 0,
		0, 36, 296, 1, 0, 0, 0, 38, 328, 1, 0, 0, 0, 40, 330, 1, 0, 0, 0, 42, 332,
		1, 0, 0, 0, 44, 334, 1, 0, 0, 0, 46, 338, 1, 0, 0, 0, 48, 340, 1, 0, 0,
		0, 50, 347, 1, 0, 0, 0, 52, 372, 1, 0, 0, 0, 54, 56, 3, 2, 1, 0, 55, 54,
		1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0,
		58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61, 5, 0, 0, 1, 61, 1, 1, 0,
		0, 0, 62, 66, 3, 4, 2, 0, 63, 66, 3, 6, 3, 0, 64, 66, 5, 27, 0, 0, 65,
		62, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 64, 1, 0, 0, 0, 66, 3, 1, 0, 0,
		0, 67, 68, 5, 26, 0, 0, 68, 69, 5, 2, 0, 0, 69, 70, 3, 38, 19, 0, 70, 71,
		5, 27, 0, 0, 71, 5, 1, 0, 0, 0, 72, 73, 3, 8, 4, 0, 73, 74, 5, 27, 0, 0,
		74, 76, 1, 0, 0, 0, 75, 72, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1,
		0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 80, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80,
		81, 5, 1, 0, 0, 81, 85, 5, 26, 0, 0, 82, 84, 5, 27, 0, 0, 83, 82, 1, 0,
		0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 88,
		1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 92, 5, 19, 0, 0, 89, 91, 5, 27, 0,
		0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93,
		1, 0, 0, 0, 93, 98, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 97, 3, 10, 5, 0,
		96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1,
		0, 0, 0, 99, 104, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 103, 5, 27, 0,
		0, 102, 101, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104,
		105, 1, 0, 0, 0, 105, 107, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 107, 108,
		5, 20, 0, 0, 108, 7, 1, 0, 0, 0, 109, 110, 3, 18, 9, 0, 110, 9, 1, 0, 0,
		0, 111, 115, 5, 23, 0, 0, 112, 114, 5, 27, 0, 0, 113, 112, 1, 0, 0, 0,
		114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116,
		118, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 120, 3, 12, 6, 0, 119, 121,
		7, 0, 0, 0, 120, 119, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 11, 1, 0,
		0, 0, 122, 124, 3, 14, 7, 0, 123, 122, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0,
		125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 13, 1, 0, 0, 0, 127, 125,
		1, 0, 0, 0, 128, 134, 3, 20, 10, 0, 129, 134, 3, 22, 11, 0, 130, 134, 3,
		24, 12, 0, 131, 134, 3, 32, 16, 0, 132, 134, 3, 16, 8, 0, 133, 128, 1,
		0, 0, 0, 133, 129, 1, 0, 0, 0, 133, 130, 1, 0, 0, 0, 133, 131, 1, 0, 0,
		0, 133, 132, 1, 0, 0, 0, 134, 15, 1, 0, 0, 0, 135, 137, 7, 1, 0, 0, 136,
		135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139,
		1, 0, 0, 0, 139, 143, 1, 0, 0, 0, 140, 142, 3, 38, 19, 0, 141, 140, 1,
		0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0,
		0, 144, 17, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146, 147, 5, 21, 0, 0, 147,
		149, 5, 26, 0, 0, 148, 150, 3, 38, 19, 0, 149, 148, 1, 0, 0, 0, 149, 150,
		1, 0, 0, 0, 150, 19, 1, 0, 0, 0, 151, 152, 5, 30, 0, 0, 152, 153, 3, 42,
		21, 0, 153, 154, 5, 16, 0, 0, 154, 155, 5, 38, 0, 0, 155, 156, 3, 12, 6,
		0, 156, 21, 1, 0, 0, 0, 157, 158, 5, 31, 0, 0, 158, 159, 3, 42, 21, 0,
		159, 160, 5, 16, 0, 0, 160, 161, 5, 40, 0, 0, 161, 162, 3, 12, 6, 0, 162,
		23, 1, 0, 0, 0, 163, 164, 3, 26, 13, 0, 164, 25, 1, 0, 0, 0, 165, 166,
		5, 32, 0, 0, 166, 167, 3, 38, 19, 0, 167, 171, 5, 42, 0, 0, 168, 170, 5,
		27, 0, 0, 169, 168, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 0,
		0, 171, 172, 1, 0, 0, 0, 172, 174, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174,
		178, 5, 19, 0, 0, 175, 177, 5, 27, 0, 0, 176, 175, 1, 0, 0, 0, 177, 180,
		1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 184, 1, 0,
		0, 0, 180, 178, 1, 0, 0, 0, 181, 183, 3, 10, 5, 0, 182, 181, 1, 0, 0, 0,
		183, 186, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185,
		190, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 187, 189, 5, 27, 0, 0, 188, 187,
		1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0,
		0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 203, 5, 20, 0, 0,
		194, 196, 5, 27, 0, 0, 195, 194, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197,
		195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 200, 1, 0, 0, 0, 199, 197,
		1, 0, 0, 0, 200, 202, 3, 28, 14, 0, 201, 197, 1, 0, 0, 0, 202, 205, 1,
		0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 213, 1, 0, 0,
		0, 205, 203, 1, 0, 0, 0, 206, 208, 5, 27, 0, 0, 207, 206, 1, 0, 0, 0, 208,
		211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 212,
		1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 212, 214, 3, 30, 15, 0, 213, 209, 1,
		0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 27, 1, 0, 0, 0, 215, 216, 5, 33, 0,
		0, 216, 217, 3, 38, 19, 0, 217, 221, 5, 42, 0, 0, 218, 220, 5, 27, 0, 0,
		219, 218, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221,
		222, 1, 0, 0, 0, 222, 224, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 228,
		5, 19, 0, 0, 225, 227, 5, 27, 0, 0, 226, 225, 1, 0, 0, 0, 227, 230, 1,
		0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 234, 1, 0, 0,
		0, 230, 228, 1, 0, 0, 0, 231, 233, 3, 10, 5, 0, 232, 231, 1, 0, 0, 0, 233,
		236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 240,
		1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 239, 5, 27, 0, 0, 238, 237, 1, 0,
		0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0,
		241, 243, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 244, 5, 20, 0, 0, 244,
		29, 1, 0, 0, 0, 245, 249, 5, 34, 0, 0, 246, 248, 5, 27, 0, 0, 247, 246,
		1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 250, 1, 0,
		0, 0, 250, 252, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 252, 256, 5, 19, 0, 0,
		253, 255, 5, 27, 0, 0, 254, 253, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256,
		254, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 262, 1, 0, 0, 0, 258, 256,
		1, 0, 0, 0, 259, 261, 3, 10, 5, 0, 260, 259, 1, 0, 0, 0, 261, 264, 1, 0,
		0, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 268, 1, 0, 0, 0,
		264, 262, 1, 0, 0, 0, 265, 267, 5, 27, 0, 0, 266, 265, 1, 0, 0, 0, 267,
		270, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 271,
		1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 271, 272, 5, 20, 0, 0, 272, 31, 1, 0,
		0, 0, 273, 275, 5, 35, 0, 0, 274, 276, 3, 36, 18, 0, 275, 274, 1, 0, 0,
		0, 275, 276, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 289, 5, 46, 0, 0, 278,
		280, 5, 22, 0, 0, 279, 281, 3, 36, 18, 0, 280, 279, 1, 0, 0, 0, 280, 281,
		1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 289, 5, 46, 0, 0, 283, 285, 5, 43,
		0, 0, 284, 286, 3, 36, 18, 0, 285, 284, 1, 0, 0, 0, 285, 286, 1, 0, 0,
		0, 286, 287, 1, 0, 0, 0, 287, 289, 5, 46, 0, 0, 288, 273, 1, 0, 0, 0, 288,
		278, 1, 0, 0, 0, 288, 283, 1, 0, 0, 0, 289, 33, 1, 0, 0, 0, 290, 292, 5,
		45, 0, 0, 291, 293, 3, 36, 18, 0, 292, 291, 1, 0, 0, 0, 292, 293, 1, 0,
		0, 0, 293, 294, 1, 0, 0, 0, 294, 295, 5, 46, 0, 0, 295, 35, 1, 0, 0, 0,
		296, 313, 3, 38, 19, 0, 297, 299, 5, 27, 0, 0, 298, 297, 1, 0, 0, 0, 299,
		302, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 303,
		1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 303, 307, 5, 14, 0, 0, 304, 306, 5, 27,
		0, 0, 305, 304, 1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0,
		307, 308, 1, 0, 0, 0, 308, 310, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310,
		312, 3, 38, 19, 0, 311, 300, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 311,
		1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 37, 1, 0, 0, 0, 315, 313, 1, 0,
		0, 0, 316, 329, 3, 52, 26, 0, 317, 329, 3, 32, 16, 0, 318, 329, 3, 34,
		17, 0, 319, 329, 5, 10, 0, 0, 320, 329, 5, 11, 0, 0, 321, 329, 3, 40, 20,
		0, 322, 329, 3, 42, 21, 0, 323, 329, 3, 44, 22, 0, 324, 329, 3, 46, 23,
		0, 325, 329, 3, 48, 24, 0, 326, 329, 3, 50, 25, 0, 327, 329, 5, 26, 0,
		0, 328, 316, 1, 0, 0, 0, 328, 317, 1, 0, 0, 0, 328, 318, 1, 0, 0, 0, 328,
		319, 1, 0, 0, 0, 328, 320, 1, 0, 0, 0, 328, 321, 1, 0, 0, 0, 328, 322,
		1, 0, 0, 0, 328, 323, 1, 0, 0, 0, 328, 324, 1, 0, 0, 0, 328, 325, 1, 0,
		0, 0, 328, 326, 1, 0, 0, 0, 328, 327, 1, 0, 0, 0, 329, 39, 1, 0, 0, 0,
		330, 331, 7, 2, 0, 0, 331, 41, 1, 0, 0, 0, 332, 333, 7, 3, 0, 0, 333, 43,
		1, 0, 0, 0, 334, 335, 7, 4, 0, 0, 335, 45, 1, 0, 0, 0, 336, 339, 5, 9,
		0, 0, 337, 339, 3, 40, 20, 0, 338, 336, 1, 0, 0, 0, 338, 337, 1, 0, 0,
		0, 339, 47, 1, 0, 0, 0, 340, 343, 3, 50, 25, 0, 341, 342, 5, 13, 0, 0,
		342, 344, 3, 50, 25, 0, 343, 341, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345,
		343, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 49, 1, 0, 0, 0, 347, 368, 5,
		15, 0, 0, 348, 365, 3, 38, 19, 0, 349, 351, 5, 27, 0, 0, 350, 349, 1, 0,
		0, 0, 351, 354, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0,
		353, 355, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 355, 359, 5, 14, 0, 0, 356,
		358, 5, 27, 0, 0, 357, 356, 1, 0, 0, 0, 358, 361, 1, 0, 0, 0, 359, 357,
		1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 362, 1, 0, 0, 0, 361, 359, 1, 0,
		0, 0, 362, 364, 3, 38, 19, 0, 363, 352, 1, 0, 0, 0, 364, 367, 1, 0, 0,
		0, 365, 363, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 369, 1, 0, 0, 0, 367,
		365, 1, 0, 0, 0, 368, 348, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 370,
		1, 0, 0, 0, 370, 371, 5, 16, 0, 0, 371, 51, 1, 0, 0, 0, 372, 373, 7, 5,
		0, 0, 373, 53, 1, 0, 0, 0, 45, 57, 65, 77, 85, 92, 98, 104, 115, 120, 125,
		133, 138, 143, 149, 171, 178, 184, 190, 197, 203, 209, 213, 221, 228, 234,
		240, 249, 256, 262, 268, 275, 280, 285, 288, 292, 300, 307, 313, 328, 338,
		345, 352, 359, 365, 368,
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
	ParalParserEOF                            = antlr.TokenEOF
	ParalParserTASK                           = 1
	ParalParserASSIGN                         = 2
	ParalParserSTRING                         = 3
	ParalParserSINGLE_QUOTE_STRING            = 4
	ParalParserFLOAT                          = 5
	ParalParserNUMBER                         = 6
	ParalParserBOOLEAN                        = 7
	ParalParserZERO_ONE                       = 8
	ParalParserDURATION                       = 9
	ParalParserURL                            = 10
	ParalParserPATH                           = 11
	ParalParserCOLON                          = 12
	ParalParserCOLONCOLON                     = 13
	ParalParserCOMMA                          = 14
	ParalParserLBRACK                         = 15
	ParalParserRBRACK                         = 16
	ParalParserLRBRACK                        = 17
	ParalParserRRBRACK                        = 18
	ParalParserBLOCK_START                    = 19
	ParalParserBLOCK_END                      = 20
	ParalParserAT                             = 21
	ParalParserFUNCTION_START                 = 22
	ParalParserPIPELINE_START                 = 23
	ParalParserLOOP_KEY                       = 24
	ParalParserLOOP_VALUE                     = 25
	ParalParserIDENTIFIER                     = 26
	ParalParserNEWLINE                        = 27
	ParalParserWS                             = 28
	ParalParserUNKNOWN_DEFAULT                = 29
	ParalParserPIPELINE_BUF                   = 30
	ParalParserPIPELINE_STASH                 = 31
	ParalParserPIPELINE_IF_CALL_START         = 32
	ParalParserPIPELINE_ELSEIF_CALL_START     = 33
	ParalParserPIPELINE_ELSE_CALL             = 34
	ParalParserPIPELINE_FUNCTION_CALL_START   = 35
	ParalParserPIPELINE_WS                    = 36
	ParalParserUNKNOWN_TEXT                   = 37
	ParalParserBUF_DOUBLE_BACK_ARROW          = 38
	ParalParserBUF_WS                         = 39
	ParalParserSTASH_DOUBLE_BACK_ARROW        = 40
	ParalParserSTASH_WS                       = 41
	ParalParserIF_CONDITION_END               = 42
	ParalParserEXPRESSION_FUNCTION_CALL_START = 43
	ParalParserEXPRESSION_WS                  = 44
	ParalParserNESTED_FUNCTION_START          = 45
	ParalParserFUNCTION_END                   = 46
	ParalParserFUNCTION_WS                    = 47
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
	ParalParserRULE_unknown_command     = 8
	ParalParserRULE_directive           = 9
	ParalParserRULE_buf                 = 10
	ParalParserRULE_stash               = 11
	ParalParserRULE_condition           = 12
	ParalParserRULE_if_condition        = 13
	ParalParserRULE_elseif_condition    = 14
	ParalParserRULE_else_condition      = 15
	ParalParserRULE_function            = 16
	ParalParserRULE_nested_function     = 17
	ParalParserRULE_argument_list       = 18
	ParalParserRULE_expression          = 19
	ParalParserRULE_number_expr         = 20
	ParalParserRULE_string_expr         = 21
	ParalParserRULE_boolean_expr        = 22
	ParalParserRULE_duration_expr       = 23
	ParalParserRULE_matrix_expr         = 24
	ParalParserRULE_list_expr           = 25
	ParalParserRULE_loop_variable       = 26
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&203423746) != 0 {
		{
			p.SetState(54)
			p.Statement()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Variable_assignment()
		}

	case ParalParserTASK, ParalParserAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Task_definition()
		}

	case ParalParserNEWLINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
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
		p.SetState(67)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(ParalParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Expression()
	}
	{
		p.SetState(70)
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
	IDENTIFIER() antlr.TerminalNode
	BLOCK_START() antlr.TerminalNode
	BLOCK_END() antlr.TerminalNode
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

func (s *Task_definitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *Task_definitionContext) BLOCK_START() antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_START, 0)
}

func (s *Task_definitionContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_END, 0)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserAT {
		{
			p.SetState(72)
			p.Task_directive()
		}
		{
			p.SetState(73)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(ParalParserTASK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(82)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(88)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(89)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(95)
			p.Pipeline_block()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(101)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(107)
		p.Match(ParalParserBLOCK_END)
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
		p.SetState(109)
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
	PIPELINE_START() antlr.TerminalNode
	Pipeline_content() IPipeline_contentContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	EOF() antlr.TerminalNode
	BLOCK_END() antlr.TerminalNode

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

func (s *Pipeline_blockContext) PIPELINE_START() antlr.TerminalNode {
	return s.GetToken(ParalParserPIPELINE_START, 0)
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

func (s *Pipeline_blockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNEWLINE)
}

func (s *Pipeline_blockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, i)
}

func (s *Pipeline_blockContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParalParserEOF, 0)
}

func (s *Pipeline_blockContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_END, 0)
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
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(ParalParserPIPELINE_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(112)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Pipeline_content()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(119)
			_la = p.GetTokenStream().LA(1)

			if !((int64((_la - -1)) & ^0x3f) == 0 && ((int64(1)<<(_la - -1))&270532609) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	} else if p.HasError() { // JIM
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
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(122)
				p.Pipeline_item()
			}

		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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
	Buf() IBufContext
	Stash() IStashContext
	Condition() IConditionContext
	Function() IFunctionContext
	Unknown_command() IUnknown_commandContext

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

func (s *Pipeline_itemContext) Buf() IBufContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBufContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBufContext)
}

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

func (s *Pipeline_itemContext) Unknown_command() IUnknown_commandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnknown_commandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnknown_commandContext)
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
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserPIPELINE_BUF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Buf()
		}

	case ParalParserPIPELINE_STASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Stash()
		}

	case ParalParserPIPELINE_IF_CALL_START:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Condition()
		}

	case ParalParserFUNCTION_START, ParalParserPIPELINE_FUNCTION_CALL_START, ParalParserEXPRESSION_FUNCTION_CALL_START:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.Function()
		}

	case ParalParserURL, ParalParserPATH, ParalParserIDENTIFIER, ParalParserUNKNOWN_TEXT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)
			p.Unknown_command()
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

// IUnknown_commandContext is an interface to support dynamic dispatch.
type IUnknown_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllUNKNOWN_TEXT() []antlr.TerminalNode
	UNKNOWN_TEXT(i int) antlr.TerminalNode
	AllPATH() []antlr.TerminalNode
	PATH(i int) antlr.TerminalNode
	AllURL() []antlr.TerminalNode
	URL(i int) antlr.TerminalNode

	// IsUnknown_commandContext differentiates from other interfaces.
	IsUnknown_commandContext()
}

type Unknown_commandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnknown_commandContext() *Unknown_commandContext {
	var p = new(Unknown_commandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_unknown_command
	return p
}

func InitEmptyUnknown_commandContext(p *Unknown_commandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_unknown_command
}

func (*Unknown_commandContext) IsUnknown_commandContext() {}

func NewUnknown_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unknown_commandContext {
	var p = new(Unknown_commandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_unknown_command

	return p
}

func (s *Unknown_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Unknown_commandContext) AllExpression() []IExpressionContext {
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

func (s *Unknown_commandContext) Expression(i int) IExpressionContext {
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

func (s *Unknown_commandContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ParalParserIDENTIFIER)
}

func (s *Unknown_commandContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, i)
}

func (s *Unknown_commandContext) AllUNKNOWN_TEXT() []antlr.TerminalNode {
	return s.GetTokens(ParalParserUNKNOWN_TEXT)
}

func (s *Unknown_commandContext) UNKNOWN_TEXT(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserUNKNOWN_TEXT, i)
}

func (s *Unknown_commandContext) AllPATH() []antlr.TerminalNode {
	return s.GetTokens(ParalParserPATH)
}

func (s *Unknown_commandContext) PATH(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserPATH, i)
}

func (s *Unknown_commandContext) AllURL() []antlr.TerminalNode {
	return s.GetTokens(ParalParserURL)
}

func (s *Unknown_commandContext) URL(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserURL, i)
}

func (s *Unknown_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unknown_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unknown_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterUnknown_command(s)
	}
}

func (s *Unknown_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitUnknown_command(s)
	}
}

func (p *ParalParser) Unknown_command() (localctx IUnknown_commandContext) {
	localctx = NewUnknown_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParalParserRULE_unknown_command)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(135)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137506065408) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(143)
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
				p.SetState(140)
				p.Expression()
			}

		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Expression() IExpressionContext

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

func (s *DirectiveContext) Expression() IExpressionContext {
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
	p.EnterRule(localctx, 18, ParalParserRULE_directive)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(ParalParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(ParalParserIDENTIFIER)
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
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&44014946521080) != 0 {
		{
			p.SetState(148)
			p.Expression()
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

// IBufContext is an interface to support dynamic dispatch.
type IBufContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPELINE_BUF() antlr.TerminalNode
	String_expr() IString_exprContext
	RBRACK() antlr.TerminalNode
	BUF_DOUBLE_BACK_ARROW() antlr.TerminalNode
	Pipeline_content() IPipeline_contentContext

	// IsBufContext differentiates from other interfaces.
	IsBufContext()
}

type BufContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBufContext() *BufContext {
	var p = new(BufContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_buf
	return p
}

func InitEmptyBufContext(p *BufContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_buf
}

func (*BufContext) IsBufContext() {}

func NewBufContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BufContext {
	var p = new(BufContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_buf

	return p
}

func (s *BufContext) GetParser() antlr.Parser { return s.parser }

func (s *BufContext) PIPELINE_BUF() antlr.TerminalNode {
	return s.GetToken(ParalParserPIPELINE_BUF, 0)
}

func (s *BufContext) String_expr() IString_exprContext {
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

func (s *BufContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRBRACK, 0)
}

func (s *BufContext) BUF_DOUBLE_BACK_ARROW() antlr.TerminalNode {
	return s.GetToken(ParalParserBUF_DOUBLE_BACK_ARROW, 0)
}

func (s *BufContext) Pipeline_content() IPipeline_contentContext {
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

func (s *BufContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BufContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BufContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterBuf(s)
	}
}

func (s *BufContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitBuf(s)
	}
}

func (p *ParalParser) Buf() (localctx IBufContext) {
	localctx = NewBufContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParalParserRULE_buf)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(ParalParserPIPELINE_BUF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.String_expr()
	}
	{
		p.SetState(153)
		p.Match(ParalParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(ParalParserBUF_DOUBLE_BACK_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
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

// IStashContext is an interface to support dynamic dispatch.
type IStashContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPELINE_STASH() antlr.TerminalNode
	String_expr() IString_exprContext
	RBRACK() antlr.TerminalNode
	STASH_DOUBLE_BACK_ARROW() antlr.TerminalNode
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

func (s *StashContext) STASH_DOUBLE_BACK_ARROW() antlr.TerminalNode {
	return s.GetToken(ParalParserSTASH_DOUBLE_BACK_ARROW, 0)
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
	p.EnterRule(localctx, 22, ParalParserRULE_stash)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(ParalParserPIPELINE_STASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.String_expr()
	}
	{
		p.SetState(159)
		p.Match(ParalParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(ParalParserSTASH_DOUBLE_BACK_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
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
	p.EnterRule(localctx, 24, ParalParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
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

// IIf_conditionContext is an interface to support dynamic dispatch.
type IIf_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPELINE_IF_CALL_START() antlr.TerminalNode
	Expression() IExpressionContext
	IF_CONDITION_END() antlr.TerminalNode
	BLOCK_START() antlr.TerminalNode
	BLOCK_END() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllPipeline_block() []IPipeline_blockContext
	Pipeline_block(i int) IPipeline_blockContext
	AllElseif_condition() []IElseif_conditionContext
	Elseif_condition(i int) IElseif_conditionContext
	Else_condition() IElse_conditionContext

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

func (s *If_conditionContext) PIPELINE_IF_CALL_START() antlr.TerminalNode {
	return s.GetToken(ParalParserPIPELINE_IF_CALL_START, 0)
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

func (s *If_conditionContext) IF_CONDITION_END() antlr.TerminalNode {
	return s.GetToken(ParalParserIF_CONDITION_END, 0)
}

func (s *If_conditionContext) BLOCK_START() antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_START, 0)
}

func (s *If_conditionContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_END, 0)
}

func (s *If_conditionContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNEWLINE)
}

func (s *If_conditionContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, i)
}

func (s *If_conditionContext) AllPipeline_block() []IPipeline_blockContext {
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

func (s *If_conditionContext) Pipeline_block(i int) IPipeline_blockContext {
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

func (s *If_conditionContext) AllElseif_condition() []IElseif_conditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseif_conditionContext); ok {
			len++
		}
	}

	tst := make([]IElseif_conditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseif_conditionContext); ok {
			tst[i] = t.(IElseif_conditionContext)
			i++
		}
	}

	return tst
}

func (s *If_conditionContext) Elseif_condition(i int) IElseif_conditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseif_conditionContext); ok {
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

	return t.(IElseif_conditionContext)
}

func (s *If_conditionContext) Else_condition() IElse_conditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_conditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_conditionContext)
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
		p.SetState(165)
		p.Match(ParalParserPIPELINE_IF_CALL_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Expression()
	}
	{
		p.SetState(167)
		p.Match(ParalParserIF_CONDITION_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(168)
			p.Match(ParalParserNEWLINE)
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
	}
	{
		p.SetState(174)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(175)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(181)
			p.Pipeline_block()
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(187)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(193)
		p.Match(ParalParserBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(197)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParalParserNEWLINE {
				{
					p.SetState(194)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(199)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(200)
				p.Elseif_condition()
			}

		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(206)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(211)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(212)
			p.Else_condition()
		}

	} else if p.HasError() { // JIM
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

// IElseif_conditionContext is an interface to support dynamic dispatch.
type IElseif_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPELINE_ELSEIF_CALL_START() antlr.TerminalNode
	Expression() IExpressionContext
	IF_CONDITION_END() antlr.TerminalNode
	BLOCK_START() antlr.TerminalNode
	BLOCK_END() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllPipeline_block() []IPipeline_blockContext
	Pipeline_block(i int) IPipeline_blockContext

	// IsElseif_conditionContext differentiates from other interfaces.
	IsElseif_conditionContext()
}

type Elseif_conditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseif_conditionContext() *Elseif_conditionContext {
	var p = new(Elseif_conditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_elseif_condition
	return p
}

func InitEmptyElseif_conditionContext(p *Elseif_conditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_elseif_condition
}

func (*Elseif_conditionContext) IsElseif_conditionContext() {}

func NewElseif_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Elseif_conditionContext {
	var p = new(Elseif_conditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_elseif_condition

	return p
}

func (s *Elseif_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Elseif_conditionContext) PIPELINE_ELSEIF_CALL_START() antlr.TerminalNode {
	return s.GetToken(ParalParserPIPELINE_ELSEIF_CALL_START, 0)
}

func (s *Elseif_conditionContext) Expression() IExpressionContext {
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

func (s *Elseif_conditionContext) IF_CONDITION_END() antlr.TerminalNode {
	return s.GetToken(ParalParserIF_CONDITION_END, 0)
}

func (s *Elseif_conditionContext) BLOCK_START() antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_START, 0)
}

func (s *Elseif_conditionContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_END, 0)
}

func (s *Elseif_conditionContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNEWLINE)
}

func (s *Elseif_conditionContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, i)
}

func (s *Elseif_conditionContext) AllPipeline_block() []IPipeline_blockContext {
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

func (s *Elseif_conditionContext) Pipeline_block(i int) IPipeline_blockContext {
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

func (s *Elseif_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Elseif_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Elseif_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterElseif_condition(s)
	}
}

func (s *Elseif_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitElseif_condition(s)
	}
}

func (p *ParalParser) Elseif_condition() (localctx IElseif_conditionContext) {
	localctx = NewElseif_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ParalParserRULE_elseif_condition)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(ParalParserPIPELINE_ELSEIF_CALL_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Expression()
	}
	{
		p.SetState(217)
		p.Match(ParalParserIF_CONDITION_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(218)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(224)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(225)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(231)
			p.Pipeline_block()
		}

		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(237)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(243)
		p.Match(ParalParserBLOCK_END)
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

// IElse_conditionContext is an interface to support dynamic dispatch.
type IElse_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPELINE_ELSE_CALL() antlr.TerminalNode
	BLOCK_START() antlr.TerminalNode
	BLOCK_END() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllPipeline_block() []IPipeline_blockContext
	Pipeline_block(i int) IPipeline_blockContext

	// IsElse_conditionContext differentiates from other interfaces.
	IsElse_conditionContext()
}

type Else_conditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_conditionContext() *Else_conditionContext {
	var p = new(Else_conditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_else_condition
	return p
}

func InitEmptyElse_conditionContext(p *Else_conditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_else_condition
}

func (*Else_conditionContext) IsElse_conditionContext() {}

func NewElse_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_conditionContext {
	var p = new(Else_conditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_else_condition

	return p
}

func (s *Else_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_conditionContext) PIPELINE_ELSE_CALL() antlr.TerminalNode {
	return s.GetToken(ParalParserPIPELINE_ELSE_CALL, 0)
}

func (s *Else_conditionContext) BLOCK_START() antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_START, 0)
}

func (s *Else_conditionContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_END, 0)
}

func (s *Else_conditionContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNEWLINE)
}

func (s *Else_conditionContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, i)
}

func (s *Else_conditionContext) AllPipeline_block() []IPipeline_blockContext {
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

func (s *Else_conditionContext) Pipeline_block(i int) IPipeline_blockContext {
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

func (s *Else_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterElse_condition(s)
	}
}

func (s *Else_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitElse_condition(s)
	}
}

func (p *ParalParser) Else_condition() (localctx IElse_conditionContext) {
	localctx = NewElse_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ParalParserRULE_else_condition)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(ParalParserPIPELINE_ELSE_CALL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(246)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(252)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(253)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(259)
			p.Pipeline_block()
		}

		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(265)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(271)
		p.Match(ParalParserBLOCK_END)
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

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPELINE_FUNCTION_CALL_START() antlr.TerminalNode
	FUNCTION_END() antlr.TerminalNode
	Argument_list() IArgument_listContext
	FUNCTION_START() antlr.TerminalNode
	EXPRESSION_FUNCTION_CALL_START() antlr.TerminalNode

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

func (s *FunctionContext) PIPELINE_FUNCTION_CALL_START() antlr.TerminalNode {
	return s.GetToken(ParalParserPIPELINE_FUNCTION_CALL_START, 0)
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

func (s *FunctionContext) EXPRESSION_FUNCTION_CALL_START() antlr.TerminalNode {
	return s.GetToken(ParalParserEXPRESSION_FUNCTION_CALL_START, 0)
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
	p.EnterRule(localctx, 32, ParalParserRULE_function)
	var _la int

	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserPIPELINE_FUNCTION_CALL_START:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.Match(ParalParserPIPELINE_FUNCTION_CALL_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&44014946521080) != 0 {
			{
				p.SetState(274)
				p.Argument_list()
			}

		}
		{
			p.SetState(277)
			p.Match(ParalParserFUNCTION_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserFUNCTION_START:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)
			p.Match(ParalParserFUNCTION_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&44014946521080) != 0 {
			{
				p.SetState(279)
				p.Argument_list()
			}

		}
		{
			p.SetState(282)
			p.Match(ParalParserFUNCTION_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserEXPRESSION_FUNCTION_CALL_START:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(283)
			p.Match(ParalParserEXPRESSION_FUNCTION_CALL_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&44014946521080) != 0 {
			{
				p.SetState(284)
				p.Argument_list()
			}

		}
		{
			p.SetState(287)
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
	p.EnterRule(localctx, 34, ParalParserRULE_nested_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(ParalParserNESTED_FUNCTION_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&44014946521080) != 0 {
		{
			p.SetState(291)
			p.Argument_list()
		}

	}
	{
		p.SetState(294)
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
	p.EnterRule(localctx, 36, ParalParserRULE_argument_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Expression()
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserCOMMA || _la == ParalParserNEWLINE {
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(297)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(303)
			p.Match(ParalParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(304)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(309)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(310)
			p.Expression()
		}

		p.SetState(315)
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
	PATH() antlr.TerminalNode
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

func (s *ExpressionContext) PATH() antlr.TerminalNode {
	return s.GetToken(ParalParserPATH, 0)
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
	p.EnterRule(localctx, 38, ParalParserRULE_expression)
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)
			p.Loop_variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(317)
			p.Function()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(318)
			p.Nested_function()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(319)
			p.Match(ParalParserURL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(320)
			p.Match(ParalParserPATH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(321)
			p.Number_expr()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(322)
			p.String_expr()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(323)
			p.Boolean_expr()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(324)
			p.Duration_expr()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(325)
			p.Matrix_expr()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(326)
			p.List_expr()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(327)
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
	p.EnterRule(localctx, 40, ParalParserRULE_number_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
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
	p.EnterRule(localctx, 42, ParalParserRULE_string_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
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
	p.EnterRule(localctx, 44, ParalParserRULE_boolean_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
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
	p.EnterRule(localctx, 46, ParalParserRULE_duration_expr)
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserDURATION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(336)
			p.Match(ParalParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserFLOAT, ParalParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)
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
	p.EnterRule(localctx, 48, ParalParserRULE_matrix_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.List_expr()
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ParalParserCOLONCOLON {
		{
			p.SetState(341)
			p.Match(ParalParserCOLONCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.List_expr()
		}

		p.SetState(345)
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
	p.EnterRule(localctx, 50, ParalParserRULE_list_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(ParalParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&44014946521080) != 0 {
		{
			p.SetState(348)
			p.Expression()
		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserCOMMA || _la == ParalParserNEWLINE {
			p.SetState(352)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParalParserNEWLINE {
				{
					p.SetState(349)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(354)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(355)
				p.Match(ParalParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(359)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParalParserNEWLINE {
				{
					p.SetState(356)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(361)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(362)
				p.Expression()
			}

			p.SetState(367)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(370)
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
	LOOP_KEY() antlr.TerminalNode
	LOOP_VALUE() antlr.TerminalNode

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

func (s *Loop_variableContext) LOOP_KEY() antlr.TerminalNode {
	return s.GetToken(ParalParserLOOP_KEY, 0)
}

func (s *Loop_variableContext) LOOP_VALUE() antlr.TerminalNode {
	return s.GetToken(ParalParserLOOP_VALUE, 0)
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
	p.EnterRule(localctx, 52, ParalParserRULE_loop_variable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ParalParserLOOP_KEY || _la == ParalParserLOOP_VALUE) {
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
