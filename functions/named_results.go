package main

import (
	"fmt"
	"strings"
)

func main() {
	// Named result parameters - simple division
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 ÷ 5 = %d remainder %d\n", quotient, remainder)

	// Named result parameters - person info
	name, isAdult := getPersonInfo("Alice", 25)
	fmt.Printf("Name: %s, Age: 25, Adult: %t\n", name, isAdult)

	// Named result parameters - string analysis
	upper, lower, digitCount := analyzeString("Hello123")
	fmt.Printf("String 'Hello123': Uppercase=%d, Lowercase=%d, Digits=%d\n", upper, lower, digitCount)

	// Named result parameters with default values
	fmt.Println("Testing splitString:")
	words, count := splitString("Go is awesome")
	fmt.Printf("Words: %v, Count: %d\n", words, count)
}

// Named result parameters - quotient and remainder
// The variables quotient and remainder are created at function start
func divide(dividend, divisor int) (quotient int, remainder int) {
	if divisor == 0 {
		fmt.Println("Error: Division by zero")
		return 0, dividend // Returns zeros
	}

	quotient = dividend / divisor
	remainder = dividend % divisor
	// No explicit values needed in return - returns current values of quotient and remainder
	return
}

// Named result parameters - person information
func getPersonInfo(name string, age int) (formattedName string, isAdult bool) {
	// Named parameters get zero values initially
	formattedName = strings.Title(strings.ToLower(name))
	isAdult = age >= 18
	
	// Return without explicit values
	return
}

// Named result parameters - string character analysis
func analyzeString(s string) (uppercaseCount, lowercaseCount, digitCount int) {
	for _, char := range s {
		switch {
		case char >= 'A' && char <= 'Z':
			uppercaseCount++
		case char >= 'a' && char <= 'z':
			lowercaseCount++
		case char >= '0' && char <= '9':
			digitCount++
		}
	}
	// Naked return - returns all named result parameters
	return
}

// Named result parameters with default values demonstration
func splitString(s string) (words []string, count int) {
	if s == "" {
		// Returns nil slice and 0 (zero values)
		return
	}

	words = strings.Fields(s)
	count = len(words)
	return
}