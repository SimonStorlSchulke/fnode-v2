package main

import (
	"fmt"
	"slices"
)

type FStringValue struct {
	value string
}

type NodeLink struct {
	FromNode   int
	FromOutput int
	ToNode     int
	ToInput    int
}

func (v FStringValue) get() string {
	return v.value
}

type NodeInput[t any] struct {
	Name         string
	DefaultValue t
}

type NodeOutput[t any] func(inputs []any) t

type Node interface {
	Inputs() []NodeInput[any]
	Outputs() []NodeOutput[any]
}

type SNode struct {
	Id      int
	Inputs  []NodeInput[any]
	Outputs []NodeOutput[any]
}

func NewValueNode() *SNode {
	return &SNode{
		Inputs: []NodeInput[any]{
			{
				Name:         "value",
				DefaultValue: 1.0,
			},
		},
		Outputs: []NodeOutput[any]{
			func(inputs []any) any {
				return inputs[0].(float64)
			},
		},
	}
}

var NodeLinks []NodeLink = []NodeLink{
	{
		FromNode:   0,
		FromOutput: 0,
		ToNode:     1,
		ToInput:    1,
	},
}

var Nodes map[int]*SNode = map[int]*SNode{
	0: NewValueNode().SetInputDefaultValue(0, 4.0).SetId(0),
	1: NewMathNode().SetInputDefaultValue(0, 2.0).SetId(1),
}

func NewMathNode() *SNode {
	return &SNode{
		Inputs: []NodeInput[any]{
			{
				Name:         "a",
				DefaultValue: 1.0,
			},
			{
				Name:         "b",
				DefaultValue: 1.0,
			},
		},
		Outputs: []NodeOutput[any]{
			func(inputs []any) any {
				return inputs[0].(float64) * inputs[1].(float64)
			},
		},
	}
}

func (n *SNode) SetInputDefaultValue(index int, value any) *SNode {
	n.Inputs[index].DefaultValue = value
	return n
}

func (n *SNode) SetId(id int) *SNode {
	n.Id = id
	return n
}

func (n *SNode) GetInputValue(index int) any {
	//connectedTo :=
	matchingLinkIndex := slices.IndexFunc(NodeLinks, func(link NodeLink) bool { return link.ToNode == n.Id && link.ToInput == index })

	if matchingLinkIndex == -1 {
		return n.Inputs[index].DefaultValue
	} else {
		link := NodeLinks[matchingLinkIndex]
		fmt.Println(matchingLinkIndex)
		return Nodes[link.FromNode].OutputValue(link.FromOutput)
	}

}

func (n *SNode) OutputValue(index int) any {
	inputValues := make([]any, len(n.Inputs))

	for i, _ := range n.Inputs {
		inputValues[i] = n.GetInputValue(i)
	}

	return n.Outputs[index](inputValues)
}

func main() {
	g := Nodes[1].OutputValue(0)
	fmt.Println(g)
}
