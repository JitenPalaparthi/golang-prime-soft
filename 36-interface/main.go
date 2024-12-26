package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello World")
	fw := New("data.txt")
	fmt.Fprintln(fw, "Hello PrimeSoft")
}

type FileWriter struct {
	FileName string
}

func New(fn string) *FileWriter {
	return &FileWriter{FileName: fn}
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	if fw == nil {
		return -1, errors.New("invalid or nil FileWriter")
	}
	f, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return -1, err
	}
	defer f.Close() // defer defers the execution to the end of the call stack
	return f.Write(p)
}
