package main

// This function calculates the factorial of a number.
func FindFactorial(numb int) int {
	if numb < 0 {
		return -1 // If the number is negative, we can't calculate factorial.
	} else if numb == 0 {
		return 1 // The factorial of 0 is 1.
	}

	Ans := 1
	// Loop through numbers from 1 to n and multiply them together.
	for i := 1; i <= numb; i++ {
		Ans *= i
	}
	return Ans // Return the calculated factorial.
}
