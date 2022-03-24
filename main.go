package main

import (
	"fmt"
	"strconv"
)

// call with Fix1, this one is slow
func _Fib(self FixFn1[int, int], n int) int {
	if n <= 2 {
		return 1
	}
	return self(self, n-2) + self(self, n-1)
}

// memoization, memo goes inside fix to capture memoized function
// this one's fast
var Fib = Fix1(
	Memo2(
		_Fib,
		func(_ FixFn1[int, int], n int) int {
			return n
		},
	),
)

func main() {

	// poor man's match expr
	switch x, xs := Uncons([]string(nil)); {
	case xs == nil:
		fmt.Println("Nothing", x, xs)
	default:
		fmt.Println("Just", x, xs)
	}

	// lazy linked list
	xs := AppendList(3, AppendList(2, MkList(1)))
	fmt.Println(xs, reverseList(xs, nil), Stol([]int{1, 2, 3}))

	// memo, the second one is immediate
	fmt.Println(_Fib(_Fib, 45))
	fmt.Println(Fib(90))

	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- strconv.Itoa(i)
		}
		close(c)
	}()

	for v := range FanIn(FanOut(c, 10)...) {
		fmt.Println(v)
	}

}
