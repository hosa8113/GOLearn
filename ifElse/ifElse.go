package main

import (
	"fmt"
	"math"
)

func main() {

	simpleIf()
	ifWithShortStatement()
	ifElse()

}

func ifElse() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func pow3(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else if x > n {
		fmt.Printf("%g >= %g\n", v, lim)
	} else {

	}
	// can't use v here, though
	return lim
}

func ifWithShortStatement() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func simpleIf() {
	fmt.Printf("SquareRoot=%v %v\n", sqrt(2), sqrt(-4))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
