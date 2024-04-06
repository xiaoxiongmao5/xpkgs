package xfilter

import (
	"bytes"
	"io"
	"sort"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type FilterOption[T any] interface {
	Apply(list []T) []T
}

func FilterList[T any](list []T, opts ...FilterOption[T]) []T {
	for _, opt := range opts {
		list = opt.Apply(list)
	}
	return list
}

// PageFilter 分页过滤器
type PageFilter[T any] struct {
	PageNo   int64 //从1开始
	PageSize int64 //0表示不分页
}

func (p *PageFilter[T]) Apply(list []T) []T {
	length := int64(len(list))
	if length == 0 || p.PageSize == 0 {
		return list
	}
	if p.PageNo <= 0 {
		p.PageNo = 1
	}
	if p.PageSize < 0 {
		p.PageSize = 0
	}
	start := (p.PageNo - 1) * p.PageSize
	if start >= length {
		return []T{}
	}
	end := start + p.PageSize
	if end > length {
		end = length
	}
	return list[start:end]
}

type SortFilterAble interface {
	GetName() string
	GetCreatedAt() time.Time
}

// SortFilter 排序过滤器
// 支持包含中英文数字等的字符串排序
type SortFilter[T SortFilterAble] struct {
	SortType  string //排序类型 asc/desc 默认desc
	SortField string //排序字段 name/createdTm 默认createdTm
}

func (s *SortFilter[T]) Apply(list []T) []T {
	sort.SliceStable(list, func(i, j int) bool {
		if s.SortField == "name" {
			str1, _ := UTF82GBK(list[i].GetName())
			str2, _ := UTF82GBK(list[j].GetName())
			if s.SortType == "asc" {
				return bytes.Compare(str1, str2) < 0
			} else {
				return bytes.Compare(str1, str2) > 0
			}
		} else {
			if s.SortType == "asc" {
				return list[i].GetCreatedAt().Before(list[j].GetCreatedAt())
			} else {
				return list[i].GetCreatedAt().After(list[j].GetCreatedAt())
			}
		}
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
