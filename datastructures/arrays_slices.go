package main

import (
	"fmt"
	"strings"
)

// ARRAYS vs SLICES
// ================
// Arrays: Fixed-size, value types (copied when passed)
// Slices: Dynamic-size, reference types (backed by arrays)

// ArrayBasics demonstrates fundamental array concepts
func ArrayBasics() {
	fmt.Println("\n=== ARRAY BASICS ===")

	// Arrays have fixed size, part of their type
	var arr1 [5]int // Array of 5 ints, initialized to [0, 0, 0, 0, 0]
	fmt.Printf("Empty array: %v\n", arr1)

	// Array literal initialization
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Initialized array: %v\n", arr2)

	// Let compiler count the size
	arr3 := [...]int{10, 20, 30}
	fmt.Printf("Auto-sized array: %v (length: %d)\n", arr3, len(arr3))

	// Accessing elements (zero-indexed)
	fmt.Printf("First element: %d, Last element: %d\n", arr2[0], arr2[4])

	// Arrays are VALUE types - copying creates a new array
	arr4 := arr2
	arr4[0] = 999
	fmt.Printf("Original: %v, Copy: %v (independent)\n", arr2, arr4)
}

// SliceBasics demonstrates fundamental slice concepts
func SliceBasics() {
	fmt.Println("\n=== SLICE BASICS ===")

	// Slices are dynamic and reference an underlying array
	var slice1 []int // nil slice (no underlying array yet)
	fmt.Printf("Nil slice: %v, len=%d, cap=%d, is nil? %v\n",
		slice1, len(slice1), cap(slice1), slice1 == nil)

	// Slice literal (creates underlying array automatically)
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice literal: %v, len=%d, cap=%d\n",
		slice2, len(slice2), cap(slice2))

	// Using make() - PROPER WAY to create slices
	// make([]Type, length, capacity)
	slice3 := make([]int, 5)     // length 5, capacity 5, initialized to zeros
	slice4 := make([]int, 3, 10) // length 3, capacity 10

	fmt.Printf("make([]int, 5): %v, len=%d, cap=%d\n",
		slice3, len(slice3), cap(slice3))
	fmt.Printf("make([]int, 3, 10): %v, len=%d, cap=%d\n",
		slice4, len(slice4), cap(slice4))
	slice3 = append(slice3, 100)
	slice3 = append(slice3, 101)
	slice3 = append(slice3, 102)
	slice3 = append(slice3, 103)
	slice3 = append(slice3, 104)
	slice3 = append(slice3, 105)
	fmt.Printf("make([]int, 5) UPDATED: %v, len=%d, cap=%d\n",
		slice3, len(slice3), cap(slice3))
	slice4 = append(slice4, 694)
	fmt.Printf("make([]int, 3, 10) UPDATED: %v, len=%d, cap=%d\n",
		slice4, len(slice4), cap(slice4))

	// Slices are REFERENCE types - they share the underlying array
	slice5 := slice2
	slice5[0] = 999
	fmt.Printf("Original: %v, Reference: %v (shared backing array)\n", slice2, slice5)
}

// SliceOperations demonstrates common slice operations
func SliceOperations() {
	fmt.Println("\n=== SLICE OPERATIONS ===")

	slice := []int{10, 20, 30, 40, 50}

	// Slicing syntax: slice[low:high] (low inclusive, high exclusive)
	fmt.Printf("Original: %v\n", slice)
	fmt.Printf("slice[1:3]: %v (elements at index 1, 2)\n", slice[1:3])
	fmt.Printf("slice[:3]: %v (from start to index 3)\n", slice[:3])
	fmt.Printf("slice[2:]: %v (from index 2 to end)\n", slice[2:])
	fmt.Printf("slice[:]: %v (entire slice)\n", slice[:])

	// APPEND - adds elements to a slice
	slice = append(slice, 60)
	fmt.Printf("After append(60): %v, len=%d, cap=%d\n",
		slice, len(slice), cap(slice))

	// Append multiple elements
	slice = append(slice, 70, 80, 90)
	fmt.Printf("After append(70,80,90): %v, len=%d, cap=%d\n",
		slice, len(slice), cap(slice))

	// Append another slice (note the ... operator)
	more := []int{100, 110}
	slice = append(slice, more...)
	fmt.Printf("After append(slice...): %v\n", slice)

	// COPY - copies elements between slices
	source := []int{1, 2, 3, 4, 5}
	dest := make([]int, 3)
	copied := copy(dest, source) // copies min(len(dest), len(source))
	fmt.Printf("Copied %d elements: dest=%v\n", copied, dest)

	// Delete element at index (no built-in delete for slices)
	deleteIndex := 2
	slice = append(slice[:deleteIndex], slice[deleteIndex+1:]...)
	fmt.Printf("After deleting index %d: %v\n", deleteIndex, slice)
}

