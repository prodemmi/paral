// Code generated from antlr/Paral.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Paral

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

var ParalParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func paralParserInit() {
	staticData := &ParalParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "'@'", "", "", "", "", "'='", "'('", "')'",
		"", "", "", "':'", "'::'", "'['", "']'", "','", "'\\t'",
	}
	staticData.SymbolicNames = []string{
		"", "IDENTIFIER", "STRING", "SINGLE_QUOTE_STRING", "NUMBER", "BOOLEAN",
		"DURATION", "AT", "MATRIX_REF", "REF", "ARG", "SINGLE_DASH_ARG", "EQUALS",
		"LPAREN", "RPAREN", "URL", "VALUE", "BSL_NEWLINE", "COLON", "COLONCOLON",
		"LBRACK", "RBRACK", "COMMA", "TAB", "NEWLINE", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "line", "variable_def", "matrix_def", "list_expr", "comment_def",
		"job_def", "job_directive_def", "job_directive_value", "cmd_expr", "cmd_directive",
		"cmd_directive_iden", "cmd_directive_value", "directive_arg", "value_expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 164, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 5, 0,
		32, 8, 0, 10, 0, 12, 0, 35, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 50, 8, 1, 1, 2, 1, 2, 1, 2,
		3, 2, 55, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 61, 8, 3, 10, 3, 12, 3, 64,
		9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 70, 8, 4, 10, 4, 12, 4, 73, 9, 4, 3,
		4, 75, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 5, 6, 82, 8, 6, 10, 6, 12, 6,
		85, 9, 6, 1, 6, 1, 6, 1, 6, 4, 6, 90, 8, 6, 11, 6, 12, 6, 91, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 5, 7, 99, 8, 7, 10, 7, 12, 7, 102, 9, 7, 3, 7, 104,
		8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 113, 8, 8, 1, 9,
		1, 9, 3, 9, 117, 8, 9, 1, 9, 3, 9, 120, 8, 9, 1, 9, 5, 9, 123, 8, 9, 10,
		9, 12, 9, 126, 9, 9, 1, 10, 1, 10, 1, 10, 3, 10, 131, 8, 10, 1, 10, 5,
		10, 134, 8, 10, 10, 10, 12, 10, 137, 9, 10, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 5, 12, 145, 8, 12, 10, 12, 12, 12, 148, 9, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 3, 13, 155, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 3, 14, 162, 8, 14, 1, 14, 0, 0, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 0, 2, 2, 0, 23, 23, 26, 26, 3, 0, 1, 4, 8, 12,
		15, 17, 179, 0, 33, 1, 0, 0, 0, 2, 49, 1, 0, 0, 0, 4, 51, 1, 0, 0, 0, 6,
		56, 1, 0, 0, 0, 8, 65, 1, 0, 0, 0, 10, 78, 1, 0, 0, 0, 12, 83, 1, 0, 0,
		0, 14, 93, 1, 0, 0, 0, 16, 112, 1, 0, 0, 0, 18, 114, 1, 0, 0, 0, 20, 127,
		1, 0, 0, 0, 22, 138, 1, 0, 0, 0, 24, 141, 1, 0, 0, 0, 26, 154, 1, 0, 0,
		0, 28, 161, 1, 0, 0, 0, 30, 32, 3, 2, 1, 0, 31, 30, 1, 0, 0, 0, 32, 35,
		1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 36, 1, 0, 0, 0,
		35, 33, 1, 0, 0, 0, 36, 37, 5, 0, 0, 1, 37, 1, 1, 0, 0, 0, 38, 39, 3, 4,
		2, 0, 39, 40, 5, 24, 0, 0, 40, 50, 1, 0, 0, 0, 41, 42, 3, 6, 3, 0, 42,
		43, 5, 24, 0, 0, 43, 50, 1, 0, 0, 0, 44, 45, 3, 10, 5, 0, 45, 46, 5, 24,
		0, 0, 46, 50, 1, 0, 0, 0, 47, 50, 3, 12, 6, 0, 48, 50, 5, 24, 0, 0, 49,
		38, 1, 0, 0, 0, 49, 41, 1, 0, 0, 0, 49, 44, 1, 0, 0, 0, 49, 47, 1, 0, 0,
		0, 49, 48, 1, 0, 0, 0, 50, 3, 1, 0, 0, 0, 51, 54, 5, 1, 0, 0, 52, 55, 3,
		28, 14, 0, 53, 55, 3, 8, 4, 0, 54, 52, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0,
		55, 5, 1, 0, 0, 0, 56, 57, 5, 1, 0, 0, 57, 62, 3, 8, 4, 0, 58, 59, 5, 19,
		0, 0, 59, 61, 3, 8, 4, 0, 60, 58, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60,
		1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 7, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0,
		65, 74, 5, 20, 0, 0, 66, 71, 3, 28, 14, 0, 67, 68, 5, 22, 0, 0, 68, 70,
		3, 28, 14, 0, 69, 67, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0,
		0, 71, 72, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 66,
		1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 5, 21, 0, 0,
		77, 9, 1, 0, 0, 0, 78, 79, 5, 25, 0, 0, 79, 11, 1, 0, 0, 0, 80, 82, 3,
		14, 7, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83,
		84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 87, 5, 1, 0,
		0, 87, 89, 5, 18, 0, 0, 88, 90, 3, 18, 9, 0, 89, 88, 1, 0, 0, 0, 90, 91,
		1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 13, 1, 0, 0, 0,
		93, 94, 5, 7, 0, 0, 94, 103, 5, 1, 0, 0, 95, 100, 3, 16, 8, 0, 96, 97,
		5, 22, 0, 0, 97, 99, 3, 16, 8, 0, 98, 96, 1, 0, 0, 0, 99, 102, 1, 0, 0,
		0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102,
		100, 1, 0, 0, 0, 103, 95, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1,
		0, 0, 0, 105, 106, 5, 24, 0, 0, 106, 15, 1, 0, 0, 0, 107, 113, 3, 28, 14,
		0, 108, 113, 5, 1, 0, 0, 109, 113, 5, 8, 0, 0, 110, 113, 5, 9, 0, 0, 111,
		113, 5, 16, 0, 0, 112, 107, 1, 0, 0, 0, 112, 108, 1, 0, 0, 0, 112, 109,
		1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113, 17, 1, 0,
		0, 0, 114, 116, 5, 24, 0, 0, 115, 117, 7, 0, 0, 0, 116, 115, 1, 0, 0, 0,
		116, 117, 1, 0, 0, 0, 117, 119, 1, 0, 0, 0, 118, 120, 3, 20, 10, 0, 119,
		118, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 124, 1, 0, 0, 0, 121, 123,
		7, 1, 0, 0, 122, 121, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0,
		0, 0, 124, 125, 1, 0, 0, 0, 125, 19, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0,
		127, 135, 3, 22, 11, 0, 128, 130, 5, 13, 0, 0, 129, 131, 3, 24, 12, 0,
		130, 129, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132,
		134, 5, 14, 0, 0, 133, 128, 1, 0, 0, 0, 134, 137, 1, 0, 0, 0, 135, 133,
		1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 21, 1, 0, 0, 0, 137, 135, 1, 0,
		0, 0, 138, 139, 5, 7, 0, 0, 139, 140, 5, 1, 0, 0, 140, 23, 1, 0, 0, 0,
		141, 146, 3, 26, 13, 0, 142, 143, 5, 22, 0, 0, 143, 145, 3, 26, 13, 0,
		144, 142, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146,
		147, 1, 0, 0, 0, 147, 25, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 155, 3,
		28, 14, 0, 150, 155, 5, 1, 0, 0, 151, 155, 5, 8, 0, 0, 152, 155, 5, 9,
		0, 0, 153, 155, 5, 16, 0, 0, 154, 149, 1, 0, 0, 0, 154, 150, 1, 0, 0, 0,
		154, 151, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 153, 1, 0, 0, 0, 155,
		27, 1, 0, 0, 0, 156, 162, 5, 2, 0, 0, 157, 162, 5, 3, 0, 0, 158, 162, 5,
		4, 0, 0, 159, 162, 5, 5, 0, 0, 160, 162, 5, 6, 0, 0, 161, 156, 1, 0, 0,
		0, 161, 157, 1, 0, 0, 0, 161, 158, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161,
		160, 1, 0, 0, 0, 162, 29, 1, 0, 0, 0, 19, 33, 49, 54, 62, 71, 74, 83, 91,
		100, 103, 112, 116, 119, 124, 130, 135, 146, 154, 161,
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
	staticData := &ParalParserStaticData
	staticData.once.Do(paralParserInit)
}

// NewParalParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParalParser(input antlr.TokenStream) *ParalParser {
	ParalParserInit()
	this := new(ParalParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ParalParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Paral.g4"

	return this
}

// ParalParser tokens.
const (
	ParalParserEOF                 = antlr.TokenEOF
	ParalParserIDENTIFIER          = 1
	ParalParserSTRING              = 2
	ParalParserSINGLE_QUOTE_STRING = 3
	ParalParserNUMBER              = 4
	ParalParserBOOLEAN             = 5
	ParalParserDURATION            = 6
	ParalParserAT                  = 7
	ParalParserMATRIX_REF          = 8
	ParalParserREF                 = 9
	ParalParserARG                 = 10
	ParalParserSINGLE_DASH_ARG     = 11
	ParalParserEQUALS              = 12
	ParalParserLPAREN              = 13
	ParalParserRPAREN              = 14
	ParalParserURL                 = 15
	ParalParserVALUE               = 16
	ParalParserBSL_NEWLINE         = 17
	ParalParserCOLON               = 18
	ParalParserCOLONCOLON          = 19
	ParalParserLBRACK              = 20
	ParalParserRBRACK              = 21
	ParalParserCOMMA               = 22
	ParalParserTAB                 = 23
	ParalParserNEWLINE             = 24
	ParalParserCOMMENT             = 25
	ParalParserWS                  = 26
)

