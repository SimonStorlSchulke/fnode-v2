package nodes

import (
	"fnode2/core"
	"math"
)

var mathResultOutput = core.NewNodeOutput(core.FTypeFloat, "Result",
	func(node *core.Node) any {
		a := node.GetInputFloat(0)
		b := node.GetInputFloat(1)
		switch node.Options["Mode"].SelectedOption {
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
	},
	true)

func newMathNode() *core.Node {
	node := core.NewNodeCreator(
		"Math",
		"Math",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeFloat, "a", 1.0),
			core.NewNodeInput(core.FTypeFloat, "b", 1.0),
		},
		[]*core.NodeOutput{mathResultOutput},
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
