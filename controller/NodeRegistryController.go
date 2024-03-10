package controller

import (
	"fnode2/core"
	"fnode2/nodes"
)

func (a *App) GetNodeCategories() []core.NodeCategory {
	return nodes.NodeCategories
}
