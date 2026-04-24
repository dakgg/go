package main

import "fmt"

// Celsius / Fahrenheit: 기본 타입을 기반으로 한 정의 타입 (defined type)
type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func dataTypes() {
	fmt.Println("\n=== 자료형 ===")

	// -- 정수 계열 --
	fmt.Println("-- 정수 계열 --")
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	fmt.Printf("  int8=%d int16=%d int32=%d int64=%d\n", i8, i16, i32, i64)
	fmt.Printf("  uint8=%d uint16=%d uint32=%d uint64=%d\n", u8, u16, u32, u64)

	// int / uint: 플랫폼에 따라 32비트 또는 64비트
	var n int = 42
	fmt.Printf("  int=%d (size: %d bytes)\n", n, 8) // 64비트 OS 기준

	// -- 부동소수점 --
	fmt.Println("-- 부동소수점 --")
	var f32 float32 = 3.14159265358979 // 약 7자리 정밀도
	var f64 float64 = 3.14159265358979 // 약 15자리 정밀도
	fmt.Printf("  float32=%.7f\n", f32)
	fmt.Printf("  float64=%.15f\n", f64)

	// -- 복소수 --
	fmt.Println("-- 복소수 --")
	c64 := complex(float32(1), float32(2)) // complex64
	c128 := complex(1.5, -2.5)            // complex128
	fmt.Printf("  complex64=%v  real=%.1f imag=%.1f\n", c64, real(c64), imag(c64))
	fmt.Printf("  complex128=%v real=%.1f imag=%.1f\n", c128, real(c128), imag(c128))

	// -- bool --
	fmt.Println("-- bool --")
	var b bool = true
	fmt.Printf("  bool: %t, !bool: %t\n", b, !b)
	fmt.Printf("  10 > 3: %t, 10 == 3: %t\n", 10 > 3, 10 == 3)

	// -- string --
	fmt.Println("-- string --")
	s := "Hello, 세계"                       // UTF-8 인코딩
	fmt.Printf("  string=%q\n", s)
	fmt.Printf("  len(bytes)=%d\n", len(s))         // 바이트 수 (한글 3바이트)
	fmt.Printf("  len(runes)=%d\n", len([]rune(s))) // 문자(코드포인트) 수
	fmt.Printf("  s[0]=%d (byte)\n", s[0])          // 바이트 인덱싱

	// 문자열은 불변 — 수정 시 []byte 변환 필요
	bs := []byte(s)
	bs[0] = 'h'
	fmt.Printf("  modified: %s\n", string(bs))

	// -- byte & rune --
	fmt.Println("-- byte & rune --")
	var bt byte = 'A' // byte = uint8, ASCII 문자
	var r rune = '세'  // rune = int32, 유니코드 코드포인트
	fmt.Printf("  byte='%c' (%d)  rune='%c' (U+%04X)\n", bt, bt, r, r)

	// 문자열 → 룬 슬라이스: 유니코드 문자 단위 순회
	for i, ch := range "Go언어" {
		fmt.Printf("  [byte %d] '%c' U+%04X\n", i, ch, ch)
	}

	// -- 배열 (Array) --
	// 슬라이스와 달리 크기가 타입의 일부 (고정 크기)
	fmt.Println("-- 배열 --")
	var arr [5]int // 길이 5, 기본값 0
	arr[0] = 10
	arr2 := [3]string{"go", "is", "fun"}
	arr3 := [...]float64{1.1, 2.2, 3.3} // ... 로 길이 추론
	fmt.Printf("  arr=%v\n", arr)
	fmt.Printf("  arr2=%v len=%d\n", arr2, len(arr2))
	fmt.Printf("  arr3=%v len=%d\n", arr3, len(arr3))

	// 배열은 값 타입 — 대입 시 전체 복사
	copy1 := arr2
	copy1[0] = "rust"
	fmt.Printf("  arr2 (원본): %v  copy1: %v\n", arr2, copy1)

	// -- 정의 타입 (Defined Type) --
	fmt.Println("-- 정의 타입 --")
	boiling := Celsius(100)
	fmt.Printf("  %.1f°C = %.1f°F\n", boiling, boiling.ToFahrenheit())

	// type MyInt = int  → alias: MyInt과 int는 완전히 같은 타입
	// type MyInt   int  → defined: 별도 타입, 메서드 추가 가능, 명시적 변환 필요
	type MyInt = int // alias
	var mi MyInt = n + 1
	fmt.Printf("  alias MyInt=%d\n", mi)
}
