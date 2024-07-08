package nodes

import (
	"fnode2/core"
)

var fileFromListCurrentFileOutput = core.NewNodeOutput(core.FTypeFile, "File",
	func(node *core.Node) any {
		return core.RunState.CurrentFile
	},
	true)

var fileFromListIndexOutput = core.NewNodeOutput(core.FTypeInt, "Index",
	func(node *core.Node) any { return core.RunState.CurrentIteration },
	true)

func newFileFromListNode() *core.Node {
	node := core.NewNodeCreator(
		"FileFromList",
		"File",
		[]core.NodeInput{},
		[]*core.NodeOutput{
			fileFromListCurrentFileOutput,
			fileFromListIndexOutput,
		},
	)

	return node
}
