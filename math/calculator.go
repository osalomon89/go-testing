package main

import "fmt"

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}

// Multiply two integers
func Multiply(x, y int) (res int) {
	return x * y
}

// Int Division
func Div(x, y int) (res int) {
	return x / y
}

func main() {

	add := Add(4, 5)
	fmt.Println("Division: ", add)

	subs := Subtract(4, 5)
	fmt.Println("Division: ", subs)

	multiply := Multiply(4, 5)
	fmt.Println("Division: ", multiply)

	division := Div(4, 5)
	fmt.Println("Division: ", division)
}
