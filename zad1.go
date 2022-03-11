package main

import (
	"fmt"
	"math"
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

		for j := range probs2[i] {
			probs2[i][j] = 0.0
			counterSlice2[i][j] = 0
		}
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

	//index i, j holds value of P(j|i)
	for i, k := range counterSlice2 {
		for j, l := range k {
			if l != 0 && counterSlice[i] != 0 {
				probs2[i][j] = float64(l) / float64(counterSlice[i])
			} else {
				probs2[i][j] = 0.0
			}
		}
	}

	//COUNT H(Y|X)
	HYX := 0.0
	for j := 0; j < 256; j++ {

		//COUNT H(Y|x)
		HYx := 0.0
		for i := 0; i < 256; i++ {
			Pyx := probs2[i][j]
			if Pyx != 0.0 {
				//fmt.Print(" ", Pyx)
				Iyx := -math.Log2(Pyx)
				HYx += Pyx * Iyx
			}

		}

		HYX += probs1[j] * HYx
	}

	H := 0.0

	for i := 0; i < 256; i++ {
		Px := probs1[i]
		if Px != 0.0 {
			Ix := -math.Log2(Px)
			H += Px * Ix
		}
	}

	defer f.Close()
	fmt.Println("HYX = ", HYX)
	fmt.Println("H = ", H)
}
