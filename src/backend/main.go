package main

import (
	"embed"
	"fnode2/controller"
	"fnode2/core"
	"fnode2/nodes"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"golang.org/x/net/context"
)

//go:embed frontend_build
var assets embed.FS

type App struct {
	ctx context.Context
}

func main() {

	nodes.GenerateNodeCategories()

	//treeIo.SaveToFile(&tree, "testfiles", "hello-toml")
	app := controller.NewApp()

	err := wails.Run(&options.App{
		Title:  "fnode2",
		Width:  1600,
		Height: 900,

		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 27, B: 27, A: 1},
		OnStartup: func(ctx context.Context) {
			//app.SetContext(ctx)
			core.Ctx = ctx
		},

		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
