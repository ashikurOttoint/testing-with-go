package main

import "fmt"

func main() {
	n := 7
	_, msg := isPrime(n)
	fmt.Println(msg)
}

func isPrime(n int) (bool, string) {
	// 0 and 1 is not prime by defination
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime by defination!", n)
	}

	// Negative numbers are not prime
	if n < 0 {
		return false, "negative numbers are not prime!"
	}
	for i := 2; i < n/2; i++ {
		if n%2 == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number", n)
}
