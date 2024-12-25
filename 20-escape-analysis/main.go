package main

import (
	"strconv"
)

func main() {
	slice1 := make([]int, 1000)
	slice2 := []int{1, 2, 3, 3, 4, 5, 6, 6, 7, 8, 88, 8, 88}
	str1 := "Hello World"
	any1 := any(str1)
	println(slice1)
	println(slice2)
	println(str1)
	println(any1)
	any1 = slice1
	for i := 1; i <= 10000000; i++ {
		slice2 = append(slice2, i)
	}

	map1 := make(map[string]any, 100000000)
	println(map1)
	n := 10
	n1 := sq(n)
	println(n1)

	for i := 1; i <= 1000000; i++ {
		str1 = str1 + "How are you doing->" + strconv.Itoa(i)
		//	str1 = str1 + "How are you doing->" + fmt.Sprint(i)
	}
	p1 := sqp1(n) // p is a pointer
	// a?
	println(*p1)
	o := new(int)
	sqp2(100, o)
	println(*o)

}

func sq(n int) int {
	n2 := n * n // n1 is allocated here
	return n2
	// n1 is deallocated here
}

func sqp1(n int) *int {
	//a := 100 // allocated
	p := new(int)
	*p = n * n
	//println(a)
	return p // a is deallocated here
}

func sqp2(n int, out *int) {
	//a := 100 // allocated
	if out != nil {
		*out = n * n
	}
	//println(a)
	//return p // a is deallocated here
}
