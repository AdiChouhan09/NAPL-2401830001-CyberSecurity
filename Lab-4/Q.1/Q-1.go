package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Goroutine 1 started: Calculating squares")
		sum := 0
		for i := 1; i <= 5; i++ {
			sum += i * i
		}
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 1 completed")
		ch <- fmt.Sprintf("Sum of squares: %d", sum)
	}()

	go func() {
		fmt.Println("Goroutine 2 started: Calculating cubes")
		sum := 0
		for i := 1; i <= 5; i++ {
			sum += i * i * i
		}
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 2 completed")
		ch <- fmt.Sprintf("Sum of cubes: %d", sum)
	}()

	go func() {
		fmt.Println("Goroutine 3 started: Calculating Fibonacci numbers")
		a, b := 0, 1
		for i := 0; i < 5; i++ {
			a, b = b, a+b
		}
		time.Sleep(1500 * time.Millisecond)
		fmt.Println("Goroutine 3 completed")
		ch <- fmt.Sprintf("Fibonacci number: %d", a)
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}
