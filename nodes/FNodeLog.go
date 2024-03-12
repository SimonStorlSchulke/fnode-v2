package nodes

import (
	"fnode2/core"
)

func printOutput(interactionLayer core.NodeInteractionLayer, inputs []any, _ map[string]*core.NodeOption) {
	interactionLayer.Print(inputs[0].(string))
}

func newLogNode() *core.Node {
	node := core.NewNodeCreator(
		"Log",
		"Output",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeString, "Text", ""),
		},
		[]*core.NodeOutput{},
	)

	node.ExecutiveFunction = printOutput

	return node
}
