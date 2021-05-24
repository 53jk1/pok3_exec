package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var n int
	var n2 uint
	fmt.Println("N for Leibniz?") // Question to the user
	_, err := fmt.Scan(&n)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("N for Euler?") // Question to the user
	_, err = fmt.Scan(&n2)

	if err != nil {
		log.Fatal(err)
	}

	leibnizError := leibniz(n)
	eulerError := Phi(n2)

	fmt.Printf("\nRozwój błędów schematu Leibniza dla k = (0, %d): \n %v \n", n, leibnizError)
	fmt.Printf("\nRozwój błędów schematu Eulera dla k = (0, %d): \n %v \n", n2, eulerError)

}

func leibniz(k int) float64 {
	sum := 0.0
	for i := 0; i < k; i++ {
		sum += leibniz_step(k)
	}

	return sum * 4
}

func leibniz_step(k int) float64 {
	var kf float64 = float64(k)
	return math.Pow(-1, kf) / (2*kf + 1)
}

// Euler’s Totient Function
func Phi(n uint) (result uint) {
	// Initialize result as n
	result = n

	// Consider all prime factors of n and subtract their
	// multiples from result
	var p uint
	for p = 2; p*p <= n; p++ {
		// Check if p is a prime factor.
		if n%p == 0 {
			// If yes, then update n and result
			for n%p == 0 {
				n /= p
			}
			result -= (result / p)
		}
	}

	// If n has a prime factor greater than sqrt(n)
	// (There can be at-most one such prime factor)
	if n > 1 {
		result -= (result / n)
	}
	return result
}
