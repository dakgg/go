# Go 문법 학습 예제

`main.go` 하나에 Go의 핵심 문법을 주제별로 정리한 학습용 프로젝트입니다.

## 실행

```bash
go run main.go
```

---

## 목차

| # | 주제 | 함수 |
|---|------|------|
| 1 | [변수 & 타입](#1-변수--타입) | `variablesAndTypes()` |
| 2 | [제어 흐름](#2-제어-흐름) | `controlFlow()` |
| 3 | [함수](#3-함수) | `functions()` |
| 4 | [클로저 & 가변 인수](#4-클로저--가변-인수) | `closuresAndVariadic()` |
| 5 | [포인터](#5-포인터) | `pointers()` |
| 6 | [구조체 & 메서드](#6-구조체--메서드) | `structs()` |
| 7 | [인터페이스](#7-인터페이스) | `interfaces()` |
| 8 | [에러 처리](#8-에러-처리) | `errorHandling()` |
| 9 | [슬라이스 & 맵](#9-슬라이스--맵) | `slicesAndMaps()` |
| 10 | [defer & panic/recover](#10-defer--panicrecover) | `deferAndPanic()` |
| 11 | [고루틴 & 채널](#11-고루틴--채널) | `goroutinesAndChannels()` |
| 12 | [제네릭 (1.18+)](#12-제네릭-118) | `generics()` |

---

## 1. 변수 & 타입

```go
var name string = "Go"   // var 키워드
pi := 3.14               // 단축 선언 (:=)
const MaxSize = 100      // 상수

// 타입 변환은 반드시 명시적으로
var f float64 = float64(age)
```

- Go는 암묵적 타입 변환이 없습니다. 명시적 변환만 허용됩니다.
- `:=`는 함수 내부에서만 사용 가능합니다.

---

## 2. 제어 흐름

```go
// if — 초기화 구문 포함 가능
if n := 42; n%2 == 0 { ... }

// for — Go의 유일한 반복문 (while 없음)
for i := 0; i < 5; i++ { ... }
for i, v := range slice { ... }   // range

// switch — break 불필요, fallthrough 명시
switch day {
case "Sat", "Sun":
    ...
}
```

---

## 3. 함수

```go
// 다중 반환값
func divide(a, b float64) (float64, error) { ... }

// 네임드 반환값 (naked return)
func minMax(nums []int) (min, max int) {
    ...
    return  // min, max 암묵적 반환
}
```

Go 함수는 여러 값을 반환할 수 있으며, 에러는 마지막 반환값으로 전달하는 것이 관례입니다.

---

## 4. 클로저 & 가변 인수

```go
// 클로저 — 외부 변수 캡처
func makeCounter() func() int {
    count := 0
    return func() int { count++; return count }
}

// 가변 인수
func sum(nums ...int) int { ... }
sum(nums...)  // 슬라이스 전개
```

---

## 5. 포인터

```go
func increment(n *int) { *n++ }

x := 10
increment(&x)   // 주소 전달

p := new(int)   // 포인터 반환, 0으로 초기화
*p = 42
```

Go는 포인터 산술을 지원하지 않습니다. 안전한 수준의 포인터만 사용합니다.

---

## 6. 구조체 & 메서드

```go
type Rectangle struct {
    Width, Height float64
}

// 값 수신자 — 복사본
func (r Rectangle) Area() float64 { ... }

// 포인터 수신자 — 원본 수정
func (r *Rectangle) Scale(factor float64) { ... }
```

- 수정이 필요하거나 구조체가 크면 포인터 수신자를 사용합니다.
- `String()` 메서드를 구현하면 `fmt.Println`에서 자동으로 사용됩니다.

---

## 7. 인터페이스

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

// 타입 어설션
if c, ok := s.(Circle); ok { ... }

// 타입 스위치
switch v := i.(type) {
case int:    ...
case string: ...
}
```

Go 인터페이스는 **암묵적 구현**입니다. `implements` 키워드 없이 메서드만 맞으면 됩니다.

---

## 8. 에러 처리

```go
type ValidationError struct {
    Field, Message string
}
func (e *ValidationError) Error() string { ... }

// errors.As — 타입별 에러 처리
if errors.As(err, &ve) { ... }

// errors.Is — 특정 에러 값 비교
if errors.Is(err, io.EOF) { ... }
```

---

## 9. 슬라이스 & 맵

```go
// 슬라이스
s := make([]int, 0, 5)      // len=0, cap=5
s = append(s, 1, 2, 3)
s = append(s, other...)      // 슬라이스 전개

// 맵
m := map[string]int{"a": 1}
v, ok := m["key"]            // 존재 여부 확인
delete(m, "key")
```

슬라이스는 배열의 뷰(view)입니다. `append` 시 용량이 부족하면 새 배열로 복사됩니다.

---

## 10. defer & panic/recover

```go
// defer — 함수 종료 시 LIFO 순서로 실행
defer file.Close()
defer fmt.Println("cleanup")

// panic/recover — 예외 처리
defer func() {
    if r := recover(); r != nil {
        err = fmt.Errorf("recovered: %v", r)
    }
}()
panic("something went wrong")
```

`defer`는 리소스 해제 패턴에 자주 사용됩니다.

---

## 11. 고루틴 & 채널

```go
// 고루틴 — go 키워드로 실행
go func() { ch <- 42 }()

// 채널
ch := make(chan int)        // 버퍼 없음 (동기)
bch := make(chan int, 10)   // 버퍼 있음 (비동기)
close(ch)

// select — 여러 채널 동시 대기
select {
case v := <-ch1: ...
case v := <-ch2: ...
}

// sync.WaitGroup — 고루틴 완료 대기
var wg sync.WaitGroup
wg.Add(1)
go func() { defer wg.Done(); ... }()
wg.Wait()
```

---

## 12. 제네릭 (1.18+)

```go
// 타입 파라미터
func Map[T, U any](s []T, f func(T) U) []U { ... }
func Filter[T any](s []T, f func(T) bool) []T { ... }
func Reduce[T, U any](s []T, init U, f func(U, T) U) U { ... }

// 타입 제약
type Number interface {
    ~int | ~float64   // ~: 해당 타입의 파생 타입 포함
}
```

---

## 참고 자료

- [공식 Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com)
- [Go 언어 스펙](https://go.dev/ref/spec)
