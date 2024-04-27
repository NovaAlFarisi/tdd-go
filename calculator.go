package main

import "errors"

func Calculate(operation string, x, y float64) (float64, error) {
	var result float64
	switch operation {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		result = x / y
	default:
		return 0, errors.New("invalid operation")
	}
	return result, nil
}
