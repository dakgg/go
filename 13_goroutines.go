package main

import (
	"fmt"
	"sync"
)

func goroutinesAndChannels() {
	fmt.Println("\n=== 고루틴 & 채널 ===")

	// 버퍼 없는 채널
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println("received:", <-ch)

	// 버퍼 채널
	bch := make(chan string, 3)
	bch <- "a"
	bch <- "b"
	bch <- "c"
	close(bch)
	for v := range bch {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// select
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "one"
	ch2 <- "two"
	select {
	case v := <-ch1:
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2:", v)
	}

	// sync.WaitGroup
	// Go 1.22+: 루프 변수가 반복마다 새로 생성되므로 클로저에 직접 캡처 가능
	var wg sync.WaitGroup
	results := make([]int, 5)
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = i * i
		}()
	}
	wg.Wait()
	fmt.Println("squares:", results)
}
