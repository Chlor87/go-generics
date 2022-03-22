package main

// RTL function composition
// can't do variadic type parameters, can't do generic compose/pipe nor curry

func Compose2[I, T, O any](
	fn1 func(T) O, fn2 func(I) T,
) func(I) O {
	return func(i I) O {
		return fn1(fn2(i))
	}
}

func Compose3[I, T1, T2, O any](
	f1 func(T2) O, f2 func(T1) T2, f3 func(I) T1,
) func(I) O {
	return func(i I) O {
		return f1(Compose2(f2, f3)(i))
	}
}

func Compose4[I, T1, T2, T3, O any](
	f1 func(T3) O, f2 func(T2) T3, f3 func(T1) T2, f4 func(I) T1,
) func(I) O {
	return func(i I) O {
		return f1(Compose3(f2, f3, f4)(i))
	}
}
