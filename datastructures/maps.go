package main

import (
	"fmt"
	"strings"
)

// MAPS (Hash Tables)
// ==================
// Maps are key-value pairs, similar to dictionaries/hash tables
// - Keys must be comparable types (no slices, maps, or functions)
// - Values can be any type
// - Maps are REFERENCE types (like slices)
// - Unordered collection

// MapBasics demonstrates fundamental map concepts
func MapBasics() {
	fmt.Println("\n=== MAP BASICS ===")

	// Nil map (cannot add elements to it!)
	var map1 map[string]int
	fmt.Printf("Nil map: %v, len=%d, is nil? %v\n",
		map1, len(map1), map1 == nil)

	// Map literal initialization
	map2 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 7,
	}
	fmt.Printf("Map literal: %v\n", map2)

	// Using make() - PROPER WAY to create empty maps
	map3 := make(map[string]int)
	fmt.Printf("Empty map with make(): %v, len=%d\n", map3, len(map3))

	// Can specify initial capacity hint (optimization)
	map4 := make(map[string]int, 100)
	fmt.Printf("Map with capacity hint: %v, len=%d\n", map4, len(map4))
}

// MapOperations demonstrates common map operations
func MapOperations() {
	fmt.Println("\n=== MAP OPERATIONS ===")

	// Create a map
	scores := make(map[string]int)

	// INSERT/UPDATE - same syntax
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92
	fmt.Printf("After inserts: %v\n", scores)

	// Update existing key
	scores["Alice"] = 98
	fmt.Printf("After update: %v\n", scores)

	// READ - access by key
	aliceScore := scores["Alice"]
	fmt.Printf("Alice's score: %d\n", aliceScore)

	// Reading non-existent key returns zero value
	noScore := scores["David"]
	fmt.Printf("Non-existent key returns: %d (zero value)\n", noScore)

	// CHECK EXISTENCE - the "comma ok" idiom
	if score, exists := scores["Alice"]; exists {
		fmt.Printf("Alice exists with score: %d\n", score)
	}

	if score, exists := scores["David"]; !exists {
		fmt.Printf("David doesn't exist, got zero value: %d\n", score)
	}

	// DELETE - using delete() built-in function
	delete(scores, "Bob")
	fmt.Printf("After deleting Bob: %v\n", scores)

	// Deleting non-existent key is safe (no-op)
	delete(scores, "NonExistent")
	fmt.Printf("After deleting non-existent key: %v\n", scores)

	// LENGTH
	fmt.Printf("Number of entries: %d\n", len(scores))
}

// MapIteration demonstrates how to iterate over maps
func MapIteration() {
	fmt.Println("\n=== MAP ITERATION ===")

	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
		"Diana":   28,
	}

	// Iterate over key-value pairs
	fmt.Println("Iterate over key-value pairs:")
	for name, age := range ages {
		fmt.Printf("  %s is %d years old\n", name, age)
	}

	// Iterate over keys only
	fmt.Println("\nIterate over keys only:")
	for name := range ages {
		fmt.Printf("  %s\n", name)
	}

	// Note: Map iteration order is RANDOM and not guaranteed
	fmt.Println("\nIteration order is random (run multiple times):")
	for i := 0; i < 3; i++ {
		fmt.Printf("  Run %d: ", i+1)
		for name := range ages {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}
}

// MapWithComplexTypes demonstrates maps with various key/value types
func MapWithComplexTypes() {
	fmt.Println("\n=== MAPS WITH COMPLEX TYPES ===")

	// Map with struct values
	type Person struct {
		Age  int
		City string
	}

	people := map[string]Person{
		"Alice": {Age: 30, City: "NYC"},
		"Bob":   {Age: 25, City: "LA"},
	}
	fmt.Printf("Map with struct values: %v\n", people)

	// Map with slice values
	grades := map[string][]int{
		"Alice": {95, 87, 92},
		"Bob":   {88, 91, 85},
	}
	fmt.Printf("Map with slice values: %v\n", grades)

	// Map with map values (nested maps)
	matrix := map[string]map[string]int{
		"row1": {"col1": 1, "col2": 2},
		"row2": {"col1": 3, "col2": 4},
	}
	fmt.Printf("Nested map: %v\n", matrix)

	// Accessing nested map
	if row, exists := matrix["row1"]; exists {
		if val, exists := row["col2"]; exists {
			fmt.Printf("matrix[row1][col2] = %d\n", val)
		}
	}

	// Map with integer keys
	counts := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Printf("Map with int keys: %v\n", counts)
}

// MapPatternGrouping demonstrates grouping pattern
func MapPatternGrouping() {
	fmt.Println("\n=== PATTERN: GROUPING ===")

	// Group words by first letter
	words := []string{"apple", "apricot", "banana", "blueberry", "cherry", "coconut"}

	grouped := make(map[rune][]string)
	for _, word := range words {
		firstLetter := rune(word[0])
		grouped[firstLetter] = append(grouped[firstLetter], word)
	}

	fmt.Println("Words grouped by first letter:")
	for letter, wordList := range grouped {
		fmt.Printf("  %c: %v\n", letter, wordList)
	}
}

