package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	// Consumer exits early after first value
	v := <-ch
	fmt.Println("got:", v)
	close(ch)
	wg.Wait()
}
