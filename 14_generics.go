package main

import (
	"fmt"
	"strings"
)

// Number is a type constraint for integer and float types.
// ~T: T 자체와 T를 underlying type으로 갖는 파생 타입 모두 허용
type Number interface {
	~int | ~float64
}

// Map applies f to each element of s and returns a new slice.
func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Filter returns elements of s for which f returns true.
func Filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce reduces s to a single value by applying f cumulatively.
func Reduce[T, U any](s []T, init U, f func(U, T) U) U {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func generics() {
	fmt.Println("\n=== 제네릭 ===")

	nums := []int{1, 2, 3, 4, 5}

	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println("Map *2:", doubled)

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Filter even:", evens)

	total := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Reduce sum:", total)

	words := []string{"go", "is", "awesome"}
	upper := Map(words, strings.ToUpper)
	fmt.Println("Map ToUpper:", upper)
}
