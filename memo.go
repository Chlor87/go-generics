package main

func Memo1[I, O any, K comparable](fn func(I) O, keyFn func(I) K) func(I) O {
	memo := make(map[K]O)
	return func(i I) O {
		k := keyFn(i)
		if res, ok := memo[k]; ok {
			return res
		}
		memo[k] = fn(i)
		return memo[k]
	}
}

func Memo2[A, B, O any, K comparable](
	fn func(A, B) O, keyFn func(A, B) K,
) func(A, B) O {
	memo := make(map[K]O)
	return func(a A, b B) O {
		k := keyFn(a, b)
		if res, ok := memo[k]; ok {
			return res
		}
		memo[k] = fn(a, b)
		return memo[k]
	}
}
