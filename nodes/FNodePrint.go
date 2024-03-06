package nodes

import (
	"fnode2/core"
	"fnode2/core/InteractionLayer"
)

func printOutput(interactionLayer InteractionLayer.NodeInteractionLayer, inputs []any, _ map[string]*core.NodeOption) {
	interactionLayer.Print(inputs[0].(string))
}

func newPrintNode() *core.Node {
	node := core.NewNodeCreator(
		"Print",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeString, "Text", ""),
		},
		[]*core.NodeOutput{},
	)

	node.ExecutiveFunction = printOutput

	return node
}
