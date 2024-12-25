package main

import "fmt"

func main() {

	var map1 map[string]string // defining a map

	map1 = make(map[string]string)

	map1["560086"] = "Bangalore-1"
	map1["560096"] = "Bangalore-2"
	map1["560036"] = "Bangalore-3"
	map1["560044"] = "Bangalore-4"

	//map2 := map[string]any{"560086": "Bangalore-1", "560096": "Bangalore-2"} // creating a map with data

	for k, v := range map1 { // range loop on map gives key and value
		fmt.Println("key:", k, "Value:", v)
	}

	v, ok := map1["560086"]

	if ok {
		fmt.Println(v)
	} else {
		println("key does not exist")
	}

	v1, done := map1["560046"]

	if done {
		fmt.Println(v1)
	} else {
		println("key does not exist")
	}

	delete(map1, "560086")
	for k, v := range map1 { // range loop on map gives key and value
		fmt.Println("key:", k, "Value:", v)
	}

	delete(map1, "5600126")
	for k, v := range map1 { // range loop on map gives key and value
		fmt.Println("key:", k, "Value:", v)
	}

	//_, done2 := map1["qwqwe"]

	// var any1 any = 12323

	// v, ok = any1.(bool)

	// fmt.Println(v)

}

// sha
// hashfunc
// loadfactor
// bucket
// overflow

// what type can be a key : a key can be any type on which can apply == operation
// maps are not orderd
// maps are not thread safe

func deleteKey(map1 map[string]any, key string) error {

	return nil
}

// _
