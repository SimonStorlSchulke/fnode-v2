package nodes

import (
	"errors"
	"fnode2/core"
)

type NodeCreator func() *core.Node

var nodeCreatorRegistry = map[string]func() *core.Node{
	"Print":        newPrintNode,
	"Value":        newValueNode,
	"Math":         newMathNode,
	"Random Value": newRandomValueNode,
	"Text":         newTextNode,
	"Text Replace": newTextReplaceNode,
}

func Create(nodeType string) (*core.Node, error) {
	if nodeCreatorRegistry[nodeType] == nil {
		return nil, errors.New("unknown node type")
	}

	return nodeCreatorRegistry[nodeType](), nil
}
