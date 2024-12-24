package main

import "fmt"

func main() {

	str1 := "Hello Primesoft!"

	fmt.Println("len of str1:", len(str1))
	for i := 0; i < len(str1); i++ {
		println(str1[i], "-->", string(str1[i]))
	}

	str2 := "Hello Primesoft!ㅊㅂ"
	fmt.Println("len of str2:", len(str2))

	for i := 0; i < len(str2); i++ {
		println("i:", i+1, ":->", str2[i], "-->", string(str2[i]))
	}

	// range loop

	for i, v := range str2 {
		fmt.Println("index:", i, "--> Value:", string(v))
	}
}

// range loop
// 1. used in array, string, slice, -> index and value
// 2. map --> key, value
// 3. channel -> value
