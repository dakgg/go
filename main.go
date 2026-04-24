package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

// ──────────────────────────────────────────────
// 1. 변수 선언 & 타입
// ──────────────────────────────────────────────

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

// ──────────────────────────────────────────────
// 2. 제어 흐름: if / for / switch
// ──────────────────────────────────────────────

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

// ──────────────────────────────────────────────
// 3. 함수: 다중 반환값 & 네임드 반환
// ──────────────────────────────────────────────

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

// ──────────────────────────────────────────────
// 4. 클로저 & 가변 인수
// ──────────────────────────────────────────────

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

// ──────────────────────────────────────────────
// 5. 포인터
// ──────────────────────────────────────────────

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

// ──────────────────────────────────────────────
// 6. 구조체 & 메서드
// ──────────────────────────────────────────────

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

// ──────────────────────────────────────────────
// 7. 인터페이스
// ──────────────────────────────────────────────

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

// ──────────────────────────────────────────────
// 8. 에러 처리 & 커스텀 에러
// ──────────────────────────────────────────────

// ValidationError is returned when input validation fails.
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s — %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically large"}
	}
	return nil
}

func errorHandling() {
	fmt.Println("\n=== 에러 처리 ===")

	for _, age := range []int{25, -1, 200} {
		if err := validateAge(age); err != nil {
			var ve *ValidationError
			if errors.As(err, &ve) {
				fmt.Printf("  field=%s msg=%s\n", ve.Field, ve.Message)
			}
		} else {
			fmt.Printf("  age %d is valid\n", age)
		}
	}
}

// ──────────────────────────────────────────────
// 9. 슬라이스 & 맵
// ──────────────────────────────────────────────

func slicesAndMaps() {
	fmt.Println("\n=== 슬라이스 & 맵 ===")

	// 슬라이스
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3)
	s = append(s, []int{4, 5}...)
	fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))

	// 슬라이싱
	fmt.Println("s[1:3]:", s[1:3])

	// 정렬
	data := []int{5, 3, 1, 4, 2}
	sort.Ints(data)
	fmt.Println("sorted:", data)

	// 맵
	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	scores["Charlie"] = 95

	// 존재 여부 확인
	if v, ok := scores["Bob"]; ok {
		fmt.Println("Bob's score:", v)
	}

	delete(scores, "Bob")
	fmt.Println("after delete:", scores)
}

// ──────────────────────────────────────────────
// 10. defer & panic/recover
// ──────────────────────────────────────────────

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

// ──────────────────────────────────────────────
// 11. 고루틴 & 채널
// ──────────────────────────────────────────────

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

// ──────────────────────────────────────────────
// 12. 제네릭 (Go 1.18+)
// ──────────────────────────────────────────────

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

// ──────────────────────────────────────────────
// main
// ──────────────────────────────────────────────

func main() {
	variablesAndTypes()
	controlFlow()
	functions()
	closuresAndVariadic()
	pointers()
	structs()
	interfaces()
	errorHandling()
	slicesAndMaps()
	deferAndPanic()
	goroutinesAndChannels()
	generics()
}
