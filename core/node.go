package core

import (
	"fmt"
	"slices"
)

type Node struct {
	Id      int
	Inputs  []NodeInput[any]
	Outputs []NodeOutput[any]
}

func (n *Node) SetInputDefaultValue(index int, value any) *Node {
	n.Inputs[index].DefaultValue = value
	return n
}

func (n *Node) SetId(id int) *Node {
	n.Id = id
	return n
}

var NodeLinks []NodeLink = []NodeLink{
	{
		FromNode:   0,
		FromOutput: 0,
		ToNode:     1,
		ToInput:    1,
	},
}

func (n *Node) GetInputValue(index int, nodeList []Node) any {
	//connectedTo :=
	matchingLinkIndex := slices.IndexFunc(NodeLinks, func(link NodeLink) bool { return link.ToNode == n.Id && link.ToInput == index })

	if matchingLinkIndex == -1 {
		return n.Inputs[index].DefaultValue
	} else {
		link := NodeLinks[matchingLinkIndex]
		fmt.Println(matchingLinkIndex)
		return nodeList[link.FromNode].OutputValue(link.FromOutput)
	}
}

func (n *Node) OutputValue(index int) any {
	inputValues := make([]any, len(n.Inputs))

	for i, _ := range n.Inputs {
		inputValues[i] = n.GetInputValue(i)
	}

	return n.Outputs[index](inputValues)
}
