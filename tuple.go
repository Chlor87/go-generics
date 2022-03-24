package main

type TupleFn[L, R any] func() (L, R)

func Tuple[L, R any](l L, r R) TupleFn[L, R] {
	return func() (L, R) {
		return l, r
	}
}
