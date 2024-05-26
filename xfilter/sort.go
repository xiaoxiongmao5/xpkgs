package xfilter

import (
	"sort"
	"time"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type SortByStringAble interface {
	GetStrSortField() string
}

func SortByString[T SortByStringAble](list []T, sortType string) []T {
	// 创建一个用于中文字符排序的比较器
	c := collate.New(language.Chinese, collate.Numeric)

	// 使用 sort.SliceStable 保持稳定排序
	sort.SliceStable(list, func(i, j int) bool {
		if sortType == "desc" {
			return c.CompareString(list[i].GetStrSortField(), list[j].GetStrSortField()) > 0
		}
		return c.CompareString(list[i].GetStrSortField(), list[j].GetStrSortField()) < 0
	})

	return list
}

type SortByTimeAble interface {
	GetTmSortField() time.Time
}

func SortByTime[T SortByTimeAble](list []T, sortType string) []T {
	sort.SliceStable(list, func(i, j int) bool {
		if sortType == "desc" {
			return list[i].GetTmSortField().After(list[j].GetTmSortField())
		}
		return list[i].GetTmSortField().Before(list[j].GetTmSortField())
	})
	return list
}

type SortByNumAble interface {
	GetIntSortField() int64
}

func SortByNum[T SortByNumAble](list []T, sortType string) []T {
	sort.SliceStable(list, func(i, j int) bool {
		if sortType == "desc" {
			return list[i].GetIntSortField() > list[j].GetIntSortField()
		}
		return list[i].GetIntSortField() < list[j].GetIntSortField()
	})
	return list
}
