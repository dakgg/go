package main

import "fmt"

func increment(n *int) {
	*n++
}

func pointers() {
	fmt.Println("\n=== 포인터 ===")
	x := 10
	increment(&x)
	fmt.Println("after increment:", x) // 11

	// new — 포인터 반환
	p := new(int)
	*p = 42
	fmt.Println("new int:", *p)
}
