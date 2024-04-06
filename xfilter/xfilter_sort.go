package xfilter

import (
	"bytes"
	"io"
	"sort"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

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
			str1, _ := s.UTF82GBK(list[i].GetName())
			str2, _ := s.UTF82GBK(list[j].GetName())
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

func (s *SortFilter[T]) UTF82GBK(str string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return io.ReadAll(transform.NewReader(bytes.NewReader([]byte(str)), GB18030.NewEncoder()))
}

func (s *SortFilter[T]) GBK2UTF8(str []byte) (string, error) {
	GB18030 := simplifiedchinese.All[0]
	b, err := io.ReadAll(transform.NewReader(bytes.NewReader(str), GB18030.NewDecoder()))
	return string(b), err
}
