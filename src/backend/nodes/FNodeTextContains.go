package nodes

import (
	"fnode2/core"
	"strings"
)

var textContainsOutput = core.NewNodeOutput(core.FTypeBool, "Contains",
	func(node *core.Node) any {
		text := node.GetInputString(0)
		contained := node.GetInputString(1)
		caseSensitive := node.GetInputBool(2)

		if caseSensitive {
			return strings.Contains(text, contained)
		}
		return strings.Contains(strings.ToLower(text), strings.ToLower(contained))

	},
	true)

func newTextContainsNode() *core.Node {
	node := core.NewNodeCreator(
		"TextContains",
		"Text",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeString, "Text", ""),
			core.NewNodeInput(core.FTypeString, "Contained", ""),
			core.NewNodeInput(core.FTypeBool, "Case Sensitive", true),
		},
		[]*core.NodeOutput{textContainsOutput})

	return node
}
