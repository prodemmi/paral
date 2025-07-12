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
		"", "'='", "':'", "'>'", "", "", "", "", "", "", "", "", "'('", "')'",
		"'['", "']'", "','", "'::'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "ARROW", "STRING", "SINGLE_QUOTE_STRING", "NUMBER", "BOOLEAN",
		"DURATION", "IDENTIFIER", "FLAG_NAME", "DIRECTIVE", "LRBRACK", "RRBRACK",
		"LBRACK", "RBRACK", "COMMA", "COLONCOLON", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"program", "line", "variable_def", "matrix_def", "job_def", "cmd_expr",
		"cmd_value", "directive_expr", "directive_args_expr", "flag_expr", "value_expr",
		"list_expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 142, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 37, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		43, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 52, 8, 3, 10,
		3, 12, 3, 55, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 63, 8, 4,
		11, 4, 12, 4, 64, 1, 5, 1, 5, 4, 5, 69, 8, 5, 11, 5, 12, 5, 70, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 83, 8, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 90, 8, 6, 10, 6, 12, 6, 93, 9, 6, 1, 7, 1,
		7, 5, 7, 97, 8, 7, 10, 7, 12, 7, 100, 9, 7, 1, 8, 1, 8, 3, 8, 104, 8, 8,
		1, 8, 1, 8, 5, 8, 108, 8, 8, 10, 8, 12, 8, 111, 9, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 119, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 127, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 133, 8, 11, 10,
		11, 12, 11, 136, 9, 11, 3, 11, 138, 8, 11, 1, 11, 1, 11, 1, 11, 0, 1, 12,
		12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 0, 152, 0, 27, 1, 0,
		0, 0, 2, 36, 1, 0, 0, 0, 4, 38, 1, 0, 0, 0, 6, 46, 1, 0, 0, 0, 8, 58, 1,
		0, 0, 0, 10, 66, 1, 0, 0, 0, 12, 82, 1, 0, 0, 0, 14, 94, 1, 0, 0, 0, 16,
		101, 1, 0, 0, 0, 18, 118, 1, 0, 0, 0, 20, 126, 1, 0, 0, 0, 22, 128, 1,
		0, 0, 0, 24, 26, 3, 2, 1, 0, 25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27,
		25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 30, 1, 0, 0, 0, 29, 27, 1, 0, 0,
		0, 30, 31, 5, 0, 0, 1, 31, 1, 1, 0, 0, 0, 32, 37, 3, 4, 2, 0, 33, 37, 3,
		6, 3, 0, 34, 37, 3, 8, 4, 0, 35, 37, 5, 18, 0, 0, 36, 32, 1, 0, 0, 0, 36,
		33, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 35, 1, 0, 0, 0, 37, 3, 1, 0, 0,
		0, 38, 39, 5, 9, 0, 0, 39, 42, 5, 1, 0, 0, 40, 43, 3, 20, 10, 0, 41, 43,
		3, 22, 11, 0, 42, 40, 1, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0,
		0, 44, 45, 5, 18, 0, 0, 45, 5, 1, 0, 0, 0, 46, 47, 5, 9, 0, 0, 47, 48,
		5, 1, 0, 0, 48, 53, 3, 22, 11, 0, 49, 50, 5, 17, 0, 0, 50, 52, 3, 22, 11,
		0, 51, 49, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54,
		1, 0, 0, 0, 54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 5, 18, 0, 0,
		57, 7, 1, 0, 0, 0, 58, 59, 5, 9, 0, 0, 59, 60, 5, 2, 0, 0, 60, 62, 5, 18,
		0, 0, 61, 63, 3, 10, 5, 0, 62, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64,
		62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 9, 1, 0, 0, 0, 66, 68, 5, 3, 0,
		0, 67, 69, 3, 12, 6, 0, 68, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 68,
		1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 5, 18, 0, 0,
		73, 11, 1, 0, 0, 0, 74, 75, 6, 6, -1, 0, 75, 83, 3, 20, 10, 0, 76, 83,
		3, 14, 7, 0, 77, 78, 5, 12, 0, 0, 78, 79, 3, 12, 6, 0, 79, 80, 5, 13, 0,
		0, 80, 83, 1, 0, 0, 0, 81, 83, 3, 18, 9, 0, 82, 74, 1, 0, 0, 0, 82, 76,
		1, 0, 0, 0, 82, 77, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 91, 1, 0, 0, 0,
		84, 85, 10, 3, 0, 0, 85, 86, 5, 14, 0, 0, 86, 87, 3, 12, 6, 0, 87, 88,
		5, 15, 0, 0, 88, 90, 1, 0, 0, 0, 89, 84, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0,
		91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 13, 1, 0, 0, 0, 93, 91, 1,
		0, 0, 0, 94, 98, 5, 11, 0, 0, 95, 97, 3, 16, 8, 0, 96, 95, 1, 0, 0, 0,
		97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 15, 1,
		0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 103, 5, 12, 0, 0, 102, 104, 3, 20, 10,
		0, 103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 109, 1, 0, 0, 0, 105,
		106, 5, 16, 0, 0, 106, 108, 3, 20, 10, 0, 107, 105, 1, 0, 0, 0, 108, 111,
		1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 1, 0,
		0, 0, 111, 109, 1, 0, 0, 0, 112, 113, 5, 13, 0, 0, 113, 17, 1, 0, 0, 0,
		114, 115, 5, 10, 0, 0, 115, 116, 5, 1, 0, 0, 116, 119, 3, 20, 10, 0, 117,
		119, 5, 10, 0, 0, 118, 114, 1, 0, 0, 0, 118, 117, 1, 0, 0, 0, 119, 19,
		1, 0, 0, 0, 120, 127, 5, 9, 0, 0, 121, 127, 5, 4, 0, 0, 122, 127, 5, 5,
		0, 0, 123, 127, 5, 6, 0, 0, 124, 127, 5, 7, 0, 0, 125, 127, 5, 8, 0, 0,
		126, 120, 1, 0, 0, 0, 126, 121, 1, 0, 0, 0, 126, 122, 1, 0, 0, 0, 126,
		123, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 21, 1,
		0, 0, 0, 128, 137, 5, 14, 0, 0, 129, 134, 3, 20, 10, 0, 130, 131, 5, 16,
		0, 0, 131, 133, 3, 20, 10, 0, 132, 130, 1, 0, 0, 0, 133, 136, 1, 0, 0,
		0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136,
		134, 1, 0, 0, 0, 137, 129, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139,
		1, 0, 0, 0, 139, 140, 5, 15, 0, 0, 140, 23, 1, 0, 0, 0, 15, 27, 36, 42,
		53, 64, 70, 82, 91, 98, 103, 109, 118, 126, 134, 137,
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
	ParalParserT__0                = 1
	ParalParserT__1                = 2
	ParalParserARROW               = 3
	ParalParserSTRING              = 4
	ParalParserSINGLE_QUOTE_STRING = 5
	ParalParserNUMBER              = 6
	ParalParserBOOLEAN             = 7
	ParalParserDURATION            = 8
	ParalParserIDENTIFIER          = 9
	ParalParserFLAG_NAME           = 10
	ParalParserDIRECTIVE           = 11
	ParalParserLRBRACK             = 12
	ParalParserRRBRACK             = 13
	ParalParserLBRACK              = 14
	ParalParserRBRACK              = 15
	ParalParserCOMMA               = 16
	ParalParserCOLONCOLON          = 17
	ParalParserNEWLINE             = 18
	ParalParserWS                  = 19
)

