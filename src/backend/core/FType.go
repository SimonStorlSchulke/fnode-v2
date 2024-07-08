package core

import "os"

const (
	FTypeFloat = iota
	FTypeInt
	FTypeString
	FTypeBool
	FTypeFile
	FTypeStringList
	FTypeFloatList
	FTypeIntList
)

type FFile struct {
	FullPath string
	Info     os.FileInfo
}
