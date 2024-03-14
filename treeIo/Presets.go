package treeIo

import (
	"os"
	"strings"
)

func GetPresetNames() ([]string, error) {
	files, err := os.ReadDir("./data/presets")

	if err != nil {
		return nil, err
	}

	paths := []string{}

	for _, file := range files {

		fileName := file.Name()

		if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
			paths = append(paths, fileName[:pos])
		} else {

			paths = append(paths, fileName)
		}
	}

	return paths, nil
}
