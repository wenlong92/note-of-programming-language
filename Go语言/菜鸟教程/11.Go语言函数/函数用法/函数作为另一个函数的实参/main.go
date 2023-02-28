package main

import (
	"fmt"
	"math"
)

func main() {
	// Go语言可以很灵活地创建函数，并作为另一个函数的实参
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))
}
