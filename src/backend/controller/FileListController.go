package controller

import (
	"fnode2/core"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var fileList *core.FileList = &core.FileList{
	LooseFiles: []string{},
	Directories: []struct {
		Path      string
		Recursive bool
	}{},
}

func (a *App) AddLooseFilesFromDialog() {
	filePaths, err := runtime.OpenMultipleFilesDialog(core.Ctx, runtime.OpenDialogOptions{})
	if err != nil {
		core.LogRawError(err)
		return
	}

	fileList.LooseFiles = append(fileList.LooseFiles, filePaths...)
}

func (a *App) AddDirectoriesFromDialog() {
	dir, err := runtime.OpenDirectoryDialog(core.Ctx, runtime.OpenDialogOptions{})
	if err != nil {
		core.LogRawError(err)
		return
	}

	fileList.Directories = append(fileList.Directories, struct {
		Path      string
		Recursive bool
	}{Path: dir, Recursive: false})
}

func (a *App) ClearFileList() {
	fileList = &core.FileList{LooseFiles: []string{}, Directories: []struct {
		Path      string
		Recursive bool
	}{}}
}

func (a *App) RemoveLooseFile(index int) {
	fileList.LooseFiles = append(fileList.LooseFiles[:index], fileList.LooseFiles[index+1:]...)
}

func (a *App) RemoveDirectory(index int) {
	fileList.Directories = append(fileList.Directories[:index], fileList.Directories[index+1:]...)
}

func (a *App) GetFileList() core.FileList {
	return *fileList
}
