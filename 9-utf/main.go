package main

import "unicode/utf8"

func main() {

	char := 'A' // 4 bytes

	len := utf8.RuneLen(char)
	println("Length of char is", len)

	char = '是'
	println(char)

	len = utf8.RuneLen(char)
	println("Length of char is", len)
}
