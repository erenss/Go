package main

import "fmt"

func main() {
	i := 1
	// this for loops look like While statements
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	//classic for loops
	for j := 6; j <= 10; j++ {
		fmt.Println(j)
	}

	//this example for about odd numbers
	//if your number odd when print to console
	for k := 0; k <= 10; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}

}
