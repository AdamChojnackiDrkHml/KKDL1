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

	counterSlice := make([]int, 256)
	probs1 := make([]float64, 256)
	counter := 0

	f, err := os.Open(path)
	check(err)

	currSymbol := make([]byte, 1)
	// prevSymbol := make([]byte, 1)

	for {
		control, _ := f.Read(currSymbol)
		if control == 0 {
			break
		}
		counter++
		counterSlice[currSymbol[0]]++
		// fmt.Println((currSymbol))
	}

	for i, k := range counterSlice {
		probs1[i] = float64(k) / float64(counter)
	}

	fmt.Println(counterSlice)
	fmt.Println(probs1)
	// fmt.Println((prevSymbol))

	defer f.Close()
}
