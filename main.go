// main.go
package main

import (
	"fmt"
	"test-package/math"
	"test-package/hello"
	"goodbye-module"
)

func main() {
	// Test our math package
	a := 5
	b := 3
	
	sum := math.Add(a, b)
	product := math.Multiply(a, b)
	square := math.Square(a)
	
	fmt.Printf("%d + %d = %d\n", a, b, sum)
	fmt.Printf("%d * %d = %d\n", a, b, product)
	fmt.Printf("%d squared = %d\n", a, square)
	
	// Test hello package
	hello.SayHello()
	
	// Test goodbye module
	goodbye.SayGoodbye()
}
