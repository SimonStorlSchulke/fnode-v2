package tests

import (
	"fnode2/core"
	"fnode2/nodes"
	"testing"
)

func TestNodeTree_OutputValue(t *testing.T) {

	tree := core.NodeTree{}

	vn := nodes.NewValueNode()
	vn.SetInputDefaultValue(0, 4.0)

	math1 := nodes.NewMathNode()
	math1.SetOption("Mode", "Add")
	math1.SetInputDefaultValue(1, 10.0)

	math2 := nodes.NewMathNode()
	math2.SetOption("Mode", "Multiply")
	math2.SetInputDefaultValue(0, 2.0)

	tree.AddNode(math2)
	tree.AddNode(math1)
	tree.AddNode(vn)

	tree.AddLink(&core.NodeLink{
		FromNode:   vn.Id,
		FromOutput: 0,
		ToNode:     math2.Id,
		ToInput:    1,
	})

	tree.AddLink(&core.NodeLink{
		FromNode:   math1.Id,
		FromOutput: 0,
		ToNode:     math2.Id,
		ToInput:    0,
	})

}
