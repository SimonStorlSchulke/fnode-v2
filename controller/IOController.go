package controller

import (
	"fnode2/core"
	"fnode2/treeIo"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) Save(name string) {
	treeIo.SaveToFile(&tree, "./presets", name)
}

func (a *App) SaveAs() {
	path, err := runtime.SaveFileDialog(core.Ctx, runtime.SaveDialogOptions{})
	if err != nil {
		core.LogError("Saving failed: &v", err)
	}
	treeIo.SaveToFile(&tree, "", path)
}

func (a *App) LoadFile() {
	path, err := runtime.OpenFileDialog(core.Ctx, runtime.OpenDialogOptions{})
	if err != nil {
		core.LogError("Loading failed: &v", err)
	}
	tree = *treeIo.LoadFromFile(path)
}
