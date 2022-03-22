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

func _Apply1[I, O any](fn func(I) O, i I) O {
	return fn(i)
}

func Apply1[I, O any](fn func(I) O) func(I) O {
	return Curry2(_Apply1[I, O])(fn)
}