// ParalParser rules.
const (
	ParalParserRULE_program             = 0
	ParalParserRULE_line                = 1
	ParalParserRULE_variable_def        = 2
	ParalParserRULE_matrix_def          = 3
	ParalParserRULE_job_def             = 4
	ParalParserRULE_cmd_expr            = 5
	ParalParserRULE_cmd_value           = 6
	ParalParserRULE_directive_expr      = 7
	ParalParserRULE_directive_args_expr = 8
	ParalParserRULE_flag_expr           = 9
	ParalParserRULE_value_expr          = 10
	ParalParserRULE_list_expr           = 11
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserIDENTIFIER || _la == ParalParserNEWLINE {
		{
			p.SetState(24)
			p.Line()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
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
	Matrix_def() IMatrix_defContext
	Job_def() IJob_defContext
	NEWLINE() antlr.TerminalNode

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

func (s *LineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, 0)
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
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Variable_def()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Matrix_def()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.Job_def()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(35)
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
	NEWLINE() antlr.TerminalNode
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

func (s *Variable_defContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, 0)
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
		p.SetState(38)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(ParalParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserSTRING, ParalParserSINGLE_QUOTE_STRING, ParalParserNUMBER, ParalParserBOOLEAN, ParalParserDURATION, ParalParserIDENTIFIER:
		{
			p.SetState(40)
			p.Value_expr()
		}

	case ParalParserLBRACK:
		{
			p.SetState(41)
			p.List_expr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(44)
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

// IMatrix_defContext is an interface to support dynamic dispatch.
type IMatrix_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllList_expr() []IList_exprContext
	List_expr(i int) IList_exprContext
	NEWLINE() antlr.TerminalNode
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

func (s *Matrix_defContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, 0)
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
		p.SetState(46)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(ParalParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.List_expr()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserCOLONCOLON {
		{
			p.SetState(49)
			p.Match(ParalParserCOLONCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.List_expr()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
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

// IJob_defContext is an interface to support dynamic dispatch.
type IJob_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
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

func (s *Job_defContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, 0)
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
	p.EnterRule(localctx, 8, ParalParserRULE_job_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(ParalParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(ParalParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(ParalParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ParalParserARROW {
		{
			p.SetState(61)
			p.Cmd_expr()
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

// ICmd_exprContext is an interface to support dynamic dispatch.
type ICmd_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARROW() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	AllCmd_value() []ICmd_valueContext
	Cmd_value(i int) ICmd_valueContext

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

func (s *Cmd_exprContext) ARROW() antlr.TerminalNode {
	return s.GetToken(ParalParserARROW, 0)
}

func (s *Cmd_exprContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, 0)
}

func (s *Cmd_exprContext) AllCmd_value() []ICmd_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmd_valueContext); ok {
			len++
		}
	}

	tst := make([]ICmd_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmd_valueContext); ok {
			tst[i] = t.(ICmd_valueContext)
			i++
		}
	}

	return tst
}

func (s *Cmd_exprContext) Cmd_value(i int) ICmd_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmd_valueContext); ok {
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

	return t.(ICmd_valueContext)
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
	p.EnterRule(localctx, 10, ParalParserRULE_cmd_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(ParalParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8176) != 0) {
		{
			p.SetState(67)
			p.cmd_value(0)
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
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

// ICmd_valueContext is an interface to support dynamic dispatch.
type ICmd_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value_expr() IValue_exprContext
	Directive_expr() IDirective_exprContext
	LRBRACK() antlr.TerminalNode
	AllCmd_value() []ICmd_valueContext
	Cmd_value(i int) ICmd_valueContext
	RRBRACK() antlr.TerminalNode
	Flag_expr() IFlag_exprContext
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode

	// IsCmd_valueContext differentiates from other interfaces.
	IsCmd_valueContext()
}

type Cmd_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmd_valueContext() *Cmd_valueContext {
	var p = new(Cmd_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_cmd_value
	return p
}

func InitEmptyCmd_valueContext(p *Cmd_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_cmd_value
}

func (*Cmd_valueContext) IsCmd_valueContext() {}

func NewCmd_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cmd_valueContext {
	var p = new(Cmd_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_cmd_value

	return p
}

func (s *Cmd_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Cmd_valueContext) Value_expr() IValue_exprContext {
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

func (s *Cmd_valueContext) Directive_expr() IDirective_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirective_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirective_exprContext)
}

func (s *Cmd_valueContext) LRBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserLRBRACK, 0)
}

func (s *Cmd_valueContext) AllCmd_value() []ICmd_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmd_valueContext); ok {
			len++
		}
	}

	tst := make([]ICmd_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmd_valueContext); ok {
			tst[i] = t.(ICmd_valueContext)
			i++
		}
	}

	return tst
}

func (s *Cmd_valueContext) Cmd_value(i int) ICmd_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmd_valueContext); ok {
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

	return t.(ICmd_valueContext)
}

func (s *Cmd_valueContext) RRBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRRBRACK, 0)
}

func (s *Cmd_valueContext) Flag_expr() IFlag_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFlag_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFlag_exprContext)
}

func (s *Cmd_valueContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserLBRACK, 0)
}

