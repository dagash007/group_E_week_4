package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group E's Week 4 Project!")

	// Let's say we want to find the factorial of 5.
	numb2 := 5
	Findfact := FindFactorial(numb2) // Call the factorial function.

	// Check if the result is -1, which means the input was negative.
	if Findfact == -1 {
		fmt.Println("Oops! Factorial is not defined for negative numbers.")
	} else {
		fmt.Printf("The factorial of %d is %d.\n", numb2, Findfact) // Print the result.
	}
}
