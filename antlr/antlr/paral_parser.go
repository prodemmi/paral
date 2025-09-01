// Code generated from antlr/ParalParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"", "'task'", "'='", "", "", "", "", "", "", "", "", "", "", "':'",
		"'::'", "", "", "", "", "", "", "", "'@'", "", "", "", "", "", "", "",
		"'_'",
	}
	staticData.SymbolicNames = []string{
		"", "TASK", "ASSIGN", "STRING", "SINGLE_QUOTE_STRING", "HICH", "FLOAT",
		"NUMBER", "BOOLEAN", "ZERO_ONE", "DURATION", "URL", "PATH", "COLON",
		"COLONCOLON", "COMMA", "LBRACK", "RBRACK", "LRBRACK", "RRBRACK", "BLOCK_START",
		"BLOCK_END", "AT", "FUNCTION_START", "PIPELINE_START", "ERROR", "LOOP_KEY",
		"LOOP_VALUE", "TRY", "CATCH", "UNDERSCORE", "IDENTIFIER", "NEWLINE",
		"WS", "UNKNOWN_DEFAULT", "PIPELINE_BUF", "PIPELINE_STASH", "PIPELINE_IF_CALL_START",
		"PIPELINE_ELSEIF_CALL_START", "PIPELINE_MATCH_CALL", "PIPELINE_ELSE_CALL",
		"PIPELINE_FUNCTION_CALL_START", "PIPELINE_WS", "UNKNOWN_TEXT", "BUF_DOUBLE_BACK_ARROW",
		"BUF_WS", "STASH_DOUBLE_BACK_ARROW", "STASH_WS", "EXPRESSION_FUNCTION_CALL_START",
		"EXPRESSION_WS", "NESTED_FUNCTION_START", "FUNCTION_WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "variable_assignment", "task_definition", "task_directive",
		"pipeline_block", "pipeline_content", "pipeline_item", "unknown_command",
		"directive", "buf", "stash", "condition", "if_condition", "elseif_condition",
		"else_condition", "try_catch", "try_block", "catch_block", "match_statement",
		"match_block", "match_expression", "function", "nested_function", "argument_list",
		"expression", "hich_expr", "number_expr", "string_expr", "boolean_expr",
		"duration_expr", "matrix_expr", "list_expr", "loop_variable", "error_variable",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 51, 496, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 5, 0, 72, 8, 0, 10,
		0, 12, 0, 75, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 82, 8, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 92, 8, 3, 10, 3, 12, 3,
		95, 9, 3, 1, 3, 1, 3, 1, 3, 5, 3, 100, 8, 3, 10, 3, 12, 3, 103, 9, 3, 1,
		3, 1, 3, 5, 3, 107, 8, 3, 10, 3, 12, 3, 110, 9, 3, 1, 3, 5, 3, 113, 8,
		3, 10, 3, 12, 3, 116, 9, 3, 1, 3, 5, 3, 119, 8, 3, 10, 3, 12, 3, 122, 9,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 130, 8, 5, 10, 5, 12, 5, 133,
		9, 5, 1, 5, 1, 5, 3, 5, 137, 8, 5, 1, 6, 5, 6, 140, 8, 6, 10, 6, 12, 6,
		143, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 153, 8,
		7, 1, 8, 4, 8, 156, 8, 8, 11, 8, 12, 8, 157, 1, 8, 5, 8, 161, 8, 8, 10,
		8, 12, 8, 164, 9, 8, 1, 9, 1, 9, 1, 9, 3, 9, 169, 8, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 189, 8, 13, 10, 13, 12, 13, 192,
		9, 13, 1, 13, 1, 13, 5, 13, 196, 8, 13, 10, 13, 12, 13, 199, 9, 13, 1,
		13, 5, 13, 202, 8, 13, 10, 13, 12, 13, 205, 9, 13, 1, 13, 5, 13, 208, 8,
		13, 10, 13, 12, 13, 211, 9, 13, 1, 13, 1, 13, 5, 13, 215, 8, 13, 10, 13,
		12, 13, 218, 9, 13, 1, 13, 5, 13, 221, 8, 13, 10, 13, 12, 13, 224, 9, 13,
		1, 13, 5, 13, 227, 8, 13, 10, 13, 12, 13, 230, 9, 13, 1, 13, 3, 13, 233,
		8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 239, 8, 14, 10, 14, 12, 14, 242,
		9, 14, 1, 14, 1, 14, 5, 14, 246, 8, 14, 10, 14, 12, 14, 249, 9, 14, 1,
		14, 5, 14, 252, 8, 14, 10, 14, 12, 14, 255, 9, 14, 1, 14, 5, 14, 258, 8,
		14, 10, 14, 12, 14, 261, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 267,
		8, 15, 10, 15, 12, 15, 270, 9, 15, 1, 15, 1, 15, 5, 15, 274, 8, 15, 10,
		15, 12, 15, 277, 9, 15, 1, 15, 5, 15, 280, 8, 15, 10, 15, 12, 15, 283,
		9, 15, 1, 15, 5, 15, 286, 8, 15, 10, 15, 12, 15, 289, 9, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 5, 16, 295, 8, 16, 10, 16, 12, 16, 298, 9, 16, 1, 16,
		1, 16, 5, 16, 302, 8, 16, 10, 16, 12, 16, 305, 9, 16, 1, 16, 1, 16, 5,
		16, 309, 8, 16, 10, 16, 12, 16, 312, 9, 16, 1, 16, 1, 16, 1, 16, 5, 16,
		317, 8, 16, 10, 16, 12, 16, 320, 9, 16, 1, 16, 1, 16, 5, 16, 324, 8, 16,
		10, 16, 12, 16, 327, 9, 16, 1, 16, 1, 16, 5, 16, 331, 8, 16, 10, 16, 12,
		16, 334, 9, 16, 1, 16, 1, 16, 3, 16, 338, 8, 16, 1, 17, 5, 17, 341, 8,
		17, 10, 17, 12, 17, 344, 9, 17, 1, 18, 5, 18, 347, 8, 18, 10, 18, 12, 18,
		350, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 356, 8, 19, 10, 19, 12,
		19, 359, 9, 19, 1, 19, 1, 19, 5, 19, 363, 8, 19, 10, 19, 12, 19, 366, 9,
		19, 1, 19, 5, 19, 369, 8, 19, 10, 19, 12, 19, 372, 9, 19, 1, 19, 5, 19,
		375, 8, 19, 10, 19, 12, 19, 378, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 3, 21, 387, 8, 21, 1, 22, 1, 22, 3, 22, 391, 8, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 396, 8, 22, 1, 22, 1, 22, 1, 22, 3, 22, 401, 8,
		22, 1, 22, 3, 22, 404, 8, 22, 1, 23, 1, 23, 3, 23, 408, 8, 23, 1, 23, 1,
		23, 1, 24, 1, 24, 5, 24, 414, 8, 24, 10, 24, 12, 24, 417, 9, 24, 1, 24,
		1, 24, 5, 24, 421, 8, 24, 10, 24, 12, 24, 424, 9, 24, 1, 24, 5, 24, 427,
		8, 24, 10, 24, 12, 24, 430, 9, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 446,
		8, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1,
		30, 3, 30, 458, 8, 30, 1, 31, 1, 31, 1, 31, 4, 31, 463, 8, 31, 11, 31,
		12, 31, 464, 1, 32, 1, 32, 1, 32, 5, 32, 470, 8, 32, 10, 32, 12, 32, 473,
		9, 32, 1, 32, 1, 32, 5, 32, 477, 8, 32, 10, 32, 12, 32, 480, 9, 32, 1,
		32, 5, 32, 483, 8, 32, 10, 32, 12, 32, 486, 9, 32, 3, 32, 488, 8, 32, 1,
		32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 0, 0, 35, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 0, 6, 2, 1, 21, 21, 32,
		32, 3, 0, 11, 12, 31, 31, 43, 43, 1, 0, 6, 7, 1, 0, 3, 4, 1, 0, 8, 9, 1,
		0, 26, 27, 539, 0, 73, 1, 0, 0, 0, 2, 81, 1, 0, 0, 0, 4, 83, 1, 0, 0, 0,
		6, 93, 1, 0, 0, 0, 8, 125, 1, 0, 0, 0, 10, 127, 1, 0, 0, 0, 12, 141, 1,
		0, 0, 0, 14, 152, 1, 0, 0, 0, 16, 155, 1, 0, 0, 0, 18, 165, 1, 0, 0, 0,
		20, 170, 1, 0, 0, 0, 22, 176, 1, 0, 0, 0, 24, 182, 1, 0, 0, 0, 26, 184,
		1, 0, 0, 0, 28, 234, 1, 0, 0, 0, 30, 264, 1, 0, 0, 0, 32, 292, 1, 0, 0,
		0, 34, 342, 1, 0, 0, 0, 36, 348, 1, 0, 0, 0, 38, 351, 1, 0, 0, 0, 40, 381,
		1, 0, 0, 0, 42, 386, 1, 0, 0, 0, 44, 403, 1, 0, 0, 0, 46, 405, 1, 0, 0,
		0, 48, 411, 1, 0, 0, 0, 50, 445, 1, 0, 0, 0, 52, 447, 1, 0, 0, 0, 54, 449,
		1, 0, 0, 0, 56, 451, 1, 0, 0, 0, 58, 453, 1, 0, 0, 0, 60, 457, 1, 0, 0,
		0, 62, 459, 1, 0, 0, 0, 64, 466, 1, 0, 0, 0, 66, 491, 1, 0, 0, 0, 68, 493,
		1, 0, 0, 0, 70, 72, 3, 2, 1, 0, 71, 70, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0,
		73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 73, 1,
		0, 0, 0, 76, 77, 5, 0, 0, 1, 77, 1, 1, 0, 0, 0, 78, 82, 3, 4, 2, 0, 79,
		82, 3, 6, 3, 0, 80, 82, 5, 32, 0, 0, 81, 78, 1, 0, 0, 0, 81, 79, 1, 0,
		0, 0, 81, 80, 1, 0, 0, 0, 82, 3, 1, 0, 0, 0, 83, 84, 5, 31, 0, 0, 84, 85,
		5, 2, 0, 0, 85, 86, 3, 50, 25, 0, 86, 87, 5, 32, 0, 0, 87, 5, 1, 0, 0,
		0, 88, 89, 3, 8, 4, 0, 89, 90, 5, 32, 0, 0, 90, 92, 1, 0, 0, 0, 91, 88,
		1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0,
		94, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 97, 5, 1, 0, 0, 97, 101, 5,
		31, 0, 0, 98, 100, 5, 32, 0, 0, 99, 98, 1, 0, 0, 0, 100, 103, 1, 0, 0,
		0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 104, 1, 0, 0, 0, 103,
		101, 1, 0, 0, 0, 104, 108, 5, 20, 0, 0, 105, 107, 5, 32, 0, 0, 106, 105,
		1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0,
		0, 0, 109, 114, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 113, 3, 10, 5, 0,
		112, 111, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 120, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 119,
		5, 32, 0, 0, 118, 117, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0,
		0, 0, 120, 121, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0,
		123, 124, 5, 21, 0, 0, 124, 7, 1, 0, 0, 0, 125, 126, 3, 18, 9, 0, 126,
		9, 1, 0, 0, 0, 127, 131, 5, 24, 0, 0, 128, 130, 5, 32, 0, 0, 129, 128,
		1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0,
		0, 0, 132, 134, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 136, 3, 12, 6, 0,
		135, 137, 7, 0, 0, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137,
		11, 1, 0, 0, 0, 138, 140, 3, 14, 7, 0, 139, 138, 1, 0, 0, 0, 140, 143,
		1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 13, 1, 0,
		0, 0, 143, 141, 1, 0, 0, 0, 144, 153, 3, 44, 22, 0, 145, 153, 3, 50, 25,
		0, 146, 153, 3, 22, 11, 0, 147, 153, 3, 20, 10, 0, 148, 153, 3, 24, 12,
		0, 149, 153, 3, 32, 16, 0, 150, 153, 3, 38, 19, 0, 151, 153, 3, 16, 8,
		0, 152, 144, 1, 0, 0, 0, 152, 145, 1, 0, 0, 0, 152, 146, 1, 0, 0, 0, 152,
		147, 1, 0, 0, 0, 152, 148, 1, 0, 0, 0, 152, 149, 1, 0, 0, 0, 152, 150,
		1, 0, 0, 0, 152, 151, 1, 0, 0, 0, 153, 15, 1, 0, 0, 0, 154, 156, 7, 1,
		0, 0, 155, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0,
		157, 158, 1, 0, 0, 0, 158, 162, 1, 0, 0, 0, 159, 161, 3, 50, 25, 0, 160,
		159, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163,
		1, 0, 0, 0, 163, 17, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 166, 5, 22,
		0, 0, 166, 168, 5, 31, 0, 0, 167, 169, 3, 50, 25, 0, 168, 167, 1, 0, 0,
		0, 168, 169, 1, 0, 0, 0, 169, 19, 1, 0, 0, 0, 170, 171, 5, 35, 0, 0, 171,
		172, 3, 56, 28, 0, 172, 173, 5, 17, 0, 0, 173, 174, 5, 44, 0, 0, 174, 175,
		3, 12, 6, 0, 175, 21, 1, 0, 0, 0, 176, 177, 5, 36, 0, 0, 177, 178, 3, 56,
		28, 0, 178, 179, 5, 17, 0, 0, 179, 180, 5, 46, 0, 0, 180, 181, 3, 12, 6,
		0, 181, 23, 1, 0, 0, 0, 182, 183, 3, 26, 13, 0, 183, 25, 1, 0, 0, 0, 184,
		185, 5, 37, 0, 0, 185, 186, 3, 50, 25, 0, 186, 190, 5, 19, 0, 0, 187, 189,
		5, 32, 0, 0, 188, 187, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0,
		0, 0, 190, 191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0,
		193, 197, 5, 20, 0, 0, 194, 196, 5, 32, 0, 0, 195, 194, 1, 0, 0, 0, 196,
		199, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 203,
		1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 202, 3, 10, 5, 0, 201, 200, 1, 0,
		0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0,
		204, 209, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 208, 5, 32, 0, 0, 207,
		206, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210,
		1, 0, 0, 0, 210, 212, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 212, 222, 5, 21,
		0, 0, 213, 215, 5, 32, 0, 0, 214, 213, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0,
		216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218,
		216, 1, 0, 0, 0, 219, 221, 3, 28, 14, 0, 220, 216, 1, 0, 0, 0, 221, 224,
		1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 232, 1, 0,
		0, 0, 224, 222, 1, 0, 0, 0, 225, 227, 5, 32, 0, 0, 226, 225, 1, 0, 0, 0,
		227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229,
		231, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 233, 3, 30, 15, 0, 232, 228,
		1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 27, 1, 0, 0, 0, 234, 235, 5, 38,
		0, 0, 235, 236, 3, 50, 25, 0, 236, 240, 5, 19, 0, 0, 237, 239, 5, 32, 0,
		0, 238, 237, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240,
		241, 1, 0, 0, 0, 241, 243, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 247,
		5, 20, 0, 0, 244, 246, 5, 32, 0, 0, 245, 244, 1, 0, 0, 0, 246, 249, 1,
		0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 253, 1, 0, 0,
		0, 249, 247, 1, 0, 0, 0, 250, 252, 3, 10, 5, 0, 251, 250, 1, 0, 0, 0, 252,
		255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 259,
		1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 258, 5, 32, 0, 0, 257, 256, 1, 0,
		0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0,
		260, 262, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 263, 5, 21, 0, 0, 263,
		29, 1, 0, 0, 0, 264, 268, 5, 40, 0, 0, 265, 267, 5, 32, 0, 0, 266, 265,
		1, 0, 0, 0, 267, 270, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 269, 1, 0,
		0, 0, 269, 271, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 271, 275, 5, 20, 0, 0,
		272, 274, 5, 32, 0, 0, 273, 272, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275,
		273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 281, 1, 0, 0, 0, 277, 275,
		1, 0, 0, 0, 278, 280, 3, 10, 5, 0, 279, 278, 1, 0, 0, 0, 280, 283, 1, 0,
		0, 0, 281, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 287, 1, 0, 0, 0,
		283, 281, 1, 0, 0, 0, 284, 286, 5, 32, 0, 0, 285, 284, 1, 0, 0, 0, 286,
		289, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 290,
		1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 290, 291, 5, 21, 0, 0, 291, 31, 1, 0,
		0, 0, 292, 296, 5, 28, 0, 0, 293, 295, 5, 32, 0, 0, 294, 293, 1, 0, 0,
		0, 295, 298, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297,
		299, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 299, 303, 5, 20, 0, 0, 300, 302,
		5, 32, 0, 0, 301, 300, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0,
		0, 0, 303, 304, 1, 0, 0, 0, 304, 306, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0,
		306, 310, 3, 34, 17, 0, 307, 309, 5, 32, 0, 0, 308, 307, 1, 0, 0, 0, 309,
		312, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 313,
		1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 313, 337, 5, 21, 0, 0, 314, 318, 5, 29,
		0, 0, 315, 317, 5, 32, 0, 0, 316, 315, 1, 0, 0, 0, 317, 320, 1, 0, 0, 0,
		318, 316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 321, 1, 0, 0, 0, 320,
		318, 1, 0, 0, 0, 321, 325, 5, 20, 0, 0, 322, 324, 5, 32, 0, 0, 323, 322,
		1, 0, 0, 0, 324, 327, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0,
		0, 0, 326, 328, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 328, 332, 3, 36, 18,
		0, 329, 331, 5, 32, 0, 0, 330, 329, 1, 0, 0, 0, 331, 334, 1, 0, 0, 0, 332,
		330, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 335, 1, 0, 0, 0, 334, 332,
		1, 0, 0, 0, 335, 336, 5, 21, 0, 0, 336, 338, 1, 0, 0, 0, 337, 314, 1, 0,
		0, 0, 337, 338, 1, 0, 0, 0, 338, 33, 1, 0, 0, 0, 339, 341, 3, 10, 5, 0,
		340, 339, 1, 0, 0, 0, 341, 344, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 342,
		343, 1, 0, 0, 0, 343, 35, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 345, 347, 3,
		10, 5, 0, 346, 345, 1, 0, 0, 0, 347, 350, 1, 0, 0, 0, 348, 346, 1, 0, 0,
		0, 348, 349, 1, 0, 0, 0, 349, 37, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 351,
		352, 5, 39, 0, 0, 352, 353, 3, 50, 25, 0, 353, 357, 5, 19, 0, 0, 354, 356,
		5, 32, 0, 0, 355, 354, 1, 0, 0, 0, 356, 359, 1, 0, 0, 0, 357, 355, 1, 0,
		0, 0, 357, 358, 1, 0, 0, 0, 358, 360, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0,
		360, 364, 5, 20, 0, 0, 361, 363, 5, 32, 0, 0, 362, 361, 1, 0, 0, 0, 363,
		366, 1, 0, 0, 0, 364, 362, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 370,
		1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 367, 369, 3, 40, 20, 0, 368, 367, 1,
		0, 0, 0, 369, 372, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 370, 371, 1, 0, 0,
		0, 371, 376, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 373, 375, 5, 32, 0, 0, 374,
		373, 1, 0, 0, 0, 375, 378, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 377,
		1, 0, 0, 0, 377, 379, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 379, 380, 5, 21,
		0, 0, 380, 39, 1, 0, 0, 0, 381, 382, 3, 42, 21, 0, 382, 383, 3, 10, 5,
		0, 383, 41, 1, 0, 0, 0, 384, 387, 3, 50, 25, 0, 385, 387, 5, 30, 0, 0,
		386, 384, 1, 0, 0, 0, 386, 385, 1, 0, 0, 0, 387, 43, 1, 0, 0, 0, 388, 390,
		5, 41, 0, 0, 389, 391, 3, 48, 24, 0, 390, 389, 1, 0, 0, 0, 390, 391, 1,
		0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 404, 5, 19, 0, 0, 393, 395, 5, 23,
		0, 0, 394, 396, 3, 48, 24, 0, 395, 394, 1, 0, 0, 0, 395, 396, 1, 0, 0,
		0, 396, 397, 1, 0, 0, 0, 397, 404, 5, 19, 0, 0, 398, 400, 5, 48, 0, 0,
		399, 401, 3, 48, 24, 0, 400, 399, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401,
		402, 1, 0, 0, 0, 402, 404, 5, 19, 0, 0, 403, 388, 1, 0, 0, 0, 403, 393,
		1, 0, 0, 0, 403, 398, 1, 0, 0, 0, 404, 45, 1, 0, 0, 0, 405, 407, 5, 50,
		0, 0, 406, 408, 3, 48, 24, 0, 407, 406, 1, 0, 0, 0, 407, 408, 1, 0, 0,
		0, 408, 409, 1, 0, 0, 0, 409, 410, 5, 19, 0, 0, 410, 47, 1, 0, 0, 0, 411,
		428, 3, 50, 25, 0, 412, 414, 5, 32, 0, 0, 413, 412, 1, 0, 0, 0, 414, 417,
		1, 0, 0, 0, 415, 413, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 418, 1, 0,
		0, 0, 417, 415, 1, 0, 0, 0, 418, 422, 5, 15, 0, 0, 419, 421, 5, 32, 0,
		0, 420, 419, 1, 0, 0, 0, 421, 424, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0, 422,
		423, 1, 0, 0, 0, 423, 425, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 425, 427,
		3, 50, 25, 0, 426, 415, 1, 0, 0, 0, 427, 430, 1, 0, 0, 0, 428, 426, 1,
		0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 49, 1, 0, 0, 0, 430, 428, 1, 0, 0,
		0, 431, 446, 3, 52, 26, 0, 432, 446, 3, 66, 33, 0, 433, 446, 3, 68, 34,
		0, 434, 446, 3, 44, 22, 0, 435, 446, 3, 46, 23, 0, 436, 446, 5, 11, 0,
		0, 437, 446, 5, 12, 0, 0, 438, 446, 3, 54, 27, 0, 439, 446, 3, 56, 28,
		0, 440, 446, 3, 58, 29, 0, 441, 446, 3, 60, 30, 0, 442, 446, 3, 62, 31,
		0, 443, 446, 3, 64, 32, 0, 444, 446, 5, 31, 0, 0, 445, 431, 1, 0, 0, 0,
		445, 432, 1, 0, 0, 0, 445, 433, 1, 0, 0, 0, 445, 434, 1, 0, 0, 0, 445,
		435, 1, 0, 0, 0, 445, 436, 1, 0, 0, 0, 445, 437, 1, 0, 0, 0, 445, 438,
		1, 0, 0, 0, 445, 439, 1, 0, 0, 0, 445, 440, 1, 0, 0, 0, 445, 441, 1, 0,
		0, 0, 445, 442, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 445, 444, 1, 0, 0, 0,
		446, 51, 1, 0, 0, 0, 447, 448, 5, 5, 0, 0, 448, 53, 1, 0, 0, 0, 449, 450,
		7, 2, 0, 0, 450, 55, 1, 0, 0, 0, 451, 452, 7, 3, 0, 0, 452, 57, 1, 0, 0,
		0, 453, 454, 7, 4, 0, 0, 454, 59, 1, 0, 0, 0, 455, 458, 5, 10, 0, 0, 456,
		458, 3, 54, 27, 0, 457, 455, 1, 0, 0, 0, 457, 456, 1, 0, 0, 0, 458, 61,
		1, 0, 0, 0, 459, 462, 3, 64, 32, 0, 460, 461, 5, 14, 0, 0, 461, 463, 3,
		64, 32, 0, 462, 460, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 462, 1, 0,
		0, 0, 464, 465, 1, 0, 0, 0, 465, 63, 1, 0, 0, 0, 466, 487, 5, 16, 0, 0,
		467, 484, 3, 50, 25, 0, 468, 470, 5, 32, 0, 0, 469, 468, 1, 0, 0, 0, 470,
		473, 1, 0, 0, 0, 471, 469, 1, 0, 0, 0, 471, 472, 1, 0, 0, 0, 472, 474,
		1, 0, 0, 0, 473, 471, 1, 0, 0, 0, 474, 478, 5, 15, 0, 0, 475, 477, 5, 32,
		0, 0, 476, 475, 1, 0, 0, 0, 477, 480, 1, 0, 0, 0, 478, 476, 1, 0, 0, 0,
		478, 479, 1, 0, 0, 0, 479, 481, 1, 0, 0, 0, 480, 478, 1, 0, 0, 0, 481,
		483, 3, 50, 25, 0, 482, 471, 1, 0, 0, 0, 483, 486, 1, 0, 0, 0, 484, 482,
		1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 488, 1, 0, 0, 0, 486, 484, 1, 0,
		0, 0, 487, 467, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0,
		489, 490, 5, 17, 0, 0, 490, 65, 1, 0, 0, 0, 491, 492, 7, 5, 0, 0, 492,
		67, 1, 0, 0, 0, 493, 494, 5, 25, 0, 0, 494, 69, 1, 0, 0, 0, 59, 73, 81,
		93, 101, 108, 114, 120, 131, 136, 141, 152, 157, 162, 168, 190, 197, 203,
		209, 216, 222, 228, 232, 240, 247, 253, 259, 268, 275, 281, 287, 296, 303,
		310, 318, 325, 332, 337, 342, 348, 357, 364, 370, 376, 386, 390, 395, 400,
		403, 407, 415, 422, 428, 445, 457, 464, 471, 478, 484, 487,
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
	ParalParserHICH                           = 5
	ParalParserFLOAT                          = 6
	ParalParserNUMBER                         = 7
	ParalParserBOOLEAN                        = 8
	ParalParserZERO_ONE                       = 9
	ParalParserDURATION                       = 10
	ParalParserURL                            = 11
	ParalParserPATH                           = 12
	ParalParserCOLON                          = 13
	ParalParserCOLONCOLON                     = 14
	ParalParserCOMMA                          = 15
	ParalParserLBRACK                         = 16
	ParalParserRBRACK                         = 17
	ParalParserLRBRACK                        = 18
	ParalParserRRBRACK                        = 19
	ParalParserBLOCK_START                    = 20
	ParalParserBLOCK_END                      = 21
	ParalParserAT                             = 22
	ParalParserFUNCTION_START                 = 23
	ParalParserPIPELINE_START                 = 24
	ParalParserERROR                          = 25
	ParalParserLOOP_KEY                       = 26
	ParalParserLOOP_VALUE                     = 27
	ParalParserTRY                            = 28
	ParalParserCATCH                          = 29
	ParalParserUNDERSCORE                     = 30
	ParalParserIDENTIFIER                     = 31
	ParalParserNEWLINE                        = 32
	ParalParserWS                             = 33
	ParalParserUNKNOWN_DEFAULT                = 34
	ParalParserPIPELINE_BUF                   = 35
	ParalParserPIPELINE_STASH                 = 36
	ParalParserPIPELINE_IF_CALL_START         = 37
	ParalParserPIPELINE_ELSEIF_CALL_START     = 38
	ParalParserPIPELINE_MATCH_CALL            = 39
	ParalParserPIPELINE_ELSE_CALL             = 40
	ParalParserPIPELINE_FUNCTION_CALL_START   = 41
	ParalParserPIPELINE_WS                    = 42
	ParalParserUNKNOWN_TEXT                   = 43
	ParalParserBUF_DOUBLE_BACK_ARROW          = 44
	ParalParserBUF_WS                         = 45
	ParalParserSTASH_DOUBLE_BACK_ARROW        = 46
	ParalParserSTASH_WS                       = 47
	ParalParserEXPRESSION_FUNCTION_CALL_START = 48
	ParalParserEXPRESSION_WS                  = 49
	ParalParserNESTED_FUNCTION_START          = 50
	ParalParserFUNCTION_WS                    = 51
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
	ParalParserRULE_match_statement     = 19
	ParalParserRULE_match_block         = 20
	ParalParserRULE_match_expression    = 21
	ParalParserRULE_function            = 22
	ParalParserRULE_nested_function     = 23
	ParalParserRULE_argument_list       = 24
	ParalParserRULE_expression          = 25
	ParalParserRULE_hich_expr           = 26
	ParalParserRULE_number_expr         = 27
	ParalParserRULE_string_expr         = 28
	ParalParserRULE_boolean_expr        = 29
	ParalParserRULE_duration_expr       = 30
	ParalParserRULE_matrix_expr         = 31
	ParalParserRULE_list_expr           = 32
	ParalParserRULE_loop_variable       = 33
	ParalParserRULE_error_variable      = 34
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
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6446645250) != 0 {
		{
			p.SetState(70)
			p.Statement()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Variable_assignment()
		}

	case ParalParserTASK, ParalParserAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Task_definition()
		}

	case ParalParserNEWLINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
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
		p.SetState(83)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(ParalParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Expression()
	}
	{
		p.SetState(86)
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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserAT {
		{
			p.SetState(88)
			p.Task_directive()
		}
		{
			p.SetState(89)
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
		p.Match(ParalParserTASK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(98)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(104)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(108)
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
				p.SetState(105)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(111)
			p.Pipeline_block()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(117)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(123)
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
		p.SetState(125)
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
		p.SetState(127)
		p.Match(ParalParserPIPELINE_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
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
				p.SetState(128)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(133)
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
		p.SetState(134)
		p.Pipeline_content()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(135)
			_la = p.GetTokenStream().LA(1)

			if !((int64((_la - -1)) & ^0x3f) == 0 && ((int64(1)<<(_la - -1))&8594128897) != 0) {
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
	p.SetState(141)
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
				p.SetState(138)
				p.Pipeline_item()
			}

		}
		p.SetState(143)
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
	Function() IFunctionContext
	Expression() IExpressionContext
	Stash() IStashContext
	Buf() IBufContext
	Condition() IConditionContext
	Try_catch() ITry_catchContext
	Match_statement() IMatch_statementContext
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

func (s *Pipeline_itemContext) Expression() IExpressionContext {
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

func (s *Pipeline_itemContext) Match_statement() IMatch_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatch_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatch_statementContext)
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
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Function()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Expression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.Stash()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(147)
			p.Buf()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(148)
			p.Condition()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(149)
			p.Try_catch()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(150)
			p.Match_statement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(151)
			p.Unknown_command()
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
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(154)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8798240512000) != 0) {
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

		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(162)
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
				p.SetState(159)
				p.Expression()
			}

		}
		p.SetState(164)
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
		p.SetState(165)
		p.Match(ParalParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(ParalParserIDENTIFIER)
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1409576297635832) != 0 {
		{
			p.SetState(167)
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
		p.SetState(170)
		p.Match(ParalParserPIPELINE_BUF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.String_expr()
	}
	{
		p.SetState(172)
		p.Match(ParalParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(ParalParserBUF_DOUBLE_BACK_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
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
		p.SetState(176)
		p.Match(ParalParserPIPELINE_STASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.String_expr()
	}
	{
		p.SetState(178)
		p.Match(ParalParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Match(ParalParserSTASH_DOUBLE_BACK_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
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
		p.SetState(182)
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
	RRBRACK() antlr.TerminalNode
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

func (s *If_conditionContext) RRBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRRBRACK, 0)
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
		p.SetState(184)
		p.Match(ParalParserPIPELINE_IF_CALL_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Expression()
	}
	{
		p.SetState(186)
		p.Match(ParalParserRRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(197)
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
				p.SetState(194)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(200)
			p.Pipeline_block()
		}

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
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
		p.Match(ParalParserBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(222)
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
			p.SetState(216)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParalParserNEWLINE {
				{
					p.SetState(213)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(218)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(219)
				p.Elseif_condition()
			}

		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(225)
				p.Match(ParalParserNEWLINE)
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
		}
		{
			p.SetState(231)
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
	RRBRACK() antlr.TerminalNode
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

func (s *Elseif_conditionContext) RRBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRRBRACK, 0)
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
		p.SetState(234)
		p.Match(ParalParserPIPELINE_ELSEIF_CALL_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.Expression()
	}
	{
		p.SetState(236)
		p.Match(ParalParserRRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(247)
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
				p.SetState(244)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(250)
			p.Pipeline_block()
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(256)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(262)
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
		p.SetState(264)
		p.Match(ParalParserPIPELINE_ELSE_CALL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
		p.Match(ParalParserBLOCK_START)
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(272)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(278)
			p.Pipeline_block()
		}

		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(284)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(290)
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
		p.SetState(292)
		p.Match(ParalParserTRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(293)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(299)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(303)
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
				p.SetState(300)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(305)
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
		p.SetState(306)
		p.Try_block()
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(307)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(313)
		p.Match(ParalParserBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ParalParserCATCH {
		{
			p.SetState(314)
			p.Match(ParalParserCATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(315)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(320)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(321)
			p.Match(ParalParserBLOCK_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(325)
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
					p.SetState(322)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(327)
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
			p.SetState(328)
			p.Catch_block()
		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(329)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(334)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(335)
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
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(339)
			p.Pipeline_block()
		}

		p.SetState(344)
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
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserPIPELINE_START {
		{
			p.SetState(345)
			p.Pipeline_block()
		}

		p.SetState(350)
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

// IMatch_statementContext is an interface to support dynamic dispatch.
type IMatch_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPELINE_MATCH_CALL() antlr.TerminalNode
	Expression() IExpressionContext
	RRBRACK() antlr.TerminalNode
	BLOCK_START() antlr.TerminalNode
	BLOCK_END() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllMatch_block() []IMatch_blockContext
	Match_block(i int) IMatch_blockContext

	// IsMatch_statementContext differentiates from other interfaces.
	IsMatch_statementContext()
}

type Match_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatch_statementContext() *Match_statementContext {
	var p = new(Match_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_match_statement
	return p
}

func InitEmptyMatch_statementContext(p *Match_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_match_statement
}

func (*Match_statementContext) IsMatch_statementContext() {}

func NewMatch_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Match_statementContext {
	var p = new(Match_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_match_statement

	return p
}

func (s *Match_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Match_statementContext) PIPELINE_MATCH_CALL() antlr.TerminalNode {
	return s.GetToken(ParalParserPIPELINE_MATCH_CALL, 0)
}

func (s *Match_statementContext) Expression() IExpressionContext {
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

func (s *Match_statementContext) RRBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRRBRACK, 0)
}

func (s *Match_statementContext) BLOCK_START() antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_START, 0)
}

func (s *Match_statementContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(ParalParserBLOCK_END, 0)
}

func (s *Match_statementContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNEWLINE)
}

func (s *Match_statementContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, i)
}

func (s *Match_statementContext) AllMatch_block() []IMatch_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatch_blockContext); ok {
			len++
		}
	}

	tst := make([]IMatch_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatch_blockContext); ok {
			tst[i] = t.(IMatch_blockContext)
			i++
		}
	}

	return tst
}

func (s *Match_statementContext) Match_block(i int) IMatch_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatch_blockContext); ok {
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

	return t.(IMatch_blockContext)
}

func (s *Match_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Match_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Match_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterMatch_statement(s)
	}
}

func (s *Match_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitMatch_statement(s)
	}
}

func (p *ParalParser) Match_statement() (localctx IMatch_statementContext) {
	localctx = NewMatch_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ParalParserRULE_match_statement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Match(ParalParserPIPELINE_MATCH_CALL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(352)
		p.Expression()
	}
	{
		p.SetState(353)
		p.Match(ParalParserRRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(354)
			p.Match(ParalParserNEWLINE)
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
	}
	{
		p.SetState(360)
		p.Match(ParalParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(361)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1409577371377656) != 0 {
		{
			p.SetState(367)
			p.Match_block()
		}

		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserNEWLINE {
		{
			p.SetState(373)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(379)
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

// IMatch_blockContext is an interface to support dynamic dispatch.
type IMatch_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Match_expression() IMatch_expressionContext
	Pipeline_block() IPipeline_blockContext

	// IsMatch_blockContext differentiates from other interfaces.
	IsMatch_blockContext()
}

type Match_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatch_blockContext() *Match_blockContext {
	var p = new(Match_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_match_block
	return p
}

func InitEmptyMatch_blockContext(p *Match_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_match_block
}

func (*Match_blockContext) IsMatch_blockContext() {}

func NewMatch_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Match_blockContext {
	var p = new(Match_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_match_block

	return p
}

func (s *Match_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Match_blockContext) Match_expression() IMatch_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatch_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatch_expressionContext)
}

func (s *Match_blockContext) Pipeline_block() IPipeline_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipeline_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipeline_blockContext)
}

func (s *Match_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Match_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Match_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterMatch_block(s)
	}
}

func (s *Match_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitMatch_block(s)
	}
}

func (p *ParalParser) Match_block() (localctx IMatch_blockContext) {
	localctx = NewMatch_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ParalParserRULE_match_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match_expression()
	}
	{
		p.SetState(382)
		p.Pipeline_block()
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

// IMatch_expressionContext is an interface to support dynamic dispatch.
type IMatch_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	UNDERSCORE() antlr.TerminalNode

	// IsMatch_expressionContext differentiates from other interfaces.
	IsMatch_expressionContext()
}

type Match_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatch_expressionContext() *Match_expressionContext {
	var p = new(Match_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_match_expression
	return p
}

func InitEmptyMatch_expressionContext(p *Match_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_match_expression
}

func (*Match_expressionContext) IsMatch_expressionContext() {}

func NewMatch_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Match_expressionContext {
	var p = new(Match_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_match_expression

	return p
}

func (s *Match_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Match_expressionContext) Expression() IExpressionContext {
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

func (s *Match_expressionContext) UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(ParalParserUNDERSCORE, 0)
}

func (s *Match_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Match_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Match_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterMatch_expression(s)
	}
}

func (s *Match_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitMatch_expression(s)
	}
}

func (p *ParalParser) Match_expression() (localctx IMatch_expressionContext) {
	localctx = NewMatch_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ParalParserRULE_match_expression)
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserSTRING, ParalParserSINGLE_QUOTE_STRING, ParalParserHICH, ParalParserFLOAT, ParalParserNUMBER, ParalParserBOOLEAN, ParalParserZERO_ONE, ParalParserDURATION, ParalParserURL, ParalParserPATH, ParalParserLBRACK, ParalParserFUNCTION_START, ParalParserERROR, ParalParserLOOP_KEY, ParalParserLOOP_VALUE, ParalParserIDENTIFIER, ParalParserPIPELINE_FUNCTION_CALL_START, ParalParserEXPRESSION_FUNCTION_CALL_START, ParalParserNESTED_FUNCTION_START:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(384)
			p.Expression()
		}

	case ParalParserUNDERSCORE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(385)
			p.Match(ParalParserUNDERSCORE)
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

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPELINE_FUNCTION_CALL_START() antlr.TerminalNode
	RRBRACK() antlr.TerminalNode
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

func (s *FunctionContext) RRBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRRBRACK, 0)
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
	p.EnterRule(localctx, 44, ParalParserRULE_function)
	var _la int

	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserPIPELINE_FUNCTION_CALL_START:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(388)
			p.Match(ParalParserPIPELINE_FUNCTION_CALL_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1409576297635832) != 0 {
			{
				p.SetState(389)
				p.Argument_list()
			}

		}
		{
			p.SetState(392)
			p.Match(ParalParserRRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserFUNCTION_START:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)
			p.Match(ParalParserFUNCTION_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1409576297635832) != 0 {
			{
				p.SetState(394)
				p.Argument_list()
			}

		}
		{
			p.SetState(397)
			p.Match(ParalParserRRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserEXPRESSION_FUNCTION_CALL_START:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(398)
			p.Match(ParalParserEXPRESSION_FUNCTION_CALL_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(400)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1409576297635832) != 0 {
			{
				p.SetState(399)
				p.Argument_list()
			}

		}
		{
			p.SetState(402)
			p.Match(ParalParserRRBRACK)
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
	RRBRACK() antlr.TerminalNode
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

func (s *Nested_functionContext) RRBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRRBRACK, 0)
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
	p.EnterRule(localctx, 46, ParalParserRULE_nested_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)
		p.Match(ParalParserNESTED_FUNCTION_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1409576297635832) != 0 {
		{
			p.SetState(406)
			p.Argument_list()
		}

	}
	{
		p.SetState(409)
		p.Match(ParalParserRRBRACK)
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
	p.EnterRule(localctx, 48, ParalParserRULE_argument_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.Expression()
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserCOMMA || _la == ParalParserNEWLINE {
		p.SetState(415)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(412)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(417)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(418)
			p.Match(ParalParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(422)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserNEWLINE {
			{
				p.SetState(419)
				p.Match(ParalParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(424)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(425)
			p.Expression()
		}

		p.SetState(430)
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
	Hich_expr() IHich_exprContext
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

func (s *ExpressionContext) Hich_expr() IHich_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHich_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHich_exprContext)
}

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
	p.EnterRule(localctx, 50, ParalParserRULE_expression)
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
			p.Hich_expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(432)
			p.Loop_variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(433)
			p.Error_variable()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(434)
			p.Function()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(435)
			p.Nested_function()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(436)
			p.Match(ParalParserURL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(437)
			p.Match(ParalParserPATH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(438)
			p.Number_expr()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(439)
			p.String_expr()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(440)
			p.Boolean_expr()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(441)
			p.Duration_expr()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(442)
			p.Matrix_expr()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(443)
			p.List_expr()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(444)
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

// IHich_exprContext is an interface to support dynamic dispatch.
type IHich_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HICH() antlr.TerminalNode

	// IsHich_exprContext differentiates from other interfaces.
	IsHich_exprContext()
}

type Hich_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHich_exprContext() *Hich_exprContext {
	var p = new(Hich_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_hich_expr
	return p
}

func InitEmptyHich_exprContext(p *Hich_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_hich_expr
}

func (*Hich_exprContext) IsHich_exprContext() {}

func NewHich_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Hich_exprContext {
	var p = new(Hich_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_hich_expr

	return p
}

func (s *Hich_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Hich_exprContext) HICH() antlr.TerminalNode {
	return s.GetToken(ParalParserHICH, 0)
}

func (s *Hich_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Hich_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Hich_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.EnterHich_expr(s)
	}
}

func (s *Hich_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalParserListener); ok {
		listenerT.ExitHich_expr(s)
	}
}

func (p *ParalParser) Hich_expr() (localctx IHich_exprContext) {
	localctx = NewHich_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ParalParserRULE_hich_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
		p.Match(ParalParserHICH)
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
	p.EnterRule(localctx, 54, ParalParserRULE_number_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(449)
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
	p.EnterRule(localctx, 56, ParalParserRULE_string_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
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
	p.EnterRule(localctx, 58, ParalParserRULE_boolean_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
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
	p.EnterRule(localctx, 60, ParalParserRULE_duration_expr)
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserDURATION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(455)
			p.Match(ParalParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserFLOAT, ParalParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(456)
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
	p.EnterRule(localctx, 62, ParalParserRULE_matrix_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(459)
		p.List_expr()
	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ParalParserCOLONCOLON {
		{
			p.SetState(460)
			p.Match(ParalParserCOLONCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.List_expr()
		}

		p.SetState(464)
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
	p.EnterRule(localctx, 64, ParalParserRULE_list_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
		p.Match(ParalParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1409576297635832) != 0 {
		{
			p.SetState(467)
			p.Expression()
		}
		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserCOMMA || _la == ParalParserNEWLINE {
			p.SetState(471)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParalParserNEWLINE {
				{
					p.SetState(468)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(473)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(474)
				p.Match(ParalParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(478)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParalParserNEWLINE {
				{
					p.SetState(475)
					p.Match(ParalParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(480)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(481)
				p.Expression()
			}

			p.SetState(486)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(489)
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
	p.EnterRule(localctx, 66, ParalParserRULE_loop_variable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
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
	p.EnterRule(localctx, 68, ParalParserRULE_error_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)
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
