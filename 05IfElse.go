package main

import "fmt"

var exam01 int
var exam02 int
var avarage int

func main() {

	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even", i)
		} else {
			fmt.Printf("%d is odd", i)
		}
		fmt.Println()
	}

	exam01, exam02 = 30, 90
	avarage = (exam01 + exam02) / 2

	if avarage > 0 && avarage < 40 {
		fmt.Println("Bad avarage!")
	} else if avarage > 41 && avarage < 70 {
		fmt.Println("Nice avarage!")
	} else {
		fmt.Print("Very well avarage")
	}

}
