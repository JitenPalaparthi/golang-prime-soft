package main

type interger = int // creating an alias
//type rune = int32

func main() {
	var day uint8 = 4
	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("noday")
	}

	// empty switch

	num := -100

	switch {

	case num >= 0 && num <= 50:
		{
			println(num, " is between 0 and 50")
		}
	case num > 50 && num <= 100:
		{
			println(num, " is between 50 and 100")
		}
	case num > 100:
		{
			println(num, " is greater than 100")
		}
	default:
		println(num, " is less than 0")
	}

	var char rune = 'A'

	switch char {
	//case char=='A' || char=='E' || char=='I':
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		println(string(char), " is vovel")
	default:
		println(string(char), "is either a consonent or a special character")
	}
	// fallthrough is no different than avoiding break in other langauges

	num = 48
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
		// default:
		// 	println(num, "is not dibisible by 8 or 4 or 2")
	}

	num = 6

	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}

	// var e1 empty

	// e1 = "Hello World"

	// str1 := e1.(string)

}

// empty is same as any or interface{}. In all aspects three of them are same
type empty = interface{}

// switch case
