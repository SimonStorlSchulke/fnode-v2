package core

import (
	"fmt"
	"os"
	"strconv"
)

// typeConversions is a dictionary of functions to convert from one outputType type to another inputType First Key: From Type second Key: To Type
var typeConversions map[int]map[int]func(input any) any = map[int]map[int]func(input any) any{
	FTypeFloat: {
		FTypeString: func(input any) any {
			return fmt.Sprintf("%.2f", input.(float64))
		},
		FTypeInt: func(input any) any {
			return int(input.(float64)) //discards the fraction
		},
		FTypeBool: func(input any) any {
			if input.(float64) > 0 {
				return true
			}
			return false
		},
	},
	FTypeInt: {
		FTypeString: func(input any) any {
			return fmt.Sprintf("%v", input.(int))
		},
		FTypeFloat: func(input any) any {
			return float64(input.(int)) //discards the fraction
		},
		FTypeBool: func(input any) any {
			if input.(int) > 0 {
				return true
			}
			return false
		},
	},
	FTypeString: {
		FTypeBool: func(input any) any {
			if input.(string) == "true" {
				return true
			}
			return false
		},
		FTypeFloat: func(input any) any {
			f, err := strconv.ParseFloat(input.(string), 64)
			if err != nil {
				return 0.0
			}
			return f
		},
		FTypeInt: func(input any) any {
			i, err := strconv.ParseInt(input.(string), 1, 32)
			if err != nil {
				return 0
			}
			return int(i)
		},
		FTypeFile: func(input any) any {
			fileInfo, _ := os.Lstat(input.(string))
			return FFile{
				FullPath: input.(string),
				Info:     fileInfo,
			}
		},
	},
	FTypeBool: {
		FTypeString: func(input any) any {
			if input.(bool) {
				return "true"
			}
			return "false"
		},
		FTypeFloat: func(input any) any {
			if input.(bool) {
				return 1.0
			}
			return 0.0
		},
		FTypeInt: func(input any) any {
			if input.(bool) {
				return 1
			}
			return 0
		},
	},
	FTypeFile: {
		FTypeString: func(input any) any {
			if input == nil {
				return ""
			}
			return input.(FFile).FullPath
		},
	},
}

func fallbackValue(ofType int) any {
	switch ofType {
	case FTypeFloat:
		return 0.0
	case FTypeInt:
		return 0
	case FTypeString:
		return ""
	case FTypeBool:
		return false
	}
	LogPanic("unknown FType in fallbackValue()")
	return nil
}

func AutoConvertTypes(fromType int, toType int, value any) any {
	if fromType == toType {
		return value
	}

	if typeConversions[fromType] == nil || typeConversions[fromType][toType] == nil {
		return fallbackValue(toType)
	}

	return typeConversions[fromType][toType](value)
}
