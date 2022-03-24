package main

import (
	"fmt"
	"strings"
)

// List is a type for lazy linked list
type List[T any] func() (T, List[T])

// String implements stringer interface
func (l List[T]) String() string {
	var tmp []string
	WalkList(func(i T) {
		tmp = append(tmp, fmt.Sprintf("%v", i))
	}, l)
	tmp = append(tmp, "nil")
	return strings.Join(tmp, " => ")
}

func MkList[T any](head T) List[T] {
	return func() (T, List[T]) {
		return head, nil
	}
}

func MapList[I, O any](fn func(I) O, xs List[I]) List[O] {
	if xs == nil {
		return nil
	}
	return func() (O, List[O]) {
		x, xs := xs()
		return fn(x), MapList(fn, xs)
	}
}

func WalkList[T any](fn func(T), xs List[T]) {
	if xs == nil {
		return
	}
	x, xs := xs()
	fn(x)
	WalkList(fn, xs)
}

func AppendList[T any](v T, xs List[T]) List[T] {
	if xs == nil {
		return MkList(v)
	}
	return func() (T, List[T]) {
		x, xs := xs()
		return x, AppendList(v, xs)
	}
}

func PrependList[T any](v T, xs List[T]) List[T] {
	return func() (T, List[T]) {
		return v, xs
	}
}

func insertList[T any](
	fn func(T) bool, v T, inserter func(v T, x T, xs List[T],
	) (T, List[T]), xs List[T],
) List[T] {
	if xs == nil {
		return xs
	}
	return func() (T, List[T]) {
		x, xs := xs()
		if fn(x) {
			return inserter(v, x, xs)
		}
		return x, insertList(fn, v, inserter, xs)
	}
}

func InsertBeforeList[T any](
	fn func(T) bool, v T, xs List[T],
) List[T] {
	return insertList(fn, v, func(v, x T, xs List[T]) (T, List[T]) {
		return v, AppendList(x, xs)
	}, xs)
}

func InsertAfterList[T any](
	fn func(T) bool, v T, xs List[T],
) List[T] {
	return insertList(fn, v, func(v, x T, xs List[T]) (T, List[T]) {
		return x, AppendList(v, xs)
	}, xs)
}

// Ltos - list to slice
func Ltos[T any](xs List[T]) (res []T) {
	WalkList(func(t T) {
		res = append(res, t)
	}, xs)
	return
}

func Stol[T any](xs []T) List[T] {
	x, xs := Uncons(xs)
	if xs == nil {
		return nil
	}
	return func() (T, List[T]) {
		return x, Stol(xs)
	}
}

func reverseList[T any](xs, ys List[T]) List[T] {
	if xs == nil {
		return ys
	}
	x, xs := xs()
	return reverseList(xs, PrependList(x, ys))
}

func ReverseList[T any](xs List[T]) List[T] {
	return reverseList(xs, nil)
}
