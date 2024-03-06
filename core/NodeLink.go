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
