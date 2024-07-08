package core

import (
	"fmt"
	"slices"

	"github.com/beevik/guid"
)

type Node struct {
	Type   string
	Id     string
	Inputs []NodeInput

	/*by default, one of each repeatableInputs will be added to Inputs in NewNodeCreator()
	more can be added via SetRepeatableInputsAmount */
	repeatableInputs       []NodeInput
	repeatableInputsAmount int
	Outputs                []*NodeOutput
	Tree                   *NodeTree
	Options                map[string]*NodeOption
	ExecutiveFunction      ExecutiveFunction
	Meta                   NodeMeta
	cachedOutputResults    []any
}

type ExecutiveFunction func(interactionLayer NodeInteractionLayer, inputs []any, Options map[string]*NodeOption)

type NodeMeta struct {
	PosX     int
	PosY     int
	Category string
}

type NodeOption struct {
	Choices        []string
	SelectedOption string
	Callback       func(node *Node, selectedChoice string)
}

type SerializableNodeOption struct {
	Choices        []string
	SelectedOption string
}

func (nodeOption *NodeOption) ToSerializable() SerializableNodeOption {
	return SerializableNodeOption{Choices: nodeOption.Choices, SelectedOption: nodeOption.SelectedOption}
}

func (node *Node) AddRepeatableInputGroup(inputs []NodeInput) {
	node.repeatableInputs = inputs
}

func (node *Node) SetRepeatableInputsAmount(number int) {
	if number < 1 {
		LogWarn("Cannot set amount of repeatableInputGroupSize < 1")
		return
	}

	difference := number - node.repeatableInputsAmount

	if difference > 0 {
		node.addRepeatableInputs(difference)
	}

	if difference < 0 {
		node.removeRepeatableInputs(difference)
	}
}

func (node *Node) addRepeatableInputs(difference int) {
	for i := 0; i < difference; i++ {
		for _, r := range node.repeatableInputs {
			node.Inputs = append(node.Inputs, r)
		}
	}
}

func (node *Node) removeRepeatableInputs(difference int) {
	inputsToSlice := -difference * len(node.repeatableInputs)
	deletionStartIndex := len(node.Inputs) - 1 - inputsToSlice
	node.Inputs = slices.Delete(node.Inputs, deletionStartIndex, len(node.Inputs)-1)
}

func (node *Node) SetInputDefaultValue(index int, value any) {
	node.Inputs[index].DefaultValue = value
	LogInfo(
		"Updated default_input of %s input '%v' to '%v'",
		node.Id, node.Inputs[index].Name, value)
}

func (node *Node) AddOption(key string, choices []string) {
	if node.Options == nil {
		node.Options = map[string]*NodeOption{}
	}
	node.Options[key] = &NodeOption{Choices: choices}
	node.SetOptionCallback(key, nil)
	node.Options[key].SelectedOption = choices[0]
}

// SetOptionCallback is triggered when an Option with the given key is changed
func (node *Node) SetOptionCallback(key string, callback func(node *Node, selectedChoice string)) {
	node.Options[key].Callback = func(node *Node, selectedChoice string) {
		LogInfo("Set Option of %s %s to %s", node.Id, key, node.Options[key].SelectedOption)
		if callback != nil {
			callback(node, selectedChoice)
		}
	}
}

func (node *Node) SetOption(key string, choice string) error {
	if node.Options[key] == nil {
		return fmt.Errorf("\n%s is not a valid option key", key)
	}

	idx := slices.IndexFunc(node.Options[key].Choices, func(choice string) bool { return choice == choice })
	if idx == -1 {
		return fmt.Errorf("%s is not a valid choice for option with key %s", choice, key)
	}

	node.Options[key].SelectedOption = choice
	node.Options[key].Callback(node, choice)
	return nil
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

func (node *Node) GetInputFloat(inputId int) float64 {
	return node.GetInputValue(inputId).(float64)
}

func (node *Node) GetInputInt(inputId int) int64 {
	return node.GetInputValue(inputId).(int64)
}

func (node *Node) GetInputString(inputId int) string {
	return node.GetInputValue(inputId).(string)
}

func (node *Node) GetInputBool(inputId int) bool {
	return node.GetInputValue(inputId).(bool)
}

func (node *Node) GetInputFile(inputId int) FFile {
	return node.GetInputValue(inputId).(FFile)
}

func (node *Node) GetInputValue(inputId int) any {
	inputLink := node.findLinkOfInput(inputId)
	if inputLink == nil {
		return node.Inputs[inputId].DefaultValue
	}

	connectedNode := node.Tree.Nodes[inputLink.FromNode]

	if connectedNode == nil {
		LogPanic("Node not found in NodeTree")
		return node.Inputs[inputId].DefaultValue
	}

	outputType := connectedNode.Outputs[inputLink.FromOutput].Type
	inputType := node.Inputs[inputId].Type

	outputValue := connectedNode.GetOutputResult(inputLink.FromOutput)
	return AutoConvertTypes(outputType, inputType, outputValue)
}

func (node *Node) GetOutputResult(index int) any {
	hasCache := node.Outputs[index].CacheEnabled && node.cachedOutputResults[index] != nil

	if hasCache {
		return node.cachedOutputResults[index]
	}

	inputValues := make([]any, len(node.Inputs))

	for i, _ := range node.Inputs {
		inputValues[i] = node.GetInputValue(i)
	}

	result := node.Outputs[index].GetResult(node)

	if node.Outputs[index].CacheEnabled {
		node.cachedOutputResults[index] = result
	}

	return result
}

func (node *Node) RemoveCaches() {
	for i, _ := range node.cachedOutputResults {
		node.cachedOutputResults[i] = nil
	}
}

func NewNodeCreator(nodeType string, category string, inputs []NodeInput, outputs []*NodeOutput) *Node {
	id := nodeType + "_" + guid.New().String()

	return &Node{
		Type:                nodeType,
		Id:                  id,
		Inputs:              inputs,
		Outputs:             outputs,
		cachedOutputResults: make([]any, len(outputs)),
		Meta:                NodeMeta{Category: category},
	}
}

type SerializableNode struct {
	Type    string
	Id      string
	Inputs  []NodeInput
	Outputs []SerializableNodeOutput
	Options map[string]SerializableNodeOption
	Meta    NodeMeta
}

func (node *Node) ToSerializable() SerializableNode {

	options := make(map[string]SerializableNodeOption, len(node.Options))

	for i, option := range node.Options {
		options[i] = option.ToSerializable()
	}

	outputs := make([]SerializableNodeOutput, len(node.Outputs))
	for i, output := range node.Outputs {
		outputs[i] = output.ToSerializable()
	}

	return SerializableNode{
		Type:    node.Type,
		Id:      node.Id,
		Inputs:  node.Inputs,
		Options: options,
		Outputs: outputs,
		Meta:    node.Meta,
	}
}
