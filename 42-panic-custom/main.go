package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fn1 := getFullName1("Jiten", "")
	// fn1, err := getFullName2("Jiten", "")
	// if err != nil {
	// 	Fatal("Fatal:", err.Error())
	// }
	fmt.Println(fn1)
}

func Fatal(msgs ...any) {
	fmt.Println(msgs...)
	os.Exit(1) // exiting the application abruptly is fatal.
	// once os.Exit is called , cannot recover from it
}

// we dont panic usually for the situations like below, rather we handle or return errors

func getFullName1(firstname, lastname string) string {
	if firstname == "" {
		panic("firstname cannot be empty")
	}
	if lastname == "" {
		panic("lastname cannot be empty")
	}
	return fmt.Sprint(firstname, " ", lastname)
}

func getFullName2(firstname, lastname string) (string, error) {
	if firstname == "" {
		return "", errors.New("invalid firstname")
	}
	if lastname == "" {
		return "", errors.New("invalid lastname")
	}

	return fmt.Sprint(firstname, " ", lastname), nil
}

// any non zero exit is considered as a failure exit of the application.
