package fun

// Ordered encompasses commonly used types that can be used in eq/gt/lt operations.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

// ForEach applies the specified return-less function to every item in a slice.
func ForEach[T any](items []T, fn func(T)) {
	for _, item := range items {
		fn(item)
	}
}

// Map applies the specified function to each member of a slice and returns a new slice.
func Map[T any](items []T, fn func(T) T) []T {
	for idx, item := range items {
		items[idx] = fn(item)
	}
	return items
}

// Filter takes only those items in the slice that satisfy the supplied condition function.
func Filter[T any](items []T, fn func(T) bool) []T {
	is := make([]T, 0)

	for _, item := range items {
		if fn(item) {
			is = append(is, item)
		}
	}

	return is
}

// Concat combines two slices into a single slice, preserving order.
func Concat[T any](items1, items2 []T) []T {
	for _, item := range items2 {
		items1 = append(items1, item)
	}
	return items1
}

// Every determines whether each item in a slice satisfies the supplied condition.
func Every[T any](items []T, fn func(T) bool) bool {
	for _, item := range items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Any determines whether any item in a slice satisfies the supplied condition.
func Any[T any](items []T, fn func(T) bool) bool {
	var satisfied bool
	for _, item := range items {
		if fn(item) {
			return true
		}
	}
	return satisfied
}

// Reduce applies the supplied function rightward through the slice and returns the result.
func Reduce[T any](items []T, fn func(T, T) T) T {
	var value T
	for _, item := range items {
		value = fn(value, item)
	}
	return value
}

// Reverse returns the supplied slice in reverse order.
func Reverse[T any](items []T) []T {
	is := make([]T, 0)
	for idx := len(items) - 1; idx >= 0; idx-- {
		is = append(is, items[idx])
	}
	return is
}

// Pop returns a tuple of the item at the specified index and the remaining slice after that item is removed.
func Pop[T any](items []T, idx int) (T, []T) {
	n := items[idx]
	return n, append(items[:idx], items[idx+1:]...)[:len(items)-1]
}

// Push adds an item to the end of the slice.
func Push[T any](items []T, item T) []T {
	return append(items, item)
}

// Find takes a condition and returns either a pointer to the first value found using that satisfies the condition or nil if no such item is found.
func Find[T comparable](items []T, fn func(T) bool) *T {
	for _, item := range items {
		if fn(item) {
			return &item
		}
	}
	return nil
}

// Includes determines whether the supplied item is included in the slice.
func Includes[T comparable](items []T, wanted T) bool {
	var includes bool
	for _, item := range items {
		if item == wanted {
			return true
		}
	}
	return includes
}

// Index returns the index of the first instance of the supplied item found or -1 if none is found.
func Index[T comparable](items []T, wanted T) int {
	for idx, item := range items {
		if item == wanted {
			return idx
		}
	}

	return -1
}

// LastIndexOf returns the index of the last instance of the supplied item found or -1 if none is found.
func LastIndexOf[T comparable](items []T, wanted T) int {
	last := -1

	for idx, item := range items {
		if item == wanted {
			last = idx
		}
	}

	return last
}

// Replace replaces the first instance of the supplied value with the desired replacement value.
func Replace[T comparable](items []T, discard, replace T) []T {
	for i := 0; i < len(items)-1; i++ {
		if items[i] == discard {
			items[i] = replace
			break
		}
	}

	return items
}

// Replace replaces all instances of the supplied value with the desired replacement value.
func ReplaceAll[T comparable](items []T, discard, replace T) []T {
	for idx, item := range items {
		if item == discard {
			items[idx] = replace
		}
	}
	return items
}

// Delete removes the first instance of an item from a slice and returns the resulting slice.
func Delete[T comparable](items []T, discard T) []T {
	for idx, item := range items {
		if item == discard {
			items = append(items[:idx], items[idx+1:]...)
			return items
		}
	}

	return items
}

// Delete removes all instances of an item from a slice and returns the resulting slice.
func DeleteAll[T comparable](items []T, discard T) []T {
	is := make([]T, 0)

	for _, item := range items {
		if item != discard {
			is = append(is, item)
		}
	}

	return is
}

// Max returns the maximum value of the slice of Ordered items.
func Max[T Ordered](items []T) T {
	max := items[0]
	for _, item := range items {
		if item > max {
			max = item
		}
	}
	return max
}

// Min returns the minimum value of the slice of Ordered items.
func Min[T Ordered](items []T) T {
	min := items[0]
	for _, item := range items {
		if item < min {
			min = item
		}
	}
	return min
}
