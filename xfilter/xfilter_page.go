package xfilter

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