// ParalParser rules.
const (
	ParalParserRULE_program             = 0
	ParalParserRULE_line                = 1
	ParalParserRULE_variable_def        = 2
	ParalParserRULE_matrix_def          = 3
	ParalParserRULE_list_expr           = 4
	ParalParserRULE_comment_def         = 5
	ParalParserRULE_job_def             = 6
	ParalParserRULE_job_directive_def   = 7
	ParalParserRULE_job_directive_value = 8
	ParalParserRULE_cmd_expr            = 9
	ParalParserRULE_cmd_directive       = 10
	ParalParserRULE_cmd_directive_iden  = 11
	ParalParserRULE_cmd_directive_value = 12
	ParalParserRULE_directive_arg       = 13
	ParalParserRULE_value_expr          = 14
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllLine() []ILineContext
	Line(i int) ILineContext

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

func (s *ProgramContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
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

	return t.(ILineContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *ParalParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParalParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&50331778) != 0 {
		{
			p.SetState(30)
			p.Line()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
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

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable_def() IVariable_defContext
	NEWLINE() antlr.TerminalNode
	Matrix_def() IMatrix_defContext
	Comment_def() IComment_defContext
	Job_def() IJob_defContext

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_line
	return p
}

func InitEmptyLineContext(p *LineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_line
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Variable_def() IVariable_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_defContext)
}

func (s *LineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, 0)
}

func (s *LineContext) Matrix_def() IMatrix_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_defContext)
}

func (s *LineContext) Comment_def() IComment_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComment_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComment_defContext)
}

func (s *LineContext) Job_def() IJob_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJob_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJob_defContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *ParalParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParalParserRULE_line)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Variable_def()
		}
		{
			p.SetState(39)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.Matrix_def()
		}
		{
			p.SetState(42)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.Comment_def()
		}
		{
			p.SetState(45)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(47)
			p.Job_def()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(48)
			p.Match(ParalParserNEWLINE)
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

// IVariable_defContext is an interface to support dynamic dispatch.
type IVariable_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Value_expr() IValue_exprContext
	List_expr() IList_exprContext

	// IsVariable_defContext differentiates from other interfaces.
	IsVariable_defContext()
}

type Variable_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariable_defContext() *Variable_defContext {
	var p = new(Variable_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_variable_def
	return p
}

func InitEmptyVariable_defContext(p *Variable_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_variable_def
}

func (*Variable_defContext) IsVariable_defContext() {}

func NewVariable_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_defContext {
	var p = new(Variable_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_variable_def

	return p
}

func (s *Variable_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_defContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *Variable_defContext) Value_expr() IValue_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_exprContext)
}

func (s *Variable_defContext) List_expr() IList_exprContext {
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

func (s *Variable_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterVariable_def(s)
	}
}

func (s *Variable_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitVariable_def(s)
	}
}

func (p *ParalParser) Variable_def() (localctx IVariable_defContext) {
	localctx = NewVariable_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParalParserRULE_variable_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserSTRING, ParalParserSINGLE_QUOTE_STRING, ParalParserNUMBER, ParalParserBOOLEAN, ParalParserDURATION:
		{
			p.SetState(52)
			p.Value_expr()
		}

	case ParalParserLBRACK:
		{
			p.SetState(53)
			p.List_expr()
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

// IMatrix_defContext is an interface to support dynamic dispatch.
type IMatrix_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllList_expr() []IList_exprContext
	List_expr(i int) IList_exprContext
	AllCOLONCOLON() []antlr.TerminalNode
	COLONCOLON(i int) antlr.TerminalNode

	// IsMatrix_defContext differentiates from other interfaces.
	IsMatrix_defContext()
}

type Matrix_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrix_defContext() *Matrix_defContext {
	var p = new(Matrix_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_matrix_def
	return p
}

func InitEmptyMatrix_defContext(p *Matrix_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_matrix_def
}

func (*Matrix_defContext) IsMatrix_defContext() {}

func NewMatrix_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matrix_defContext {
	var p = new(Matrix_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_matrix_def

	return p
}

func (s *Matrix_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Matrix_defContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *Matrix_defContext) AllList_expr() []IList_exprContext {
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

func (s *Matrix_defContext) List_expr(i int) IList_exprContext {
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

func (s *Matrix_defContext) AllCOLONCOLON() []antlr.TerminalNode {
	return s.GetTokens(ParalParserCOLONCOLON)
}

func (s *Matrix_defContext) COLONCOLON(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserCOLONCOLON, i)
}

func (s *Matrix_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matrix_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matrix_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterMatrix_def(s)
	}
}

func (s *Matrix_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitMatrix_def(s)
	}
}

func (p *ParalParser) Matrix_def() (localctx IMatrix_defContext) {
	localctx = NewMatrix_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParalParserRULE_matrix_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.List_expr()
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserCOLONCOLON {
		{
			p.SetState(58)
			p.Match(ParalParserCOLONCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.List_expr()
		}

		p.SetState(64)
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
	AllValue_expr() []IValue_exprContext
	Value_expr(i int) IValue_exprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *List_exprContext) AllValue_expr() []IValue_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValue_exprContext); ok {
			len++
		}
	}

	tst := make([]IValue_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValue_exprContext); ok {
			tst[i] = t.(IValue_exprContext)
			i++
		}
	}

	return tst
}