// SliceCapacityAndGrowth demonstrates how slices grow
func SliceCapacityAndGrowth() {
	fmt.Println("\n=== SLICE CAPACITY & GROWTH ===")

	// Start with empty slice
	var slice []int
	fmt.Printf("Initial: len=%d, cap=%d\n", len(slice), cap(slice))

	// Watch how capacity grows as we append
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		fmt.Printf("After append(%d): len=%d, cap=%d\n", i, len(slice), cap(slice))
	}

	// Pre-allocating capacity for performance
	fmt.Println("\nPre-allocated slice:")
	optimized := make([]int, 0, 10) // length 0, capacity 10
	for i := 0; i < 10; i++ {
		optimized = append(optimized, i)
		fmt.Printf("After append(%d): len=%d, cap=%d (no reallocation!)\n",
			i, len(optimized), cap(optimized))
	}
}

// SlicePatternFilter demonstrates filtering pattern
func SlicePatternFilter() {
	fmt.Println("\n=== PATTERN: FILTERING ===")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter even numbers
	var evens []int
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Even numbers: %v\n", evens)

	// Filter in-place (modifies original slice, more efficient)
	filtered := numbers[:0] // reuse backing array
	for _, num := range numbers {
		if num > 5 {
			filtered = append(filtered, num)
		}
	}
	fmt.Printf("Numbers > 5 (in-place): %v\n", filtered)
}

// SlicePatternMap demonstrates mapping pattern
func SlicePatternMap() {
	fmt.Println("\n=== PATTERN: MAPPING ===")

	numbers := []int{1, 2, 3, 4, 5}

	// Double each number
	doubled := make([]int, len(numbers))
	for i, num := range numbers {
		doubled[i] = num * 2
	}
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Doubled: %v\n", doubled)

	// Transform to strings
	strings := make([]string, len(numbers))
	for i, num := range numbers {
		strings[i] = fmt.Sprintf("Number-%d", num)
	}
	fmt.Printf("As strings: %v\n", strings)
}

// SlicePatternReduce demonstrates reduction pattern
func SlicePatternReduce() {
	fmt.Println("\n=== PATTERN: REDUCING ===")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Sum all numbers
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Sum: %d\n", sum)

	// Find maximum
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Printf("Maximum: %d\n", max)

	// Count elements matching condition
	count := 0
	for _, num := range numbers {
		if num > 5 {
			count++
		}
	}
	fmt.Printf("Count of numbers > 5: %d\n", count)
}

// SliceGotchas demonstrates common pitfalls
func SliceGotchas() {
	fmt.Println("\n=== COMMON GOTCHAS ===")

	// Gotcha 1: Appending to a slice after slicing
	fmt.Println("\nGotcha 1: Shared backing arrays")
	original := []int{1, 2, 3, 4, 5}
	sub := original[0:2] // {1, 2}
	fmt.Printf("Original: %v, Sub: %v\n", original, sub)

	sub = append(sub, 999) // This modifies original's backing array!
	fmt.Printf("After append to sub:\n")
	fmt.Printf("Original: %v (MODIFIED!)\n", original)
	fmt.Printf("Sub: %v\n", sub)

	// Solution: Use full slice expression to limit capacity
	fmt.Println("\nSolution: Limit capacity with [low:high:max]")
	original2 := []int{1, 2, 3, 4, 5}
	sub2 := original2[0:2:2] // length 2, capacity 2
	fmt.Printf("Original: %v, Sub: %v, cap(sub)=%d\n", original2, sub2, cap(sub2))

	sub2 = append(sub2, 999) // Forces new array allocation
	fmt.Printf("After append to sub:\n")
	fmt.Printf("Original: %v (UNCHANGED)\n", original2)
	fmt.Printf("Sub: %v\n", sub2)

	// Gotcha 2: Range loop with pointer references
	fmt.Println("\nGotcha 2: Range variable reuse")
	numbers := []int{1, 2, 3}
	var pointers []*int

	for _, num := range numbers {
		// BUG: 'num' variable is reused in each iteration
		pointers = append(pointers, &num)
	}

	fmt.Printf("Values via pointers (WRONG): ")
	for _, p := range pointers {
		fmt.Printf("%d ", *p) // All point to same address!
	}
	fmt.Println()

	// Solution: Create a new variable
	var pointers2 []*int
	for _, num := range numbers {
		num := num // Create new variable in this scope
		pointers2 = append(pointers2, &num)
	}

	fmt.Printf("Values via pointers (CORRECT): ")
	for _, p := range pointers2 {
		fmt.Printf("%d ", *p)
	}
	fmt.Println()
}

// RunArraysSlices runs all arrays and slices examples
func RunArraysSlices() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("ARRAYS AND SLICES IN GO")
	fmt.Println(strings.Repeat("=", 60))

	ArrayBasics()
	SliceBasics()
	SliceOperations()
	SliceCapacityAndGrowth()
	SlicePatternFilter()
	SlicePatternMap()
	SlicePatternReduce()
	SliceGotchas()
}
