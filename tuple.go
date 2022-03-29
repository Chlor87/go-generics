package main

type Tuple[L, R any] func() (L, R)

func MkTuple[L, R any](l L, r R) Tuple[L, R] {
	return func() (L, R) {
		return l, r
	}
}
