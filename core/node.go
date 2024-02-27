package core

import "github.com/beevik/guid"

type Node struct {
	Type    string
	Id      string
	Inputs  []NodeInput[any]
	Outputs []NodeOutput[any]
}

func (node *Node) SetInputDefaultValue(index int, value any) *Node {
	node.Inputs[index].DefaultValue = value
	return n
}

/* func (n *Node) GetInputValue(index int, nodeList []Node) any {
	//connectedTo :=
	matchingLinkIndex := slices.IndexFunc(NodeLinks, func(link NodeLink) bool { return link.ToNode == n.Id && link.ToInput == index })

	if matchingLinkIndex == -1 {
		return n.Inputs[index].DefaultValue
	} else {
		link := NodeLinks[matchingLinkIndex]
		fmt.Println(matchingLinkIndex)
		return nodeList[link.FromNode].OutputValue(link.FromOutput)
	}
} */

/* func (n *Node) OutputValue(index int) any {
	inputValues := make([]any, len(n.Inputs))

	for i, _ := range n.Inputs {
		inputValues[i] = n.GetInputValue(i)
	}

	return n.Outputs[index](inputValues)
} */

func NewNode(nodeType string, inputs []NodeInput[any], outputs []NodeOutput[any]) *Node {
	id := nodeType + "_" + guid.New().String()
	return &Node{
		Type:    nodeType,
		Id:      id,
		Inputs:  inputs,
		Outputs: outputs,
	}
}
