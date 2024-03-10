package controller

import (
	"fnode2/core"
	"fnode2/nodes"
)

func (a *App) AddNode(ofType string, posX int, posY int) {
	node, err := nodes.Create(ofType)

	if err != nil {
		core.Log("Cannot add Node of type "+ofType, core.LogLevelError)
		return
	}
	node.Meta.PosX = posX
	node.Meta.PosY = posY
	tree.AddNode(node)
}
