package core

type NodeLink struct {
	FromNode   string
	FromOutput int
	ToNode     string
	ToInput    int
}

func NewLink(fromNode string, fromOutput int, toNode string, toInput int) *NodeLink {
	return &NodeLink{
		FromNode:   fromNode,
		FromOutput: fromOutput,
		ToNode:     toNode,
		ToInput:    toInput,
	}
}

func (link1 *NodeLink) Equals(link2 *NodeLink) bool {
	return link1.FromNode == link2.FromNode &&
		link1.FromOutput == link2.FromOutput &&
		link1.ToNode == link2.ToNode &&
		link1.ToInput == link2.ToInput
}
