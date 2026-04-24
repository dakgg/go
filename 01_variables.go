package main

import "fmt"

func variablesAndTypes() {
	// var 키워드
	var name string = "Go"
	var age int = 10

	// 단축 선언 (:=)
	pi := 3.14

	// 복수 선언
	var x, y int = 1, 2

	// 상수
	const MaxSize = 100

	fmt.Println("=== 변수와 타입 ===")
	fmt.Printf("name=%s age=%d pi=%.2f x=%d y=%d MaxSize=%d\n", name, age, pi, x, y, MaxSize)

	// 타입 변환 (암묵적 변환 없음, 명시적으로만)
	var f float64 = float64(age)
	fmt.Printf("int→float64: %v\n", f)
}
