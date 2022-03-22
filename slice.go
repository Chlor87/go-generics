package main

// Cons is functional append
func Cons[T any](x T, xs []T) []T {
	return append(xs, x)
}

func Snoc[T any](xs []T, x T) []T {
	return Cons(x, xs)
}

func Uncons[T any](xs []T) (T, []T) {
	if len(xs) > 0 {
		return xs[0], xs[1:]
	}
	var t T
	return t, xs
}

func _Concat[T any](xs, ys []T) []T {
	return append(xs, ys...)
}

func Concat[T any](xs []T) func(ys []T) []T {
	return Curry2(_Concat[T])(xs)
}

// transducers

func _MapT[I, O, A any](fn func(I) O, step func(A, O) A) func(A, I) A {
	return func(p A, c I) A {
		return step(p, fn(c))
	}
}

func MapT[I, O, A any](fn func(I) O) func(func(A, O) A) func(A, I) A {
	return Curry2(_MapT[I, O, A])(fn)
}

func _FilterT[I, A any](fn func(I) bool, step func(A, I) A) func(A, I) A {
	return func(p A, c I) A {
		if fn(c) {
			return step(p, c)
		}
		return p
	}
}

func FilterT[I, A any](fn func(I) bool) func(func(A, I) A) func(A, I) A {
	return Curry2(_FilterT[I, A])(fn)
}

// Reduce, Map and Filter with transducers

func _Reduce[I, A any](fn func(A, I) A, acc A, xs []I) A {
	for _, x := range xs {
		acc = fn(acc, x)
	}
	return acc
}

func Reduce[I, A any](fn func(A, I) A) func(acc A) func(xs []I) A {
	return Curry3(_Reduce[I, A])(fn)
}

func _Map[I, O any](fn func(I) O, xs []I) []O {
	return Reduce(MapT[I, O, []O](fn)(Snoc[O]))([]O{})(xs)
}

func Map[I, O any](fn func(I) O) func(xs []I) []O {
	return Curry2(_Map[I, O])(fn)
}

func _Filter[T any](fn func(T) bool, xs []T) []T {
	return Reduce(FilterT[T, []T](fn)(Snoc[T]))([]T{})(xs)
}

func Filter[T any](fn func(T) bool) func([]T) []T {
	return Curry2(_Filter[T])(fn)
}

// GroupBy with MapT transducer
func _GroupBy[T any, K comparable](fn func(T) K, xs []T) map[K][]T {
	return Reduce(
		MapT[T, T, map[K][]T](Id[T])(func(p map[K][]T, c T) map[K][]T {
			k := fn(c)
			if _, ok := p[k]; !ok {
				p[k] = []T{}
			}
			p[k] = append(p[k], c)
			return p
		}))(make(map[K][]T))(xs)
}

func GroupBy[T any, K comparable](fn func(T) K) func([]T) map[K][]T {
	return Curry2(_GroupBy[T, K])(fn)
}
