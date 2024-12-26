package main

import (
	"fmt"
	"unsafe"
)

func main() {
	t1 := T1{b1: true, r1: 'A', l1: 12312312, b2: true}
	fmt.Println("Size of t1:", unsafe.Sizeof(t1))
	t2 := T2{b1: true, r1: 'A', l1: 12312312, b2: true}
	fmt.Println("Size of t2:", unsafe.Sizeof(t2))
}

type T1 struct {
	b1 bool  // 1
	r1 rune  // 4
	l1 int64 // 8
	b2 bool  // 1
}

type T2 struct {
	l1 int64 // 8
	r1 rune  // 4
	b1 bool  // 1
	b2 bool  // 1
}

type T3 struct {
	l1 int32 // 4
	r1 rune  // 4
	b1 bool  // 1
	b2 bool  // 1
	l2 int64 // 8
	b3 bool  // 1
}
