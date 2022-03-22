package main

import "sync"

func FanIn[T any](chs ...chan T) chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(chs))

	for _, c := range chs {
		go func(c chan T) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func FanOut[T any](in chan T, n int) []chan T {
	chs := make([]chan T, n)
	for i := 0; i < n; i++ {
		chs[i] = make(chan T)
	}
	var wg sync.WaitGroup
	// this one waits for the whole read loop to exit. Without it
	// the second goroutine closes the out channels immediately
	wg.Add(1)

	go func() {
		defer wg.Done()
		for v := range in {
			for _, c := range chs {
				wg.Add(1)
				go func(c chan T, v T) {
					c <- v
					wg.Done()
				}(c, v)
			}
		}
	}()

	go func() {
		wg.Wait()
		for _, c := range chs {
			close(c)
		}
	}()

	return chs
}
