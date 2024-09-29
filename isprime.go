package main

// Function to check if a number is prime
func IsPrime(num int) bool {
	if num <= 1 {
		return false // Numbers less than or equal to 1 are not prime
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false // If num is divisible by i, it's not prime
		}
	}
	return true // If no divisors were found, num is prime
}
