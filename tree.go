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
	WalkTree(func(x, p Tree[T], i, d int) {
		res = append(res, fmt.Sprintf(
			"%s%s( %v )",
			strings.Repeat(".", d),
			strings.Repeat(" ", int(math.Min(float64(d), 1))),
			Value(x),
		))
	}, t)
	return strings.Join(res, "\n")
}

func DFS[T any](fn func(T), t Tree[T]) {
	if t == nil {
		return
	}
	x, xs := t()
	WalkL(func(x Tree[T]) { DFS(fn, x) }, xs)
	fn(x)
}

func BFS[T any](fn func(T), t Tree[T]) {
	if t == nil {
		return
	}
	x, xs := t()
	fn(x)
	WalkL(func(x Tree[T]) { BFS(fn, x) }, xs)
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
	if t == nil {
		var t T
		return t
	}
	x, _ := t()
	return x
}

func Children[T any](t Tree[T]) List[Tree[T]] {
	if t == nil {
		return nil
	}
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

// walkCB documents walk callback's arguments
type walkCB[T any] func(node, parent Tree[T], index, depth int)

func walkTree[T any](fn walkCB[T], i, d int, p, t Tree[T]) {
	fn(t, p, i, d)
	WalkL(func(c Tree[T]) { walkTree(fn, i, d+1, t, c); i++ }, Children(t))
}

func WalkTree[T any](fn func(Tree[T], Tree[T], int, int), t Tree[T]) {
	walkTree(fn, 0, 0, nil, t)
}

func pathTree[T any](fn func(Tree[T]) bool, path []int, t Tree[T]) []int {
	if t == nil {
		return nil
	}

	if fn(t) {
		return path
	}

	xs := Children(t)
	if xs == nil {
		return nil
	}

	var i int
	for {
		var x Tree[T]
		x, xs = xs()
		if p := pathTree(fn, append(path, i), x); p != nil {
			return p
		}
		if xs == nil {
			break
		}
		i++
	}
	return nil
}

func PathTree[T any](fn func(Tree[T]) bool, t Tree[T]) []int {
	return pathTree(fn, nil, t)
}

func ByPathTree[T any](path []int, t Tree[T]) (Tree[T], bool) {
	if t == nil {
		return nil, false
	}

	curr, ok := NthL(path[0], Children(t))
	if !ok {
		return nil, false
	}

	if len(path) == 1 {
		return curr, true
	}

	return ByPathTree(path[1:], curr)
}

func TreeTest() {

	t := AddChild(Cmp(2), 3, AddChild(Cmp(1), 2, MkTree(1)))
	t = AddChild(Cmp(3), 5, AddChild(Cmp(1), 4, t))

	WalkTree(func(t, p Tree[int], i, d int) {
		fmt.Println(Value(t), Value(p), i, d)
	}, t)

	fmt.Println(ByPathTree(PathTree(func(t Tree[int]) bool { return Value(t) == 4 }, t), t))
	fmt.Println(t)
	return

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
