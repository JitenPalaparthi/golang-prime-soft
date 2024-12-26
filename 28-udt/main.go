package main

import (
	"errors"
	"fmt"
)

func main() {

	var m1 mymap
	m1 = make(mymap)
	//m1 = make((map[string]any)(mymap))

	m1["560086"] = "Blr-1"
	m1["560096"] = "Blr-2"
	m1["522001"] = "Gnt-1"
	m1["522002"] = "Gnt-2"

	for k, v := range m1 {
		fmt.Println("Key:", k, "Value:", v)
	}

	if keys, err := m1.GetKeys(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Keys:", keys)
	}

	if values, err := m1.GetValues(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Values:", values)
	}

	if err := m1.Delete("343242"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("successfully deleted")
	}

	if err := m1.Delete("560086"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("successfully deleted")
	}

	fmt.Println("normal map")
	m2 := make(map[string]any, 4)
	m2["560086"] = "Blr-1"
	m2["560096"] = "Blr-2"
	m2["522001"] = "Gnt-1"
	m2["522002"] = "Gnt-2"

	for k, v := range m2 {
		fmt.Println("Key:", k, "Value:", v)
	}

	if keys, err := mymap(m2).GetKeys(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("keys:", keys)
	}

	if values, err := mymap(m2).GetValues(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("values:", values)
	}

	// m3 := make(map[string]string)

	// mymap(m3).GetKeys()

}

type mymap map[string]any // make, delete, range ,clear

func (m mymap) GetKeys() ([]string, error) {
	if m == nil {
		return nil, errors.New("invalid map or nil map")
	}
	keys := make([]string, len(m)) // keys:=[    ]
	index := 0
	for k := range m {
		keys[index] = k
		index++
	}
	return keys, nil
}

func (m mymap) GetValues() ([]any, error) {
	if m == nil {
		return nil, errors.New("invalid map or nil map")
	}
	var values []any // ptr=nil, cap=0 len=0
	for _, v := range m {
		values = append(values, v)
	}
	return values, nil
}

func (m mymap) Delete(key string) error {
	if m == nil {
		return errors.New("invalid map or nil map")
	}
	_, ok := m[key]
	if !ok {
		return fmt.Errorf("%v is not found", key)
	}
	delete(m, key)
	return nil
}
