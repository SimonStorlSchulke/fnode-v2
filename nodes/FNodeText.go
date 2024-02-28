package nodes

import "fnode2/core"

var textOutput *core.NodeOutput = core.NewNodeOutput(core.FTypeString, "Text",
	func(inputs []any, _ map[string]*core.NodeOption) any {
		return inputs[0].(string)
	})

func NewTextNode() *core.Node {
	return core.NewNode(
		"Value",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeString, "value", ""),
		},
		[]*core.NodeOutput{
			textOutput,
		},
	)
}
