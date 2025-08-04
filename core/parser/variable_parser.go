package parser

import (
	"fmt"
	parser "paral/antlr/antlr"
	"paral/core"
	"paral/core/metadata"
	"paral/core/variable"
	"strconv"
)

func (p *Parser) parseVariable(ctx parser.IVariable_assignmentContext) *variable.Variable {
	name := ctx.IDENTIFIER().GetText()
	filename := ctx.GetStop().GetInputStream().GetSourceName()
	line := ctx.GetStop().GetLine()
	column := ctx.GetStop().GetColumn()

	val := ctx.Expression()
	mt := &metadata.Metadata{
		Filename: filename,
		Content:  ctx.GetText(),
		Line:     line,
		Column:   column,
	}

	if val.Number_expr() != nil {
		numText := val.GetText()
		num, err := strconv.Atoi(numText)
		if err != nil {
			fnum, err := strconv.ParseFloat(numText, 64)
			if err != nil {
				p.Runtime.Reporter.ThrowSyntaxError(fmt.Sprintf("Invalid number: %s", numText), mt)
			}
			return &variable.Variable{
				VarBase: variable.VarBase{Name: name},
				Value:   variable.FloatValue{VarBase: variable.VarBase{Name: name}, Value: fnum},
			}
		} else {
			return &variable.Variable{
				VarBase: variable.VarBase{Name: name},
				Value:   variable.IntValue{VarBase: variable.VarBase{Name: name}, Value: num},
			}
		}
	} else if val.String_expr() != nil {
		return &variable.Variable{
			VarBase: variable.VarBase{Name: name},
			Value:   variable.StringValue{VarBase: variable.VarBase{Name: name}, Value: core.TrimQuotes(val.GetText())},
		}
	} else if val.Boolean_expr() != nil {
		boolText := val.GetText()
		return &variable.Variable{
			VarBase: variable.VarBase{Name: name},
			Value:   variable.BoolValue{VarBase: variable.VarBase{Name: name}, Value: boolText == "true"},
		}
	} else if val.Duration_expr() != nil {
		return &variable.Variable{
			VarBase: variable.VarBase{Name: name},
			Value:   variable.DurationValue{VarBase: variable.VarBase{Name: name}, Value: val.GetText()},
		}
	} else if val.URL() != nil {
		return &variable.Variable{
			VarBase: variable.VarBase{Name: name},
			Value:   variable.StringValue{VarBase: variable.VarBase{Name: name}, Value: val.GetText()},
		}
	} else if val.Matrix_expr() != nil {
		var matrixValues [][]interface{}
		for _, list := range val.Matrix_expr().AllList_expr() {
			row := core.ExtractListValues(list)
			if len(row) == 0 {
				p.Runtime.Reporter.Warn(fmt.Sprintf("Empty list in matrix %s", name), mt)
			}
			matrixValues = append(matrixValues, row)
		}
		return &variable.Variable{
			VarBase: variable.VarBase{Name: name},
			Value:   variable.MatrixValue{VarBase: variable.VarBase{Name: name}, Value: matrixValues},
		}
	} else if val.List_expr() != nil {
		listValues := core.ExtractListValues(val.List_expr())
		return &variable.Variable{
			VarBase: variable.VarBase{Name: name},
			Value:   variable.ListValue{VarBase: variable.VarBase{Name: name}, Value: listValues},
		}
	} else if val.IDENTIFIER() != nil {
		refName := val.IDENTIFIER().GetText()
		refVar := p.Runtime.GetVariable(refName)
		if refVar == nil {
			p.Runtime.Reporter.ThrowSyntaxError(fmt.Sprintf("Undefined variable: %s", refName), mt)
		}
		return &variable.Variable{
			VarBase: variable.VarBase{Name: name},
			Value:   refVar.Value,
		}
	} else if val.Nested_function() != nil {
		nestedFunc := p.parseNestedFunction(nil, val.Nested_function())
		result, _ := nestedFunc.Call()
		return &variable.Variable{
			VarBase: variable.VarBase{Name: name},
			Value:   result,
		}
	} else if val.Function() != nil {
		function := p.parseFunction(nil, val.Function())
		result, _ := function.Call()
		return &variable.Variable{
			VarBase: variable.VarBase{Name: name},
			Value:   result,
		}
	} else {
		fmt.Println("val", val.GetText())
		p.Runtime.Reporter.ThrowSyntaxError("Unknown value type in variable definition", mt)
	}

	return nil
}
