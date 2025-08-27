package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}
	for e := range ch {
		fmt.Println(e)
	}

	ch, cancel := countTo(10)
	for e := range ch {
		fmt.Println(e)
	}
	cancel()

	fmt.Println("test_select")
	test_search_data("test", []func(string) []string{
		func(s string) []string {
			return []string{s}
		},
	})

	test_select()

}

func countTo(n int) (<-chan int, func()) {

	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
		close(ch)
	}()
	return ch, cancel
}

func test_search_data(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	results := make(chan []string)
	for _, searcher := range searchers {
		go func(s string) {
			select {
			case <-done:
			case results <- searcher(s):
			}
		}(s)
	}
	r := <-results
	close(done)
	return r
}

func test_select() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(v, v2)
}
