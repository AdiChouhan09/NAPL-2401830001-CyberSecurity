package main

import "fmt"

func main() {
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
}
