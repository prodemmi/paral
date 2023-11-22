// Code generated from /home/emad/Documents/paral/ParalExpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"start", "variables", "executeables", "variable", "execute",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 6, 43, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		1, 0, 1, 0, 1, 0, 5, 0, 14, 8, 0, 10, 0, 12, 0, 17, 9, 0, 1, 0, 3, 0, 20,
		8, 0, 1, 1, 4, 1, 23, 8, 1, 11, 1, 12, 1, 24, 1, 2, 5, 2, 28, 8, 2, 10,
		2, 12, 2, 31, 9, 2, 1, 2, 3, 2, 34, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 0, 0, 5, 0, 2, 4, 6, 8, 0, 0, 42, 0, 19, 1, 0, 0, 0,
		2, 22, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 35, 1, 0, 0, 0, 8, 39, 1, 0, 0,
		0, 10, 11, 3, 2, 1, 0, 11, 12, 5, 3, 0, 0, 12, 14, 1, 0, 0, 0, 13, 10,
		1, 0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0,
		16, 20, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 20, 3, 4, 2, 0, 19, 15, 1,
		0, 0, 0, 19, 18, 1, 0, 0, 0, 20, 1, 1, 0, 0, 0, 21, 23, 3, 6, 3, 0, 22,
		21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0,
		0, 25, 3, 1, 0, 0, 0, 26, 28, 3, 8, 4, 0, 27, 26, 1, 0, 0, 0, 28, 31, 1,
		0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 34, 1, 0, 0, 0, 31,
		29, 1, 0, 0, 0, 32, 34, 1, 0, 0, 0, 33, 29, 1, 0, 0, 0, 33, 32, 1, 0, 0,
		0, 34, 5, 1, 0, 0, 0, 35, 36, 5, 5, 0, 0, 36, 37, 5, 1, 0, 0, 37, 38, 5,
		2, 0, 0, 38, 7, 1, 0, 0, 0, 39, 40, 5, 6, 0, 0, 40, 41, 5, 2, 0, 0, 41,
		9, 1, 0, 0, 0, 5, 15, 19, 24, 29, 33,
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
	ParalExprParserRULE_start        = 0
	ParalExprParserRULE_variables    = 1
	ParalExprParserRULE_executeables = 2
	ParalExprParserRULE_variable     = 3
	ParalExprParserRULE_execute      = 4
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariables() []IVariablesContext
	Variables(i int) IVariablesContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	Executeables() IExecuteablesContext

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

func (s *StartContext) AllVariables() []IVariablesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariablesContext); ok {
			len++
		}
	}

	tst := make([]IVariablesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariablesContext); ok {
			tst[i] = t.(IVariablesContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Variables(i int) IVariablesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariablesContext); ok {
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

	return t.(IVariablesContext)
}

func (s *StartContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ParalExprParserNEWLINE)
}

func (s *StartContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ParalExprParserNEWLINE, i)
}

func (s *StartContext) Executeables() IExecuteablesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExecuteablesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExecuteablesContext)
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

	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalExprParserVARIABLE {
			{
				p.SetState(10)
				p.Variables()
			}
			{
				p.SetState(11)
				p.Match(ParalExprParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(17)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(18)
			p.Executeables()
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

// IVariablesContext is an interface to support dynamic dispatch.
type IVariablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariable() []IVariableContext
	Variable(i int) IVariableContext

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

func (s *VariablesContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *VariablesContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
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
	p.EnterRule(localctx, 2, ParalExprParserRULE_variables)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ParalExprParserVARIABLE {
		{
			p.SetState(21)
			p.Variable()
		}

		p.SetState(24)
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

// IExecuteablesContext is an interface to support dynamic dispatch.
type IExecuteablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExecute() []IExecuteContext
	Execute(i int) IExecuteContext

	// IsExecuteablesContext differentiates from other interfaces.
	IsExecuteablesContext()
}

type ExecuteablesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecuteablesContext() *ExecuteablesContext {
	var p = new(ExecuteablesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalExprParserRULE_executeables
	return p
}

func InitEmptyExecuteablesContext(p *ExecuteablesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParalExprParserRULE_executeables
}

func (*ExecuteablesContext) IsExecuteablesContext() {}

func NewExecuteablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecuteablesContext {
	var p = new(ExecuteablesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParalExprParserRULE_executeables

	return p
}

func (s *ExecuteablesContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecuteablesContext) AllExecute() []IExecuteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExecuteContext); ok {
			len++
		}
	}

	tst := make([]IExecuteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExecuteContext); ok {
			tst[i] = t.(IExecuteContext)
			i++
		}
	}

	return tst
}

func (s *ExecuteablesContext) Execute(i int) IExecuteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExecuteContext); ok {
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

	return t.(IExecuteContext)
}

func (s *ExecuteablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecuteablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecuteablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.EnterExecuteables(s)
	}
}

func (s *ExecuteablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParalExprListener); ok {
		listenerT.ExitExecuteables(s)
	}
}

func (p *ParalExprParser) Executeables() (localctx IExecuteablesContext) {
	localctx = NewExecuteablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParalExprParserRULE_executeables)
	var _la int

	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParalExprParserEXECUTE {
			{
				p.SetState(26)
				p.Execute()
			}

			p.SetState(31)
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
	p.EnterRule(localctx, 6, ParalExprParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(ParalExprParserVARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.Match(ParalExprParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
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
	p.EnterRule(localctx, 8, ParalExprParserRULE_execute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(ParalExprParserEXECUTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
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