func (s *List_exprContext) Value_expr(i int) IValue_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_exprContext); ok {
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

	return t.(IValue_exprContext)
}

func (s *List_exprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ParalParserCOMMA)
}

func (s *List_exprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserCOMMA, i)
}

func (s *List_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterList_expr(s)
	}
}

func (s *List_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitList_expr(s)
	}
}

func (p *ParalParser) List_expr() (localctx IList_exprContext) {
	localctx = NewList_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ParalParserRULE_list_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(ParalParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&124) != 0 {
		{
			p.SetState(66)
			p.Value_expr()
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserCOMMA {
			{
				p.SetState(67)
				p.Match(ParalParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(68)
				p.Value_expr()
			}

			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(76)
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

// IComment_defContext is an interface to support dynamic dispatch.
type IComment_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENT() antlr.TerminalNode

	// IsComment_defContext differentiates from other interfaces.
	IsComment_defContext()
}

type Comment_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComment_defContext() *Comment_defContext {
	var p = new(Comment_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_comment_def
	return p
}

func InitEmptyComment_defContext(p *Comment_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_comment_def
}

func (*Comment_defContext) IsComment_defContext() {}

func NewComment_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comment_defContext {
	var p = new(Comment_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_comment_def

	return p
}

func (s *Comment_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Comment_defContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(ParalParserCOMMENT, 0)
}

func (s *Comment_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comment_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comment_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterComment_def(s)
	}
}

func (s *Comment_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitComment_def(s)
	}
}

func (p *ParalParser) Comment_def() (localctx IComment_defContext) {
	localctx = NewComment_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ParalParserRULE_comment_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(ParalParserCOMMENT)
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

// IJob_defContext is an interface to support dynamic dispatch.
type IJob_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllJob_directive_def() []IJob_directive_defContext
	Job_directive_def(i int) IJob_directive_defContext
	AllCmd_expr() []ICmd_exprContext
	Cmd_expr(i int) ICmd_exprContext

	// IsJob_defContext differentiates from other interfaces.
	IsJob_defContext()
}

type Job_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJob_defContext() *Job_defContext {
	var p = new(Job_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_job_def
	return p
}

func InitEmptyJob_defContext(p *Job_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_job_def
}

func (*Job_defContext) IsJob_defContext() {}

func NewJob_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Job_defContext {
	var p = new(Job_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_job_def

	return p
}

func (s *Job_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Job_defContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *Job_defContext) COLON() antlr.TerminalNode {
	return s.GetToken(ParalParserCOLON, 0)
}

func (s *Job_defContext) AllJob_directive_def() []IJob_directive_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJob_directive_defContext); ok {
			len++
		}
	}

	tst := make([]IJob_directive_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJob_directive_defContext); ok {
			tst[i] = t.(IJob_directive_defContext)
			i++
		}
	}

	return tst
}

func (s *Job_defContext) Job_directive_def(i int) IJob_directive_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJob_directive_defContext); ok {
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

	return t.(IJob_directive_defContext)
}

func (s *Job_defContext) AllCmd_expr() []ICmd_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmd_exprContext); ok {
			len++
		}
	}

	tst := make([]ICmd_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmd_exprContext); ok {
			tst[i] = t.(ICmd_exprContext)
			i++
		}
	}

	return tst
}

func (s *Job_defContext) Cmd_expr(i int) ICmd_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmd_exprContext); ok {
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

	return t.(ICmd_exprContext)
}

func (s *Job_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Job_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Job_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterJob_def(s)
	}
}

func (s *Job_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitJob_def(s)
	}
}

