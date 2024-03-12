package core

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	LogLevelInfo = iota
	LogLevelWarning
	LogLevelError
	LogLevelPanic
)

func Log(format string, level int, args ...any) {
	formatString := fmt.Sprintf(format, args...)
	levelString := []string{"Info: ", "Warn: ", "Error: ", "Panic: "}[level]
	fmt.Println(levelString + formatString)
	if level == LogLevelPanic {
		panic(formatString)
	}
	runtime.EventsEmit(Ctx, "log", levelString+formatString)
}

func LogErr(err error) {
	fmt.Println("Error: " + err.Error())
}
