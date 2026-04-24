# Go 문법 학습 예제

`main.go` 하나에 Go의 핵심 문법을 주제별로 정리한 학습용 프로젝트입니다.

## 실행

```bash
go run .
```

---

## 파일 구조

| 파일 | 섹션 |
|------|------|
| `main.go` | 진입점 (`main` 함수만 포함) |
| `01_variables.go` | 1. 변수 & 타입 |
| `02_types.go` | 2. 자료형 |
| `03_data_structures.go` | 3. 자료구조 |
| `04_control_flow.go` | 4. 제어 흐름 |
| `05_functions.go` | 5. 함수 |
| `06_closures.go` | 6. 클로저 & 가변 인수 |
| `07_pointers.go` | 7. 포인터 |
| `08_structs.go` | 8. 구조체 & 메서드 |
| `09_interfaces.go` | 9. 인터페이스 |
| `10_errors.go` | 10. 에러 처리 |
| `11_slices_maps.go` | 11. 슬라이스 & 맵 |
| `12_defer_panic.go` | 12. defer & panic/recover |
| `13_goroutines.go` | 13. 고루틴 & 채널 |
| `14_generics.go` | 14. 제네릭 |

---

## 목차

| # | 주제 | 함수 |
|---|------|------|
| 1 | [변수 & 타입](#1-변수--타입) | `variablesAndTypes()` |
| 2 | [자료형](#2-자료형) | `dataTypes()` |
| 3 | [자료구조](#3-자료구조) | `dataStructures()` |
| 4 | [제어 흐름](#4-제어-흐름) | `controlFlow()` |
| 5 | [함수](#5-함수) | `functions()` |
| 6 | [클로저 & 가변 인수](#6-클로저--가변-인수) | `closuresAndVariadic()` |
| 7 | [포인터](#7-포인터) | `pointers()` |
| 8 | [구조체 & 메서드](#8-구조체--메서드) | `structs()` |
| 9 | [인터페이스](#9-인터페이스) | `interfaces()` |
| 10 | [에러 처리](#10-에러-처리) | `errorHandling()` |
| 11 | [슬라이스 & 맵](#11-슬라이스--맵) | `slicesAndMaps()` |
| 12 | [defer & panic/recover](#12-defer--panicrecover) | `deferAndPanic()` |
| 13 | [고루틴 & 채널](#13-고루틴--채널) | `goroutinesAndChannels()` |
| 14 | [제네릭 (1.18+)](#14-제네릭-118) | `generics()` |

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

## 2. 자료형

Go의 모든 내장 자료형을 범주별로 정리합니다.

### 정수 계열

| 타입 | 크기 | 범위 |
|------|------|------|
| `int8` | 1바이트 | -128 ~ 127 |
| `int16` | 2바이트 | -32768 ~ 32767 |
| `int32` | 4바이트 | ±2.1억 |
| `int64` | 8바이트 | ±9.2×10¹⁸ |
| `uint8` | 1바이트 | 0 ~ 255 |
| `uint16` | 2바이트 | 0 ~ 65535 |
| `uint32` | 4바이트 | 0 ~ 4.2억 |
| `uint64` | 8바이트 | 0 ~ 1.8×10¹⁹ |
| `int` / `uint` | 플랫폼 의존 | 64비트 OS에서 8바이트 |

### 부동소수점 & 복소수

```go
var f32 float32 = 3.14   // 약 7자리 정밀도
var f64 float64 = 3.14   // 약 15자리 정밀도 (기본)

c := complex(1.5, -2.5)  // complex128
real(c)   // 1.5
imag(c)   // -2.5
```

### bool

```go
var b bool = true
// 비교 연산자(==, !=, <, >) 결과가 bool
```

### string, byte, rune

```go
s := "Hello, 세계"          // UTF-8 인코딩
len(s)                      // 바이트 수 (한글 1자 = 3바이트)
len([]rune(s))              // 유니코드 문자 수

var bt byte = 'A'           // byte = uint8
var r  rune = '세'           // rune = int32, 유니코드 코드포인트

for i, ch := range s { ... } // rune 단위 순회
```

- `string`은 불변(immutable)입니다. 수정하려면 `[]byte`로 변환합니다.
- `byte`는 `uint8`의, `rune`은 `int32`의 별칭입니다.

### 배열 (Array)

```go
var arr [5]int               // 크기가 타입의 일부 — [5]int ≠ [6]int
arr2 := [...]float64{1, 2}   // ... 로 길이 추론
```

- 배열은 **값 타입**: 대입하면 전체 복사됩니다.
- 크기가 고정되어 있으며, 슬라이스와 달리 `append`가 없습니다.

### 정의 타입 vs 타입 별칭

```go
type Celsius float64    // 정의 타입: 별도 타입, 메서드 추가 가능
type MyInt  = int       // 타입 별칭: int와 완전히 동일, 혼용 가능
```

- 정의 타입끼리는 명시적 변환이 필요합니다: `Fahrenheit(celsius * 9/5 + 32)`.
- 타입 별칭은 기존 타입에 다른 이름을 붙인 것으로 호환됩니다.

---

## 3. 자료구조

제네릭을 활용해 타입 안전한 자료구조를 직접 구현합니다.

### Stack (LIFO)

```go
type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(v T)          { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() (T, bool)    { /* 마지막 요소 제거 & 반환 */ }
func (s *Stack[T]) Peek() (T, bool)   { /* 제거 없이 조회 */ }
```

### Queue (FIFO)

```go
type Queue[T any] struct{ items []T }

func (q *Queue[T]) Enqueue(v T)       { q.items = append(q.items, v) }
func (q *Queue[T]) Dequeue() (T, bool){ /* 앞에서 제거 & 반환 */ }
```

### Linked List (단방향)

```go
type LinkedList[T any] struct{ head *node[T]; size int }

ll.Prepend(0)   // O(1) — 앞 삽입
ll.Append(1)    // O(n) — 뒤 삽입
ll.ToSlice()    // 전체 조회
```

### Binary Search Tree

```go
bst.Insert(5)
bst.Contains(4)  // true
bst.InOrder()    // 오름차순 정렬된 슬라이스
```

- 중위 순회(In-Order)로 BST의 모든 값을 오름차순으로 가져올 수 있습니다.
- 재귀적 삽입으로 BST 불변식을 유지합니다.

### Set (집합)

```go
s := NewSet[string]()
s.Add("go")
s.Contains("go")  // true
s.Remove("go")
s.Len()           // 중복 없는 원소 수
```

- `map[T]struct{}` 기반: 빈 struct는 메모리를 차지하지 않습니다.
- `comparable` 제약: 맵 키로 사용 가능한 타입만 허용합니다.

---

## 4. 제어 흐름

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

## 5. 함수

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

## 6. 클로저 & 가변 인수

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

## 7. 포인터

```go
func increment(n *int) { *n++ }

x := 10
increment(&x)   // 주소 전달

p := new(int)   // 포인터 반환, 0으로 초기화
*p = 42
```

Go는 포인터 산술을 지원하지 않습니다. 안전한 수준의 포인터만 사용합니다.

---

## 8. 구조체 & 메서드

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

## 9. 인터페이스

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

## 10. 에러 처리

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

## 11. 슬라이스 & 맵

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

## 12. defer & panic/recover

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

## 13. 고루틴 & 채널

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

## 14. 제네릭 (1.18+)

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
