package main

import v2 "math/rand/v2"

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println()
	j := 1
	for {
		if j > 10 {
			break
		}
		print(j, " ")
		j++
	}

	println()

	// done := false
	// for i := 1; i <= 10; i++ {
	// 	if done {
	// 		break
	// 	}
	// 	for j := 5; j <= 10; j++ {
	// 		if i == j {
	// 			done = true
	// 			break
	// 		}
	// 		println("i:i", i, "j:", j)
	// 	}
	// }
out:
	for k := 1; k <= 5; k++ {
		for i := 1; i <= 10; i++ {
			for j := 5; j <= 10; j++ {
				if i == j && j == k {
					break out
				}
				println("i:i", i, "j:", j, "k:", k)
			}
		}
	}

	i := 1
outer:
	for {
		num := v2.IntN(100)
		switch {
		case num%2 == 0:
			{
				println(num, "is even number")
				if i == 10 {
					break outer
				}
			}
		default:
			println(num, "is even number")
			if i == 10 {
				break outer
			}
		}
		i++
	}

	for i := 1; ; i++ {
		if i > 10 {
			break
		}
		if i%2 == 0 {
			continue
		}
		println(i)

	}

	k := 1

	for k <= 10 { // like a while loop
		println(k)
		k++
	}

	for i, j, ok := 1, 10, true; j >= i && ok && true && (true || false); i, j, ok = i+1, j-1, true {
		println(i, "-->", j, "condition:", j >= i && ok && true && (true || false))
	}

}
