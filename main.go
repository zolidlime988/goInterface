package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	readArg := os.Args[1]
	fmt.Println("Open File:", readArg)
	contentWrites, err := os.Open(readArg)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, contentWrites)
	fmt.Println()
}