func (p *ParalParser) Job_def() (localctx IJob_defContext) {
	localctx = NewJob_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ParalParserRULE_job_def)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserAT {
		{
			p.SetState(80)
			p.Job_directive_def()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(ParalParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(88)
				p.Cmd_expr()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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

// IJob_directive_defContext is an interface to support dynamic dispatch.
type IJob_directive_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	AllJob_directive_value() []IJob_directive_valueContext
	Job_directive_value(i int) IJob_directive_valueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsJob_directive_defContext differentiates from other interfaces.
	IsJob_directive_defContext()
}

type Job_directive_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJob_directive_defContext() *Job_directive_defContext {
	var p = new(Job_directive_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_job_directive_def
	return p
}

func InitEmptyJob_directive_defContext(p *Job_directive_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_job_directive_def
}

func (*Job_directive_defContext) IsJob_directive_defContext() {}

func NewJob_directive_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Job_directive_defContext {
	var p = new(Job_directive_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_job_directive_def

	return p
}

func (s *Job_directive_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Job_directive_defContext) AT() antlr.TerminalNode {
	return s.GetToken(ParalParserAT, 0)
}

func (s *Job_directive_defContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *Job_directive_defContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, 0)
}

func (s *Job_directive_defContext) AllJob_directive_value() []IJob_directive_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJob_directive_valueContext); ok {
			len++
		}
	}

	tst := make([]IJob_directive_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJob_directive_valueContext); ok {
			tst[i] = t.(IJob_directive_valueContext)
			i++
		}
	}

	return tst
}

func (s *Job_directive_defContext) Job_directive_value(i int) IJob_directive_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJob_directive_valueContext); ok {
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

	return t.(IJob_directive_valueContext)
}

func (s *Job_directive_defContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ParalParserCOMMA)
}

func (s *Job_directive_defContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserCOMMA, i)
}

func (s *Job_directive_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Job_directive_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Job_directive_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterJob_directive_def(s)
	}
}

func (s *Job_directive_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitJob_directive_def(s)
	}
}

func (p *ParalParser) Job_directive_def() (localctx IJob_directive_defContext) {
	localctx = NewJob_directive_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParalParserRULE_job_directive_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(ParalParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(ParalParserIDENTIFIER)
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66430) != 0 {
		{
			p.SetState(95)
			p.Job_directive_value()
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserCOMMA {
			{
				p.SetState(96)
				p.Match(ParalParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(97)
				p.Job_directive_value()
			}

			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(105)
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

// IJob_directive_valueContext is an interface to support dynamic dispatch.
type IJob_directive_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value_expr() IValue_exprContext
	IDENTIFIER() antlr.TerminalNode
	MATRIX_REF() antlr.TerminalNode
	REF() antlr.TerminalNode
	VALUE() antlr.TerminalNode

	// IsJob_directive_valueContext differentiates from other interfaces.
	IsJob_directive_valueContext()
}

type Job_directive_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJob_directive_valueContext() *Job_directive_valueContext {
	var p = new(Job_directive_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_job_directive_value
	return p
}

func InitEmptyJob_directive_valueContext(p *Job_directive_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_job_directive_value
}

func (*Job_directive_valueContext) IsJob_directive_valueContext() {}

func NewJob_directive_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Job_directive_valueContext {
	var p = new(Job_directive_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_job_directive_value

	return p
}

func (s *Job_directive_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Job_directive_valueContext) Value_expr() IValue_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_exprContext)
}

func (s *Job_directive_valueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *Job_directive_valueContext) MATRIX_REF() antlr.TerminalNode {
	return s.GetToken(ParalParserMATRIX_REF, 0)
}

func (s *Job_directive_valueContext) REF() antlr.TerminalNode {
	return s.GetToken(ParalParserREF, 0)
}

func (s *Job_directive_valueContext) VALUE() antlr.TerminalNode {
	return s.GetToken(ParalParserVALUE, 0)
}

func (s *Job_directive_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Job_directive_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Job_directive_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterJob_directive_value(s)
	}
}

func (s *Job_directive_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitJob_directive_value(s)
	}
}

func (p *ParalParser) Job_directive_value() (localctx IJob_directive_valueContext) {
	localctx = NewJob_directive_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParalParserRULE_job_directive_value)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserSTRING, ParalParserSINGLE_QUOTE_STRING, ParalParserNUMBER, ParalParserBOOLEAN, ParalParserDURATION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Value_expr()
		}

	case ParalParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Match(ParalParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserMATRIX_REF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.Match(ParalParserMATRIX_REF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserREF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(110)
			p.Match(ParalParserREF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserVALUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(111)
			p.Match(ParalParserVALUE)
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

// ICmd_exprContext is an interface to support dynamic dispatch.
type ICmd_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEWLINE() antlr.TerminalNode
	Cmd_directive() ICmd_directiveContext
	TAB() antlr.TerminalNode
	WS() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllSINGLE_QUOTE_STRING() []antlr.TerminalNode
	SINGLE_QUOTE_STRING(i int) antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	AllMATRIX_REF() []antlr.TerminalNode
	MATRIX_REF(i int) antlr.TerminalNode
	AllARG() []antlr.TerminalNode
	ARG(i int) antlr.TerminalNode
	AllSINGLE_DASH_ARG() []antlr.TerminalNode
	SINGLE_DASH_ARG(i int) antlr.TerminalNode
	AllEQUALS() []antlr.TerminalNode
	EQUALS(i int) antlr.TerminalNode
	AllREF() []antlr.TerminalNode
	REF(i int) antlr.TerminalNode
	AllURL() []antlr.TerminalNode
	URL(i int) antlr.TerminalNode
	AllVALUE() []antlr.TerminalNode
	VALUE(i int) antlr.TerminalNode
	AllBSL_NEWLINE() []antlr.TerminalNode
	BSL_NEWLINE(i int) antlr.TerminalNode

	// IsCmd_exprContext differentiates from other interfaces.
	IsCmd_exprContext()
}

type Cmd_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmd_exprContext() *Cmd_exprContext {
	var p = new(Cmd_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_cmd_expr
	return p
}

func InitEmptyCmd_exprContext(p *Cmd_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_cmd_expr
}

func (*Cmd_exprContext) IsCmd_exprContext() {}

func NewCmd_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cmd_exprContext {
	var p = new(Cmd_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_cmd_expr

	return p
}

func (s *Cmd_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Cmd_exprContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, 0)
}

func (s *Cmd_exprContext) Cmd_directive() ICmd_directiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmd_directiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmd_directiveContext)
}

func (s *Cmd_exprContext) TAB() antlr.TerminalNode {
	return s.GetToken(ParalParserTAB, 0)
}

func (s *Cmd_exprContext) WS() antlr.TerminalNode {
	return s.GetToken(ParalParserWS, 0)
}

func (s *Cmd_exprContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ParalParserIDENTIFIER)
}

func (s *Cmd_exprContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, i)
}

