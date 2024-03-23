package tests

import (
	"fnode2/core"
	"fnode2/nodes"
	"fnode2/treeIo"
	"testing"
)

func TestNodeTree_OutputValue(t *testing.T) {
	tree := core.NodeTree{}

	vn, _ := nodes.Create("Value")
	vn.SetInputDefaultValue(0, 4.0)

	math1, _ := nodes.Create("Math")
	math1.SetOption("Mode", "Add")
	math1.SetInputDefaultValue(0, 5.0)
	math1.SetInputDefaultValue(1, 10.0)

	math2, _ := nodes.Create("Math")
	math2.SetOption("Mode", "Multiply")
	math2.SetInputDefaultValue(1, 2.0)

	printer, _ := nodes.Create("Print")

	tree.AddNode(math2)
	tree.AddNode(math1)
	tree.AddNode(vn)
	tree.AddNode(printer)

	tree.Link(vn.Id, 0, math2.Id, 1)
	tree.Link(math1.Id, 0, math2.Id, 0)
	tree.Link(math2.Id, 0, printer.Id, 0)

	il := &InteractionLayerMock{}
	tree.Parse(il, &core.FileList{LooseFiles: []string{"testfile"}})

	treeIo.SaveToFile(&tree, "./assets", "testfile1")

	il.VerifyPrinted(t, "60.00")
}
