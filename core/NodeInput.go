package core

const (
	FTypeFloat = iota
	FTypeInt
	FTypeString
	FTypeAny
	FTypeStringList
	FTypeFloatList
	FTypeIntList
)

type NodeInput struct {
	Name         string
	Type         int
	DefaultValue any
}

func NewFloatInput(name string, initialDefaultValue float64) NodeInput {
	return NodeInput{
		Name:         name,
		Type:         FTypeFloat,
		DefaultValue: initialDefaultValue,
	}
}

func NewStringInput(name string, initialDefaultValue string) NodeInput {
	return NodeInput{
		Name:         name,
		Type:         FTypeString,
		DefaultValue: initialDefaultValue,
	}
}
