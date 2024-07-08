package controller

import (
	"fnode2/core"
	"fnode2/treeIo"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) Save(name string) {
	treeIo.SaveToFile(&tree, "./presets", name)
}

func (a *App) SaveAs() {
	path, err := runtime.SaveFileDialog(core.Ctx, runtime.SaveDialogOptions{})

	directory := filepath.Dir(path)
	fileName := filepath.Base(path)

	if err != nil {
		core.LogError("Saving failed: &v", err)
	}
	err = treeIo.SaveToFile(&tree, directory, fileName)
	if err != nil {
		core.LogError("Saving failed: &v", err)
	}
}

func (a *App) LoadFile() {
	path, err := runtime.OpenFileDialog(core.Ctx, runtime.OpenDialogOptions{})
	if err != nil {
		core.LogError("Loading failed: &v", err)
	}
	tree = *treeIo.LoadFromFile(path)
}