func (s *Cmd_valueContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRBRACK, 0)
}

func (s *Cmd_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cmd_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cmd_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterCmd_value(s)
	}
}

func (s *Cmd_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitCmd_value(s)
	}
}

func (p *ParalParser) Cmd_value() (localctx ICmd_valueContext) {
	return p.cmd_value(0)
}

func (p *ParalParser) cmd_value(_p int) (localctx ICmd_valueContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCmd_valueContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICmd_valueContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, ParalParserRULE_cmd_value, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserSTRING, ParalParserSINGLE_QUOTE_STRING, ParalParserNUMBER, ParalParserBOOLEAN, ParalParserDURATION, ParalParserIDENTIFIER:
		{
			p.SetState(75)
			p.Value_expr()
		}

	case ParalParserDIRECTIVE:
		{
			p.SetState(76)
			p.Directive_expr()
		}

	case ParalParserLRBRACK:
		{
			p.SetState(77)
			p.Match(ParalParserLRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.cmd_value(0)
		}
		{
			p.SetState(79)
			p.Match(ParalParserRRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserFLAG_NAME:
		{
			p.SetState(81)
			p.Flag_expr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(91)
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
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCmd_valueContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ParalParserRULE_cmd_value)
			p.SetState(84)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(85)
				p.Match(ParalParserLBRACK)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(86)
				p.cmd_value(0)
			}
			{
				p.SetState(87)
				p.Match(ParalParserRBRACK)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDirective_exprContext is an interface to support dynamic dispatch.
type IDirective_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DIRECTIVE() antlr.TerminalNode
	AllDirective_args_expr() []IDirective_args_exprContext
	Directive_args_expr(i int) IDirective_args_exprContext

	// IsDirective_exprContext differentiates from other interfaces.
	IsDirective_exprContext()
}

type Directive_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirective_exprContext() *Directive_exprContext {
	var p = new(Directive_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_directive_expr
	return p
}

func InitEmptyDirective_exprContext(p *Directive_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_directive_expr
}

func (*Directive_exprContext) IsDirective_exprContext() {}

func NewDirective_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Directive_exprContext {
	var p = new(Directive_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_directive_expr

	return p
}

func (s *Directive_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Directive_exprContext) DIRECTIVE() antlr.TerminalNode {
	return s.GetToken(ParalParserDIRECTIVE, 0)
}

func (s *Directive_exprContext) AllDirective_args_expr() []IDirective_args_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDirective_args_exprContext); ok {
			len++
		}
	}

	tst := make([]IDirective_args_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDirective_args_exprContext); ok {
			tst[i] = t.(IDirective_args_exprContext)
			i++
		}
	}

	return tst
}

func (s *Directive_exprContext) Directive_args_expr(i int) IDirective_args_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirective_args_exprContext); ok {
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

	return t.(IDirective_args_exprContext)
}

