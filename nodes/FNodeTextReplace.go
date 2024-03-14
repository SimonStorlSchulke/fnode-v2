package nodes

import (
	"fnode2/core"
	"regexp"
	"strings"
)

var textReplaceOutput = core.NewNodeOutput(core.FTypeString, "Result",
	func(node *core.Node) any {
		text := node.GetInputString(0)
		oldSegment := node.GetInputString(1)
		newSegment := node.GetInputString(2)

		switch node.Options["Mode"].SelectedOption {
		case "String":
			return strings.ReplaceAll(text, oldSegment, newSegment)
		case "Regex":
			re, err := regexp.Compile(oldSegment)
			if err != nil {
				return "invalid regex"
			}
			re.ReplaceAllString(text, newSegment)
		}

		return ""
	},
	true)

func newTextReplaceNode() *core.Node {
	node := core.NewNodeCreator(
		"TextReplace",
		"Text",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeString, "text", ""),
			core.NewNodeInput(core.FTypeString, "old", ""),
			core.NewNodeInput(core.FTypeString, "new", ""),
		},
		[]*core.NodeOutput{textReplaceOutput})

	node.AddOption("Mode", []string{
		"String",
		"Regex",
	})

	return node
}
