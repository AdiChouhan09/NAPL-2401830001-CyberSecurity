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

	fmt.Println("===== SIMPLE CALCULATOR =====")
	fmt.Println("Hello! \"Aditya\"")

	for {
		var choice int

		fmt.Println("\nChoose an option:")
		fmt.Println("1. Perform operations using integers")
		fmt.Println("2. Perform operations using float values")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		fmt.Scan(&choice)

		if choice == 1 {
			var a, b int

			fmt.Print("\nEnter two integers: ")
			fmt.Scan(&a, &b)

			fmt.Println("\n--- Integer Operations ---")
			fmt.Println("Addition Values =", a+b)
			fmt.Println("Subtraction Values =", a-b)
			fmt.Println("Multiplication Values =", a*b)

		} else if choice == 2 {
			var x, y float64

			fmt.Print("\nEnter two float values: ")
			fmt.Scan(&x, &y)

			fmt.Println("\n--- Floating-Point Operations ---")
			fmt.Println("Addition Values =", x+y)
			fmt.Println("Subtraction Values =", x-y)
			fmt.Println("Multiplication Values =", x*y)

		} else if choice == 3 {
			fmt.Println("\nThank you for using the calculator!")
			break

		} else {
			fmt.Println("\nInvalid choice! Please select 1, 2 or 3.")
		}
	}

	fmt.Println("\n===== Program Completed =====")
}
