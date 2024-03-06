package nodes

import "fnode2/core"

var textOutput = core.NewNodeOutput(core.FTypeString, "Text",
	func(inputs []any, _ map[string]*core.NodeOption) any {
		return inputs[0].(string)
	},
	true)

func newTextNode() *core.Node {
	return core.NewNodeCreator(
		"Value",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeString, "value", ""),
		},
		[]*core.NodeOutput{
			textOutput,
		},
	)
}
