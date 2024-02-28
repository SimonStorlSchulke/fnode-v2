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
		[]core.NodeInput{
			core.NewStringInput("Text", ""),
		},
		[]*core.NodeOutput{},
	)

	node.ExecutiveFunction = printOutput

	return node
}
