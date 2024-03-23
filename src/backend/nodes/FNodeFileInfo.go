package nodes

import (
	"fnode2/core"
	"path/filepath"
	"strings"
)

var fileInfoOutputName = core.NewNodeOutput(core.FTypeString, "Name",
	func(node *core.Node) any {
		fileName := node.GetInputFile(0).Info.Name()

		if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
			return fileName[:pos]
		}
		return fileName
	},
	true)

var fileInfoOutputExtension = core.NewNodeOutput(core.FTypeString, "Extension",
	func(node *core.Node) any {
		return filepath.Ext(node.GetInputFile(0).Info.Name())
	},
	true)

var fileInfoOutputNameWithExtension = core.NewNodeOutput(core.FTypeString, "Name with Extension",
	func(node *core.Node) any {
		return node.GetInputFile(0).Info.Name()
	},
	true)

var fileInfoOutputBaseDir = core.NewNodeOutput(core.FTypeString, "Base Dir",
	func(node *core.Node) any {
		return filepath.Dir(node.GetInputFile(0).FullPath)
	},
	true)

func newFileInfoNode() *core.Node {
	node := core.NewNodeCreator(
		"FileInfo",
		"File",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeFile, "File", nil),
		},
		[]*core.NodeOutput{
			fileInfoOutputName,
			fileInfoOutputExtension,
			fileInfoOutputNameWithExtension,
			fileInfoOutputBaseDir,
		},
	)

	return node
}
