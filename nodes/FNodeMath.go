package nodes

import (
	"fnode2/core"
	"math"
)

func mathOutput(inputs []any, Options map[string]*core.NodeOption) any {

	a := inputs[0].(float64)
	b := inputs[1].(float64)
	switch Options["Mode"].SelectedOption {
	case "Add":
		return a + b
	case "Subtract":
		return a - b
	case "Multiply":
		return a * b
	case "Divide":
		return a / b
	case "Power":
		return math.Pow(a, b)
	case "Square Root":
		return math.Sqrt(a)
	case "Sin":
		return math.Sin(a)
	case "Cos":
		return math.Cos(a)
	}

	return 0.0
}

func NewMathNode() *core.Node {

	node := core.NewNode(
		"Math",
		[]core.NodeInput{
			core.NewFloatInput("a", 1),
			core.NewFloatInput("b", 1),
		},
		[]core.NodeOutputFuncFunc{
			mathOutput,
		},
	)

	node.AddOption("Mode", []string{
		"Add",
		"Subtract",
		"Multiply",
		"Divide",
		"Power",
		"Square Root",
		"Sin",
		"Cos"})

	return node
}
