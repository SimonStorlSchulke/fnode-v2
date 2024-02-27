package nodes

import "fnode2/core"

func mathOutput(inputs []any) any {
	return inputs[0].(float64) * inputs[1].(float64)
}

func NewMathNode() *core.Node {
	return core.NewNode(
		"Math",
		[]core.NodeInput[any]{
			{
				Name:         "a",
				DefaultValue: 1.0,
			},
			{
				Name:         "b",
				DefaultValue: 1.0,
			},
		},
		[]core.NodeOutput[any]{
			mathOutput,
		},
	)
}
