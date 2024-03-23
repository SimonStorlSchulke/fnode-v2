package core

import (
	"fmt"

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
