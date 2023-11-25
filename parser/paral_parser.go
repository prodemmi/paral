// Code generated from Paral.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
	staticData.SymbolicNames = []string{
		"", "IDENT", "VALUE", "NEWLINE", "WS", "VARIABLE", "EXECUTE",
	}
	staticData.RuleNames = []string{
		"start", "prog", "variable", "execute",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 6, 28, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1, 0,
		1, 0, 5, 0, 12, 8, 0, 10, 0, 12, 0, 15, 9, 0, 1, 1, 1, 1, 3, 1, 19, 8,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 0, 0, 4, 0, 2, 4, 6,
		0, 0, 25, 0, 13, 1, 0, 0, 0, 2, 18, 1, 0, 0, 0, 4, 20, 1, 0, 0, 0, 6, 24,
		1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 10, 5, 3, 0, 0, 10, 12, 1, 0, 0, 0, 11,
		8, 1, 0, 0, 0, 12, 15, 1, 0, 0, 0, 13, 11, 1, 0, 0, 0, 13, 14, 1, 0, 0,
		0, 14, 1, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 16, 19, 3, 4, 2, 0, 17, 19, 3,
		6, 3, 0, 18, 16, 1, 0, 0, 0, 18, 17, 1, 0, 0, 0, 19, 3, 1, 0, 0, 0, 20,
		21, 5, 5, 0, 0, 21, 22, 5, 1, 0, 0, 22, 23, 5, 2, 0, 0, 23, 5, 1, 0, 0,
		0, 24, 25, 5, 6, 0, 0, 25, 26, 5, 2, 0, 0, 26, 7, 1, 0, 0, 0, 2, 13, 18,
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
	ParalParserEOF      = antlr.TokenEOF
	ParalParserIDENT    = 1
	ParalParserVALUE    = 2
	ParalParserNEWLINE  = 3
	ParalParserWS       = 4
	ParalParserVARIABLE = 5
	ParalParserEXECUTE  = 6
)

// ParalParser rules.
const (
	ParalParserRULE_start    = 0
	ParalParserRULE_prog     = 1
	ParalParserRULE_variable = 2
	ParalParserRULE_execute  = 3
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllProg() []IProgContext
	Prog(i int) IProgContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) AllProg() []IProgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProgContext); ok {
			len++
		}
	}

	tst := make([]IProgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProgContext); ok {
			tst[i] = t.(IProgContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Prog(i int) IProgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgContext); ok {
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

	return t.(IProgContext)
}

func (s *StartContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalParserNEWLINE)
}

func (s *StartContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalParserNEWLINE, i)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *ParalParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParalParserRULE_start)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParalParserVARIABLE || _la == ParalParserEXECUTE {
		{
			p.SetState(8)
			p.Prog()
		}
		{
			p.SetState(9)
			p.Match(ParalParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(15)
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

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext
	Execute() IExecuteContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ProgContext) Execute() IExecuteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExecuteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExecuteContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *ParalParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParalParserRULE_prog)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalParserVARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Variable()
		}

	case ParalParserEXECUTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(17)
			p.Execute()
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

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARIABLE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	VALUE() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(ParalParserVARIABLE, 0)
}

func (s *VariableContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ParalParserIDENT, 0)
}

func (s *VariableContext) VALUE() antlr.TerminalNode {
	return s.GetToken(ParalParserVALUE, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *ParalParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParalParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Match(ParalParserVARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(21)
		p.Match(ParalParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.Match(ParalParserVALUE)
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

// IExecuteContext is an interface to support dynamic dispatch.
type IExecuteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXECUTE() antlr.TerminalNode
	VALUE() antlr.TerminalNode

	// IsExecuteContext differentiates from other interfaces.
	IsExecuteContext()
}

type ExecuteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecuteContext() *ExecuteContext {
	var p = new(ExecuteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_execute
	return p
}

func InitEmptyExecuteContext(p *ExecuteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalParserRULE_execute
}

func (*ExecuteContext) IsExecuteContext() {}

func NewExecuteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecuteContext {
	var p = new(ExecuteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalParserRULE_execute

	return p
}

func (s *ExecuteContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecuteContext) EXECUTE() antlr.TerminalNode {
	return s.GetToken(ParalParserEXECUTE, 0)
}

func (s *ExecuteContext) VALUE() antlr.TerminalNode {
	return s.GetToken(ParalParserVALUE, 0)
}

func (s *ExecuteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecuteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecuteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.EnterExecute(s)
	}
}

func (s *ExecuteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalListener); ok {
		listenerT.ExitExecute(s)
	}
}

func (p *ParalParser) Execute() (localctx IExecuteContext) {
	localctx = NewExecuteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParalParserRULE_execute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(ParalParserEXECUTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(25)
		p.Match(ParalParserVALUE)
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
