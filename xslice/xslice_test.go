package xslice

import (
	"reflect"
	"testing"
)

// 判断切片中是否存在某个元素
func TestInSlice(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	ok := InSlice[int](list, 1)
	Equal(t, []bool{ok}, []bool{true})

	ok = InSlice[int](list, 6)
	Equal(t, []bool{ok}, []bool{false})

	list1 := []string{"a", "b", "c", "d", "e"}
	ok = InSlice[string](list1, "a")
	Equal(t, []bool{ok}, []bool{true})

	ok = InSlice[string](list1, "f")
	Equal(t, []bool{ok}, []bool{false})

	list2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	ok = InSlice[float64](list2, 1.1)
	Equal(t, []bool{ok}, []bool{true})

	ok = InSlice[float64](list2, 6.6)
	Equal(t, []bool{ok}, []bool{false})
}

// 判断切片中是否存在某些元素（或|且）
func TestInSliceByLogic(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	ok, _ := InSliceByLogic[int](list, []int{1, 2, 3}, "or")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[int](list, []int{1, 2, 3}, "and")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[int](list, []int{1, 2, 6}, "or")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[int](list, []int{1, 2, 6}, "and")
	Equal(t, []bool{ok}, []bool{false})

	ok, _ = InSliceByLogic[int](list, []int{6, 7, 8}, "or")
	Equal(t, []bool{ok}, []bool{false})

	ok, _ = InSliceByLogic[int](list, []int{6, 7, 8}, "and")
	Equal(t, []bool{ok}, []bool{false})

	ok, _ = InSliceByLogic[int](list, []int{1, 2, 3, 4, 5}, "or")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[int](list, []int{1, 2, 3, 4, 5}, "and")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[int](list, []int{1, 2, 3, 4, 5, 6}, "or")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[int](list, []int{1, 2, 3, 4, 5, 6}, "and")
	Equal(t, []bool{ok}, []bool{false})

	list1 := []string{"a", "b", "c", "d", "e"}
	ok, _ = InSliceByLogic[string](list1, []string{"a", "b", "c"}, "or")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[string](list1, []string{"a", "b", "c"}, "and")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[string](list1, []string{"a", "b", "f"}, "or")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[string](list1, []string{"a", "b", "f"}, "and")
	Equal(t, []bool{ok}, []bool{false})

	ok, _ = InSliceByLogic[string](list1, []string{"f", "g", "h"}, "or")
	Equal(t, []bool{ok}, []bool{false})

	ok, _ = InSliceByLogic[string](list1, []string{"f", "g", "h"}, "and")
	Equal(t, []bool{ok}, []bool{false})

	ok, _ = InSliceByLogic[string](list1, []string{"a", "b", "c", "d", "e"}, "or")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[string](list1, []string{"a", "b", "c", "d", "e"}, "and")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[string](list1, []string{"a", "b", "c", "d", "e", "f"}, "or")
	Equal(t, []bool{ok}, []bool{true})

	ok, _ = InSliceByLogic[string](list1, []string{"a", "b", "c", "d", "e", "f"}, "and")
	Equal(t, []bool{ok}, []bool{false})
}

// 从切片中删除部分元素
func TestRemoveSomeFromSlice(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	list = RemoveSomeFromSlice[int](list, []int{1, 2, 3})
	Equal(t, list, []int{4, 5})

	list1 := []string{"a", "b", "c", "d", "e"}
	list1 = RemoveSomeFromSlice[string](list1, []string{"a", "b", "c"})
	Equal(t, list1, []string{"d", "e"})

	list2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	list2 = RemoveSomeFromSlice[float64](list2, []float64{1.1, 2.2, 3.3})
	Equal(t, list2, []float64{4.4, 5.5})

	list3 := []int{1, 2, 3, 4, 5}
	list3 = RemoveSomeFromSlice[int](list3, []int{6, 7, 8})
	Equal(t, list3, []int{1, 2, 3, 4, 5})

	list4 := []int{1, 2, 3, 4, 5}
	list4 = RemoveSomeFromSlice[int](list4, []int{1, 2, 3, 4, 5})
	Equal(t, list4, []int{})

	list5 := []int{1, 2, 3, 4, 5}
	list5 = RemoveSomeFromSlice[int](list5, []int{1, 2, 3, 4, 5, 6})
	Equal(t, list5, []int{})

	list6 := []int{1, 2, 3, 4, 5}
	list6 = RemoveSomeFromSlice[int](list6, []int{6, 7, 8, 9, 10})
	Equal(t, list6, []int{1, 2, 3, 4, 5})

	list7 := []string{"a", "b", "c", "d", "e"}
	list7 = RemoveSomeFromSlice[string](list7, []string{"f", "g", "h"})
	Equal(t, list7, []string{"a", "b", "c", "d", "e"})

	list8 := []string{"a", "b", "c", "d", "e"}
	list8 = RemoveSomeFromSlice[string](list8, []string{"a", "b", "c", "d", "e"})
	Equal(t, list8, []string{})

	list9 := []string{"a", "b", "c", "d", "e"}
	list9 = RemoveSomeFromSlice[string](list9, []string{"a", "b", "c", "d", "e", "f"})
	Equal(t, list9, []string{})
}

// 切片去重
func TestRemoveDuplicates(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 3, 2, 1}
	list = RemoveDuplicates[int](list)
	Equal(t, list, []int{1, 2, 3, 4, 5})

	list1 := []string{"a", "b", "c", "d", "e", "c", "b", "a"}
	list1 = RemoveDuplicates[string](list1)
	Equal(t, list1, []string{"a", "b", "c", "d", "e"})

	list2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 3.3, 2.2, 1.1}
	list2 = RemoveDuplicates[float64](list2)
	Equal(t, list2, []float64{1.1, 2.2, 3.3, 4.4, 5.5})

	list3 := []int{1, 2, 3, 4, 5}
	list3 = RemoveDuplicates[int](list3)
	Equal(t, list3, []int{1, 2, 3, 4, 5})
}

func Equal[T any](t *testing.T, a, b []T) {
	if len(a) != len(b) {
		t.Errorf("Expected length to be %d, but got %d", len(a), len(b))
	}
	for i := range a {
		v1 := reflect.ValueOf(a[i])
		v2 := reflect.ValueOf(b[i])
		switch v1.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if v1.Int() != v2.Int() {
				t.Errorf("Expected value to be %v, but got %v", v1, v2)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if v1.Uint() != v2.Uint() {
				t.Errorf("Expected value to be %v, but got %v", v1, v2)
			}
		case reflect.Float32, reflect.Float64:
			if v1.Float() != v2.Float() {
				t.Errorf("Expected value to be %v, but got %v", v1, v2)
			}
		case reflect.String:
			if v1.String() != v2.String() {
				t.Errorf("Expected value to be %v, but got %v", v1, v2)
			}
		case reflect.Bool:
			if v1.Bool() != v2.Bool() {
				t.Errorf("Expected value to be %v, but got %v", v1, v2)
			}
		default:
			t.Errorf("Unsupported type: %v", v1.Kind())
		}
	}
}
