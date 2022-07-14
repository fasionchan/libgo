/*
 * Author: fasion
 * Created time: 2022-04-15 11:15:54
 * Last Modified by: fasion
 * Last Modified time: 2022-07-14 14:42:53
 */

package generics

import (
	"golang.org/x/exp/constraints"
)

type Slice[T any] []T

func (slice Slice[T]) Len() int {
	return len(slice)
}

func (slice Slice[T]) Empty() bool {
	return slice.Len() == 0
}

func (slice Slice[T]) First() (value T) {
	if !slice.Empty() {
		value = slice[0]
	}
	return
}

func (slice Slice[T]) Last() (value T) {
	n := slice.Len()
	if n > 0 {
		value = slice[n-1]
	}
	return
}

func (slice Slice[T]) Dup() Slice[T] {
	return make(Slice[T], 0, slice.Len()).Append(slice...)
}

func (slice Slice[T]) Reverse() Slice[T] {
	n := slice.Len()
	for i := 0; i < n/2; i++ {
		slice[i], slice[n-i-1] = slice[n-i-1], slice[i]
	}
	return slice
}

func (slice Slice[T]) Append(values ...T) Slice[T] {
	return append(slice, values...)
}

func (slice Slice[T]) Concat(others ...Slice[T]) Slice[T] {
	for _, values := range others {
		slice = slice.Append(values...)
	}
	return slice
}

func (slice Slice[T]) ForEach(f func(i int, value T)) Slice[T] {
	for i, value := range slice {
		f(i, value)
	}
	return slice
}

func (slice Slice[T]) ForEachLite(f func(value T)) Slice[T] {
	return slice.ForEach(func(i int, value T) {
		f(value)
	})
}

func (slice Slice[T]) Filter(f func(i int, value T) bool) Slice[T] {
	result := make(Slice[T], 0, slice.Len())
	for i, value := range slice {
		if f(i, value) {
			result = result.Append(value)
		}
	}
	return result
}

func (slice Slice[T]) FilterLite(f func(value T) bool) Slice[T] {
	return slice.Filter(func(i int, value T) bool {
		return f(value)
	})
}

func (slice Slice[T]) Purge(f func(i int, value T) bool) Slice[T] {
	return slice.Filter(func(i int, value T) bool {
		return !f(i, value)
	})
}

func (slice Slice[T]) PurgeLite(f func(value T) bool) Slice[T] {
	return slice.Purge(func(i int, value T) bool {
		return f(value)
	})
}

func (slice Slice[T]) Strings(f func(i int, value T) string) StringSlice {
	result := make(StringSlice, 0, len(slice))
	for i, value := range slice {
		result = StringSlice(result.Slice().Append(f(i, value)))
	}
	return result
}

func (slice Slice[T]) StringsLite(f func(value T) string) StringSlice {
	return slice.Strings(func(i int, value T) string {
		return f(value)
	})
}

type ComparableValueSlice[T comparable] Slice[T]

func (slice ComparableValueSlice[T]) Slice() Slice[T] {
	return Slice[T](slice)
}

func (slice ComparableValueSlice[T]) Filter(value T) ComparableValueSlice[T] {
	result := make(ComparableValueSlice[T], 0, slice.Slice().Len())
	for _, v := range slice {
		if v == value {
			result = append(result, v)
		}
	}
	return result
}

func (slice ComparableValueSlice[T]) Purge(value T) ComparableValueSlice[T] {
	result := make(ComparableValueSlice[T], 0, slice.Slice().Len())
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

func (slice ComparableValueSlice[T]) PurgeZero() ComparableValueSlice[T] {
	var zero T
	return slice.Purge(zero)
}

type OrderedSlice[T constraints.Ordered] ComparableValueSlice[T]

func (slice OrderedSlice[T]) Native() []T {
	return []T(slice)
}

func (slice OrderedSlice[T]) Slice() Slice[T] {
	return (Slice[T])(slice)
}

func (slice OrderedSlice[T]) ComparableValueSlice() ComparableValueSlice[T] {
	return (ComparableValueSlice[T])(slice)
}

func (slice OrderedSlice[T]) Compare(other OrderedSlice[T]) int {
	otherLength := other.Slice().Len()
	for i, s := range slice {
		if i >= otherLength {
			return 1
		}

		o := other[i]
		switch {
		case s < o:
			return -1
		case s > o:
			return 1
		}
	}

	if slice.Slice().Len() < otherLength {
		return -1
	}

	return 0
}

func (slice OrderedSlice[T]) Less(other OrderedSlice[T]) bool {
	return slice.Compare(other) < 0
}

func (slice OrderedSlice[T]) Greater(other OrderedSlice[T]) bool {
	return slice.Compare(other) > 0
}

func (slice OrderedSlice[T]) Equal(other OrderedSlice[T]) bool {
	return slice.Compare(other) == 0
}

type AnyPointer = *any

type PointerSlice[P AnyPointer] ComparableValueSlice[P]

func (slice PointerSlice[P]) Slice() Slice[P] {
	return Slice[P](slice)
}

func (slice PointerSlice[P]) ComparableValueSlice() ComparableValueSlice[P] {
	return ComparableValueSlice[P](slice)
}

func (slice PointerSlice[P]) PurgeNil() PointerSlice[P] {
	return PointerSlice[P](slice.ComparableValueSlice().Purge(nil))
}

type StringSlice ComparableValueSlice[string]

func (slice StringSlice) Slice() Slice[string] {
	return Slice[string](slice)
}

func (slice StringSlice) ComparableValueSlice() ComparableValueSlice[string] {
	return ComparableValueSlice[string](slice)
}

type IntSlice = OrderedSlice[int]
