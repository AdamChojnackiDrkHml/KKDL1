package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	scanFile(os.Args[1])
}

func scanFile(path string) {
	f, err := os.Open(path)
	check(err)

	currSymbol := make([]byte, 8)
	prevSymbol := make([]byte, 8)

	f.Read(currSymbol)
	f.Read(prevSymbol)

	fmt.Println(string(currSymbol))
	fmt.Println(string(prevSymbol))
}
