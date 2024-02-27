package core

import (
	"fmt"
	"slices"
)

type NodeTree struct {
	Nodes map[string]*Node
	Links []*NodeLink
}

func (tree *NodeTree) findLinksOfOutput(nodeId string, outputId int) []*NodeLink {
	var matchingLinks []*NodeLink
	for _, link := range tree.Links {
		if link.FromNode == nodeId && link.FromOutput == outputId {
			matchingLinks = append(matchingLinks, link)
		}
	}
	return matchingLinks
}

func (tree *NodeTree) findLinkOfInput(nodeId string, inputId int) *NodeLink {
	matchingLinkIndex := slices.IndexFunc(tree.Links, func(link *NodeLink) bool { return link.ToNode == nodeId && link.ToInput == inputId })
	if matchingLinkIndex == -1 {
		return nil
	}
	return tree.Links[matchingLinkIndex]
}

func (tree *NodeTree) GetConnectedOutput(ofNodeId string, ofInputId int) NodeOutput[any] {
	link := tree.findLinkOfInput(ofNodeId, ofInputId)
	return tree.Nodes[link.FromNode].Outputs[link.FromOutput]
}

func (tree *NodeTree) GetInputValue(ofNodeId string, ofInputId int) any {
	//connectedTo :=
/* 	matchingLinkIndex := slices.IndexFunc(NodeLinks, func(link NodeLink) bool { return link.ToNode == tree.Id && link.ToInput == index })

	if matchingLinkIndex == -1 {
		return tree.Inputs[index].DefaultValue
	} else {
		link := NodeLinks[matchingLinkIndex]
		fmt.Println(matchingLinkIndex)
		return Nodes[link.FromNode].OutputValue(link.FromOutput)
	} */
		

}

func (tree *NodeTree) GetOutputValue(ofNodeId string, ofOutputId int) any {
	output := tree.GetConnectedOutput(ofNodeId, ofOutputId)

	return output
}

func (tree *NodeTree) AddNode(node *Node) {
	tree.Nodes[node.Id] = node
}

func (tree *NodeTree) RemoveNode(node *Node) {
	delete(tree.Nodes, node.Id)
}
