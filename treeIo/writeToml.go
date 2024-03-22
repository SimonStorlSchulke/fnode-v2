package treeIo

import (
	"fmt"
	"fnode2/core"
	"os"
	"path"
	"strings"

	"github.com/pelletier/go-toml"
)

func serialize(tree *core.NodeTree) ([]byte, error) {
	serializable := tree.ToSerializable()
	data, err := toml.Marshal(serializable)
	return data, err
}

func createDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		return err
	}
	return nil
}

func SaveToFile(tree *core.NodeTree, directory string, fileName string) {
	data, err := serialize(tree)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createDirIfNotExists(directory)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !strings.HasSuffix(fileName, ".fn") {
		fileName = fileName + ".fn"
	}

	fullPath := path.Join(directory, fileName)

	core.LogInfo("Saving Tree to File %s", fullPath)

	err = os.WriteFile(fullPath, data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
