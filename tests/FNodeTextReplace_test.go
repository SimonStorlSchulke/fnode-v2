package tests

import (
	"fnode2/core"
	"fnode2/nodes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFNodeTestReplace(t *testing.T) {
	tree := core.NodeTree{}

	testSubject, _ := nodes.Create("Text Replace")

	tree.AddNode(testSubject)

	testSubject.SetInputDefaultValue(0, "Hello World! World!")
	testSubject.SetInputDefaultValue(1, "World")
	testSubject.SetInputDefaultValue(2, "Mars")
	testSubject.SetOption("Mode", "String")

	result := testSubject.GetOutputResult(0)

	assert.Equal(t, "Hello Mars! Mars!", result)
}
