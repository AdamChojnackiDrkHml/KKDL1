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
	scanFile("test")
}

func scanFile(path string) {

	counterSlice := make([]int, 256)

	f, err := os.Open(path)
	check(err)

	currSymbol := make([]byte, 1)
	// prevSymbol := make([]byte, 1)

	for {
		control, _ := f.Read(currSymbol)
		if control == 0 {
			break
		}
		counterSlice[currSymbol[0]]++
		fmt.Println((currSymbol))
	}

	// f.Read(prevSymbol)

	// fmt.Println((prevSymbol))

	defer f.Close()
}
