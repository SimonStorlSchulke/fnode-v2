package nodes

import "fnode2/core"

func valueOutput(inputs []any, _ map[string]*core.NodeOption) any {
	return inputs[0].(float64)
}

func NewValueNode() *core.Node {
	return core.NewNode(
		"Value",
		[]core.NodeInput[any]{
			{
				Name:         "value",
				DefaultValue: 1.0,
			},
		},
		[]core.NodeOutput[any]{
			valueOutput,
		},
	)
}
