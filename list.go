package main

import (
	"fmt"
	"strings"
)

type List[T any] Tuple[T, List[T]]

func (l List[T]) String() string {
	var res []string
	WalkL(func(t T) { res = append(res, fmt.Sprintf("%v", t)) }, l)
	res = append(res, "nil")
	return strings.Join(res, " => ")
}

func MkL[T any](x T) List[T] {
	return ConsL(x, nil)
}

func WalkL[T any](fn func(T), xs List[T]) {
	if xs == nil {
		return
	}
	x, xs := xs()
	fn(x)
	WalkL(fn, xs)
}

func ConsL[T any](v T, xs List[T]) List[T] {
	return func() (T, List[T]) {
		return v, xs
	}
}

func SnocL[T any](xs List[T], v T) List[T] {
	if xs == nil {
		return MkL(v)
	}
	x, xs := xs()
	return ConsL(x, SnocL(xs, v))
}

func UnconsL[T any](xs List[T]) (T, List[T]) {
	return xs()
}

func MapL[I, O any](fn func(I) O, xs List[I]) List[O] {
	if xs == nil {
		return nil
	}
	x, xs := xs()
	return ConsL(fn(x), MapL(fn, xs))
}

func insertL[T any](
	fn func(T) bool,
	v T,
	inserter func(T, T, List[T]) List[T],
	xs List[T],
) List[T] {
	if xs == nil {
		return xs
	}
	x, xs := xs()
	if fn(x) {
		return inserter(v, x, xs)
	}
	return ConsL(x, insertL(fn, v, inserter, xs))
}

func InsertBeforeL[T any](fn func(T) bool, v T, xs List[T]) List[T] {
	return insertL(fn, v, func(v, x T, xs List[T]) List[T] {
		return ConsL(v, ConsL(x, xs))
	}, xs)
}

func InsertAfterL[T any](fn func(T) bool, v T, xs List[T]) List[T] {
	return insertL(fn, v, func(v, x T, xs List[T]) List[T] {
		return ConsL(x, ConsL(v, xs))
	}, xs)
}

func RemoveL[T any](fn func(T) bool, xs List[T]) List[T] {
	if xs == nil {
		return xs
	}
	x, xs := xs()
	if fn(x) {
		return xs
	}
	return ConsL(x, RemoveL(fn, xs))
}

func ReverseL[T any](xs List[T]) List[T] {
	if xs == nil {
		return xs
	}
	x, xs := xs()
	return SnocL(ReverseL(xs), x)
}

func FindL[T any](fn func(T) bool, xs List[T]) (T, bool) {
	if xs == nil {
		var t T
		return t, false
	}
	x, xs := xs()
	if fn(x) {
		return x, true
	}
	return FindL(fn, xs)
}

func Ltos[T any](xs List[T]) []T {
	if xs == nil {
		return nil
	}
	x, xs := xs()
	return append([]T{x}, Ltos(xs)...)
}

func Stol[T any](xs []T) List[T] {
	x, xs := Uncons(xs)
	if xs == nil {
		return nil
	}
	return ConsL(x, Stol(xs))
}

func HeadL[T any](xs List[T]) T {
	x, _ := xs()
	return x
}

func TailL[T any](xs List[T]) List[T] {
	if xs == nil {
		return nil
	}
	_, xs = xs()
	return xs
}

func NthL[T any](idx int, xs List[T]) (T, bool) {
	if xs == nil {
		var t T
		return t, false
	}
	x, xs := xs()
	if idx == 0 {
		return x, true
	}
	return NthL(idx-1, xs)
}

func ConcatL[T any](xss ...List[T]) List[T] {
	if len(xss) == 0 || xss[0] == nil {
		return nil
	}
	x, xs := xss[0]()
	if xs == nil {
		return ConsL(x, ConcatL(xss[1:]...))
	}
	return ConsL(x, ConcatL(append([]List[T]{xs}, xss[1:]...)...))
}
