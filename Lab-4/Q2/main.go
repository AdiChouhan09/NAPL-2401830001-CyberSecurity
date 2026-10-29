package main

import (
	"fmt"
	"math"
	"strings"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative number")
	}
	return math.Sqrt(n), nil
}

func checkInput(s string) error {
	if strings.TrimSpace(s) == "" {
		return fmt.Errorf("empty input")
	}
	return nil
}

func main() {
	_, e1 := divide(10, 0)
	fmt.Println(e1)

	_, e2 := squareRoot(-4)
	fmt.Println(e2)

	e3 := checkInput("")
	fmt.Println(e3)
}
