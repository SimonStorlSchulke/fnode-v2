package tests

import (
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
