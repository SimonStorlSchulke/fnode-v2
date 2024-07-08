package core

import (
	"os"
	"path/filepath"
)

type FileList struct {
	LooseFiles  []string
	Directories []struct {
		Path      string
		Recursive bool
	}
}

/* GetFlatList returns a flat list of filepaths from the FileList */
func (list *FileList) GetFlatList() []FFile {
	fileInfos := []FFile{}

	for _, filePath := range list.LooseFiles {
		info, err := os.Lstat(filePath)
		if err != nil {
			continue
		}
		fileInfos = append(fileInfos, FFile{FullPath: filePath, Info: info})
	}

	for _, dir := range list.Directories {

		if dir.Recursive {
			fileInfos = append(fileInfos, getDirFilesRecursive(dir.Path)...)
		} else {
			fileInfos = append(fileInfos, getDirFilesFlat(dir.Path)...)
		}
	}

	return fileInfos
}

// getDirFilesRecursive returns all files in a folder and its subdirectories
func getDirFilesRecursive(path string) []FFile {
	fileInfos := []FFile{}

	err := filepath.Walk(path,
		func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() { //TODO - what to do with directories?
				fileInfos = append(fileInfos, FFile{FullPath: filePath, Info: info})
			}
			return nil
		})
	if err != nil {
		LogWarn("File not found: %s", path)
	}
	return fileInfos
}

// getDirFilesFlat gets all files in a directory, NOT including files in subdirectories
func getDirFilesFlat(path string) []FFile {
	fileInfos := []FFile{}

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		LogWarn("File not found: %s", path)
	}
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			fullPath := filepath.Join(path, entry.Name())
			fileInfos = append(fileInfos, FFile{FullPath: fullPath, Info: info})
		}
	}
	return fileInfos
}
