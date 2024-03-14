package nodes

import "fnode2/core"

var textOutput = core.NewNodeOutput(core.FTypeString, "Text",
	func(node *core.Node) any {
		return node.GetInputString(0)
	},
	true)

func newTextNode() *core.Node {
	return core.NewNodeCreator(
		"Text",
		"Math",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeString, "value", ""),
		},
		[]*core.NodeOutput{
			textOutput,
		},
	)
}
