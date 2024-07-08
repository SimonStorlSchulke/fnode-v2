package core

import (
	"fmt"
	"path"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type InteractionLayerPreview struct {
	printedStrings []string
}

func (l *InteractionLayerPreview) Print(text string) {
	l.printedStrings = append(l.printedStrings, text)
	fmt.Println(text)
	l.emitOutput(text)
}

func (l *InteractionLayerPreview) RemoveFile(path string) error {
	l.emitOutput("Would remove File: " + path)
	return nil
}

func (l *InteractionLayerPreview) emitOutput(output string) {
	runtime.EventsEmit(Ctx, "output", output)
}

func (l *InteractionLayerPreview) RenameFile(oldPath string, newName string) string {
	l.emitOutput(fmt.Sprintf("Would renamed File '%s' to '%s'", oldPath, newName))
	return path.Join(path.Dir(oldPath), newName)
}

func (l *InteractionLayerPreview) MoveFile(oldPath string, toFolder string) string {
	l.emitOutput(fmt.Sprintf("Would move File '%s' to '%s'", oldPath, toFolder))
	return toFolder
}
