package nodes

import (
	"fnode2/core"
	"regexp"
	"strings"
)

var textReplaceOutput = core.NewNodeOutput(core.FTypeString, "Result",
	func(inputs []any, Options map[string]*core.NodeOption) any {
		text := inputs[0].(string)
		oldSegment := inputs[1].(string)
		newSegment := inputs[2].(string)

		switch Options["Mode"].SelectedOption {
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
		"Text Replace",
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
