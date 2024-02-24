package core

import "slices"

type NodeTree struct {
	Nodes []*Node
	Links []*NodeLink
}

func (t *NodeTree) GetConnectedOutput(ofNodeId, ofInputId int) NodeOutput[any] {
	matchingLinkIndex := slices.IndexFunc(t.Links, func(link *NodeLink) bool { return link.ToNode == ofNodeId && link.ToInput == ofInputId })

	link := t.Links[matchingLinkIndex]

	return t.Nodes[link.FromNode].Outputs[link.FromOutput]
}
