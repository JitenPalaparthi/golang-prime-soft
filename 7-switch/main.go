package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	var any1 any

	// what is zero value - nil
	// type ?  - nil

	fmt.Println("Value:", any1, "Type:", reflect.TypeOf(any1))
	fmt.Printf("Value: %v Type: %T\n", any1, any1)
	// if num, err := strconv.Atoi("23213"); err != nil {

	// }

	any1 = int(100)
	switch v := any1.(type) { // type switch , used only using any or empty interface
	case int:
		println(v * v)
	case int8:
		println(any1.(int8) * any1.(int8))
	case int16:
		println(any1.(int16) * any1.(int16))
	case int32:
		println(any1.(int32) * any1.(int32))
	case int64:
		println(any1.(int64) * any1.(int64))
	case uint:
		println(any1.(uint) * any1.(uint))
	case uint8:
		println(any1.(uint8) * any1.(uint8))
	case uint16:
		println(any1.(uint16) * any1.(uint16))
	case uint32:
		println(any1.(uint32) * any1.(uint32))
	case uint64:
		fmt.Println(any1.(uint64) * any1.(uint64))
	case float32:
		fmt.Println(any1.(float32) * any1.(float32))
	case float64:
		fmt.Println(any1.(float64) * any1.(float64))
	default:
		println(" not a number type so cannot square")
	}

	var any2 any = rune('A')
	if r, err := GetSquare(any2); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Square of any2: %v\n", r)
	}

}

func GetSquare(any1 any) (any, error) {
	if any1 == nil {
		return nil, fmt.Errorf("invalid input or nil input")
	}
	var result any
	switch v := any1.(type) { // type switch , used only using any or empty interface
	case int:
		result = v * v
	case int8:
		result = v * v
	case int16:
		result = v * v
	case int32:
		result = v * v
	case int64:
		result = v * v
	case uint:
		result = v * v
	case uint8:
		result = v * v
	case uint16:
		result = v * v
	case uint32:
		result = v * v
	case uint64:
		result = v * v
	case float32:
		result = v * v
	case float64:
		result = v * v
	default:
		return nil, errors.New("input is not a number")
	}

	return result, nil
}

func IsNumber(any1 any) (bool, error) {
	if any1 == nil {
		return false, fmt.Errorf("invalid input or nil input")
	}
	switch any1.(type) { // type switch , used only using any or empty interface
	case int, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		return true, nil
	default:
		return false, errors.New("input is not a number")
	}
}
