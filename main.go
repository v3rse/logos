package main

import (
	"fmt"
	"os"
)

func ReadPostFile(path string) []byte {
	file, _ := os.ReadFile(path)

	return file
}

func main() {
	input := os.Args[1]

	output := ReadPostFile(input)

	fmt.Println(string(output))
}
