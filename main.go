package main

import (
	"embed"
	"fnode2/controller"
	"fnode2/core"
	"fnode2/tests"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {

	tree := core.NodeTree{}

	tree.Parse(&tests.InteractionLayerMock{})

	//treeIo.SaveToFile(&tree, "testfiles", "hello-toml")
	app := controller.NewApp()

	err := wails.Run(&options.App{
		Title:     "fnode2",
		Width:     1024,
		Height:    768,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 27, B: 27, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
