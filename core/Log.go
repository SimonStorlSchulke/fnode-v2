package core

import "fmt"

const (
	LogLevelInfo = iota
	LogLevelWarning
	LogLevelError
)

func Log(format string, level int, args ...any) {
	formatString := fmt.Sprintf(format, args...)
	levelString := []string{"Info: ", "Warning: ", "Error: "}[level]
	fmt.Println(levelString + formatString)
}
