package controller

import "fnode2/core"

func (a *App) AddLink(link core.NodeLink) {
	tree.AddLink(&link)
}
