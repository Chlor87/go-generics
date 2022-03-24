package main

// Id - identity
func Id[T any](i T) T {
	return i
}

func Flip[A, B, O any](fn func(A, B) O) func(B, A) O {
	return func(b B, a A) O {
		return fn(a, b)
	}
}

func Const[T any](v T) func() T {
	return func() T {
		return v
	}
}

func Cmp_[T comparable](a, b T) bool {
	return a == b
}

func Cmp[T comparable](a T) func(T) bool {
	return Curry2(Cmp_[T])(a)
}

func Apply1_[I, O any](fn func(I) O, i I) O {
	return fn(i)
}

func Apply1[I, O any](fn func(I) O) func(I) O {
	return Curry2(Apply1_[I, O])(fn)
}
