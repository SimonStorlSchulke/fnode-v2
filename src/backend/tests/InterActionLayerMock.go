package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type InteractionLayerMock struct {
	printedStrings []string
	removedFiles   []string
}

func (l *InteractionLayerMock) Print(text string) {
	l.printedStrings = append(l.printedStrings, text)
}

func (l *InteractionLayerMock) VerifyPrinted(t *testing.T, text string) {
	assert.Contains(t, l.printedStrings, text)
}

func (l *InteractionLayerMock) RemoveFile(path string) error {
	l.removedFiles = append(l.removedFiles, path)
	return nil
}

func (l *InteractionLayerMock) RenameFile(oldPath string, newName string) string {
	return fmt.Sprintf("Would renamed File '%s' to '%s'", oldPath, newName)
}

func (l *InteractionLayerMock) MoveFile(oldPath string, newPath string) string {
	return fmt.Sprintf("Would move File '%s' to '%s'", oldPath, newPath)
}
