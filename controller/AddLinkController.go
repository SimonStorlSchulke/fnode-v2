package controller

import "fnode2/core"

func (a *App) AddLink(link core.NodeLink) {
	if &link == nil {
		core.Log("Attempting to create nil Link", core.LogLevelError)
		return
	}
	tree.AddLink(&link)
}

func (a *App) RemoveLink(link core.NodeLink) {
	tree.RemoveLink(&link)
}
