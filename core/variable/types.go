package variable

type MatrixValue struct {
	VarBase
	Value [][]interface{}
}

type ListValue struct {
	VarBase
	Value []interface{}
}

type IntValue struct {
	VarBase
	Value int
}

type FloatValue struct {
	VarBase
	Value float64
}

type StringValue struct {
	VarBase
	Value string
}

type BoolValue struct {
	VarBase
	Value bool
}

type DurationValue struct {
	VarBase
	Value interface{}
}
