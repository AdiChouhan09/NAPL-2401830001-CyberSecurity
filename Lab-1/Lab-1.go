package main

import "fmt"

func main() {
	fmt.Println("===== SIMPLE CALCULATOR =====")
	fmt.Println("Hello! \"Aditya\"")

	for {
		var choice int

		fmt.Println("\nChoose an option:")
		fmt.Println("1. Perform operations using integers")
		fmt.Println("2. Perform operations using float values")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("\nInvalid input. Please enter a number.")
			return
		}

		switch choice {
		case 1:
			var a, b int
			fmt.Print("\nEnter two integers: ")
			if _, err := fmt.Scan(&a, &b); err != nil {
				fmt.Println("\nInvalid integer input.")
				return
			}

			fmt.Println("\n--- Integer Operations ---")
			fmt.Println("Addition Values =", a+b)
			fmt.Println("Subtraction Values =", a-b)
			fmt.Println("Multiplication Values =", a*b)

		case 2:
			var x, y float64
			fmt.Print("\nEnter two float values: ")
			if _, err := fmt.Scan(&x, &y); err != nil {
				fmt.Println("\nInvalid float input.")
				return
			}

			fmt.Println("\n--- Floating-Point Operations ---")
			fmt.Println("Addition Values =", x+y)
			fmt.Println("Subtraction Values =", x-y)
			fmt.Println("Multiplication Values =", x*y)

		case 3:
			fmt.Println("\nThank you for using the calculator!")
			fmt.Println("\n===== Program Completed =====")
			return

		default:
			fmt.Println("\nInvalid choice! Please select 1, 2 or 3.")
		}
	}
}
