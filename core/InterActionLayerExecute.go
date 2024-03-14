package core

import (
	"fmt"
	"os"

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
