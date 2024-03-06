package treeIo

import (
	"fnode2/core"
	"fnode2/nodes"
	"github.com/pelletier/go-toml"
	"os"
)

func deserialize(data []byte) (*core.NodeTree, error) {
	var serializableTree = &core.SerializableTree{}
	err := toml.Unmarshal(data, serializableTree)
	if err != nil {
		return nil, err
	}

	tree := &core.NodeTree{}

	for _, serializedNode := range serializableTree.Nodes {
		node, err := deserializeNode(serializedNode)
		if err != nil {
			core.Log(err.Error(), core.LogLevelError)
			return nil, err
		}
		tree.AddNode(node)
	}

	for _, link := range serializableTree.Links {
		tree.AddLink(link)
	}

	return tree, nil
}

func deserializeNode(serializedNode core.SerializableNode) (*core.Node, error) {
	node, err := nodes.Create(serializedNode.Type)
	if err != nil {
		return nil, err
	}
	node.Id = serializedNode.Id
	node.Meta = serializedNode.Meta

	for i, serializedInput := range serializedNode.Inputs {
		node.Inputs[i].DefaultValue = serializedInput.DefaultValue
	}

	for i, serializedOption := range serializedNode.Options {
		node.Options[i].SelectedOption = serializedOption.SelectedOption
	}

	return node, nil
}

func LoadFromFile(path string) *core.NodeTree {
	data, _ := os.ReadFile(path)
	tree, _ := deserialize(data)
	return tree
}
