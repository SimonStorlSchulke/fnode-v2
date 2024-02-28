package core

import (
	"fmt"
	"github.com/beevik/guid"
	"slices"
)

type NodeOutputFunc func(inputs []any, Options map[string]*NodeOption) any
type ExecutiveFunction func(inputs []any, Options map[string]*NodeOption)

type NodeOption struct {
	Choices        []string
	SelectedOption string
	Callback       func(node *Node)
}

type SerializableNodeOption struct {
	SelectedOption string
}

func (nodeOption *NodeOption) ToSerializable() SerializableNodeOption {
	return SerializableNodeOption{SelectedOption: nodeOption.SelectedOption}
}

type Node struct {
	Type              string
	Id                string
	Inputs            []NodeInput
	Outputs           []*NodeOutput
	Tree              *NodeTree
	Options           map[string]*NodeOption
	ExecutiveFunction ExecutiveFunction
}

func (node *Node) SetInputDefaultValue(index int, value any) {
	node.Inputs[index].DefaultValue = value
}

func (node *Node) AddOption(key string, choices []string) {
	if node.Options == nil {
		node.Options = map[string]*NodeOption{}
	}
	node.Options[key] = &NodeOption{Choices: choices}
	node.SetOptionCallback(key, nil)
	node.Options[key].SelectedOption = choices[0]
}

func (node *Node) SetOptionCallback(key string, callback func(node *Node)) {
	node.Options[key].Callback = func(node *Node) {
		Log("Set Option of %s %s to %s", LogLevelInfo, node.Id, key, node.Options[key].SelectedOption)
		if callback != nil {
			callback(node)
		}
	}
}

func (node *Node) SetOption(key string, choice string) {
	if node.Options[key] == nil {
		fmt.Printf("\n%s is not a valid option key", key)
		return
	}

	idx := slices.IndexFunc(node.Options[key].Choices, func(choice string) bool { return choice == choice })
	if idx == -1 {
		Log("%s is not a valid choice for option with key %s", LogLevelError, choice, key)
		return
	}

	node.Options[key].SelectedOption = choice
	node.Options[key].Callback(node)
}

func (node *Node) findLinksOfOutput(outputId int) []*NodeLink {
	var matchingLinks []*NodeLink
	for _, link := range node.Tree.Links {
		if link.FromNode == node.Id && link.FromOutput == outputId {
			matchingLinks = append(matchingLinks, link)
		}
	}
	return matchingLinks
}

func (node *Node) findLinkOfInput(inputId int) *NodeLink {
	matchingLinkIndex := slices.IndexFunc(node.Tree.Links, func(link *NodeLink) bool { return link.ToNode == node.Id && link.ToInput == inputId })
	if matchingLinkIndex == -1 {
		return nil
	}
	return node.Tree.Links[matchingLinkIndex]
}

func (node *Node) GetInputValue(inputId int) any {
	inputLink := node.findLinkOfInput(inputId)
	if inputLink == nil {
		return node.Inputs[inputId].DefaultValue
	}

	connectedNode := node.Tree.Nodes[inputLink.FromNode]

	if connectedNode == nil {
		Log("Node not found in NodeTree", LogLevelError)
		return node.Inputs[inputId].DefaultValue
	}

	/*	if()

		node.Inputs[inputId]*/

	return connectedNode.OutputValue(inputLink.FromOutput)
}

func (node *Node) OutputValue(index int) any {
	inputValues := make([]any, len(node.Inputs))

	for i, _ := range node.Inputs {
		inputValues[i] = node.GetInputValue(i)
	}

	return node.Outputs[index].Operator(inputValues, node.Options)
}

func NewNode(nodeType string, inputs []NodeInput, outputs []*NodeOutput) *Node {
	id := nodeType + "_" + guid.New().String()
	return &Node{
		Type:    nodeType,
		Id:      id,
		Inputs:  inputs,
		Outputs: outputs,
	}
}

type SerializableNode struct {
	Type               string
	Id                 string
	InputDefaultValues []any
	Options            map[string]SerializableNodeOption
}

func (node *Node) ToSerializable() SerializableNode {

	defaultValues := make([]any, len(node.Inputs))
	options := make(map[string]SerializableNodeOption, len(node.Options))

	for i, input := range node.Inputs {
		defaultValues[i] = input.DefaultValue
	}

	for i, option := range node.Options {
		options[i] = option.ToSerializable()
	}

	return SerializableNode{
		Type:               node.Type,
		Id:                 node.Id,
		InputDefaultValues: defaultValues,
		Options:            options,
	}
}
