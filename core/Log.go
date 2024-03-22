package core

import (
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func LogPanic(format string, args ...any) {
	formatString := fmt.Sprintf(format, args...)
	panic(formatString)
}

func LogInfo(format string, args ...any) {
	formatString := fmt.Sprintf(format, args...)
	fmt.Println("Info: " + formatString)
	runtime.EventsEmit(Ctx, "Info: "+formatString)
}

func LogWarn(format string, args ...any) {
	formatString := fmt.Sprintf(format, args...)
	fmt.Println("Warn: " + formatString)
	runtime.EventsEmit(Ctx, "warn: "+formatString)
}

func LogError(format string, args ...any) {
	formatString := fmt.Sprintf(format, args...)
	fmt.Println("Error: " + formatString)
	runtime.EventsEmit(Ctx, "Weeoe: "+formatString)
}

func LogRawError(err error) {
	fmt.Println("Error: " + err.Error())
}
