package main

import (
	"fmt"
	"math"
)

func main() {
	// Multiple return values - get both sum and difference
	sum, diff := addAndSubtract(10, 5)
	fmt.Printf("Sum: %d, Difference: %d\n", sum, diff)

	// Multiple return values - calculate area and perimeter
	area, perimeter := rectangle(4, 6)
	fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", area, perimeter)

	// Multiple return values - check if number is prime and get its factors
	isPrime, factors := analyzeNumber(17)
	fmt.Printf("Number 17 - Prime: %t, Factors: %v\n", isPrime, factors)

	// Using underscore to ignore one return value
	_, perimeter2 := rectangle(3, 8)
	fmt.Printf("Rectangle perimeter only: %.2f\n", perimeter2)
}

// addAndSubtract returns both sum and difference of two numbers
func addAndSubtract(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// rectangle returns both area and perimeter of a rectangle
func rectangle(width, height float64) (float64, float64) {
	area := width * height
	perimeter := 2 * (width + height)
	return area, perimeter
}

// analyzeNumber returns if number is prime and its factors
func analyzeNumber(n int) (bool, []int) {
	if n <= 1 {
		return false, []int{}
	}

	if n == 2 {
		return true, []int{1, 2}
	}

	// Check for primality
	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	// Get all factors
	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}

	return isPrime, factors
}