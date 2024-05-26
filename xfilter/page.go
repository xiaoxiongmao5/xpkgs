package xfilter

func PageFilterList[T any](list []T, pageNo, pageSize int64) []T {
	length := int64(len(list))
	if length == 0 || pageSize == 0 {
		return list
	}
	if pageNo < 0 {
		pageNo = 0
	}
	if pageSize < 0 {
		pageSize = 0
	}
	offset := pageNo * pageSize
	if offset >= length {
		return []T{}
	}
	end := offset + pageSize
	if end > length {
		end = length
	}
	return list[offset:end]
}
