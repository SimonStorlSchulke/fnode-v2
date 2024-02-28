package main

import (
	"fnode2/core"
	"fnode2/nodes"
)

func main() {

	tree := core.NodeTree{}

	mn := nodes.NewMathNode()
	vn := nodes.NewValueNode()

	vn.SetInputDefaultValue(0, 4.0)

	mn.SetInputDefaultValue(0, 2.0)

	mn2 := nodes.NewMathNode()

	mn2.SetInputDefaultValue(1, 10.0)

	mn.SetOption("Mode", "Multiply")
	mn2.SetOption("Mode", "Add")

	tn := nodes.NewTextNode()
	tn.SetInputDefaultValue(0, "12.130")

	pn := nodes.NewPrintNode()

	tree.AddNode(mn)
	tree.AddNode(mn2)
	tree.AddNode(vn)
	tree.AddNode(tn)
	tree.AddNode(pn)

	tree.AddLink(&core.NodeLink{
		FromNode:   vn.Id,
		FromOutput: 0,
		ToNode:     mn.Id,
		ToInput:    1,
	})

	tree.AddLink(&core.NodeLink{
		FromNode:   tn.Id,
		FromOutput: 0,
		ToNode:     mn.Id,
		ToInput:    0,
	})

	tree.AddLink(&core.NodeLink{
		FromNode:   mn.Id,
		FromOutput: 0,
		ToNode:     pn.Id,
		ToInput:    0,
	})

	tree.Parse()

	//treeIo.SaveToFile(&tree, "testfiles", "hello-toml")
}
