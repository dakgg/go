package main

import "fmt"

func controlFlow() {
	fmt.Println("\n=== 제어 흐름 ===")

	// if — 초기화 구문 포함
	if n := 42; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	}

	// for (Go에는 while이 없고 for만 있음)
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("sum 1..5 =", sum)

	// range over integer (Go 1.22+)
	for i := range 3 {
		fmt.Printf("  range %d\n", i)
	}

	// range — 슬라이스 순회
	fruits := []string{"apple", "banana", "cherry"}
	for i, v := range fruits {
		fmt.Printf("  [%d] %s\n", i, v)
	}

	// switch — break 불필요, fallthrough 명시
	day := "Mon"
	switch day {
	case "Sat", "Sun":
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday:", day)
	}
}
