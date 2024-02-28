package core

import (
	"fmt"
	"strconv"
)

// First Key: From Type second Key: To Type
var typeConversions map[int]map[int]func(input any) any = map[int]map[int]func(input any) any{
	FTypeFloat: {
		FTypeString: func(input any) any {
			return fmt.Sprintf("%.2f", input.(float64))
		},
		FTypeInt: func(input any) any {
			return int(input.(float64)) //discards the fraction
		},
	},
	FTypeString: {
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
	default:
		panic("unknown FType in fallbackValue()")
	}
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
