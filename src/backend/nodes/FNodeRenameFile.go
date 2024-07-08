package nodes

import (
	"fnode2/core"
)

func renameFileOutput(interactionLayer core.NodeInteractionLayer, inputs []any, _ map[string]*core.NodeOption) {
	text := inputs[0].(string)
	if len(text) > 0 {
		interactionLayer.Print(text)
	}
}

func newRenameFileNode() *core.Node {
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
