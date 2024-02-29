package main

import (
	"embed"
	"fnode2/core"
	"fnode2/nodes"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {

	tree := core.NodeTree{}

	mn := nodes.NewMathNode()
	vn := nodes.NewValueNode()

	vn.SetInputDefaultValue(0, 4.0)

	mn.SetInputDefaultValue(0, 2.0)

	mn2 := nodes.NewMathNode()

	mn2.SetInputDefaultValue(1, 10.0)

	mn.SetOption("Mode", "Multiply")
	mn2.SetOption("Mode", "Add")

	tn := nodes.NewTextNode()
	tn.SetInputDefaultValue(0, "12.130")

	pn := nodes.NewPrintNode()

	tree.AddNode(mn)
	tree.AddNode(mn2)
	tree.AddNode(vn)
	tree.AddNode(tn)
	tree.AddNode(pn)

	tree.AddLink(&core.NodeLink{
		FromNode:   vn.Id,
		FromOutput: 0,
		ToNode:     mn.Id,
		ToInput:    1,
	})

	tree.AddLink(&core.NodeLink{
		FromNode:   tn.Id,
		FromOutput: 0,
		ToNode:     mn.Id,
		ToInput:    0,
	})

	tree.AddLink(&core.NodeLink{
		FromNode:   mn.Id,
		FromOutput: 0,
		ToNode:     pn.Id,
		ToInput:    0,
	})

	tree.Parse()

	//treeIo.SaveToFile(&tree, "testfiles", "hello-toml")
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "fnode2",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
