package nodes

import (
	"fmt"
	"fnode2/core"
)

func printOutput(inputs []any, _ map[string]*core.NodeOption) {
	fmt.Println(inputs[0])
}

func NewPrintNode() *core.Node {
	node := core.NewNode(
		"Print",
		[]core.NodeInput[any]{
			{
				Name:         "Text",
				DefaultValue: "",
			},
		},
		[]core.NodeOutput[any]{},
	)

	node.ExecutiveFunction = printOutput

	return node
}
