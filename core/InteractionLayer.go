package core

// NodeInteractionLayer is the layer used while parsing the nodetree for actually
// doing things like printing to the console, writing, renaming or deleting files
// (or mocking these actions)
type NodeInteractionLayer interface {
	Print(text string)
}
