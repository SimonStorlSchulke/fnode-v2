package core

import (
	"fmt"
	"os"
	"path"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type InteractionLayerExecute struct {
	printedStrings []string
}

func (l *InteractionLayerExecute) Print(text string) {
	l.printedStrings = append(l.printedStrings, text)
	fmt.Println(text)
	l.emitOutput(text)
}

func (l *InteractionLayerExecute) RemoveFile(path string) error {
	return os.Remove(path)
}

func (l *InteractionLayerExecute) emitOutput(output string) {
	runtime.EventsEmit(Ctx, "output", output)
}

func (I *InteractionLayerExecute) RenameFile(oldPath string, newName string) string {
	fullNewPath := path.Join(path.Dir(oldPath), newName)
	err := os.Rename(oldPath, fullNewPath)
	if err != nil {
		fmt.Println(err)
	}
	return fullNewPath
}

func (I *InteractionLayerExecute) MoveFile(oldPath string, toFolder string) string {
	fullNewPath := path.Join(toFolder, path.Base(oldPath))
	err := os.Rename(oldPath, fullNewPath)
	if err != nil {
		fmt.Println(err)
	}
	return toFolder
}

func (I *InteractionLayerExecute) OpenFile(path string, appLocation string) {

}
