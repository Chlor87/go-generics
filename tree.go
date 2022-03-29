package main

import (
	"fmt"
	"math"
	"strings"
)

type Tree[T any] Tuple[T, List[Tree[T]]]

func MkTree[T any](v T) Tree[T] {
	return Tree[T](MkTuple[T, List[Tree[T]]](v, nil))
}

func (t Tree[T]) String() string {
	var res []string
	bfsDepth(func(x T, d int) {
		res = append(res, fmt.Sprintf(
			"%s%s( %v )",
			strings.Repeat(".", d),
			strings.Repeat(" ", int(math.Min(float64(d), 1))),
			x,
		),
		)
	}, t, 0)
	return strings.Join(res, "\n")
}

func DFS[T any](fn func(T), t Tree[T]) {
	if t == nil {
		return
	}
	x, xs := t()
	WalkL(func(x Tree[T]) {
		DFS(fn, x)
	}, xs)
	fn(x)
}

func BFS[T any](fn func(T), t Tree[T]) {
	if t == nil {
		return
	}
	x, xs := t()
	fn(x)
	WalkL(func(x Tree[T]) {
		BFS(fn, x)
	}, xs)
}

func bfsDepth[T any](fn func(T, int), t Tree[T], depth int) {
	if t == nil {
		return
	}
	x, xs := t()
	fn(x, depth)
	WalkL(func(x Tree[T]) {
		bfsDepth(fn, x, depth+1)
	}, xs)
}

func AddChild[T any](fn func(T) bool, v T, t Tree[T]) Tree[T] {
	if t == nil {
		return t
	}
	x, xs := t()
	if fn(x) {
		return Tree[T](MkTuple(x, SnocL(xs, MkTree(v))))
	}
	return Tree[T](MkTuple(x, MapL(Curry3(AddChild[T])(fn)(v), xs)))
}

func RemoveChild[T any](fn func(T) bool, t Tree[T]) Tree[T] {
	if t == nil {
		return nil
	}
	x, xs := t()
	if fn(x) {
		return nil
	}
	return Tree[T](MkTuple(x, MapL(Curry2(RemoveChild[T])(fn), xs)))
}

func Value[T any](t Tree[T]) T {
	x, _ := t()
	return x
}

func Children[T any](t Tree[T]) List[Tree[T]] {
	_, ts := t()
	return ts
}

func Parent[T any](fn func(T) bool, t, p Tree[T]) Tree[T] {
	if t == nil {
		return nil
	}
	x, xs := t()
	if fn(x) {
		return p
	}
	for {
		var x Tree[T]
		x, xs = xs()
		if p := Parent(fn, x, t); p != nil {
			return p
		}
		if xs == nil {
			break
		}
	}
	return nil
}

func WalkParent[T any](fn func(Tree[T], Tree[T]), t Tree[T]) {
	if t == nil {
		return
	}
	xs := Children(t)
	for {
		var x Tree[T]
		x, xs = xs()
		fn(x, t)
		if xs == nil {
			break
		}
		WalkParent(fn, x)
	}
}

// IterTree returns a flat list (BFS)
// TODO: figure out how to emit parent and depth
func IterTree[T any](t Tree[T]) List[Tree[T]] {
	if t == nil {
		return nil
	}
	_, xs := t()
	if xs == nil {
		return MkL(t)
	}
	res := MkL(t)
	for {
		var x Tree[T]
		x, xs = xs()
		res = ConcatL(res, IterTree(x))
		if xs == nil {
			break
		}
	}
	return res
}

func TreeTest() {

	t := AddChild(Cmp(2), 3, AddChild(Cmp(1), 2, MkTree(1)))
	t = AddChild(Cmp(3), 5, AddChild(Cmp(1), 4, t))

	next := IterTree(t)

	for {
		var c Tree[int]
		c, next = next()
		fmt.Println(Value(c))
		if next == nil {
			break
		}
	}

	DFS(func(i int) { fmt.Println(i) }, t)
	BFS(func(i int) { fmt.Println(i) }, t)

	fmt.Println(RemoveChild(Cmp(2), t))
	fmt.Println(Parent(Cmp(3), t, nil))

}
