package nodes

import "fnode2/core"

var valueOutput = core.NewNodeOutput(core.FTypeFloat, "value",
	func(node *core.Node) any {
		return node.GetInputFloat(0)
	},
	true)

func newValueNode() *core.Node {
	return core.NewNodeCreator(
		"Value",
		"Math",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeFloat, "value", 1.0),
		},
		[]*core.NodeOutput{
			valueOutput,
		},
	)
}
