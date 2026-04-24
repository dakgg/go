package main

import (
	"fmt"
	"math"
)

// Shape defines the interface for 2D geometric shapes.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle represents a circle with a given radius.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

// Perimeter returns the perimeter of r.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func printShape(s Shape) {
	fmt.Printf("  area=%.2f perimeter=%.2f\n", s.Area(), s.Perimeter())
}

func interfaces() {
	fmt.Println("\n=== 인터페이스 ===")

	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 3},
	}
	for _, s := range shapes {
		fmt.Printf("%T:", s)
		printShape(s)
	}

	// 타입 어설션
	var s Shape = Circle{Radius: 3}
	if c, ok := s.(Circle); ok {
		fmt.Printf("Circle radius: %.1f\n", c.Radius)
	}

	// 타입 스위치 — any는 interface{}의 alias (Go 1.18+)
	describe := func(i any) string {
		switch v := i.(type) {
		case int:
			return fmt.Sprintf("int(%d)", v)
		case string:
			return fmt.Sprintf("string(%q)", v)
		default:
			return fmt.Sprintf("unknown(%T)", v)
		}
	}
	fmt.Println(describe(42), describe("hello"), describe(true))
}
