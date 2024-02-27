package treeIo

import (
	"fmt"
	"fnode2/core"
	"github.com/pelletier/go-toml"
	"os"
	"path"
)

func Serialize(tree *core.NodeTree) ([]byte, error) {
	serializable := tree.ToSerializable()
	data, err := toml.Marshal(serializable)
	return data, err
}

func createDirIfNotExists(path string) error {
	var err error = nil
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
	}
	return err
}

func SaveToFile(tree *core.NodeTree, directory string, fileName string) {
	data, err := Serialize(tree)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createDirIfNotExists(directory)
	if err != nil {
		fmt.Println(err)
		return
	}

	fullPath := path.Join(directory, fileName+".fn")

	err = os.WriteFile(fullPath, data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
