package core

type NodeInput struct {
	Name         string
	Type         int
	DefaultValue any
}

func NewNodeInput(Type int, name string, initialDefaultValue any) NodeInput {
	return NodeInput{
		Name:         name,
		Type:         Type,
		DefaultValue: initialDefaultValue,
	}
}
