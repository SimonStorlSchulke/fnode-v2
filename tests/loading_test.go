package tests

import (
	"fnode2/treeIo"
	"testing"
)

func TestLoadFromaFile_Loads(t *testing.T) {
	tree := treeIo.LoadFromFile("assets/testfile1.fn")
	il := &InteractionLayerMock{}

	tree.Parse(il)
	il.VerifyPrinted(t, "60.00")
}
