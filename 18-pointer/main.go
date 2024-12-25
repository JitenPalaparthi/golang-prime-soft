package main

func main() {
	var num1 int // zero value - 0

	var ptr1 *int // zero value - nil

	num1 = 100

	ptr1 = &num1 // reference
	*ptr1 = 200  // deference

	var ptr2 *int

	// *ptr2 = 200 // what is the problem here? derefenceing a nil pointer
	ptr2 = &num1

	*ptr2 = 400

	*&num1 = 600

	var ptr3 **int = &ptr2
	var ptr4 ***int = &ptr3

	***ptr4 = 700

	println(num1)

	ptr5 := new(int) // it allocates some(based on input argument) memory , that address is assigned to the pointer
	//make --?
	*ptr5 = 12312

	var ptr6 *int = new(int)

	*ptr6 = 12312

	// var ptr7 *int

	// *ptr7 = 23434

	var ptr8 *bool = new(bool)
	println(*ptr8)

}

// what is a pointer?
// what does it contain?
// what is a zero value of a pointer?
// what is  dereference?

// 1. What happend if you derenfence a nil pointer
