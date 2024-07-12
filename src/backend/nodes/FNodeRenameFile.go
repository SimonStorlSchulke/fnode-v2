package nodes

import (
	"fnode2/core"
)

func renameFileOutput(interactionLayer core.NodeInteractionLayer, inputs []any, _ map[string]*core.NodeOption) {
	from := inputs[0].(string)
	to := inputs[1].(string)
	interactionLayer.RenameFile(from, to)
}

func newRenameFileNode() *core.Node {
	node := core.NewNodeCreator(
		"RenameFile",
		"File",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeString, "Text", ""),
			core.NewNodeInput(core.FTypeString, "Text", ""),
		},
		[]*core.NodeOutput{},
	)

	node.ExecutiveFunction = renameFileOutput

	return node
}
