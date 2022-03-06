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
	counterSlice2 := make([][]int, 256)
	probs1 := make([]float64, 256)
	probs2 := make([][]float64, 256)

	for i := range probs2 {
		probs2[i] = make([]float64, 256)
		counterSlice2[i] = make([]int, 256)
	}
	counter := 0

	f, err := os.Open(path)
	check(err)

	var prevSymbol byte
	currSymbol := make([]byte, 1)
	// prevSymbol := make([]byte, 1)

	for {
		prevSymbol = currSymbol[0]
		control, _ := f.Read(currSymbol)
		if control == 0 {
			break
		}
		counter++
		counterSlice[currSymbol[0]]++
		counterSlice2[prevSymbol][currSymbol[0]]++
		// fmt.Println((currSymbol))
	}

	for i, k := range counterSlice {
		probs1[i] = float64(k) / float64(counter)
	}

	for _, k := range counterSlice2 {
		fmt.Println(k)
		fmt.Println()
	}
	//fmt.Println(probs1)
	// fmt.Println((prevSymbol))

	defer f.Close()
}
