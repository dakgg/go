package main

import (
	"errors"
	"fmt"
)

// ErrDivisionByZero is returned when dividing by zero.
// sentinel error: 패키지 레벨에서 변수로 선언하여 errors.Is 비교 가능
var ErrDivisionByZero = errors.New("division by zero")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// minMax returns the minimum and maximum values in nums.
// 네임드 반환값 (naked return)
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}

func functions() {
	fmt.Println("\n=== 함수 ===")

	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10/3 = %.4f\n", result)
	}

	_, err = divide(5, 0)
	// errors.Is — sentinel error 값 비교
	fmt.Println("is ErrDivisionByZero:", errors.Is(err, ErrDivisionByZero))

	min, max := minMax([]int{3, 1, 4, 1, 5, 9, 2, 6})
	fmt.Printf("min=%d max=%d\n", min, max)
}
