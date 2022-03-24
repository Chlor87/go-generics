package main

// currying requries a function per number of combinations of input and output
// argument counts

func Curry2[A, B, O any](
	fn func(A, B) O,
) func(A) func(B) O {
	return func(a A) func(B) O {
		return func(b B) O {
			return fn(a, b)
		}
	}
}

// Curry2_2 takes 2 inputs and returns 2 outputs
// ugly as hell
func Curry2_2[I1, I2, O1, O2 any](
	fn func(I1, I2) (O1, O2),
) func(I1) func(I2) (O1, O2) {
	return func(i1 I1) func(I2) (O1, O2) {
		return func(i2 I2) (O1, O2) {
			return fn(i1, i2)
		}
	}
}

func Curry3[A, B, C, O any](
	fn func(A, B, C) O,
) func(A) func(B) func(C) O {
	return func(a A) func(B) func(C) O {
		return Curry2(func(b B, c C) O {
			return fn(a, b, c)
		})
	}
}

func Curry4[A, B, C, D, O any](
	fn func(A, B, C, D) O,
) func(A) func(B) func(C) func(D) O {
	return func(a A) func(B) func(C) func(D) O {
		return Curry3(func(b B, c C, d D) O {
			return fn(a, b, c, d)
		})
	}
}
