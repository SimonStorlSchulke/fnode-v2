package nodes

import (
	"fnode2/core"
)

func deleteFileExecutor(interactionLayer core.NodeInteractionLayer, inputs []any, _ map[string]*core.NodeOption) {
	if inputs[0] == nil {
		return
	}
	filePath := inputs[0].(core.FFile).FullPath
	err := interactionLayer.RemoveFile(filePath)
	if err != nil {
		core.LogErr(err)
	}
}

var deleteFileSuccessOutput = core.NewNodeOutput(core.FTypeBool, "Success",
	func(node *core.Node) any {
		matches := node.GetInputBool(1)
		if matches {
			return node.GetInputFile(0)
		}
		return nil
	},
	true)

func newDeleteFileNode() *core.Node {
	node := core.NewNodeCreator(
		"DeleteFile",
		"File",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeFile, "File", nil),
		},
		[]*core.NodeOutput{
			deleteFileSuccessOutput,
		},
	)

	node.ExecutiveFunction = deleteFileExecutor

	return node
}
