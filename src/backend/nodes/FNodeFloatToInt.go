package nodes

import (
	"fnode2/core"
	"math"
)

var floatToIntOutput = core.NewNodeOutput(core.FTypeInt, "value",
	func(node *core.Node) any {

		switch node.Options["Mode"].SelectedOption {
		case "Round":
			return int(math.Round(node.GetInputFloat(0)))
		case "Ceil":
			return int(math.Ceil(node.GetInputFloat(0)))
		case "Floor":
			return int(math.Floor(node.GetInputFloat(0)))
		}
		return int(math.Floor(node.GetInputFloat(0)))
	},
	true)

func newFloatToIntNode() *core.Node {
	node := core.NewNodeCreator(
		"FloatToInt",
		"Math",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeFloat, "Float Value", 1.0),
		},
		[]*core.NodeOutput{
			floatToIntOutput,
		},
	)

	node.AddOption("Mode", []string{
		"Round",
		"Ceil",
		"Floor"})

	return node
}
