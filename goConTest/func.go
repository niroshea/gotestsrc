package goConTest

import (
	"errors"
)

// Add xxx
func Add(a, b int) int {
	return a + b
}

// Subtract xxx
func Subtract(a, b int) int {
	return a - b
}

// Multiply xxx
func Multiply(a, b int) int {
	return a * b
}

// Division xxx
func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0 。")
	}
	return a / b, nil
}
