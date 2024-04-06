package xslice

import "fmt"

type UtilsSliceAble interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

// 判断切片中是否存在某个元素
func InSlice[T UtilsSliceAble](slice []T, val T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// 判断切片中是否存在某些元素（或|且）
func InSliceByLogic[T UtilsSliceAble](slice []T, vals []T, logic string) (bool, error) {
	if logic != "or" && logic != "and" {
		return false, fmt.Errorf("InSliceByLogic: logic must be 'or' or 'and'")
	}

	for _, v := range vals {
		if logic == "or" {
			if ok := InSlice(slice, v); ok {
				return true, nil
			}
		} else {
			if ok := InSlice(slice, v); !ok {
				return false, nil
			}
		}
	}
	return logic == "and", nil
}

// 从切片中删除部分元素
func RemoveSomeFromSlice[T UtilsSliceAble](slice []T, vals []T) []T {
	if len(slice) == 0 || len(vals) == 0 {
		return slice
	}

	result := make([]T, 0)
	for _, v := range slice {
		if ok := InSlice(vals, v); ok {
			continue
		}
		result = append(result, v)
	}

	return result
}

// 切片去重
func RemoveDuplicates[T UtilsSliceAble](slice []T) []T {
	encountered := make(map[T]struct{})
	result := make([]T, 0)

	for _, v := range slice {
		if _, ok := encountered[v]; !ok {
			encountered[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
