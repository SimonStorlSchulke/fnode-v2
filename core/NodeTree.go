package core

import (
	"fmt"
	"slices"
)

type TreeRunState struct {
	CurrentIteration int
	CurrentFile      FFile
}

var RunState TreeRunState = TreeRunState{}

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

func (tree *NodeTree) FindNodeById(id string) (*Node, error) {
	foundNode := tree.Nodes[id]
	if foundNode == nil {
		return nil, fmt.Errorf("Can't find Node with ID %s.", id)
	}
	return foundNode, nil
}

func (tree *NodeTree) Parse(layer NodeInteractionLayer, fileList *FileList) {

	LogInfo("Parsing NodeTree")

	filesList := fileList.GetFlatList()

	tree.removeOutputCaches()
	executives := tree.findExecutiveNodes()

	for i, file := range filesList {
		RunState.CurrentIteration = i
		RunState.CurrentFile = file

		tree.parseIteration(executives, layer)
		tree.removeOutputCaches()
	}

	RunState = TreeRunState{}
}

func (tree *NodeTree) parseIteration(executiveNodes []*Node, layer NodeInteractionLayer) {
	for _, executiveNode := range executiveNodes {

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

	for i := len(tree.Links) - 1; i >= 0; i-- {
		connectedToRemovedNode := tree.Links[i].FromNode == node.Id || tree.Links[i].ToNode == node.Id
		if connectedToRemovedNode {
			tree.Links = append(tree.Links[:i], tree.Links[i+1:]...)
		}
	}
}

func (tree *NodeTree) AddLink(newLink *NodeLink) {
	if newLink.FromNode == newLink.ToNode {
		LogError("\ncannot connect sockets of the same Node %s", newLink.FromNode)
		return
	}

	if slices.ContainsFunc(tree.Links, func(link *NodeLink) bool {
		return link.Equals(newLink)
	}) {
		LogInfo("Link already exists:", newLink)
		return
	}

	LogInfo("Creating Link from %s [%v] tot %s [%v]:", newLink.FromNode, newLink.FromOutput, newLink.ToNode, newLink.ToInput)

	//If a link already exists for the requested input, find and remove it
	tree.Links = slices.DeleteFunc(tree.Links, func(link *NodeLink) bool {
		return link.ToNode == newLink.ToNode && link.ToInput == newLink.ToInput
	})

	tree.Links = append(tree.Links, newLink)
}

func (tree *NodeTree) RemoveLink(toRemove *NodeLink) {
	LogInfo("Removing Link", toRemove)
	tree.Links = slices.DeleteFunc(tree.Links, func(link *NodeLink) bool {
		return link.Equals(toRemove)
	})
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
