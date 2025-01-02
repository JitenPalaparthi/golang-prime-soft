package main

import (
	"demo/mystrings"
	"fmt"
)

func main() {
	// Example usage of both functions
	n := 1000
	fmt.Println(mystrings.ConcatenateUsingPlusOperator(n))
	fmt.Println(mystrings.ConcatenateUsingBuilder(n))
}
