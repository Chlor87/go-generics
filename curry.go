package main

// currying requries a type and a function per number of combinations of counts
// of input and output arguments
func Curry2[A, B, O any](fn func(A, B) O) func(A) func(B) O {
	return func(a A) func(B) O {
		return func(b B) O {
			return fn(a, b)
		}
	}
}

func Curry3[A, B, C, O any](fn func(A, B, C) O) func(A) func(B) func(C) O {
	return func(a A) func(B) func(C) O {
		return Curry2(func(b B, c C) O {
			return fn(a, b, c)
		})
	}
}

func Curry4[A, B, C, D, O any](fn func(A, B, C, D) O) func(A) func(B) func(C) func(D) O {
	return func(a A) func(B) func(C) func(D) O {
		return Curry3(func(b B, c C, d D) O {
			return fn(a, b, c, d)
		})
	}
}
