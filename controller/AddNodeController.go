package controller

import (
	"fnode2/core"
	"fnode2/nodes"
)

func (a *App) AddNode(ofType string, posX int, posY int) string {
	node, err := nodes.Create(ofType)

	if err != nil {
		core.LogError("Cannot add Node of type " + ofType)
		return ""
	}
	node.Meta.PosX = posX
	node.Meta.PosY = posY
	tree.AddNode(node)
	return node.Id
}

// AddConnectedNode adds a new Node next to the one with id 'connectedNodeId'
func (a *App) AddConnectedNode(ofType string, connectedNodeId string) string {
	node, err := nodes.Create(ofType)

	if err != nil {
		core.LogError("Cannot add Node of type " + ofType)
		return ""
	}

	var connectedNode *core.Node = nil

	if connectedNodeId != "" {
		connectedNode, err = tree.FindNodeById(connectedNodeId)
	}

	tree.AddNode(node)

	if connectedNode == nil {
		core.LogError("Cannot find Node to connect to " + ofType)
		node.Meta.PosX = 100
		node.Meta.PosY = 100
	} else {
		node.Meta.PosX = connectedNode.Meta.PosX + 200
		node.Meta.PosY = connectedNode.Meta.PosY

		if len(connectedNode.Outputs) > 0 && len(node.Inputs) > 0 {
			tree.AddLink(&core.NodeLink{
				FromNode:   connectedNodeId,
				FromOutput: 0,
				ToNode:     node.Id,
				ToInput:    0,
			})
		}
	}
	return node.Id

}

func (a *App) RemoveNode(nodeId string) {

	node, err := tree.FindNodeById(nodeId)

	if err != nil {
		core.LogError("Cannot Remove Node '%s' - Not found ", nodeId)
		return
	}

	tree.RemoveNode(node)
}
