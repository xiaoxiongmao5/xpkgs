package xfilter

import (
	"bytes"
	"io"
	"reflect"
	"sort"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// SortFilterByReflect 反射排序过滤器
// 支持包含中英文数字等的字符串排序
type SortFilterByReflect[T any] struct {
	SortType  string //排序类型 asc/desc 默认desc
	SortField string //排序字段
}

func (s *SortFilterByReflect[T]) Apply(list []T) []T {
	return SortByReflect(list, s.SortField, s.SortType)
}

func SortByReflect[T any](list []T, field string, sortType string) []T {
	sort.SliceStable(list, func(i, j int) bool {
		// 通过反射获取字段值
		v1 := reflect.ValueOf(list[i])
		v2 := reflect.ValueOf(list[j])
		if v1.Kind() == reflect.Ptr {
			v1 = v1.Elem()
			v2 = v2.Elem()
		}
		if v1.Kind() != reflect.Struct {
			return false
		}
		f1 := v1.FieldByName(field)
		f2 := v2.FieldByName(field)

		switch f1.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if sortType == "asc" {
				return f1.Int() < f2.Int()
			} else {
				return f1.Int() > f2.Int()
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if sortType == "asc" {
				return f1.Uint() < f2.Uint()
			} else {
				return f1.Uint() > f2.Uint()
			}
		case reflect.Float32, reflect.Float64:
			if sortType == "asc" {
				return f1.Float() < f2.Float()
			} else {
				return f1.Float() > f2.Float()
			}
		case reflect.String:
			str1, _ := UTF82GBK(f1.String())
			str2, _ := UTF82GBK(f2.String())
			if sortType == "asc" {
				return bytes.Compare(str1, str2) < 0
			} else {
				return bytes.Compare(str1, str2) > 0
			}
		case reflect.Struct:
			if f1.Type().String() == "time.Time" {
				if sortType == "asc" {
					return f1.Interface().(time.Time).Before(f2.Interface().(time.Time))
				} else {
					return f1.Interface().(time.Time).After(f2.Interface().(time.Time))
				}
			}
		}
		return false
	})
	return list
}

func UTF82GBK(str string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return io.ReadAll(transform.NewReader(bytes.NewReader([]byte(str)), GB18030.NewEncoder()))
}

func GBK2UTF8(str []byte) (string, error) {
	GB18030 := simplifiedchinese.All[0]
	b, err := io.ReadAll(transform.NewReader(bytes.NewReader(str), GB18030.NewDecoder()))
	return string(b), err
}
