package components

import "strconv"

func TurnNumbertoString(a interface{}) string{
	switch a.(type){
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return strconv.Itoa(a.(int))
	case float64:
		return strconv.FormatFloat(a.(float64),'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(a.(float32)),'f', -1, 32)
	}
}
