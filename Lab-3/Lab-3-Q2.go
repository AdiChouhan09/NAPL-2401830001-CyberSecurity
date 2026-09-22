package main

import "fmt"

func modifyValue(num *int) {
	*num = *num + 10
}

type Student struct {
	Name  string
	Marks float64
	Age   int
}

func main() {

	var x int = 20
	ptr := &x

	fmt.Println("Original value:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Pointer value:", ptr)
	fmt.Println("Value using dereferencing:", *ptr)

	fmt.Println("\nBefore modification:", x)
	modifyValue(ptr)
	fmt.Println("After modification:", x)

	student := new(Student)

	student.Name = "Aditya"
	student.Marks = 85.5
	student.Age = 20

	fmt.Println("\nStudent details:")
	fmt.Println("Name:", student.Name)
	fmt.Println("Marks:", student.Marks)
	fmt.Println("Age:", student.Age)

	student.Marks = 92.5
	student.Age = 21

	fmt.Println("\nAfter modifying marks and age:")
	fmt.Println("Name:", student.Name)
	fmt.Println("Marks:", student.Marks)
	fmt.Println("Age:", student.Age)
}
