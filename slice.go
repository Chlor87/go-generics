package main

// Cons_ is functional prepend
func Cons_[T any](x T, xs []T) []T {
	return append([]T{x}, xs...)
}

func Cons[T any](x T) func(xs []T) []T {
	return Curry2(Cons_[T])(x)
}

func Snoc_[T any](xs []T, x T) []T {
	return append(xs, x)
}

func Snoc[T any](xs []T) func(x T) []T {
	return Curry2(Snoc_[T])(xs)
}

func Uncons[T any](xs []T) (T, []T) {
	if len(xs) > 0 {
		return xs[0], xs[1:]
	}
	// zero value for type T
	var t T
	return t, nil
}

func Concat_[T any](xs, ys []T) []T {
	return append(xs, ys...)
}

func Concat[T any](xs []T) func(ys []T) []T {
	return Curry2(Concat_[T])(xs)
}

// transducers

func MapT_[I, O, A any](fn func(I) O, step func(A, O) A) func(A, I) A {
	return func(p A, c I) A {
		return step(p, fn(c))
	}
}

func MapT[I, O, A any](fn func(I) O) func(func(A, O) A) func(A, I) A {
	return Curry2(MapT_[I, O, A])(fn)
}

func FilterT_[I, A any](fn func(I) bool, step func(A, I) A) func(A, I) A {
	return func(p A, c I) A {
		if fn(c) {
			return step(p, c)
		}
		return p
	}
}

func FilterT[I, A any](fn func(I) bool) func(func(A, I) A) func(A, I) A {
	return Curry2(FilterT_[I, A])(fn)
}

// Reduce, Map and Filter with transducers

func Reduce_[I, A any](fn func(A, I) A, acc A, xs []I) A {
	for _, x := range xs {
		acc = fn(acc, x)
	}
	return acc
}

func Reduce[I, A any](fn func(A, I) A) func(acc A) func(xs []I) A {
	return Curry3(Reduce_[I, A])(fn)
}

func Map_[I, O any](fn func(I) O, xs []I) []O {
	return Reduce(MapT[I, O, []O](fn)(Snoc_[O]))([]O{})(xs)
}

func Map[I, O any](fn func(I) O) func(xs []I) []O {
	return Curry2(Map_[I, O])(fn)
}

func Filter_[T any](fn func(T) bool, xs []T) []T {
	return Reduce(FilterT[T, []T](fn)(Snoc_[T]))([]T{})(xs)
}

func Filter[T any](fn func(T) bool) func([]T) []T {
	return Curry2(Filter_[T])(fn)
}

// FindFirst_ doesn't need to be recursive, this implementation is suboptimal
func FindFirst_[T any](fn func(T) bool, xs []T) (T, bool) {
	switch x, xs := Uncons(xs); {
	case fn(x):
		return x, true
	case xs == nil:
		var t T
		return t, false
	default:
		return FindFirst_(fn, xs)
	}
}

func FindFirst[T any](fn func(T) bool) func([]T) (T, bool) {
	return Curry2_2(FindFirst_[T])(fn)
}

func FindLast_[T any](fn func(T) bool, xs []T) (T, bool) {
	for i := len(xs) - 1; i >= 0; i-- {
		if fn(xs[i]) {
			return xs[i], true
		}
	}
	var t T
	return t, false
}

func FindLast[T any](fn func(T) bool) func([]T) (T, bool) {
	return Curry2_2(FindLast_[T])(fn)
}

func Contains_[T comparable](v T, xs []T) bool {
	switch x, xs := Uncons(xs); {
	case x == v:
		return true
	case xs == nil:
		return false
	default:
		return Contains_(v, xs)
	}
}

func Contains[T comparable](a T) func([]T) bool {
	return Curry2(Contains_[T])(a)
}

func Any_[T any](fn func(T) bool, xs []T) bool {
	switch x, xs := Uncons(xs); {
	case fn(x):
		return true
	case xs == nil:
		return false
	default:
		return Any_(fn, xs)
	}
}

func Any[T any](fn func(T) bool) func([]T) bool {
	return Curry2(Any_[T])(fn)
}

func All_[T any](fn func(T) bool, xs []T) bool {
	switch x, xs := Uncons(xs); {
	case xs == nil:
		return true
	case !fn(x):
		return false
	default:
		return All_(fn, xs)
	}
}

func All[T any](fn func(T) bool) func([]T) bool {
	return Curry2(All_[T])(fn)
}

func GroupBy_[T any, K comparable](fn func(T) K, xs []T) map[K][]T {
	return Reduce(
		MapT[T, TupleFn[K, T], map[K][]T](
			func(i T) TupleFn[K, T] {
				return Tuple(fn(i), i)
			})(func(p map[K][]T, c TupleFn[K, T]) map[K][]T {
			k, v := c()
			p[k] = Snoc_(p[k], v)
			return p
		}))(map[K][]T{})(xs)
}

func GroupBy[T any, K comparable](fn func(T) K) func([]T) map[K][]T {
	return Curry2(GroupBy_[T, K])(fn)
}
