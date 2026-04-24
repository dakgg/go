package main

import "fmt"

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func closuresAndVariadic() {
	fmt.Println("\n=== 클로저 & 가변 인수 ===")

	counter := makeCounter()
	fmt.Println(counter(), counter(), counter()) // 1 2 3

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("sum:", sumAll(nums...)) // 슬라이스 전개
}
