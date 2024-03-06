package nodes

import "fnode2/core"

var valueOutput = core.NewNodeOutput(core.FTypeFloat, "value",
	func(inputs []any, _ map[string]*core.NodeOption) any {
		return inputs[0].(float64)
	},
	true)

func newValueNode() *core.Node {
	return core.NewNodeCreator(
		"Value",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeFloat, "value", 1.0),
		},
		[]*core.NodeOutput{
			valueOutput,
		},
	)
}
