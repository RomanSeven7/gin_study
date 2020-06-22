package util

import (
	"reflect"
	"strconv"
)

func Number2Float64(v interface{}, kind reflect.Kind) (float64, bool) {
	switch kind {
	case reflect.Int:
		return float64(v.(int)), true
	case reflect.Int8:
		return float64(v.(int8)), true
	case reflect.Int16:
		return float64(v.(int16)), true
	case reflect.Int32:
		return float64(v.(int32)), true
	case reflect.Int64:
		return float64(v.(int64)), true
	case reflect.Uint:
		return float64(v.(uint)), true
	case reflect.Uint8:
		return float64(v.(uint8)), true
	case reflect.Uint16:
		return float64(v.(uint16)), true
	case reflect.Uint32:
		return float64(v.(uint32)), true
	case reflect.Uint64:
		return float64(v.(uint64)), true
	case reflect.Float32:
		return float64(v.(float32)), true
	case reflect.Float64:
		return float64(v.(float64)), true
	}
	return float64(0), false
}

/**
基础类型转string
*/
func Primary2String(v interface{}, kind reflect.Kind) (string, bool) {
	switch kind {
	case reflect.Int:
		return strconv.Itoa(v.(int)), true
	case reflect.Int8:
		return strconv.FormatInt(int64(v.(int8)), 10), true
	case reflect.Int16:
		return strconv.FormatInt(int64(v.(int16)), 10), true
	case reflect.Int32:
		return strconv.FormatInt(int64(v.(int32)), 10), true
	case reflect.Int64:
		return strconv.FormatInt(int64(v.(int64)), 10), true
	case reflect.Uint:
		return strconv.FormatUint(uint64(v.(uint)), 10), true
	case reflect.Uint8:
		return strconv.FormatUint(uint64(v.(uint8)), 10), true
	case reflect.Uint16:
		return strconv.FormatUint(uint64(v.(uint16)), 10), true
	case reflect.Uint32:
		return strconv.FormatUint(uint64(v.(uint32)), 10), true
	case reflect.Uint64:
		return strconv.FormatUint(uint64(v.(uint64)), 10), true
	case reflect.Float32:
		return strconv.FormatFloat(float64(v.(float32)), 'f', 6, 64), true
	case reflect.Float64:
		return strconv.FormatFloat(float64(v.(float64)), 'f', 6, 64), true
	case reflect.String:
		return v.(string), true
	}
	return "", false
}
