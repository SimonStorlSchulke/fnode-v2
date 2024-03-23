package nodes

import (
	"fnode2/core"
)

var fileFilterOutput = core.NewNodeOutput(core.FTypeFile, "Filtered Files",
	func(node *core.Node) any {
		matches := node.GetInputBool(1)
		if matches {
			return node.GetInputFile(0)
		}
		return nil
	},
	true)

func newFileFilterNode() *core.Node {
	node := core.NewNodeCreator(
		"FilterFiles",
		"File",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeFile, "File", nil),
			core.NewNodeInput(core.FTypeBool, "Matched", false),
		},
		[]*core.NodeOutput{
			fileFilterOutput,
		},
	)

	return node
}
