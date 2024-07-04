package nodes

import (
	"errors"
	"fnode2/core"
	"slices"
	"strings"
)

type NodeCreator func() *core.Node

var nodeCreatorRegistry = map[string]func() *core.Node{
	"Output.Log":        newLogNode,
	"Math.Value":        newValueNode,
	"Math.Math":         newMathNode,
	"Math.RandomValue":  newRandomValueNode,
	"Math.FloatToInt":   newFloatToIntNode,
	"Text.Text":         newTextNode,
	"Text.TextReplace":  newTextReplaceNode,
	"Text.TextContains": newTextContainsNode,
	"Control.If":        newIfNode,
	"File.FileFromList": newFileFromListNode,
	"File.FileInfo":     newFileInfoNode,
	"File.FilterFiles":  newFileFilterNode,
	"File.DeleteFile":   newDeleteFileNode,
}

func Create(nodeType string) (*core.Node, error) {
	if nodeCreatorRegistry[nodeType] == nil {
		return nil, errors.New("unknown node type")
	}

	return nodeCreatorRegistry[nodeType](), nil
}

// NodeCategory is a list of Categories containing the categoryName and a list of NodeNames
var NodeCategories []core.NodeCategory = make([]core.NodeCategory, 0)

func GenerateNodeCategories() {
	for key, _ := range nodeCreatorRegistry {
		arr := strings.Split(key, ".")

		category := arr[0]
		nodeType := arr[1]

		idxCategory := slices.IndexFunc(NodeCategories, func(c core.NodeCategory) bool { return c.Name == category })

		if idxCategory == -1 {
			newCategory := core.NodeCategory{
				Name:      category,
				NodeTypes: []string{nodeType},
			}
			NodeCategories = append(NodeCategories, newCategory)
		} else {
			NodeCategories[idxCategory].NodeTypes = append(NodeCategories[idxCategory].NodeTypes, nodeType)
		}
	}
}