func (s *Directive_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Directive_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Directive_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterDirective_expr(s)
	}
}

func (s *Directive_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitDirective_expr(s)
	}
}

func (p *ParalParser) Directive_expr() (localctx IDirective_exprContext) {
	localctx = NewDirective_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParalParserRULE_directive_expr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(ParalParserDIRECTIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(95)
				p.Directive_args_expr()
			}

		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
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

// IDirective_args_exprContext is an interface to support dynamic dispatch.
type IDirective_args_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LRBRACK() antlr.TerminalNode
	RRBRACK() antlr.TerminalNode
	AllValue_expr() []IValue_exprContext
	Value_expr(i int) IValue_exprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDirective_args_exprContext differentiates from other interfaces.
	IsDirective_args_exprContext()
}

type Directive_args_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirective_args_exprContext() *Directive_args_exprContext {
	var p = new(Directive_args_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_directive_args_expr
	return p
}

func InitEmptyDirective_args_exprContext(p *Directive_args_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_directive_args_expr
}

func (*Directive_args_exprContext) IsDirective_args_exprContext() {}

func NewDirective_args_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Directive_args_exprContext {
	var p = new(Directive_args_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_directive_args_expr

	return p
}

func (s *Directive_args_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Directive_args_exprContext) LRBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserLRBRACK, 0)
}

func (s *Directive_args_exprContext) RRBRACK() antlr.TerminalNode {
	return s.GetToken(ParalParserRRBRACK, 0)
}

func (s *Directive_args_exprContext) AllValue_expr() []IValue_exprContext {
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

func (s *Directive_args_exprContext) Value_expr(i int) IValue_exprContext {
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

func (s *Directive_args_exprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ParalParserCOMMA)
}

func (s *Directive_args_exprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserCOMMA, i)
}

