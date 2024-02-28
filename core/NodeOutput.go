package core

type NodeOutput struct {
	Type     int
	Name     string
	Operator NodeOutputFunc
}

func NewNodeOutput(outputType int, name string, operator NodeOutputFunc) *NodeOutput {
	return &NodeOutput{
		Type:     outputType,
		Name:     name,
		Operator: operator,
	}
}
