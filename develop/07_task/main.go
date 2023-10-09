package main

import (
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("Done after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	var out = make(chan interface{})
	for i := 0; i < len(channels); i++ {
		ch := channels[i]
		go func(ch <-chan interface{}) {
			select {
			case <-ch:
				close(out)
			}
		}(ch)
	}
	return out
}
