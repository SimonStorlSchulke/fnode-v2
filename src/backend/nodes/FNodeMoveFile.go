package nodes

import (
	"fnode2/core"
)

func moveFileOutput(interactionLayer core.NodeInteractionLayer, inputs []any, _ map[string]*core.NodeOption) {
	file := inputs[0].(string)
	toFolder := inputs[1].(string)
	interactionLayer.MoveFile(file, toFolder)
}

func newMoveFileNode() *core.Node {
	node := core.NewNodeCreator(
		"MoveFile",
		"File",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeFile, "File", ""),
			core.NewNodeInput(core.FTypeFile, "To Folder", ""),
		},
		[]*core.NodeOutput{},
	)

	node.ExecutiveFunction = moveFileOutput

	return node
}
