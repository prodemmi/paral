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
		"AT", "FUNCTION_START", "PIPELINE_START", "ERROR", "LOOP_KEY", "LOOP_VALUE",
		"TRY", "CATCH", "IDENTIFIER", "NEWLINE", "WS", "UNKNOWN_DEFAULT", "PIPELINE_BUF",
		"PIPELINE_STASH", "PIPELINE_IF_CALL_START", "PIPELINE_ELSEIF_CALL_START",
		"PIPELINE_ELSE_CALL", "PIPELINE_FUNCTION_CALL_START", "PIPELINE_WS",
		"UNKNOWN_TEXT", "BUF_DOUBLE_BACK_ARROW", "BUF_WS", "STASH_DOUBLE_BACK_ARROW",
		"STASH_WS", "IF_CONDITION_END", "EXPRESSION_FUNCTION_CALL_START", "EXPRESSION_WS",
		"NESTED_FUNCTION_START", "FUNCTION_END", "FUNCTION_WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "variable_assignment", "task_definition", "task_directive",
		"pipeline_block", "pipeline_content", "pipeline_item", "unknown_command",
		"directive", "buf", "stash", "condition", "if_condition", "elseif_condition",
		"else_condition", "try_catch", "try_block", "catch_block", "function",
		"nested_function", "argument_list", "expression", "number_expr", "string_expr",
		"boolean_expr", "duration_expr", "matrix_expr", "list_expr", "loop_variable",
		"error_variable",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 446, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 5,
		0, 64, 8, 0, 10, 0, 12, 0, 67, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1,
		74, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 84, 8,
		3, 10, 3, 12, 3, 87, 9, 3, 1, 3, 1, 3, 1, 3, 5, 3, 92, 8, 3, 10, 3, 12,
		3, 95, 9, 3, 1, 3, 1, 3, 5, 3, 99, 8, 3, 10, 3, 12, 3, 102, 9, 3, 1, 3,
		5, 3, 105, 8, 3, 10, 3, 12, 3, 108, 9, 3, 1, 3, 5, 3, 111, 8, 3, 10, 3,
		12, 3, 114, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 122, 8, 5,
		10, 5, 12, 5, 125, 9, 5, 1, 5, 1, 5, 3, 5, 129, 8, 5, 1, 6, 5, 6, 132,
		8, 6, 10, 6, 12, 6, 135, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		143, 8, 7, 1, 8, 4, 8, 146, 8, 8, 11, 8, 12, 8, 147, 1, 8, 5, 8, 151, 8,
		8, 10, 8, 12, 8, 154, 9, 8, 1, 9, 1, 9, 1, 9, 3, 9, 159, 8, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 179, 8, 13, 10, 13, 12,
		13, 182, 9, 13, 1, 13, 1, 13, 5, 13, 186, 8, 13, 10, 13, 12, 13, 189, 9,
		13, 1, 13, 5, 13, 192, 8, 13, 10, 13, 12, 13, 195, 9, 13, 1, 13, 5, 13,
		198, 8, 13, 10, 13, 12, 13, 201, 9, 13, 1, 13, 1, 13, 5, 13, 205, 8, 13,
		10, 13, 12, 13, 208, 9, 13, 1, 13, 5, 13, 211, 8, 13, 10, 13, 12, 13, 214,
		9, 13, 1, 13, 5, 13, 217, 8, 13, 10, 13, 12, 13, 220, 9, 13, 1, 13, 3,
		13, 223, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 229, 8, 14, 10, 14,
		12, 14, 232, 9, 14, 1, 14, 1, 14, 5, 14, 236, 8, 14, 10, 14, 12, 14, 239,
		9, 14, 1, 14, 5, 14, 242, 8, 14, 10, 14, 12, 14, 245, 9, 14, 1, 14, 5,
		14, 248, 8, 14, 10, 14, 12, 14, 251, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		5, 15, 257, 8, 15, 10, 15, 12, 15, 260, 9, 15, 1, 15, 1, 15, 5, 15, 264,
		8, 15, 10, 15, 12, 15, 267, 9, 15, 1, 15, 5, 15, 270, 8, 15, 10, 15, 12,
		15, 273, 9, 15, 1, 15, 5, 15, 276, 8, 15, 10, 15, 12, 15, 279, 9, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 5, 16, 285, 8, 16, 10, 16, 12, 16, 288, 9, 16,
		1, 16, 1, 16, 5, 16, 292, 8, 16, 10, 16, 12, 16, 295, 9, 16, 1, 16, 1,
		16, 5, 16, 299, 8, 16, 10, 16, 12, 16, 302, 9, 16, 1, 16, 1, 16, 1, 16,
		5, 16, 307, 8, 16, 10, 16, 12, 16, 310, 9, 16, 1, 16, 1, 16, 5, 16, 314,
		8, 16, 10, 16, 12, 16, 317, 9, 16, 1, 16, 1, 16, 5, 16, 321, 8, 16, 10,
		16, 12, 16, 324, 9, 16, 1, 16, 1, 16, 3, 16, 328, 8, 16, 1, 17, 5, 17,
		331, 8, 17, 10, 17, 12, 17, 334, 9, 17, 1, 18, 5, 18, 337, 8, 18, 10, 18,
		12, 18, 340, 9, 18, 1, 19, 1, 19, 3, 19, 344, 8, 19, 1, 19, 1, 19, 1, 19,
		3, 19, 349, 8, 19, 1, 19, 1, 19, 1, 19, 3, 19, 354, 8, 19, 1, 19, 3, 19,
		357, 8, 19, 1, 20, 1, 20, 3, 20, 361, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		5, 21, 367, 8, 21, 10, 21, 12, 21, 370, 9, 21, 1, 21, 1, 21, 5, 21, 374,
		8, 21, 10, 21, 12, 21, 377, 9, 21, 1, 21, 5, 21, 380, 8, 21, 10, 21, 12,
		21, 383, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 398, 8, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 3, 26, 408, 8, 26, 1, 27, 1, 27,
		1, 27, 4, 27, 413, 8, 27, 11, 27, 12, 27, 414, 1, 28, 1, 28, 1, 28, 5,
		28, 420, 8, 28, 10, 28, 12, 28, 423, 9, 28, 1, 28, 1, 28, 5, 28, 427, 8,
		28, 10, 28, 12, 28, 430, 9, 28, 1, 28, 5, 28, 433, 8, 28, 10, 28, 12, 28,
		436, 9, 28, 3, 28, 438, 8, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30,
		1, 30, 0, 0, 31, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 0, 6, 2,
		1, 20, 20, 30, 30, 3, 0, 10, 11, 29, 29, 40, 40, 1, 0, 5, 6, 1, 0, 3, 4,
		1, 0, 7, 8, 1, 0, 25, 26, 485, 0, 65, 1, 0, 0, 0, 2, 73, 1, 0, 0, 0, 4,
		75, 1, 0, 0, 0, 6, 85, 1, 0, 0, 0, 8, 117, 1, 0, 0, 0, 10, 119, 1, 0, 0,
		0, 12, 133, 1, 0, 0, 0, 14, 142, 1, 0, 0, 0, 16, 145, 1, 0, 0, 0, 18, 155,
		1, 0, 0, 0, 20, 160, 1, 0, 0, 0, 22, 166, 1, 0, 0, 0, 24, 172, 1, 0, 0,
		0, 26, 174, 1, 0, 0, 0, 28, 224, 1, 0, 0, 0, 30, 254, 1, 0, 0, 0, 32, 282,
		1, 0, 0, 0, 34, 332, 1, 0, 0, 0, 36, 338, 1, 0, 0, 0, 38, 356, 1, 0, 0,
		0, 40, 358, 1, 0, 0, 0, 42, 364, 1, 0, 0, 0, 44, 397, 1, 0, 0, 0, 46, 399,
		1, 0, 0, 0, 48, 401, 1, 0, 0, 0, 50, 403, 1, 0, 0, 0, 52, 407, 1, 0, 0,
		0, 54, 409, 1, 0, 0, 0, 56, 416, 1, 0, 0, 0, 58, 441, 1, 0, 0, 0, 60, 443,
		1, 0, 0, 0, 62, 64, 3, 2, 1, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0,
		65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 65, 1,
		0, 0, 0, 68, 69, 5, 0, 0, 1, 69, 1, 1, 0, 0, 0, 70, 74, 3, 4, 2, 0, 71,
		74, 3, 6, 3, 0, 72, 74, 5, 30, 0, 0, 73, 70, 1, 0, 0, 0, 73, 71, 1, 0,
		0, 0, 73, 72, 1, 0, 0, 0, 74, 3, 1, 0, 0, 0, 75, 76, 5, 29, 0, 0, 76, 77,
		5, 2, 0, 0, 77, 78, 3, 44, 22, 0, 78, 79, 5, 30, 0, 0, 79, 5, 1, 0, 0,
		0, 80, 81, 3, 8, 4, 0, 81, 82, 5, 30, 0, 0, 82, 84, 1, 0, 0, 0, 83, 80,
		1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0,
		86, 88, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 89, 5, 1, 0, 0, 89, 93, 5,
		29, 0, 0, 90, 92, 5, 30, 0, 0, 91, 90, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0,
		93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 96, 1, 0, 0, 0, 95, 93, 1,
		0, 0, 0, 96, 100, 5, 19, 0, 0, 97, 99, 5, 30, 0, 0, 98, 97, 1, 0, 0, 0,
		99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 106,
		1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 105, 3, 10, 5, 0, 104, 103, 1, 0,
		0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0,
		107, 112, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 111, 5, 30, 0, 0, 110,
		109, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113,
		1, 0, 0, 0, 113, 115, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 5, 20,
		0, 0, 116, 7, 1, 0, 0, 0, 117, 118, 3, 18, 9, 0, 118, 9, 1, 0, 0, 0, 119,
		123, 5, 23, 0, 0, 120, 122, 5, 30, 0, 0, 121, 120, 1, 0, 0, 0, 122, 125,
		1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0,
		0, 0, 125, 123, 1, 0, 0, 0, 126, 128, 3, 12, 6, 0, 127, 129, 7, 0, 0, 0,
		128, 127, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 11, 1, 0, 0, 0, 130, 132,
		3, 14, 7, 0, 131, 130, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0,
		0, 0, 133, 134, 1, 0, 0, 0, 134, 13, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0,
		136, 143, 3, 20, 10, 0, 137, 143, 3, 22, 11, 0, 138, 143, 3, 24, 12, 0,
		139, 143, 3, 32, 16, 0, 140, 143, 3, 38, 19, 0, 141, 143, 3, 16, 8, 0,
		142, 136, 1, 0, 0, 0, 142, 137, 1, 0, 0, 0, 142, 138, 1, 0, 0, 0, 142,
		139, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 141, 1, 0, 0, 0, 143, 15, 1,
		0, 0, 0, 144, 146, 7, 1, 0, 0, 145, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0,
		0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 152, 1, 0, 0, 0, 149,
		151, 3, 44, 22, 0, 150, 149, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150,
		1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 17, 1, 0, 0, 0, 154, 152, 1, 0,
		0, 0, 155, 156, 5, 21, 0, 0, 156, 158, 5, 29, 0, 0, 157, 159, 3, 44, 22,
		0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 19, 1, 0, 0, 0, 160,
		161, 5, 33, 0, 0, 161, 162, 3, 48, 24, 0, 162, 163, 5, 16, 0, 0, 163, 164,
		5, 41, 0, 0, 164, 165, 3, 12, 6, 0, 165, 21, 1, 0, 0, 0, 166, 167, 5, 34,
		0, 0, 167, 168, 3, 48, 24, 0, 168, 169, 5, 16, 0, 0, 169, 170, 5, 43, 0,
		0, 170, 171, 3, 12, 6, 0, 171, 23, 1, 0, 0, 0, 172, 173, 3, 26, 13, 0,
		173, 25, 1, 0, 0, 0, 174, 175, 5, 35, 0, 0, 175, 176, 3, 44, 22, 0, 176,
		180, 5, 45, 0, 0, 177, 179, 5, 30, 0, 0, 178, 177, 1, 0, 0, 0, 179, 182,
		1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 183, 1, 0,
		0, 0, 182, 180, 1, 0, 0, 0, 183, 187, 5, 19, 0, 0, 184, 186, 5, 30, 0,
		0, 185, 184, 1, 0, 0, 0, 186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187,
		188, 1, 0, 0, 0, 188, 193, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 192,
		3, 10, 5, 0, 191, 190, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0,
		0, 0, 193, 194, 1, 0, 0, 0, 194, 199, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0,
		196, 198, 5, 30, 0, 0, 197, 196, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199,
		197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 202, 1, 0, 0, 0, 201, 199,
		1, 0, 0, 0, 202, 212, 5, 20, 0, 0, 203, 205, 5, 30, 0, 0, 204, 203, 1,
		0, 0, 0, 205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0,
		0, 207, 209, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 211, 3, 28, 14, 0,
		210, 206, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212,
		213, 1, 0, 0, 0, 213, 222, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215, 217,
		5, 30, 0, 0, 216, 215, 1, 0, 0, 0, 217, 220, 1, 0, 0, 0, 218, 216, 1, 0,
		0, 0, 218, 219, 1, 0, 0, 0, 219, 221, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0,
		221, 223, 3, 30, 15, 0, 222, 218, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223,
		27, 1, 0, 0, 0, 224, 225, 5, 36, 0, 0, 225, 226, 3, 44, 22, 0, 226, 230,
		5, 45, 0, 0, 227, 229, 5, 30, 0, 0, 228, 227, 1, 0, 0, 0, 229, 232, 1,
		0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 233, 1, 0, 0,
		0, 232, 230, 1, 0, 0, 0, 233, 237, 5, 19, 0, 0, 234, 236, 5, 30, 0, 0,
		235, 234, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237,
		238, 1, 0, 0, 0, 238, 243, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 240, 242,
		3, 10, 5, 0, 241, 240, 1, 0, 0, 0, 242, 245, 1, 0, 0, 0, 243, 241, 1, 0,
		0, 0, 243, 244, 1, 0, 0, 0, 244, 249, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0,
		246, 248, 5, 30, 0, 0, 247, 246, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249,
		247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 252, 1, 0, 0, 0, 251, 249,
		1, 0, 0, 0, 252, 253, 5, 20, 0, 0, 253, 29, 1, 0, 0, 0, 254, 258, 5, 37,
		0, 0, 255, 257, 5, 30, 0, 0, 256, 255, 1, 0, 0, 0, 257, 260, 1, 0, 0, 0,
		258, 256, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 261, 1, 0, 0, 0, 260,
		258, 1, 0, 0, 0, 261, 265, 5, 19, 0, 0, 262, 264, 5, 30, 0, 0, 263, 262,
		1, 0, 0, 0, 264, 267, 1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 265, 266, 1, 0,
		0, 0, 266, 271, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 268, 270, 3, 10, 5, 0,
		269, 268, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271,
		272, 1, 0, 0, 0, 272, 277, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 276,
		5, 30, 0, 0, 275, 274, 1, 0, 0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0,
		0, 0, 277, 278, 1, 0, 0, 0, 278, 280, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0,
		280, 281, 5, 20, 0, 0, 281, 31, 1, 0, 0, 0, 282, 286, 5, 27, 0, 0, 283,
		285, 5, 30, 0, 0, 284, 283, 1, 0, 0, 0, 285, 288, 1, 0, 0, 0, 286, 284,
		1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 289, 1, 0, 0, 0, 288, 286, 1, 0,
		0, 0, 289, 293, 5, 19, 0, 0, 290, 292, 5, 30, 0, 0, 291, 290, 1, 0, 0,
		0, 292, 295, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294,
		296, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 300, 3, 34, 17, 0, 297, 299,
		5, 30, 0, 0, 298, 297, 1, 0, 0, 0, 299, 302, 1, 0, 0, 0, 300, 298, 1, 0,
		0, 0, 300, 301, 1, 0, 0, 0, 301, 303, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0,
		303, 327, 5, 20, 0, 0, 304, 308, 5, 28, 0, 0, 305, 307, 5, 30, 0, 0, 306,
		305, 1, 0, 0, 0, 307, 310, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 309,
		1, 0, 0, 0, 309, 311, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 311, 315, 5, 19,
		0, 0, 312, 314, 5, 30, 0, 0, 313, 312, 1, 0, 0, 0, 314, 317, 1, 0, 0, 0,
		315, 313, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 318, 1, 0, 0, 0, 317,
		315, 1, 0, 0, 0, 318, 322, 3, 36, 18, 0, 319, 321, 5, 30, 0, 0, 320, 319,
		1, 0, 0, 0, 321, 324, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 322, 323, 1, 0,
		0, 0, 323, 325, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 325, 326, 5, 20, 0, 0,
		326, 328, 1, 0, 0, 0, 327, 304, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328,
		33, 1, 0, 0, 0, 329, 331, 3, 10, 5, 0, 330, 329, 1, 0, 0, 0, 331, 334,
		1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 35, 1, 0,
		0, 0, 334, 332, 1, 0, 0, 0, 335, 337, 3, 10, 5, 0, 336, 335, 1, 0, 0, 0,
		337, 340, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339,
		37, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 341, 343, 5, 38, 0, 0, 342, 344,
		3, 42, 21, 0, 343, 342, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 345, 1,
		0, 0, 0, 345, 357, 5, 49, 0, 0, 346, 348, 5, 22, 0, 0, 347, 349, 3, 42,
		21, 0, 348, 347, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0,
		350, 357, 5, 49, 0, 0, 351, 353, 5, 46, 0, 0, 352, 354, 3, 42, 21, 0, 353,
		352, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 357,
		5, 49, 0, 0, 356, 341, 1, 0, 0, 0, 356, 346, 1, 0, 0, 0, 356, 351, 1, 0,
		0, 0, 357, 39, 1, 0, 0, 0, 358, 360, 5, 48, 0, 0, 359, 361, 3, 42, 21,
		0, 360, 359, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362,
		363, 5, 49, 0, 0, 363, 41, 1, 0, 0, 0, 364, 381, 3, 44, 22, 0, 365, 367,
		5, 30, 0, 0, 366, 365, 1, 0, 0, 0, 367, 370, 1, 0, 0, 0, 368, 366, 1, 0,
		0, 0, 368, 369, 1, 0, 0, 0, 369, 371, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0,
		371, 375, 5, 14, 0, 0, 372, 374, 5, 30, 0, 0, 373, 372, 1, 0, 0, 0, 374,
		377, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 378,
		1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 380, 3, 44, 22, 0, 379, 368, 1,
		0, 0, 0, 380, 383, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 381, 382, 1, 0, 0,
		0, 382, 43, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 384, 398, 3, 58, 29, 0, 385,
		398, 3, 60, 30, 0, 386, 398, 3, 38, 19, 0, 387, 398, 3, 40, 20, 0, 388,
		398, 5, 10, 0, 0, 389, 398, 5, 11, 0, 0, 390, 398, 3, 46, 23, 0, 391, 398,
		3, 48, 24, 0, 392, 398, 3, 50, 25, 0, 393, 398, 3, 52, 26, 0, 394, 398,
		3, 54, 27, 0, 395, 398, 3, 56, 28, 0, 396, 398, 5, 29, 0, 0, 397, 384,
		1, 0, 0, 0, 397, 385, 1, 0, 0, 0, 397, 386, 1, 0, 0, 0, 397, 387, 1, 0,
		0, 0, 397, 388, 1, 0, 0, 0, 397, 389, 1, 0, 0, 0, 397, 390, 1, 0, 0, 0,
		397, 391, 1, 0, 0, 0, 397, 392, 1, 0, 0, 0, 397, 393, 1, 0, 0, 0, 397,
		394, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 397, 396, 1, 0, 0, 0, 398, 45, 1,
		0, 0, 0, 399, 400, 7, 2, 0, 0, 400, 47, 1, 0, 0, 0, 401, 402, 7, 3, 0,
		0, 402, 49, 1, 0, 0, 0, 403, 404, 7, 4, 0, 0, 404, 51, 1, 0, 0, 0, 405,
		408, 5, 9, 0, 0, 406, 408, 3, 46, 23, 0, 407, 405, 1, 0, 0, 0, 407, 406,
		1, 0, 0, 0, 408, 53, 1, 0, 0, 0, 409, 412, 3, 56, 28, 0, 410, 411, 5, 13,
		0, 0, 411, 413, 3, 56, 28, 0, 412, 410, 1, 0, 0, 0, 413, 414, 1, 0, 0,
		0, 414, 412, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 55, 1, 0, 0, 0, 416,
		437, 5, 15, 0, 0, 417, 434, 3, 44, 22, 0, 418, 420, 5, 30, 0, 0, 419, 418,
		1, 0, 0, 0, 420, 423, 1, 0, 0, 0, 421, 419, 1, 0, 0, 0, 421, 422, 1, 0,
		0, 0, 422, 424, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0, 424, 428, 5, 14, 0, 0,
		425, 427, 5, 30, 0, 0, 426, 425, 1, 0, 0, 0, 427, 430, 1, 0, 0, 0, 428,
		426, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 431, 1, 0, 0, 0, 430, 428,
		1, 0, 0, 0, 431, 433, 3, 44, 22, 0, 432, 421, 1, 0, 0, 0, 433, 436, 1,
		0, 0, 0, 434, 432, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 438, 1, 0, 0,
		0, 436, 434, 1, 0, 0, 0, 437, 417, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438,
		439, 1, 0, 0, 0, 439, 440, 5, 16, 0, 0, 440, 57, 1, 0, 0, 0, 441, 442,
		7, 5, 0, 0, 442, 59, 1, 0, 0, 0, 443, 444, 5, 24, 0, 0, 444, 61, 1, 0,
		0, 0, 54, 65, 73, 85, 93, 100, 106, 112, 123, 128, 133, 142, 147, 152,
		158, 180, 187, 193, 199, 206, 212, 218, 222, 230, 237, 243, 249, 258, 265,
		271, 277, 286, 293, 300, 308, 315, 322, 327, 332, 338, 343, 348, 353, 356,
		360, 368, 375, 381, 397, 407, 414, 421, 428, 434, 437,
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
	ParalParserERROR                          = 24
	ParalParserLOOP_KEY                       = 25
	ParalParserLOOP_VALUE                     = 26
	ParalParserTRY                            = 27
	ParalParserCATCH                          = 28
	ParalParserIDENTIFIER                     = 29
	ParalParserNEWLINE                        = 30
	ParalParserWS                             = 31
	ParalParserUNKNOWN_DEFAULT                = 32
	ParalParserPIPELINE_BUF                   = 33
	ParalParserPIPELINE_STASH                 = 34
	ParalParserPIPELINE_IF_CALL_START         = 35
	ParalParserPIPELINE_ELSEIF_CALL_START     = 36
	ParalParserPIPELINE_ELSE_CALL             = 37
	ParalParserPIPELINE_FUNCTION_CALL_START   = 38
	ParalParserPIPELINE_WS                    = 39
	ParalParserUNKNOWN_TEXT                   = 40
	ParalParserBUF_DOUBLE_BACK_ARROW          = 41
	ParalParserBUF_WS                         = 42
	ParalParserSTASH_DOUBLE_BACK_ARROW        = 43
	ParalParserSTASH_WS                       = 44
	ParalParserIF_CONDITION_END               = 45
	ParalParserEXPRESSION_FUNCTION_CALL_START = 46
	ParalParserEXPRESSION_WS                  = 47
	ParalParserNESTED_FUNCTION_START          = 48
	ParalParserFUNCTION_END                   = 49
	ParalParserFUNCTION_WS                    = 50
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
	ParalParserRULE_try_catch           = 16
	ParalParserRULE_try_block           = 17
	ParalParserRULE_catch_block         = 18
	ParalParserRULE_function            = 19
	ParalParserRULE_nested_function     = 20
	ParalParserRULE_argument_list       = 21
	ParalParserRULE_expression          = 22
	ParalParserRULE_number_expr         = 23
	ParalParserRULE_string_expr         = 24
	ParalParserRULE_boolean_expr        = 25
	ParalParserRULE_duration_expr       = 26
	ParalParserRULE_matrix_expr         = 27
	ParalParserRULE_list_expr           = 28
	ParalParserRULE_loop_variable       = 29
	ParalParserRULE_error_variable      = 30
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1612709890) != 0 {
		{
			p.SetState(62)
			p.Statement()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
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
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Variable_assignment()
		}

	case ParalParserTASK, ParalParserAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Task_definition()
		}

	case ParalParserNEWLINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
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
		p.SetState(75)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(ParalParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Expression()
	}
	{
		p.SetState(78)
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
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserAT {
		{
			p.SetState(80)
			p.Task_directive()
		}
		{
			p.SetState(81)
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
		p.Match(ParalParserTASK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(90)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(100)
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
				p.SetState(97)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(103)
			p.Pipeline_block()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(109)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
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
		p.SetState(117)
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
		p.SetState(119)
		p.Match(ParalParserPIPELINE_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(123)
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
				p.SetState(120)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(125)
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
		p.SetState(126)
		p.Pipeline_content()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(127)
			_la = p.GetTokenStream().LA(1)

			if !((int64((_la - -1)) & ^0x3f) == 0 && ((int64(1)<<(_la - -1))&2149580801) != 0) {
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
	p.SetState(133)
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
				p.SetState(130)
				p.Pipeline_item()
			}

		}
		p.SetState(135)
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
	Try_catch() ITry_catchContext
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

func (s *Pipeline_itemContext) Try_catch() ITry_catchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITry_catchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITry_catchContext)
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
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserPIPELINE_BUF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Buf()
		}

	case ParalParserPIPELINE_STASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Stash()
		}

	case ParalParserPIPELINE_IF_CALL_START:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(138)
			p.Condition()
		}

	case ParalParserTRY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(139)
			p.Try_catch()
		}

	case ParalParserFUNCTION_START, ParalParserPIPELINE_FUNCTION_CALL_START, ParalParserEXPRESSION_FUNCTION_CALL_START:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(140)
			p.Function()
		}

	case ParalParserURL, ParalParserPATH, ParalParserIDENTIFIER, ParalParserUNKNOWN_TEXT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(141)
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
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(144)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1100048501760) != 0) {
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

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(152)
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
				p.SetState(149)
				p.Expression()
			}

		}
		p.SetState(154)
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
		p.SetState(155)
		p.Match(ParalParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(ParalParserIDENTIFIER)
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&352119257337848) != 0 {
		{
			p.SetState(157)
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
		p.SetState(160)
		p.Match(ParalParserPIPELINE_BUF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.String_expr()
	}
	{
		p.SetState(162)
		p.Match(ParalParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(ParalParserBUF_DOUBLE_BACK_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
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
		p.SetState(166)
		p.Match(ParalParserPIPELINE_STASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.String_expr()
	}
	{
		p.SetState(168)
		p.Match(ParalParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(ParalParserSTASH_DOUBLE_BACK_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
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
		p.SetState(172)
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
		p.SetState(174)
		p.Match(ParalParserPIPELINE_IF_CALL_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Expression()
	}
	{
		p.SetState(176)
		p.Match(ParalParserIF_CONDITION_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(177)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(183)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
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
				p.SetState(184)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(190)
			p.Pipeline_block()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(196)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(202)
		p.Match(ParalParserBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(212)
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
			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParalParserNEWLINE {
				{
					p.SetState(203)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(208)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(209)
				p.Elseif_condition()
			}

		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(215)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(221)
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
		p.SetState(224)
		p.Match(ParalParserPIPELINE_ELSEIF_CALL_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Expression()
	}
	{
		p.SetState(226)
		p.Match(ParalParserIF_CONDITION_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(227)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(233)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(237)
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
				p.SetState(234)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(240)
			p.Pipeline_block()
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
		p.SetState(254)
		p.Match(ParalParserPIPELINE_ELSE_CALL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(255)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(261)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(265)
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
				p.SetState(262)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(268)
			p.Pipeline_block()
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(274)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(280)
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

// ITry_catchContext is an interface to support dynamic dispatch.
type ITry_catchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRY() antlr.TerminalNode
	AllBLOCK_START() []antlr.TerminalNode
	BLOCK_START(i int) antlr.TerminalNode
	Try_block() ITry_blockContext
	AllBLOCK_END() []antlr.TerminalNode
	BLOCK_END(i int) antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	CATCH() antlr.TerminalNode
	Catch_block() ICatch_blockContext

	// IsTry_catchContext differentiates from other interfaces.
	IsTry_catchContext()
}

type Try_catchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTry_catchContext() *Try_catchContext {
	var p = new(Try_catchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_try_catch
	return p
}

func InitEmptyTry_catchContext(p *Try_catchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_try_catch
}

func (*Try_catchContext) IsTry_catchContext() {}

func NewTry_catchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Try_catchContext {
	var p = new(Try_catchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_try_catch

	return p
}

func (s *Try_catchContext) GetParser() antlr.Parser { return s.parser }

func (s *Try_catchContext) TRY() antlr.TerminalNode {
	return s.GetToken(ParalParserTRY, 0)
}

func (s *Try_catchContext) AllBLOCK_START() []antlr.TerminalNode {
	return s.GetTokens(ParalParserBLOCK_START)
}

func (s *Try_catchContext) BLOCK_START(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_START, i)
}

func (s *Try_catchContext) Try_block() ITry_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITry_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITry_blockContext)
}

func (s *Try_catchContext) AllBLOCK_END() []antlr.TerminalNode {
	return s.GetTokens(ParalParserBLOCK_END)
}

func (s *Try_catchContext) BLOCK_END(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_END, i)
}

func (s *Try_catchContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNEWLINE)
}

func (s *Try_catchContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, i)
}

func (s *Try_catchContext) CATCH() antlr.TerminalNode {
	return s.GetToken(ParalParserCATCH, 0)
}

func (s *Try_catchContext) Catch_block() ICatch_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatch_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatch_blockContext)
}

func (s *Try_catchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Try_catchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Try_catchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterTry_catch(s)
	}
}

func (s *Try_catchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitTry_catch(s)
	}
}

func (p *ParalParser) Try_catch() (localctx ITry_catchContext) {
	localctx = NewTry_catchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ParalParserRULE_try_catch)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(ParalParserTRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(283)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(289)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(290)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Try_block()
	}
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
		p.Match(ParalParserBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ParalParserCATCH {
		{
			p.SetState(304)
			p.Match(ParalParserCATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(305)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(310)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(311)
			p.Match(ParalParserBLOCK_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(312)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(317)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.Catch_block()
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(319)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(324)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(325)
			p.Match(ParalParserBLOCK_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// ITry_blockContext is an interface to support dynamic dispatch.
type ITry_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPipeline_block() []IPipeline_blockContext
	Pipeline_block(i int) IPipeline_blockContext

	// IsTry_blockContext differentiates from other interfaces.
	IsTry_blockContext()
}

type Try_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTry_blockContext() *Try_blockContext {
	var p = new(Try_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_try_block
	return p
}

func InitEmptyTry_blockContext(p *Try_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_try_block
}

func (*Try_blockContext) IsTry_blockContext() {}

func NewTry_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Try_blockContext {
	var p = new(Try_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_try_block

	return p
}

func (s *Try_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Try_blockContext) AllPipeline_block() []IPipeline_blockContext {
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

func (s *Try_blockContext) Pipeline_block(i int) IPipeline_blockContext {
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

func (s *Try_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Try_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Try_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterTry_block(s)
	}
}

func (s *Try_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitTry_block(s)
	}
}

func (p *ParalParser) Try_block() (localctx ITry_blockContext) {
	localctx = NewTry_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ParalParserRULE_try_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(329)
			p.Pipeline_block()
		}

		p.SetState(334)
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

// ICatch_blockContext is an interface to support dynamic dispatch.
type ICatch_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPipeline_block() []IPipeline_blockContext
	Pipeline_block(i int) IPipeline_blockContext

	// IsCatch_blockContext differentiates from other interfaces.
	IsCatch_blockContext()
}

type Catch_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatch_blockContext() *Catch_blockContext {
	var p = new(Catch_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_catch_block
	return p
}

func InitEmptyCatch_blockContext(p *Catch_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_catch_block
}

func (*Catch_blockContext) IsCatch_blockContext() {}

func NewCatch_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Catch_blockContext {
	var p = new(Catch_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_catch_block

	return p
}

func (s *Catch_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Catch_blockContext) AllPipeline_block() []IPipeline_blockContext {
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

func (s *Catch_blockContext) Pipeline_block(i int) IPipeline_blockContext {
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

func (s *Catch_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Catch_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Catch_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterCatch_block(s)
	}
}

func (s *Catch_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitCatch_block(s)
	}
}

func (p *ParalParser) Catch_block() (localctx ICatch_blockContext) {
	localctx = NewCatch_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ParalParserRULE_catch_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(335)
			p.Pipeline_block()
		}

		p.SetState(340)
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
	p.EnterRule(localctx, 38, ParalParserRULE_function)
	var _la int

	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserPIPELINE_FUNCTION_CALL_START:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)
			p.Match(ParalParserPIPELINE_FUNCTION_CALL_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&352119257337848) != 0 {
			{
				p.SetState(342)
				p.Argument_list()
			}

		}
		{
			p.SetState(345)
			p.Match(ParalParserFUNCTION_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserFUNCTION_START:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.Match(ParalParserFUNCTION_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&352119257337848) != 0 {
			{
				p.SetState(347)
				p.Argument_list()
			}

		}
		{
			p.SetState(350)
			p.Match(ParalParserFUNCTION_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserEXPRESSION_FUNCTION_CALL_START:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(351)
			p.Match(ParalParserEXPRESSION_FUNCTION_CALL_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&352119257337848) != 0 {
			{
				p.SetState(352)
				p.Argument_list()
			}

		}
		{
			p.SetState(355)
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
	p.EnterRule(localctx, 40, ParalParserRULE_nested_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.Match(ParalParserNESTED_FUNCTION_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&352119257337848) != 0 {
		{
			p.SetState(359)
			p.Argument_list()
		}

	}
	{
		p.SetState(362)
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
	p.EnterRule(localctx, 42, ParalParserRULE_argument_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Expression()
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserCOMMA || _la == ParalParserNEWLINE {
		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(365)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(370)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(371)
			p.Match(ParalParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(372)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(377)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(378)
			p.Expression()
		}

		p.SetState(383)
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
	Error_variable() IError_variableContext
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

func (s *ExpressionContext) Error_variable() IError_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IError_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IError_variableContext)
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
	p.EnterRule(localctx, 44, ParalParserRULE_expression)
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(384)
			p.Loop_variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(385)
			p.Error_variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(386)
			p.Function()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(387)
			p.Nested_function()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(388)
			p.Match(ParalParserURL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(389)
			p.Match(ParalParserPATH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(390)
			p.Number_expr()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(391)
			p.String_expr()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(392)
			p.Boolean_expr()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(393)
			p.Duration_expr()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(394)
			p.Matrix_expr()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(395)
			p.List_expr()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(396)
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
	p.EnterRule(localctx, 46, ParalParserRULE_number_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
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
	p.EnterRule(localctx, 48, ParalParserRULE_string_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)
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
	p.EnterRule(localctx, 50, ParalParserRULE_boolean_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
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
	p.EnterRule(localctx, 52, ParalParserRULE_duration_expr)
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserDURATION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.Match(ParalParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserFLOAT, ParalParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)
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
	p.EnterRule(localctx, 54, ParalParserRULE_matrix_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.List_expr()
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ParalParserCOLONCOLON {
		{
			p.SetState(410)
			p.Match(ParalParserCOLONCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.List_expr()
		}

		p.SetState(414)
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
	p.EnterRule(localctx, 56, ParalParserRULE_list_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(416)
		p.Match(ParalParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&352119257337848) != 0 {
		{
			p.SetState(417)
			p.Expression()
		}
		p.SetState(434)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserCOMMA || _la == ParalParserNEWLINE {
			p.SetState(421)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParalParserNEWLINE {
				{
					p.SetState(418)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(423)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(424)
				p.Match(ParalParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(428)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParalParserNEWLINE {
				{
					p.SetState(425)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(430)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(431)
				p.Expression()
			}

			p.SetState(436)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(439)
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
	p.EnterRule(localctx, 58, ParalParserRULE_loop_variable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
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

// IError_variableContext is an interface to support dynamic dispatch.
type IError_variableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ERROR() antlr.TerminalNode

	// IsError_variableContext differentiates from other interfaces.
	IsError_variableContext()
}

type Error_variableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyError_variableContext() *Error_variableContext {
	var p = new(Error_variableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_error_variable
	return p
}

func InitEmptyError_variableContext(p *Error_variableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_error_variable
}

func (*Error_variableContext) IsError_variableContext() {}

func NewError_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Error_variableContext {
	var p = new(Error_variableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_error_variable

	return p
}

func (s *Error_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Error_variableContext) ERROR() antlr.TerminalNode {
	return s.GetToken(ParalParserERROR, 0)
}

func (s *Error_variableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Error_variableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Error_variableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterError_variable(s)
	}
}

func (s *Error_variableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitError_variable(s)
	}
}

func (p *ParalParser) Error_variable() (localctx IError_variableContext) {
	localctx = NewError_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ParalParserRULE_error_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(443)
		p.Match(ParalParserERROR)
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