// MapPatternCounting demonstrates counting pattern
func MapPatternCounting() {
	fmt.Println("\n=== PATTERN: COUNTING ===")

	// Count word occurrences
	text := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}

	counts := make(map[string]int)
	for _, word := range text {
		counts[word]++
	}

	fmt.Println("Word counts:")
	for word, count := range counts {
		fmt.Printf("  %s: %d\n", word, count)
	}
}

// MapPatternSet demonstrates set implementation using maps
func MapPatternSet() {
	fmt.Println("\n=== PATTERN: SET (Using Maps) ===")

	// Go doesn't have a built-in set type
	// Use map[T]bool or map[T]struct{} for sets

	// Using map[string]bool
	set1 := make(map[string]bool)
	set1["apple"] = true
	set1["banana"] = true
	set1["apple"] = true // Duplicate, no effect

	fmt.Printf("Set (using map[string]bool): %v\n", set1)

	// Check membership
	if set1["apple"] {
		fmt.Println("  'apple' is in the set")
	}

	// Using map[string]struct{} (more memory efficient)
	set2 := make(map[string]struct{})
	set2["apple"] = struct{}{}
	set2["banana"] = struct{}{}

	fmt.Println("\nSet (using map[string]struct{}):")
	for item := range set2 {
		fmt.Printf("  - %s\n", item)
	}

	// Set operations
	setA := map[string]bool{"a": true, "b": true, "c": true}
	setB := map[string]bool{"b": true, "c": true, "d": true}

	// Union
	union := make(map[string]bool)
	for k := range setA {
		union[k] = true
	}
	for k := range setB {
		union[k] = true
	}
	fmt.Printf("\nUnion of {a,b,c} and {b,c,d}: %v\n", union)

	// Intersection
	intersection := make(map[string]bool)
	for k := range setA {
		if setB[k] {
			intersection[k] = true
		}
	}
	fmt.Printf("Intersection: %v\n", intersection)
}

// MapPatternCache demonstrates caching pattern
func MapPatternCache() {
	fmt.Println("\n=== PATTERN: CACHING/MEMOIZATION ===")

	// Cache expensive computation results
	cache := make(map[int]int)

	// Fibonacci with memoization
	var fib func(int) int
	fib = func(n int) int {
		// Check cache first
		if result, exists := cache[n]; exists {
			fmt.Printf("  Cache hit for fib(%d)\n", n)
			return result
		}

		// Base cases
		if n <= 1 {
			return n
		}

		// Compute and cache
		result := fib(n-1) + fib(n-2)
		cache[n] = result
		return result
	}

	fmt.Println("Computing fib(10) with caching:")
	result := fib(10)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Cache contents: %v\n", cache)
}

// MapGotchas demonstrates common pitfalls
func MapGotchas() {
	fmt.Println("\n=== COMMON GOTCHAS ===")

	// Gotcha 1: Nil map panic
	fmt.Println("\nGotcha 1: Cannot assign to nil map")
	var nilMap map[string]int
	fmt.Printf("nilMap is nil: %v\n", nilMap == nil)
	// nilMap["key"] = 1 // This would panic!
	fmt.Println("  (Attempting to assign would cause panic)")

	// Solution: Initialize with make()
	nilMap = make(map[string]int)
	nilMap["key"] = 1
	fmt.Printf("After make(): %v\n", nilMap)

	// Gotcha 2: Maps are not safe for concurrent access
	fmt.Println("\nGotcha 2: Maps are not concurrent-safe")
	fmt.Println("  (Need sync.Mutex or sync.Map for concurrent access)")

	// Gotcha 3: Can't take address of map element
	fmt.Println("\nGotcha 3: Cannot take address of map elements")
	type Point struct{ X, Y int }
	points := map[string]Point{"origin": {0, 0}}

	// points["origin"].X = 10 // Compilation error!
	// &points["origin"] // Compilation error!

	// Solution: Read, modify, write back
	p := points["origin"]
	p.X = 10
	points["origin"] = p
	fmt.Printf("Modified point: %v\n", points)

	// Or use pointers as values
	pointPtrs := map[string]*Point{"origin": {0, 0}}
	pointPtrs["origin"].X = 10 // This works!
	fmt.Printf("With pointer values: %v\n", pointPtrs)

	// Gotcha 4: Checking existence
	fmt.Println("\nGotcha 4: Zero values vs non-existent keys")
	scores := map[string]int{"Alice": 0}

	aliceScore := scores["Alice"] // 0 (exists, value is 0)
	bobScore := scores["Bob"]     // 0 (doesn't exist, zero value)

	fmt.Printf("Alice: %d, Bob: %d (both are 0!)\n", aliceScore, bobScore)

	// Must use comma-ok idiom to distinguish
	if _, exists := scores["Alice"]; exists {
		fmt.Println("  Alice exists")
	}
	if _, exists := scores["Bob"]; !exists {
		fmt.Println("  Bob doesn't exist")
	}
}

// RunMaps runs all map examples
func RunMaps() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("MAPS IN GO")
	fmt.Println(strings.Repeat("=", 60))

	MapBasics()
	MapOperations()
	MapIteration()
	MapWithComplexTypes()
	MapPatternGrouping()
	MapPatternCounting()
	MapPatternSet()
	MapPatternCache()
	MapGotchas()
}