func (s *Directive_args_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Directive_args_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Directive_args_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterDirective_args_expr(s)
	}
}

func (s *Directive_args_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitDirective_args_expr(s)
	}
}

func (p *ParalParser) Directive_args_expr() (localctx IDirective_args_exprContext) {
	localctx = NewDirective_args_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParalParserRULE_directive_args_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(ParalParserLRBRACK)
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1008) != 0 {
		{
			p.SetState(102)
			p.Value_expr()
		}

	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserCOMMA {
		{
			p.SetState(105)
			p.Match(ParalParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Value_expr()
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(112)
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

// IFlag_exprContext is an interface to support dynamic dispatch.
type IFlag_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFlag_exprContext differentiates from other interfaces.
	IsFlag_exprContext()
}

type Flag_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlag_exprContext() *Flag_exprContext {
	var p = new(Flag_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_flag_expr
	return p
}

func InitEmptyFlag_exprContext(p *Flag_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_flag_expr
}

func (*Flag_exprContext) IsFlag_exprContext() {}

func NewFlag_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Flag_exprContext {
	var p = new(Flag_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_flag_expr

	return p
}

func (s *Flag_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Flag_exprContext) CopyAll(ctx *Flag_exprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Flag_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Flag_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FlagWithValueContext struct {
	Flag_exprContext
}

func NewFlagWithValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FlagWithValueContext {
	var p = new(FlagWithValueContext)

	InitEmptyFlag_exprContext(&p.Flag_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Flag_exprContext))

	return p
}

func (s *FlagWithValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlagWithValueContext) FLAG_NAME() antlr.TerminalNode {
	return s.GetToken(ParalParserFLAG_NAME, 0)
}

func (s *FlagWithValueContext) Value_expr() IValue_exprContext {
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

func (s *FlagWithValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterFlagWithValue(s)
	}
}

func (s *FlagWithValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitFlagWithValue(s)
	}
}

type FlagAloneContext struct {
	Flag_exprContext
}

func NewFlagAloneContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FlagAloneContext {
	var p = new(FlagAloneContext)

	InitEmptyFlag_exprContext(&p.Flag_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Flag_exprContext))

	return p
}

func (s *FlagAloneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlagAloneContext) FLAG_NAME() antlr.TerminalNode {
	return s.GetToken(ParalParserFLAG_NAME, 0)
}

func (s *FlagAloneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterFlagAlone(s)
	}
}

func (s *FlagAloneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitFlagAlone(s)
	}
}

func (p *ParalParser) Flag_expr() (localctx IFlag_exprContext) {
	localctx = NewFlag_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ParalParserRULE_flag_expr)
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFlagWithValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Match(ParalParserFLAG_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Match(ParalParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Value_expr()
		}

	case 2:
		localctx = NewFlagAloneContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(ParalParserFLAG_NAME)
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

type IdentifierContext struct {
	Value_exprContext
}

func NewIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierContext {
	var p = new(IdentifierContext)

	InitEmptyValue_exprContext(&p.Value_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Value_exprContext))

	return p
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENTIFIER, 0)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitIdentifier(s)
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
	p.EnterRule(localctx, 20, ParalParserRULE_value_expr)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserIDENTIFIER:
		localctx = NewIdentifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(ParalParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.Match(ParalParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserSINGLE_QUOTE_STRING:
		localctx = NewSingleQuoteStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)
			p.Match(ParalParserSINGLE_QUOTE_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)
			p.Match(ParalParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserBOOLEAN:
		localctx = NewBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(124)
			p.Match(ParalParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParalParserDURATION:
		localctx = NewDurationContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(125)
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
	p.EnterRule(localctx, 22, ParalParserRULE_list_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(ParalParserLBRACK)
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1008) != 0 {
		{
			p.SetState(129)
			p.Value_expr()
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalParserCOMMA {
			{
				p.SetState(130)
				p.Match(ParalParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(131)
				p.Value_expr()
			}

			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(139)
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

func (p *ParalParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *Cmd_valueContext = nil
		if localctx != nil {
			t = localctx.(*Cmd_valueContext)
		}
		return p.Cmd_value_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ParalParser) Cmd_value_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
