package components

import (
	"strconv"
	"errors"
)

func TurnNumberToString(a interface{}) (string, error){
	switch a.(type){
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return strconv.Itoa(a.(int)), nil
	case float64:
		return strconv.FormatFloat(a.(float64),'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(a.(float32)),'f', -1, 32), nil
	default:
		err := errors.New("当前类型无法转为数字")
		return "", err
	}
}
