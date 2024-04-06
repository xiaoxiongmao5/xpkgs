package xfilter

type FilterOption[T any] interface {
	Apply([]T) []T
}

func FilterList[T any](list []T, opts ...FilterOption[T]) []T {
	for _, opt := range opts {
		list = opt.Apply(list)
	}
	return list
}