func (s *Cmd_exprContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ParalParserSTRING)
}

func (s *Cmd_exprContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserSTRING, i)
}

func (s *Cmd_exprContext) AllSINGLE_QUOTE_STRING() []antlr.TerminalNode {
	return s.GetTokens(ParalParserSINGLE_QUOTE_STRING)
}

func (s *Cmd_exprContext) SINGLE_QUOTE_STRING(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserSINGLE_QUOTE_STRING, i)
}

func (s *Cmd_exprContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNUMBER)
}

func (s *Cmd_exprContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNUMBER, i)
}

func (s *Cmd_exprContext) AllMATRIX_REF() []antlr.TerminalNode {
	return s.GetTokens(ParalParserMATRIX_REF)
}

func (s *Cmd_exprContext) MATRIX_REF(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserMATRIX_REF, i)
}

func (s *Cmd_exprContext) AllARG() []antlr.TerminalNode {
	return s.GetTokens(ParalParserARG)
}

func (s *Cmd_exprContext) ARG(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserARG, i)
}

func (s *Cmd_exprContext) AllSINGLE_DASH_ARG() []antlr.TerminalNode {
	return s.GetTokens(ParalParserSINGLE_DASH_ARG)
}

func (s *Cmd_exprContext) SINGLE_DASH_ARG(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserSINGLE_DASH_ARG, i)
}

func (s *Cmd_exprContext) AllEQUALS() []antlr.TerminalNode {
	return s.GetTokens(ParalParserEQUALS)
}

func (s *Cmd_exprContext) EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserEQUALS, i)
}

func (s *Cmd_exprContext) AllREF() []antlr.TerminalNode {
	return s.GetTokens(ParalParserREF)
}

func (s *Cmd_exprContext) REF(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserREF, i)
}

func (s *Cmd_exprContext) AllURL() []antlr.TerminalNode {
	return s.GetTokens(ParalParserURL)
}

func (s *Cmd_exprContext) URL(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserURL, i)
}

func (s *Cmd_exprContext) AllVALUE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserVALUE)
}

func (s *Cmd_exprContext) VALUE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserVALUE, i)
}

func (s *Cmd_exprContext) AllBSL_NEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserBSL_NEWLINE)
}

func (s *Cmd_exprContext) BSL_NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserBSL_NEWLINE, i)
}

func (s *Cmd_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cmd_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cmd_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterCmd_expr(s)
	}
}

func (s *Cmd_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitCmd_expr(s)
	}
}

func (p *ParalParser) Cmd_expr() (localctx ICmd_exprContext) {
	localctx = NewCmd_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ParalParserRULE_cmd_expr)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(ParalParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ParalParserTAB || _la == ParalParserWS {
		{
			p.SetState(115)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ParalParserTAB || _la == ParalParserWS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(118)
			p.Cmd_directive()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(121)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&237342) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
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

// ICmd_directiveContext is an interface to support dynamic dispatch.
type ICmd_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Cmd_directive_iden() ICmd_directive_idenContext
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	AllCmd_directive_value() []ICmd_directive_valueContext
	Cmd_directive_value(i int) ICmd_directive_valueContext

	// IsCmd_directiveContext differentiates from other interfaces.
	IsCmd_directiveContext()
}

type Cmd_directiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmd_directiveContext() *Cmd_directiveContext {
	var p = new(Cmd_directiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_cmd_directive
	return p
}

func InitEmptyCmd_directiveContext(p *Cmd_directiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_cmd_directive
}

func (*Cmd_directiveContext) IsCmd_directiveContext() {}

func NewCmd_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cmd_directiveContext {
	var p = new(Cmd_directiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_cmd_directive

	return p
}

func (s *Cmd_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Cmd_directiveContext) Cmd_directive_iden() ICmd_directive_idenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmd_directive_idenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmd_directive_idenContext)
}

func (s *Cmd_directiveContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(ParalParserLPAREN)
}

func (s *Cmd_directiveContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserLPAREN, i)
}

