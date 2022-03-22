package main

import (
	"fmt"
	"strconv"
)

// call with Fix1
func _Fib(self FixFn1[int, int], n int) int {
	if n <= 2 {
		return 1
	}
	return self(self, n-2) + self(self, n-1)
}

// memoization, memo goes inside fix to capture memoized function
var Fib = Fix1(
	Memo2(
		_Fib,
		func(_ FixFn1[int, int], n int) int {
			return n
		},
	),
)

func main() {

	fmt.Println(Fix1(_Fib)(45))
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
