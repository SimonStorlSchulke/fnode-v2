package core

//TODO - should represent reusable nodetrees with in- and outputs that can be used like a node

type NodeGroup struct {
	Inputs []NodeInput
	tree   *NodeTree
}
