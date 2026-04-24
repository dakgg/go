package main

import "fmt"

// Rectangle represents a 2D rectangle.
type Rectangle struct {
	Width, Height float64
}

// Area returns the area of r. (값 수신자 — 복사본에서 실행)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Scale multiplies both dimensions by factor. (포인터 수신자 — 원본 수정)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// String implements fmt.Stringer.
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.1f×%.1f)", r.Width, r.Height)
}

func structs() {
	fmt.Println("\n=== 구조체 & 메서드 ===")
	rect := Rectangle{Width: 4, Height: 3}
	fmt.Println(rect, "area:", rect.Area())

	rect.Scale(2)
	fmt.Println("after Scale(2):", rect, "area:", rect.Area())
}
