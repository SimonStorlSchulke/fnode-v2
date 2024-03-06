package core

type NodeOutputFunc func(inputs []any, Options map[string]*NodeOption) any
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

func (output *NodeOutput) GetResult(inputs []any, options map[string]*NodeOption) any {
	result := output.Operator(inputs, options)
	return result
}

func (output *NodeOutput) ToSerializable() SerializableNodeOutput {
	return SerializableNodeOutput{
		Type: output.Type,
		Name: output.Name,
	}
}
