package main

// FixFn1 is a type for fixed point combinator function, not truly generic
type FixFn1[I, O any] func(FixFn1[I, O], I) O

// Fix1 creates a self referencing function (fixed point combinator)
func Fix1[I, O any](fn FixFn1[I, O]) func(I) O {
	return func(in I) O {
		return fn(fn, in)
	}
}

// ... that's not great, there must exist a variant per combination of input and
// output parameter counts we want to handle

type FixFn2[I1, I2, O any] func(FixFn2[I1, I2, O], I1, I2) O

func Fix2[I1, I2, O any](fn FixFn2[I1, I2, O]) func(I1, I2) O {
	return func(a I1, b I2) O {
		return fn(fn, a, b)
	}
}
