package core

import (
	"fmt"
)

const (
	LogLevelInfo = iota
	LogLevelWarning
	LogLevelError
	LogLevelPanic
)

func Log(format string, level int, args ...any) {
	formatString := fmt.Sprintf(format, args...)
	levelString := []string{"Info: ", "Warning: ", "Error: ", "Panic: "}[level]
	fmt.Println(levelString + formatString)
	if level == LogLevelPanic {
		panic(formatString)
	}
}
