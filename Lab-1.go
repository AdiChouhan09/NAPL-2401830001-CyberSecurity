package main

import "fmt"

func main() {

	/*
		fmt.Println("===== GO CALCULATOR =====")
		fmt.Println("Hello! \"Aditya\"")

		a := 10
		b := 5

		fmt.Println("\nInteger Values:")
		fmt.Println("a =", a)
		fmt.Println("b =", b)
		fmt.Println("Addition Values =", a+b)
		fmt.Println("Subtraction Values =", a-b)
		fmt.Println("Multiplication Values =", a*b)

		x := 10.5
		y := 5.5

		fmt.Println("\nFloat Values:")
		fmt.Println("x =", x)
		fmt.Println("y =", y)
		fmt.Println("Addition Values =", x+y)
		fmt.Println("Subtraction Values =", x-y)
		fmt.Println("Multiplication Values =", x*y)

		fmt.Println("\n===== Program Completed =====")
	*/

	// Question 2 - Simple Calculator

	fmt.Println("===== SIMPLE CALCULATOR =====")

	var a, b int

	for {
		fmt.Print("\nEnter two integers (1-100): ")
		_, err := fmt.Scan(&a, &b)

		if err == nil && a >= 1 && a <= 100 && b >= 1 && b <= 100 {
			break
		}

		fmt.Println("Invalid input! Enter numbers between 1 and 100.")
		fmt.Scanln()
	}

	fmt.Println("\nInteger Values:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("Addition Values =", a+b)
	fmt.Println("Subtraction Values =", a-b)
	fmt.Println("Multiplication Values =", a*b)

	var x, y float64

	for {
		fmt.Print("\nEnter two decimal numbers (1-100): ")
		_, err := fmt.Scan(&x, &y)

		if err == nil && x >= 1 && x <= 100 && y >= 1 && y <= 100 {
			break
		}

		fmt.Println("Invalid input! Enter numbers between 1 and 100.")
		fmt.Scanln()
	}

	fmt.Println("\nFloat Values:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("Addition Values =", x+y)
	fmt.Println("Subtraction Values =", x-y)
	fmt.Println("Multiplication Values =", x*y)

	fmt.Println("\n===== Program Completed =====")
}