func (s *Cmd_directiveContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(ParalParserRPAREN)
}

func (s *Cmd_directiveContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserRPAREN, i)
}

func (s *Cmd_directiveContext) AllCmd_directive_value() []ICmd_directive_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmd_directive_valueContext); ok {
			len++
		}
	}

	tst := make([]ICmd_directive_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmd_directive_valueContext); ok {
			tst[i] = t.(ICmd_directive_valueContext)
			i++
		}
	}

	return tst
}

func (s *Cmd_directiveContext) Cmd_directive_value(i int) ICmd_directive_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmd_directive_valueContext); ok {
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

	return t.(ICmd_directive_valueContext)
}

func (s *Cmd_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cmd_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cmd_directiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterCmd_directive(s)
	}
}

func (s *Cmd_directiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitCmd_directive(s)
	}
}

func (p *ParalParser) Cmd_directive() (localctx ICmd_directiveContext) {
	localctx = NewCmd_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParalParserRULE_cmd_directive)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Cmd_directive_iden()
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserLPAREN {
		{
			p.SetState(128)
			p.Match(ParalParserLPAREN)
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

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66430) != 0 {
			{
				p.SetState(129)
				p.Cmd_directive_value()
			}

		}
		{
			p.SetState(132)
			p.Match(ParalParserRPAREN)
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

// ICmd_directive_idenContext is an interface to support dynamic dispatch.
type ICmd_directive_idenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsCmd_directive_idenContext differentiates from other interfaces.
	IsCmd_directive_idenContext()
}

type Cmd_directive_idenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmd_directive_idenContext() *Cmd_directive_idenContext {
	var p = new(Cmd_directive_idenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_cmd_directive_iden
	return p
}

func InitEmptyCmd_directive_idenContext(p *Cmd_directive_idenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_cmd_directive_iden
}

func (*Cmd_directive_idenContext) IsCmd_directive_idenContext() {}

func NewCmd_directive_idenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cmd_directive_idenContext {
	var p = new(Cmd_directive_idenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_cmd_directive_iden

	return p
}

func (s *Cmd_directive_idenContext) GetParser() antlr.Parser { return s.parser }

func (s *Cmd_directive_idenContext) AT() antlr.TerminalNode {
	return s.GetToken(ParalParserAT, 0)
}

func (s *Cmd_directive_idenContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *Cmd_directive_idenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cmd_directive_idenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cmd_directive_idenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterCmd_directive_iden(s)
	}
}

func (s *Cmd_directive_idenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitCmd_directive_iden(s)
	}
}

func (p *ParalParser) Cmd_directive_iden() (localctx ICmd_directive_idenContext) {
	localctx = NewCmd_directive_idenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ParalParserRULE_cmd_directive_iden)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(ParalParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(ParalParserIDENTIFIER)
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

// ICmd_directive_valueContext is an interface to support dynamic dispatch.
type ICmd_directive_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDirective_arg() []IDirective_argContext
	Directive_arg(i int) IDirective_argContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCmd_directive_valueContext differentiates from other interfaces.
	IsCmd_directive_valueContext()
}

type Cmd_directive_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmd_directive_valueContext() *Cmd_directive_valueContext {
	var p = new(Cmd_directive_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_cmd_directive_value
	return p
}

func InitEmptyCmd_directive_valueContext(p *Cmd_directive_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_cmd_directive_value
}

func (*Cmd_directive_valueContext) IsCmd_directive_valueContext() {}

func NewCmd_directive_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cmd_directive_valueContext {
	var p = new(Cmd_directive_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_cmd_directive_value

	return p
}

func (s *Cmd_directive_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Cmd_directive_valueContext) AllDirective_arg() []IDirective_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDirective_argContext); ok {
			len++
		}
	}

	tst := make([]IDirective_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDirective_argContext); ok {
			tst[i] = t.(IDirective_argContext)
			i++
		}
	}

	return tst
}

func (s *Cmd_directive_valueContext) Directive_arg(i int) IDirective_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirective_argContext); ok {
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

	return t.(IDirective_argContext)
}

func (s *Cmd_directive_valueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ParalParserCOMMA)
}

func (s *Cmd_directive_valueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserCOMMA, i)
}

func (s *Cmd_directive_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cmd_directive_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cmd_directive_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterCmd_directive_value(s)
	}
}

func (s *Cmd_directive_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitCmd_directive_value(s)
	}
}

