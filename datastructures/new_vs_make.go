package main

import (
	"fmt"
	"strings"
)

// NEW vs MAKE
// ===========
// Go has two memory allocation primitives: new() and make()
// They serve different purposes and are used for different types

// QUICK REFERENCE:
// - new(T):  Allocates memory, returns *T (pointer), zeroed memory, works with ANY type
// - make(T): Initializes and returns T (not pointer), only for slices, maps, channels

// NewBasics demonstrates new() function
func NewBasics() {
	fmt.Println("\n=== new() FUNCTION ===")

	// new() allocates memory and returns a pointer
	// The memory is zeroed (set to zero value of the type)

	// new() with basic types
	intPtr := new(int)
	fmt.Printf("new(int): %v, value: %d, type: %T\n", intPtr, *intPtr, intPtr)

	*intPtr = 42
	fmt.Printf("After assignment: %d\n", *intPtr)


	// new() with strings
	strPtr := new(string)
	fmt.Printf("new(string): %v, value: %q, type: %T\n", strPtr, *strPtr, strPtr)

	// new() with structs
	type Person struct {
		Name string
		Age  int
	}

	personPtr := new(Person)
	fmt.Printf("new(Person): %v, value: %+v, type: %T\n", personPtr, *personPtr, personPtr)

	// Can access fields directly (automatic dereferencing) (can't access through pointer)
	personPtr.Name = "Alice"
	personPtr.Age = 30
	fmt.Printf("After setting fields: %+v\n", *personPtr)

	// new() with slices - creates pointer to nil slice
	slicePtr := new([]int)
	fmt.Printf("new([]int): %v, value: %v, is nil? %v\n",
		slicePtr, *slicePtr, *slicePtr == nil)
}

