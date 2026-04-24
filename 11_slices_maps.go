package main

import (
	"fmt"
	"sort"
)

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
