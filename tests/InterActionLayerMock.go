package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type InteractionLayerMock struct {
	printedStrings []string
}

func (l *InteractionLayerMock) Print(text string) {
	l.printedStrings = append(l.printedStrings, text)
}

func (l InteractionLayerMock) VerifyPrinted(t *testing.T, text string) {
	assert.Contains(t, l.printedStrings, text)
}
