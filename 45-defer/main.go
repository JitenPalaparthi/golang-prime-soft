package main

func main() {
	defer println("End of main")
	num := 100
	func() { // func1
		defer println("end of func1")
		func() { // func2
			defer println("end of func-2")
			println("start of func-2")
			for i := 10; i >= 0; i-- {
				print(num/i, " ") // panics
			}
			defer println("defer after panic --???")
			println("graceful exit-->?")
		}()
		println("do something in func1")
	}()

	str := "hello Wrold"
	for _, v := range str {
		println(string(v))
	}
	println("graceful exit of main-->?")
}
