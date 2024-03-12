package nodes

import (
	"fnode2/core"
)

var ifOutput = core.NewNodeOutput(core.FTypeFloat, "Result",
	func(inputs []any, _ map[string]*core.NodeOption) any {
		if inputs[0].(bool) {
			return inputs[2]
		} else {
			return inputs[1]
		}
	},
	true)

func newIfNode() *core.Node {
	node := core.NewNodeCreator(
		"If",
		"Control",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeBool, "Condition", false),
			core.NewNodeInput(core.FTypeFloat, "If false", 0.0),
			core.NewNodeInput(core.FTypeFloat, "If true", 1.0),
		},
		[]*core.NodeOutput{ifOutput},
	)

	node.AddOption("Type", []string{
		"Float",
		"Int",
		"Text",
	})

	node.SetOptionCallback("Type", func(node *core.Node, selectedChoice string) {
		typeToSet := core.FTypeFloat
		var value1ToSet any = 0.0
		var value2ToSet any = 1.0
		switch selectedChoice {
		case "Float":
			typeToSet = core.FTypeFloat
			value1ToSet = 0.0
			value2ToSet = 1.0
		case "Int":
			typeToSet = core.FTypeInt
			value1ToSet = 0
			value2ToSet = 1
		case "Text":
			typeToSet = core.FTypeString
			value1ToSet = "false"
			value2ToSet = "true"
		}
		node.Inputs[1].Type = typeToSet
		node.Inputs[2].Type = typeToSet
		node.Outputs[0].Type = typeToSet
		node.SetInputDefaultValue(1, value1ToSet)
		node.SetInputDefaultValue(2, value2ToSet)
	})

	return node
}
