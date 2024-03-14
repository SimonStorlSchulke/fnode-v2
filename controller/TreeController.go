package controller

import (
	"context"
	"fnode2/core"
	"os"
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

// Startup is called when the app starts. TLeah Gotti

var tree core.NodeTree = core.NodeTree{}

func (a *App) ParseTree(list *core.FileList) {
	il := core.InteractionLayerExecute{}
	tree.Parse(&il, list)
}

func (a *App) ParseTreePreview(list *core.FileList) {
	il := core.InteractionLayerPreview{}
	tree.Parse(&il, list)
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
	case core.FTypeFile:
		val = value.(string)

		fileInfo, _ := os.Lstat(value.(string))
		val = core.FFile{
			FullPath: value.(string),
			Info:     fileInfo,
		}
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

func (a *App) GetTree() core.SerializableTree {
	return tree.ToSerializable()
}
