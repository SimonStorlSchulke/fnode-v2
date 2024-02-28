package core

const (
	FTypeFloat = iota
	FTypeInt
	FTypeString
	FTypeStringList
	FTypeFloatList
	FTypeIntList
)

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
