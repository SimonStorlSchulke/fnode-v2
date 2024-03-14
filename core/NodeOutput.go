package core

type NodeOutputFunc func(node *Node) any

type NodeOutput struct {
	Type         int
	Name         string
	Operator     NodeOutputFunc
	CacheEnabled bool
}

type SerializableNodeOutput struct {
	Type int
	Name string
}

func NewNodeOutput(outputType int, name string, operator NodeOutputFunc, cacheEnabled bool) *NodeOutput {
	return &NodeOutput{
		Type:         outputType,
		Name:         name,
		Operator:     operator,
		CacheEnabled: cacheEnabled,
	}
}

func (output *NodeOutput) GetResult(node *Node) any {
	result := output.Operator(node)
	return result
}

func (output *NodeOutput) ToSerializable() SerializableNodeOutput {
	return SerializableNodeOutput{
		Type: output.Type,
		Name: output.Name,
	}
}
