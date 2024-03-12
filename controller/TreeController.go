package controller

import (
	"context"
	"fnode2/core"
	"strconv"
)

type App struct {
	ctx context.Context
}

var appContext context.Context

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) SetContext(ctx context.Context) {
	appContext = a.ctx
}

var tree core.NodeTree = core.NodeTree{}

func (a *App) ParseTree() {
	il := core.InteractionLayerExecute{}
	tree.Parse(&il)
}

func (a *App) ClearTree() {
	tree = core.NodeTree{}
}

func (a *App) UpdateNodePosition(nodeId string, posX int, posY int) {
	node, err := tree.FindNodeById(nodeId)
	if err != nil {
		core.Log("error updating Node position %v", core.LogLevelError, err)
	}
	node.Meta.PosX = posX
	node.Meta.PosY = posY
}

func (a *App) UpdateInputDefaultValue(nodeId string, inputIndex int, value any, valueType int) {
	var val any
	switch valueType {
	case core.FTypeBool:
		val = value.(bool)
	case core.FTypeFloat:
		val, _ = strconv.ParseFloat(value.(string), 64)
	case core.FTypeInt:
		intVal, _ := strconv.ParseInt(value.(string), 10, 64)
		val = int(intVal) // TODO check int overflow
	case core.FTypeString:
		val = value.(string)
	}

	node, err := tree.FindNodeById(nodeId)
	if err != nil {
		core.Log("Node with ID %s could not be found in tree", core.LogLevelError, nodeId)
		return
	}
	node.SetInputDefaultValue(inputIndex, val)
}

// UpdateUption returns true on success
func (a *App) UpdateUption(nodeId string, key string, selectedChoice string) bool {
	node, err := tree.FindNodeById(nodeId)
	if err != nil {
		core.Log("Node with ID %s could not be found in tree", core.LogLevelError, nodeId)
		return false
	}
	err = node.SetOption(key, selectedChoice)
	if err != nil {
		core.LogErr(err)
	}
	return true
}

func (a *App) GetTestTree() core.SerializableTree {
	/*vn, _ := nodes.Create("Math.Value")
	vn.SetInputDefaultValue(0, 4.0)

	math1, _ := nodes.Create("Math.Math")
	math1.SetOption("Mode", "Add")
	math1.SetInputDefaultValue(1, 10.0)

	math2, _ := nodes.Create("Math.Math")
	math2.SetOption("Mode", "Multiply")
	math2.SetInputDefaultValue(0, 2.0)

	printer, _ := nodes.Create("Output.Print")

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
	tree.Link(math2.Id, 0, printer.Id, 0)*/

	return tree.ToSerializable()
}
