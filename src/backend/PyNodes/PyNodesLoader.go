package PyNodes

import (
	"bufio"
	"fnode2/core"
	"os"
	"path"
	"strings"
)

func pythonOutput(funcName string) {

}

func PyToNode(pythonScriptPath string) {
	file, err := os.Open(pythonScriptPath)

	if err != nil {
		//	log.Print(err)
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			//	log.Print(err)
			return
		}
	}()

	node := core.NewNodeCreator(path.Base(pythonScriptPath), "custom", make([]core.NodeInput, 0), make([]*core.NodeOutput, 0))

	scanner := bufio.NewScanner(file)

	lineIsInput := false
	lineIsOutput := false
	lineIsOption := false

	for scanner.Scan() {
		line := scanner.Text()

		if lineIsInput {
			node.Inputs = append(node.Inputs, core.NewNodeInput(core.FTypeFloat, "test", 1.0))
		}
		if lineIsOutput {
			node.Outputs = append(node.Outputs, core.NewNodeOutput(core.FTypeFloat, "test", func(node *core.Node) any {
				pythonOutput("bla")
				return 0.0 // --> run python script
			}, false)) //todo eval cache here
		}

		if lineIsOption {
			node.AddOption("", []string{})
		}

		lineIsInput = strings.HasPrefix(line, "# FInput")
		lineIsOutput = strings.HasPrefix(line, "# FOutput")
		lineIsOption = strings.HasPrefix(line, "# FOption")

	}

}
