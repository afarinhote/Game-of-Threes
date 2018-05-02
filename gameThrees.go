package main

import (
	"fmt"
)

func main() {
	var numberToDivide, oldNumber int
	fmt.Println("Enter number: ")
	fmt.Scanln(&numberToDivide)
	for numberToDivide != 1 {
		if numberToDivide == 2 {
			oldNumber = numberToDivide
			numberToDivide = numberToDivide + 1
			fmt.Println(oldNumber, "+ 1 :", numberToDivide)
		}
		if numberToDivide%3 == 0 {
			oldNumber = numberToDivide
			numberToDivide = numberToDivide / 3
			fmt.Println(oldNumber, "/ 3 :", numberToDivide)
		} else {
			oldNumber = numberToDivide
			numberToDivide = numberToDivide - 1
			fmt.Println(oldNumber, "- 1 :", numberToDivide)
		}
	}

}
