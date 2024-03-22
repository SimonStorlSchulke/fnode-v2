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

func (a *App) ParseTree() {
	il := core.InteractionLayerExecute{}
	tree.Parse(&il, fileList)
}

func (a *App) ParseTreePreview() {
	il := core.InteractionLayerPreview{}
	tree.Parse(&il, fileList)
}

func (a *App) ClearTree() {
	tree = core.NodeTree{}
}

func (a *App) UpdateNodePosition(nodeId string, posX int, posY int) {
	node, err := tree.FindNodeById(nodeId)
	if err != nil {
		core.LogError("error updating Node position %v", err)
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
		core.LogError("Node with ID %s could not be found in tree", nodeId)
		return
	}
	node.SetInputDefaultValue(inputIndex, val)
}

// UpdateUption returns true on success
func (a *App) UpdateUption(nodeId string, key string, selectedChoice string) bool {
	node, err := tree.FindNodeById(nodeId)
	if err != nil {
		core.LogError("Node with ID %s could not be found in tree", nodeId)
		return false
	}
	err = node.SetOption(key, selectedChoice)
	if err != nil {
		core.LogRawError(err)
	}
	return true
}

func (a *App) GetTree() core.SerializableTree {
	return tree.ToSerializable()
}
