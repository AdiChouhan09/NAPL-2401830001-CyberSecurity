package main

import (
	"fmt"

	"CustomPackage/myutils"
)

func main() {
	text := "Hello World"

	fmt.Println("Original String:", text)
	fmt.Println("Reversed String:", myutils.Reverse(text))
	fmt.Println("Number of Vowels:", myutils.CountVowels(text))

	fmt.Println("Factorial of 5:", myutils.Factorial(5))
	fmt.Println("2 raised to the power 5:", myutils.Power(2, 5))
}
