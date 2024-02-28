package nodes

import "fnode2/core"

var valueOutput *core.NodeOutput = core.NewNodeOutput(core.FTypeFloat, "Result",
	func(inputs []any, _ map[string]*core.NodeOption) any {
		return inputs[0].(float64)
	})

func NewValueNode() *core.Node {
	return core.NewNode(
		"Value",
		[]core.NodeInput{
			core.NewFloatInput("value", 1),
		},
		[]*core.NodeOutput{
			valueOutput,
		},
	)
}