// MakeBasics demonstrates make() function
func MakeBasics() {
	fmt.Println("\n=== make() FUNCTION ===")

	// make() ONLY works with slices, maps, and channels
	// It initializes and returns the type itself (not a pointer)

	// make() with slices
	fmt.Println("\nSlices:")
	slice1 := make([]int, 5)     // length 5, capacity 5
	slice2 := make([]int, 3, 10) // length 3, capacity 10

	fmt.Printf("make([]int, 5): %v, len=%d, cap=%d, type: %T\n",
		slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("make([]int, 3, 10): %v, len=%d, cap=%d\n",
		slice2, len(slice2), cap(slice2))

	// make() with maps
	fmt.Println("\nMaps:")
	map1 := make(map[string]int)
	fmt.Printf("make(map[string]int): %v, len=%d, type: %T\n",
		map1, len(map1), map1)

	map2 := make(map[string]int, 100) // with capacity hint
	fmt.Printf("make(map[string]int, 100): %v, len=%d\n",
		map2, len(map2))

	// Can immediately use maps created with make()
	map1["key"] = 42
	fmt.Printf("After insertion: %v\n", map1)
}

// NewVsMakeComparison directly compares new() and make()
func NewVsMakeComparison() {
	fmt.Println("\n=== new() vs make() COMPARISON ===")

	fmt.Println("\n1. WITH SLICES:")

	// Using new() with slice - creates pointer to nil slice
	sliceNew := new([]int)
	fmt.Printf("new([]int):\n")
	fmt.Printf("  Type: %T (pointer to slice)\n", sliceNew)
	fmt.Printf("  Value: %v (pointer to nil slice)\n", sliceNew)
	fmt.Printf("  Dereferenced: %v, is nil? %v\n", *sliceNew, *sliceNew == nil)
	// Can't append to nil slice without dereferencing
	*sliceNew = append(*sliceNew, 1, 2, 3)
	fmt.Printf("  After append: %v\n", *sliceNew)

	// Using make() with slice - creates ready-to-use slice
	sliceMake := make([]int, 3, 5)
	fmt.Printf("\nmake([]int, 3, 5):\n")
	fmt.Printf("  Type: %T (slice, not pointer)\n", sliceMake)
	fmt.Printf("  Value: %v\n", sliceMake)
	fmt.Printf("  Length: %d, Capacity: %d\n", len(sliceMake), cap(sliceMake))
	// Can use immediately
	sliceMake[0] = 10
	sliceMake = append(sliceMake, 20)
	fmt.Printf("  After operations: %v\n", sliceMake)

	fmt.Println("\n2. WITH MAPS:")

	// Using new() with map - creates pointer to nil map
	mapNew := new(map[string]int)
	fmt.Printf("new(map[string]int):\n")
	fmt.Printf("  Type: %T (pointer to map)\n", mapNew)
	fmt.Printf("  Value: %v (pointer to nil map)\n", mapNew)
	fmt.Printf("  Dereferenced: %v, is nil? %v\n", *mapNew, *mapNew == nil)
	// Cannot assign to nil map - must initialize first
	*mapNew = make(map[string]int)
	(*mapNew)["key"] = 42
	fmt.Printf("  After make and assign: %v\n", *mapNew)

	// Using make() with map - creates ready-to-use map
	mapMake := make(map[string]int)
	fmt.Printf("\nmake(map[string]int):\n")
	fmt.Printf("  Type: %T (map, not pointer)\n", mapMake)
	fmt.Printf("  Value: %v\n", mapMake)
	// Can use immediately
	mapMake["key"] = 42
	fmt.Printf("  After insert: %v\n", mapMake)

	fmt.Println("\n3. WITH STRUCTS:")

	type Point struct{ X, Y int }

	// Using new() with struct - creates pointer to zeroed struct
	pointNew := new(Point)
	fmt.Printf("new(Point):\n")
	fmt.Printf("  Type: %T (pointer to struct)\n", pointNew)
	fmt.Printf("  Value: %+v\n", *pointNew)
	pointNew.X = 10
	fmt.Printf("  After modification: %+v\n", *pointNew)

	// make() DOES NOT work with structs
	// pointMake := make(Point) // Compilation error!
	fmt.Printf("\nmake(Point): NOT ALLOWED (compilation error)\n")
	fmt.Printf("  make() only works with slices, maps, and channels\n")

	// For structs, use literals or new()
	pointLiteral := Point{X: 5, Y: 10}
	pointLiteralPtr := &Point{X: 5, Y: 10}
	fmt.Printf("\nStruct literal: %+v (type: %T)\n", pointLiteral, pointLiteral)
	fmt.Printf("Pointer to literal: %+v (type: %T)\n", *pointLiteralPtr, pointLiteralPtr)
}

// WhenToUseWhat provides guidance on when to use new() vs make()
func WhenToUseWhat() {
	fmt.Println("\n=== WHEN TO USE WHAT ===")

	fmt.Println("\nUse make() for:")
	fmt.Println("  ✓ Slices   - make([]T, len, cap)")
	fmt.Println("  ✓ Maps     - make(map[K]V)")
	fmt.Println("  ✓ Channels - make(chan T)")
	fmt.Println("  → Returns initialized, ready-to-use value")

	fmt.Println("\nUse new() for:")
	fmt.Println("  ✓ Any type when you need a pointer to zero value")
	fmt.Println("  ✓ Rarely used in practice")
	fmt.Println("  → Returns pointer to zeroed memory")

	fmt.Println("\nIn practice:")
	fmt.Println("  → Slices: Use make() or literal []T{}")
	fmt.Println("  → Maps: Use make() or literal map[K]V{}")
	fmt.Println("  → Structs: Use literal T{} or &T{}")
	fmt.Println("  → new() is rarely needed")
}

// PracticalExamples shows idiomatic usage patterns
func PracticalExamples() {
	fmt.Println("\n=== PRACTICAL EXAMPLES ===")

	// Example 1: Creating slices
	fmt.Println("\nCreating slices (IDIOMATIC):")

	// Empty slice - use literal or make
	var empty1 []int         // nil slice
	empty2 := []int{}        // empty slice literal
	empty3 := make([]int, 0) // empty slice with make
	fmt.Printf("  nil slice: %v\n", empty1)
	fmt.Printf("  empty literal: %v\n", empty2)
	fmt.Printf("  make empty: %v\n", empty3)

	// Slice with known size
	sized := make([]int, 10) // preferred
	fmt.Printf("  sized slice: len=%d, cap=%d\n", len(sized), cap(sized))

	// Slice with initial values
	initialized := []int{1, 2, 3, 4, 5}
	fmt.Printf("  initialized: %v\n", initialized)

	// Example 2: Creating maps
	fmt.Println("\nCreating maps (IDIOMATIC):")

	// Empty map
	var nilMap map[string]int        // nil map (can't use!)
	emptyMap := make(map[string]int) // preferred
	fmt.Printf("  nil map: %v (can't insert!)\n", nilMap)
	fmt.Printf("  empty map: %v (ready to use)\n", emptyMap)

	// Map with initial values
	initialized2 := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Printf("  initialized: %v\n", initialized2)

	// Example 3: Creating structs
	fmt.Println("\nCreating structs (IDIOMATIC):")

	type Config struct {
		Host string
		Port int
	}

	// Various ways to create structs
	var zero Config // zero value
	literal := Config{Host: "localhost", Port: 8080}
	literalPtr := &Config{Host: "example.com", Port: 443}
	withNew := new(Config) // rarely used

	fmt.Printf("  zero value: %+v\n", zero)
	fmt.Printf("  literal: %+v\n", literal)
	fmt.Printf("  pointer to literal: %+v\n", *literalPtr)
	fmt.Printf("  with new: %+v\n", *withNew)
}

// MemoryAllocationDetails shows what happens behind the scenes
func MemoryAllocationDetails() {
	fmt.Println("\n=== MEMORY ALLOCATION DETAILS ===")

	// new() allocates zeroed memory
	intPtr := new(int)
	fmt.Printf("new(int): address=%p, value=%d\n", intPtr, *intPtr)

	// Equivalent to:
	var i int
	intPtr2 := &i
	fmt.Printf("var + &:  address=%p, value=%d\n", intPtr2, *intPtr2)

	// make() for slices allocates backing array
	slice := make([]int, 3, 5)
	fmt.Printf("\nmake([]int, 3, 5):\n")
	fmt.Printf("  Slice header: len=%d, cap=%d\n", len(slice), cap(slice))
	fmt.Printf("  Backing array allocated with capacity 5\n")
	fmt.Printf("  Elements initialized to zero: %v\n", slice)

	// make() for maps allocates hash table
	m := make(map[string]int, 100)
	fmt.Printf("\nmake(map[string]int, 100):\n")
	fmt.Printf("  Hash table allocated with space for ~100 elements\n")
	fmt.Printf("  Ready for immediate use: %v\n", m)
}

// CommonMistakes shows common errors when using new() and make()
func CommonMistakes() {
	fmt.Println("\n=== COMMON MISTAKES ===")

	// Mistake 1: Using new() with maps/slices expecting them to work
	fmt.Println("\nMistake 1: Trying to use new() with maps")
	mapPtr := new(map[string]int)
	fmt.Printf("  new(map[string]int) creates: %T\n", mapPtr)
	fmt.Printf("  Value: %v (pointer to nil map)\n", mapPtr)
	// (*mapPtr)["key"] = 1 // PANIC! Assignment to nil map
	fmt.Println("  ❌ Can't insert - map is nil!")

	// Solution: Use make()
	goodMap := make(map[string]int)
	goodMap["key"] = 1
	fmt.Printf("  ✓ make(map[string]int) works: %v\n", goodMap)

	// Mistake 2: Using make() with structs
	fmt.Println("\nMistake 2: Trying to use make() with structs")
	// type Point struct{ X, Y int }
	// p := make(Point) // COMPILATION ERROR!
	fmt.Println("  ❌ make(Point) doesn't compile")
	fmt.Println("  ✓ Use Point{} or new(Point) instead")

	// Mistake 3: Confusing return types
	fmt.Println("\nMistake 3: Forgetting new() returns pointer")
	intPtr := new(int)
	// var x int = intPtr // Type error! intPtr is *int, not int
	var x int = *intPtr // Must dereference
	fmt.Printf("  new(int) returns: %T (need to dereference)\n", intPtr)
	fmt.Printf("  Dereferenced value: %d (type: %T)\n", x, x)
}

// RunNewVsMake runs all new vs make examples
func RunNewVsMake() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("new() vs make() IN GO")
	fmt.Println(strings.Repeat("=", 60))

	NewBasics()
	MakeBasics()
	NewVsMakeComparison()
	WhenToUseWhat()
	PracticalExamples()
	MemoryAllocationDetails()
	CommonMistakes()
}
