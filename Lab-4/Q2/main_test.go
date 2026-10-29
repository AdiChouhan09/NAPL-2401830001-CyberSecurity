package main

import "testing"

func TestDivide(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Error("expected division error")
	}
}

func TestSquareRoot(t *testing.T) {
	_, err := squareRoot(-4)
	if err == nil {
		t.Error("expected square root error")
	}
}

func TestInput(t *testing.T) {
	err := checkInput("")
	if err == nil {
		t.Error("expected empty input error")
	}
}
