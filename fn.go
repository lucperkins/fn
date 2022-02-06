package fn

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 | ~string
}

func ForEach[T any](items []T, forEachFn func(T)) {
	for _, item := range items {
		forEachFn(item)
	}
}

func Map[T any](items []T, mapFn func(T) T) []T {
	for idx, item := range items {
		items[idx] = mapFn(item)
	}
	return items
}

func Find[T comparable](items []T, findFn func(T) bool) *T {
	for _, item := range items {
		if findFn(item) {
			return &item
		}
	}
	return nil
}

func Filter[T any](items []T, filterFn func(T) bool) []T {
	is := make([]T, 0)

	for _, item := range items {
		if filterFn(item) {
			is = append(is, item)
		}
	}

	return is
}

func Concat[T any](items1, items2 []T) []T {
	for _, item := range items2 {
		items1 = append(items1, item)
	}
	return items1
}

func Every[T any](items []T, everyFn func(T) bool) bool {
	for _, item := range items {
		if !everyFn(item) {
			return false
		}
	}

	return true
}

func Any[T any](items []T, anyFn func(T) bool) bool {
	var satisfied bool
	for _, item := range items {
		if anyFn(item) {
			return true
		}
	}
	return satisfied
}

func Index[T comparable](items []T, wanted T) int {
	for idx, item := range items {
		if item == wanted {
			return idx
		}
	}

	return -1
}

func LastIndexOf[T comparable](items []T, wanted T) int {
	last := -1

	for idx, item := range items {
		if item == wanted {
			last = idx
		}
	}

	return last
}

func Reduce[T any](items []T, reduceFn func(T, T) T) T {
	var value T
	for _, item := range items {
		value = reduceFn(value, item)
	}
	return value
}

func Reverse[T any](items []T) []T {
	is := make([]T, 0)
	for idx := len(items) - 1; idx >= 0; idx-- {
		is = append(is, items[idx])
	}
	return is
}

func Pop[T any](items []T, idx int) (T, []T) {
	n := items[idx]
	return n, append(items[:idx], items[idx+1:]...)[:len(items)-1]
}

func Push[T any](items []T, item T) []T {
	return append(items, item)
}

func Includes[T comparable](items []T, wanted T) bool {
	var includes bool
	for _, item := range items {
		if item == wanted {
			return true
		}
	}
	return includes
}

func Max[T Ordered](items []T) T {
	max := items[0]
	for _, item := range items {
		if item > max {
			max = item
		}
	}
	return max
}

func Min[T Ordered](items []T) T {
	min := items[0]
	for _, item := range items {
		if item < min {
			min = item
		}
	}
	return min
}

func Replace[T comparable](items []T, discard, replace T) []T {
	for i := 0; i < len(items)-1; i++ {
		if items[i] == discard {
			items[i] = replace
			break
		}
	}

	return items
}

func ReplaceAll[T comparable](items []T, discard, replace T) []T {
	for idx, item := range items {
		if item == discard {
			items[idx] = replace
		}
	}
	return items
}

func Delete[T comparable](items []T, discard T) []T {
	for idx, item := range items {
		if item == discard {
			items = append(items[:idx], items[idx+1:]...)
			return items
		}
	}

	return items
}

func DeleteAll[T comparable](items []T, discard T) []T {
	is := make([]T, 0)

	for _, item := range items {
		if item != discard {
			is = append(is, item)
		}
	}

	return is
}
