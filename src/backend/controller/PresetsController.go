package controller

import (
	"fnode2/core"
	"fnode2/treeIo"
)

func (a *App) GetPresetNames() []string {
	names, err := treeIo.GetPresetNames()
	if err != nil {
		core.LogRawError(err)
		return []string{}
	}
	return names
}
