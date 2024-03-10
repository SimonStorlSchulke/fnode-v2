package InteractionLayer

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

type InteractionLayerExecute struct {
	printedStrings []string
}

var Ctx context.Context

func (l *InteractionLayerExecute) Print(text string) {
	l.printedStrings = append(l.printedStrings, text)
	fmt.Println(text)
	l.emitOutput(text)
}

func (l *InteractionLayerExecute) emitOutput(output string) {
	runtime.EventsEmit(Ctx, "output", output)
}
