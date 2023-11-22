// Code generated from ParalExpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ParalExpr

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

type ParalExprParser struct {
	*antlr.BaseParser
}

var ParalExprParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func paralexprParserInit() {
	staticData := &ParalExprParserStaticData
	staticData.SymbolicNames = []string{
		"", "IDENT", "VALUE", "NEWLINE", "WS", "VARIABLE", "EXECUTE",
	}
	staticData.RuleNames = []string{
		"start", "prog", "variables", "executables", "variable", "execute",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 6, 41, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 5, 0, 16, 8, 0, 10, 0, 12, 0, 19, 9, 0, 1,
		0, 3, 0, 22, 8, 0, 1, 1, 1, 1, 3, 1, 26, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		3, 3, 32, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 0, 0, 6,
		0, 2, 4, 6, 8, 10, 0, 0, 38, 0, 21, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 27,
		1, 0, 0, 0, 6, 31, 1, 0, 0, 0, 8, 33, 1, 0, 0, 0, 10, 37, 1, 0, 0, 0, 12,
		13, 3, 2, 1, 0, 13, 14, 5, 3, 0, 0, 14, 16, 1, 0, 0, 0, 15, 12, 1, 0, 0,
		0, 16, 19, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 22,
		1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 20, 22, 1, 0, 0, 0, 21, 17, 1, 0, 0, 0,
		21, 20, 1, 0, 0, 0, 22, 1, 1, 0, 0, 0, 23, 26, 3, 4, 2, 0, 24, 26, 3, 6,
		3, 0, 25, 23, 1, 0, 0, 0, 25, 24, 1, 0, 0, 0, 26, 3, 1, 0, 0, 0, 27, 28,
		3, 8, 4, 0, 28, 5, 1, 0, 0, 0, 29, 32, 3, 10, 5, 0, 30, 32, 1, 0, 0, 0,
		31, 29, 1, 0, 0, 0, 31, 30, 1, 0, 0, 0, 32, 7, 1, 0, 0, 0, 33, 34, 5, 5,
		0, 0, 34, 35, 5, 1, 0, 0, 35, 36, 5, 2, 0, 0, 36, 9, 1, 0, 0, 0, 37, 38,
		5, 6, 0, 0, 38, 39, 5, 2, 0, 0, 39, 11, 1, 0, 0, 0, 4, 17, 21, 25, 31,
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

// ParalExprParserInit initializes any static state used to implement ParalExprParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewParalExprParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParalExprParserInit() {
	staticData := &ParalExprParserStaticData
	staticData.once.Do(paralexprParserInit)
}

// NewParalExprParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParalExprParser(input antlr.TokenStream) *ParalExprParser {
	ParalExprParserInit()
	this := new(ParalExprParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ParalExprParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ParalExpr.g4"

	return this
}

// ParalExprParser tokens.
const (
	ParalExprParserEOF      = antlr.TokenEOF
	ParalExprParserIDENT    = 1
	ParalExprParserVALUE    = 2
	ParalExprParserNEWLINE  = 3
	ParalExprParserWS       = 4
	ParalExprParserVARIABLE = 5
	ParalExprParserEXECUTE  = 6
)

// ParalExprParser rules.
const (
	ParalExprParserRULE_start       = 0
	ParalExprParserRULE_prog        = 1
	ParalExprParserRULE_variables   = 2
	ParalExprParserRULE_executables = 3
	ParalExprParserRULE_variable    = 4
	ParalExprParserRULE_execute     = 5
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
	p.RuleIndex = ParalExprParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalExprParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalExprParserRULE_start

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
	return s.GetTokens(ParalExprParserNEWLINE)
}

func (s *StartContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalExprParserNEWLINE, i)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *ParalExprParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParalExprParserRULE_start)
	var _la int

	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&104) != 0 {
			{
				p.SetState(12)
				p.Prog()
			}
			{
				p.SetState(13)
				p.Match(ParalExprParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(19)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)

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

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variables() IVariablesContext
	Executables() IExecutablesContext

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
	p.RuleIndex = ParalExprParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalExprParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalExprParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Variables() IVariablesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariablesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariablesContext)
}

