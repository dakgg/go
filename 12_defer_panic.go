package main

import "fmt"

func riskyOperation(n int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	if n == 0 {
		panic("n must not be zero")
	}
	return 100 / n, nil
}

func deferAndPanic() {
	fmt.Println("\n=== defer & panic/recover ===")

	// defer — LIFO 순서로 실행
	defer fmt.Println("  defer 3 (last registered, first out)")
	defer fmt.Println("  defer 2")
	defer fmt.Println("  defer 1 (first registered, last out)")
	fmt.Println("  (defers run after this function returns)")

	res, err := riskyOperation(5)
	fmt.Printf("  riskyOperation(5): %d %v\n", res, err)

	res, err = riskyOperation(0)
	fmt.Printf("  riskyOperation(0): %d %v\n", res, err)
}