func (p *ParalParser) Cmd_directive_value() (localctx ICmd_directive_valueContext) {
	localctx = NewCmd_directive_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ParalParserRULE_cmd_directive_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Directive_arg()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserCOMMA {
		{
			p.SetState(142)
			p.Match(ParalParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Directive_arg()
		}

		p.SetState(148)
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

// IDirective_argContext is an interface to support dynamic dispatch.
type IDirective_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value_expr() IValue_exprContext
	IDENTIFIER() antlr.TerminalNode
	MATRIX_REF() antlr.TerminalNode
	REF() antlr.TerminalNode
	VALUE() antlr.TerminalNode

	// IsDirective_argContext differentiates from other interfaces.
	IsDirective_argContext()
}

type Directive_argContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirective_argContext() *Directive_argContext {
	var p = new(Directive_argContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_directive_arg
	return p
}

func InitEmptyDirective_argContext(p *Directive_argContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_directive_arg
}

func (*Directive_argContext) IsDirective_argContext() {}

func NewDirective_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Directive_argContext {
	var p = new(Directive_argContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_directive_arg

	return p
}

func (s *Directive_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Directive_argContext) Value_expr() IValue_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_exprContext)
}

func (s *Directive_argContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *Directive_argContext) MATRIX_REF() antlr.TerminalNode {
	return s.GetToken(ParalParserMATRIX_REF, 0)
}

func (s *Directive_argContext) REF() antlr.TerminalNode {
	return s.GetToken(ParalParserREF, 0)
}

func (s *Directive_argContext) VALUE() antlr.TerminalNode {
	return s.GetToken(ParalParserVALUE, 0)
}

func (s *Directive_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Directive_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Directive_argContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterDirective_arg(s)
	}
}

func (s *Directive_argContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitDirective_arg(s)
	}
}

func (p *ParalParser) Directive_arg() (localctx IDirective_argContext) {
	localctx = NewDirective_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ParalParserRULE_directive_arg)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserSTRING, ParalParserSINGLE_QUOTE_STRING, ParalParserNUMBER, ParalParserBOOLEAN, ParalParserDURATION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Value_expr()
		}

	case ParalParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Match(ParalParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserMATRIX_REF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.Match(ParalParserMATRIX_REF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserREF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(152)
			p.Match(ParalParserREF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserVALUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(153)
			p.Match(ParalParserVALUE)
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

// IValue_exprContext is an interface to support dynamic dispatch.
type IValue_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValue_exprContext differentiates from other interfaces.
	IsValue_exprContext()
}

type Value_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_exprContext() *Value_exprContext {
	var p = new(Value_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_value_expr
	return p
}

func InitEmptyValue_exprContext(p *Value_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_value_expr
}

func (*Value_exprContext) IsValue_exprContext() {}

func NewValue_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_exprContext {
	var p = new(Value_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_value_expr

	return p
}

func (s *Value_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_exprContext) CopyAll(ctx *Value_exprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Value_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DurationContext struct {
	Value_exprContext
}

func NewDurationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DurationContext {
	var p = new(DurationContext)

	InitEmptyValue_exprContext(&p.Value_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Value_exprContext))

	return p
}

func (s *DurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationContext) DURATION() antlr.TerminalNode {
	return s.GetToken(ParalParserDURATION, 0)
}

func (s *DurationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterDuration(s)
	}
}

func (s *DurationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitDuration(s)
	}
}

type SingleQuoteStringContext struct {
	Value_exprContext
}

func NewSingleQuoteStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleQuoteStringContext {
	var p = new(SingleQuoteStringContext)

	InitEmptyValue_exprContext(&p.Value_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Value_exprContext))

	return p
}

func (s *SingleQuoteStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleQuoteStringContext) SINGLE_QUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(ParalParserSINGLE_QUOTE_STRING, 0)
}

func (s *SingleQuoteStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterSingleQuoteString(s)
	}
}

func (s *SingleQuoteStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitSingleQuoteString(s)
	}
}

type NumberContext struct {
	Value_exprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	InitEmptyValue_exprContext(&p.Value_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Value_exprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ParalParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitNumber(s)
	}
}

type StringContext struct {
	Value_exprContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyValue_exprContext(&p.Value_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Value_exprContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(ParalParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitString(s)
	}
}

type BoolContext struct {
	Value_exprContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	InitEmptyValue_exprContext(&p.Value_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Value_exprContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ParalParserBOOLEAN, 0)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitBool(s)
	}
}

func (p *ParalParser) Value_expr() (localctx IValue_exprContext) {
	localctx = NewValue_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ParalParserRULE_value_expr)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.Match(ParalParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserSINGLE_QUOTE_STRING:
		localctx = NewSingleQuoteStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)
			p.Match(ParalParserSINGLE_QUOTE_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(158)
			p.Match(ParalParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserBOOLEAN:
		localctx = NewBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(159)
			p.Match(ParalParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserDURATION:
		localctx = NewDurationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(160)
			p.Match(ParalParserDURATION)
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