func (s *ProgContext) Executables() IExecutablesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExecutablesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExecutablesContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *ParalExprParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParalExprParserRULE_prog)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalExprParserVARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Variables()
		}

	case ParalExprParserNEWLINE, ParalExprParserEXECUTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Executables()
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

// IVariablesContext is an interface to support dynamic dispatch.
type IVariablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext

	// IsVariablesContext differentiates from other interfaces.
	IsVariablesContext()
}

type VariablesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariablesContext() *VariablesContext {
	var p = new(VariablesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalExprParserRULE_variables
	return p
}

func InitEmptyVariablesContext(p *VariablesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalExprParserRULE_variables
}

func (*VariablesContext) IsVariablesContext() {}

func NewVariablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablesContext {
	var p = new(VariablesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalExprParserRULE_variables

	return p
}

func (s *VariablesContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablesContext) Variable() IVariableContext {
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

func (s *VariablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.EnterVariables(s)
	}
}

func (s *VariablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.ExitVariables(s)
	}
}

func (p *ParalExprParser) Variables() (localctx IVariablesContext) {
	localctx = NewVariablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParalExprParserRULE_variables)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Variable()
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

// IExecutablesContext is an interface to support dynamic dispatch.
type IExecutablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Execute() IExecuteContext

	// IsExecutablesContext differentiates from other interfaces.
	IsExecutablesContext()
}

type ExecutablesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutablesContext() *ExecutablesContext {
	var p = new(ExecutablesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalExprParserRULE_executables
	return p
}

func InitEmptyExecutablesContext(p *ExecutablesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalExprParserRULE_executables
}

func (*ExecutablesContext) IsExecutablesContext() {}

func NewExecutablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutablesContext {
	var p = new(ExecutablesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalExprParserRULE_executables

	return p
}

func (s *ExecutablesContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecutablesContext) Execute() IExecuteContext {
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

func (s *ExecutablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecutablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecutablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.EnterExecutables(s)
	}
}

func (s *ExecutablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.ExitExecutables(s)
	}
}

func (p *ParalExprParser) Executables() (localctx IExecutablesContext) {
	localctx = NewExecutablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParalExprParserRULE_executables)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParalExprParserEXECUTE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Execute()
		}

	case ParalExprParserNEWLINE:
		p.EnterOuterAlt(localctx, 2)

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
	p.RuleIndex = ParalExprParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalExprParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalExprParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(ParalExprParserVARIABLE, 0)
}

func (s *VariableContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ParalExprParserIDENT, 0)
}

func (s *VariableContext) VALUE() antlr.TerminalNode {
	return s.GetToken(ParalExprParserVALUE, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *ParalExprParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ParalExprParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(ParalExprParserVARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Match(ParalExprParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
		p.Match(ParalExprParserVALUE)
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
	p.RuleIndex = ParalExprParserRULE_execute
	return p
}

func InitEmptyExecuteContext(p *ExecuteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalExprParserRULE_execute
}

func (*ExecuteContext) IsExecuteContext() {}

func NewExecuteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecuteContext {
	var p = new(ExecuteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalExprParserRULE_execute

	return p
}

func (s *ExecuteContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecuteContext) EXECUTE() antlr.TerminalNode {
	return s.GetToken(ParalExprParserEXECUTE, 0)
}

func (s *ExecuteContext) VALUE() antlr.TerminalNode {
	return s.GetToken(ParalExprParserVALUE, 0)
}

func (s *ExecuteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecuteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecuteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.EnterExecute(s)
	}
}

func (s *ExecuteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.ExitExecute(s)
	}
}

func (p *ParalExprParser) Execute() (localctx IExecuteContext) {
	localctx = NewExecuteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ParalExprParserRULE_execute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Match(ParalExprParserEXECUTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(ParalExprParserVALUE)
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
