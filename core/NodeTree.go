package core

import (
	"fmt"
	"fnode2/core/InteractionLayer"
)

type NodeTree struct {
	Nodes map[string]*Node
	Links []*NodeLink
}

func (tree *NodeTree) AddNode(node *Node) {
	node.Tree = tree
	if tree.Nodes == nil {
		tree.Nodes = map[string]*Node{}
	}
	tree.Nodes[node.Id] = node
}

func (tree *NodeTree) findExecutiveNodes() []*Node {
	var executiveNodes []*Node
	for _, node := range tree.Nodes {
		if node.ExecutiveFunction != nil {
			executiveNodes = append(executiveNodes, node)
		}
	}
	return executiveNodes
}

func (tree *NodeTree) Parse(layer InteractionLayer.NodeInteractionLayer) {
	tree.removeOutputCaches()
	executives := tree.findExecutiveNodes()

	for _, executiveNode := range executives {

		inputValues := make([]any, len(executiveNode.Inputs))

		for i, _ := range executiveNode.Inputs {
			inputValues[i] = executiveNode.GetInputValue(i)
		}
		executiveNode.ExecutiveFunction(layer, inputValues, executiveNode.Options)
	}
}

func (tree *NodeTree) removeOutputCaches() {
	for _, node := range tree.Nodes {
		node.RemoveCaches()
	}
}

func (tree *NodeTree) RemoveNode(node *Node) {
	delete(tree.Nodes, node.Id)
}

func (tree *NodeTree) AddLink(link *NodeLink) {
	if link.FromNode == link.ToNode {
		fmt.Printf("\ncannot connect sockets of the same Node %s", link.FromNode)
	}

	tree.Links = append(tree.Links, link)
}

func (tree *NodeTree) Link(fromNode string, fromOutput int, toNode string, toInput int) {
	link := &NodeLink{
		FromNode:   fromNode,
		FromOutput: fromOutput,
		ToNode:     toNode,
		ToInput:    toInput,
	}

	if link.FromNode == link.ToNode {
		fmt.Printf("\ncannot connect sockets of the same Node %s", link.FromNode)
	}
	tree.Links = append(tree.Links, link)
}

type SerializableTree struct {
	Nodes []SerializableNode
	Links []*NodeLink
}

func (tree *NodeTree) ToSerializable() SerializableTree {
	nodes := make([]SerializableNode, len(tree.Nodes))

	i := 0
	for _, node := range tree.Nodes {
		nodes[i] = node.ToSerializable()
		i++
	}

	return SerializableTree{
		Nodes: nodes,
		Links: tree.Links,
	}
}
