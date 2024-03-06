package controller

import (
	"context"
	"fmt"
	"fnode2/core"
	"fnode2/core/InteractionLayer"
	"fnode2/nodes"
	"strconv"
)

type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	fmt.Println("YAAAAAa")
}

var tree core.NodeTree

func (a *App) ParseTree() []string {
	il := InteractionLayer.InteractionLayerExecute{}
	tree.Parse(&il)
	return il.GetOutput()
}

func (a *App) UpdateInputDefaultValue(nodeId string, inputIndex int, value string, valueType int) {

	var val any

	switch valueType {
	case core.FTypeBool:
		if value == "true" {
			val = true
		} else {
			val = false
		}
		break
	case core.FTypeFloat:
		val, _ = strconv.ParseFloat(value, 64)
	case core.FTypeString:
		val = value
	}

	fmt.Println(nodeId)
	fmt.Println(inputIndex)
	fmt.Println(value)

	//kinda ugly. maybe better save nodes in a map with id as key
	for _, node := range tree.Nodes {
		if node.Id == nodeId {
			node.Inputs[inputIndex].DefaultValue = val
			return
		}
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetTestTree() core.SerializableTree {
	tree = core.NodeTree{}

	vn, _ := nodes.Create("Value")
	vn.SetInputDefaultValue(0, 4.0)

	math1, _ := nodes.Create("Math")
	math1.SetOption("Mode", "Add")
	math1.SetInputDefaultValue(1, 10.0)

	math2, _ := nodes.Create("Math")
	math2.SetOption("Mode", "Multiply")
	math2.SetInputDefaultValue(0, 2.0)

	printer, _ := nodes.Create("Print")

	printer.Meta.PosY = 200
	printer.Meta.PosX = 700

	vn.Meta.PosX = 20
	vn.Meta.PosY = 200

	math1.Meta.PosX = 250
	math1.Meta.PosY = 60

	math2.Meta.PosX = 500
	math2.Meta.PosY = 120

	tree.AddNode(math2)
	tree.AddNode(math1)
	tree.AddNode(vn)
	tree.AddNode(printer)

	tree.Link(vn.Id, 0, math2.Id, 1)
	tree.Link(math1.Id, 0, math2.Id, 0)
	tree.Link(math2.Id, 0, printer.Id, 0)

	return tree.ToSerializable()
}
