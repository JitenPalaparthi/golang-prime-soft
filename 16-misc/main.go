package main

import "fmt"

func main() {

	cx1 := 12.34 + 2.3i
	cx2 := complex(12.344, 2.3)

	cx3 := cx1 + cx2
	cx4 := cx1 - cx2

	var a, b float32 = 12.2, 1.2
	cx5 := complex(a, b)

	var cx6 complex64 = complex64(cx1) * cx5

	fmt.Println(cx1, cx2, cx3, cx4, cx5, cx6)

	r1, img1 := real(cx1), imag(cx1)
	println(r1, img1)

	a1, b1, c1 := 10, 20, 30

	// t1 := b1
	// b1 = a1
	// a1 = t1

	a1, b1 = b1, a1 // swapping

	a1, b1, c1 = b1, c1, a1 // swap of 3 variables

	a2, b2 := 1, 1 // 1 1 2 3 5 8 13 21 34 55
	println(a2, " ")
	for i := 1; i <= 10; i++ {
		a2, b2 = b2, a2+b2 // swapping
		print(a2, " ")
	}
	println()

	i := 1
loop:
	println(i)
	i++
	if i <= 10 {
		goto loop
	} else {
		goto exit
	}
exit:
	fmt.Println("loop exited")

}

// 1-50
// if the number is even number jump to a label call even and print
// if the number is odd number jump to a lebel call odd and print
// if the number is >50 it should jump to a label called exit
// if the number is <50 it should go to a label call loop
