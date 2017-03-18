package main

import (
	"fmt"
	"math"
)

const pi float32 = 3.14

func main() {
	const number = 9
	//const declares a constant value.
	//cannot change to value to pi const variable, its fault
	fmt.Println(pi)

	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt(number))

}
