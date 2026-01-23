package main

import (
	"fmt"
	"strings"
)

// STRUCTS
// =======
// Structs are typed collections of fields
// - Similar to classes in other languages (but no inheritance)
// - Can have methods
// - Value types by default (but often used with pointers)

// Person is a basic struct with exported (public) fields
type Person struct {
	Name string
	Age  int
	City string
}

// Employee demonstrates struct embedding (composition)
type Employee struct {
	Person     // Embedded struct (anonymous field)
	EmployeeID int
	Department string
}

// Point demonstrates a simple geometric struct
type Point struct {
	X, Y int // Multiple fields of same type on one line
}

// Rectangle demonstrates struct with struct fields
type Rectangle struct {
	TopLeft     Point
	BottomRight Point
}

// StructBasics demonstrates fundamental struct concepts
func StructBasics() {
	fmt.Println("\n=== STRUCT BASICS ===")

	// Zero value initialization (all fields get zero values)
	var p1 Person
	fmt.Printf("Zero value: %+v\n", p1)

	// Struct literal with field names (recommended - clear and order-independent)
	p2 := Person{
		Name: "Alice",
		Age:  30,
		City: "NYC",
	}
	fmt.Printf("With field names: %+v\n", p2)

	// Struct literal without field names (must match order, not recommended)
	p3 := Person{"Bob", 25, "LA"}
	fmt.Printf("Without field names: %+v\n", p3)

	// Partial initialization (unspecified fields get zero values)
	p4 := Person{Name: "Charlie", Age: 35}
	fmt.Printf("Partial init: %+v\n", p4)

	// Accessing fields
	fmt.Printf("\nAccessing fields:\n")
	fmt.Printf("  Name: %s\n", p2.Name)
	fmt.Printf("  Age: %d\n", p2.Age)

	// Modifying fields
	p2.Age = 31
	fmt.Printf("After modification: %+v\n", p2)
}

// StructPointers demonstrates working with struct pointers
func StructPointers() {
	fmt.Println("\n=== STRUCT POINTERS ===")

	// Create struct
	p1 := Person{Name: "Alice", Age: 30}

	// Get pointer to struct
	p2 := &p1
	fmt.Printf("Original: %+v\n", p1)
	fmt.Printf("Pointer: %p, Value: %+v\n", p2, *p2)

	// Go automatically dereferences pointers to structs
	// These are equivalent:
	(*p2).Age = 31 // Explicit dereference
	p2.Age = 32    // Automatic dereference (preferred)
	fmt.Printf("After modification via pointer: %+v\n", p1)

	// Creating struct with new() - returns pointer to zeroed struct
	p3 := new(Person)
	fmt.Printf("Created with new(): %p, Value: %+v\n", p3, *p3)
	p3.Name = "Bob" // Can access fields directly
	fmt.Printf("After setting fields: %+v\n", *p3)

	// Common pattern: pointer to struct literal
	p4 := &Person{
		Name: "Charlie",
		Age:  35,
	}
	fmt.Printf("Pointer to literal: %+v\n", *p4)
}

// StructComparison demonstrates struct comparison
func StructComparison() {
	fmt.Println("\n=== STRUCT COMPARISON ===")

	p1 := Person{Name: "Alice", Age: 30, City: "NYC"}
	p2 := Person{Name: "Alice", Age: 30, City: "NYC"}
	p3 := Person{Name: "Bob", Age: 25, City: "LA"}

	// Structs are comparable if all fields are comparable
	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p1 == p3: %v\n", p1 == p3)

	// Struct with slice (not comparable)
	type PersonWithHobbies struct {
		Name    string
		Hobbies []string // Slices are not comparable
	}

	// ph1 := PersonWithHobbies{Name: "Alice", Hobbies: []string{"reading"}}
	// ph2 := PersonWithHobbies{Name: "Alice", Hobbies: []string{"reading"}}
	// fmt.Println(ph1 == ph2) // Compilation error!
	fmt.Println("  (Structs with slices/maps cannot be compared with ==)")
}

// StructEmbedding demonstrates struct composition
func StructEmbedding() {
	fmt.Println("\n=== STRUCT EMBEDDING (Composition) ===")

	// Create employee with embedded Person
	emp := Employee{
		Person: Person{
			Name: "Alice",
			Age:  30,
			City: "NYC",
		},
		EmployeeID: 12345,
		Department: "Engineering",
	}

	fmt.Printf("Employee: %+v\n", emp)

	// Access embedded fields directly (promoted fields)
	fmt.Printf("Name (promoted): %s\n", emp.Name)
	fmt.Printf("Age (promoted): %d\n", emp.Age)
	fmt.Printf("EmployeeID: %d\n", emp.EmployeeID)

	// Can still access via embedded field name
	fmt.Printf("Name (via Person): %s\n", emp.Person.Name)

	// Embedded fields can be used for "inheritance-like" behavior
	printPerson(emp.Person) // Can pass embedded struct to functions
}

func printPerson(p Person) {
	fmt.Printf("  Person: %s, %d years old\n", p.Name, p.Age)
}

// StructMethods demonstrates methods on structs
func StructMethods() {
	fmt.Println("\n=== STRUCT METHODS ===")

	p := Point{X: 3, Y: 4}
	fmt.Printf("Point: %+v\n", p)

	// Call value receiver method
	dist := p.Distance()
	fmt.Printf("Distance from origin: %.2f\n", dist)

	// Call pointer receiver method
	p.Scale(2)
	fmt.Printf("After scaling by 2: %+v\n", p)

	// Go automatically takes address for pointer receiver methods
	p2 := Point{X: 1, Y: 1}
	p2.Scale(3) // Automatically converted to (&p2).Scale(3)
	fmt.Printf("After scaling: %+v\n", p2)
}

// Distance calculates distance from origin (value receiver)
// Value receivers work with copies - don't modify original
func (p Point) Distance() float64 {
	return float64(p.X*p.X + p.Y*p.Y) // Simplified, not actual distance
}

// Scale scales the point (pointer receiver)
// Pointer receivers can modify the original
func (p *Point) Scale(factor int) {
	p.X *= factor
	p.Y *= factor
}

// String implements fmt.Stringer interface
func (p Point) String() string {
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

// StructTags demonstrates struct tags (used by JSON, XML, etc.)
func StructTags() {
	fmt.Println("\n=== STRUCT TAGS ===")

	// Struct with tags (commonly used with encoding packages)
	type User struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email,omitempty"` // Omit if empty
		Password  string `json:"-"`               // Never serialize
		CreatedAt string `json:"created_at"`
	}

	user := User{
		ID:       1,
		Name:     "Alice",
		Email:    "",
		Password: "secret123",
	}

	fmt.Printf("User struct: %+v\n", user)
	fmt.Println("  (Tags are metadata for packages like encoding/json)")
	fmt.Println("  - `json:\"id\"` maps field to JSON key")
	fmt.Println("  - `json:\"-\"` excludes field from JSON")
	fmt.Println("  - `json:\"email,omitempty\"` omits if zero value")
}

// StructPatternConstructor demonstrates constructor pattern
func StructPatternConstructor() {
	fmt.Println("\n=== PATTERN: CONSTRUCTOR FUNCTIONS ===")

	// Constructor function (idiomatic Go pattern)
	p1 := NewPerson("Alice", 30)
	fmt.Printf("Created with constructor: %+v\n", *p1)

	// Validation in constructor
	p2 := NewPersonValidated("", -5)
	fmt.Printf("Invalid person: %+v\n", p2)
}

// NewPerson is a constructor function (returns pointer)
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
		City: "Unknown", // Set default values
	}
}

// NewPersonValidated is a constructor with validation
func NewPersonValidated(name string, age int) *Person {
	if name == "" || age < 0 {
		return nil // Return nil for invalid input
	}
	return &Person{Name: name, Age: age}
}

// StructPatternBuilder demonstrates builder pattern
func StructPatternBuilder() {
	fmt.Println("\n=== PATTERN: BUILDER (Functional Options) ===")

	// Builder pattern for complex structs
	type Config struct {
		Host    string
		Port    int
		Timeout int
		Debug   bool
	}

	// Option function type
	type ConfigOption func(*Config)

	// Option functions
	WithHost := func(host string) ConfigOption {
		return func(c *Config) {
			c.Host = host
		}
	}

	WithPort := func(port int) ConfigOption {
		return func(c *Config) {
			c.Port = port
		}
	}

	WithDebug := func(debug bool) ConfigOption {
		return func(c *Config) {
			c.Debug = debug
		}
	}

	// Constructor with options
	NewConfig := func(opts ...ConfigOption) *Config {
		// Default values
		config := &Config{
			Host:    "localhost",
			Port:    8080,
			Timeout: 30,
			Debug:   false,
		}

		// Apply options
		for _, opt := range opts {
			opt(config)
		}

		return config
	}

	// Create with custom options
	cfg := NewConfig(
		WithHost("example.com"),
		WithPort(9090),
		WithDebug(true),
	)

	fmt.Printf("Config: %+v\n", *cfg)
}

// StructPatternAnonymous demonstrates anonymous structs
func StructPatternAnonymous() {
	fmt.Println("\n=== PATTERN: ANONYMOUS STRUCTS ===")

	// Anonymous struct (no type definition needed)
	point := struct {
		X int
		Y int
	}{
		X: 10,
		Y: 20,
	}

	fmt.Printf("Anonymous struct: %+v\n", point)

	// Useful for one-off data structures
	config := struct {
		Enabled bool
		Count   int
	}{true, 42}

	fmt.Printf("Config: %+v\n", config)

	// Common in table-driven tests
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"zero", 0, 0},
		{"positive", 5, 25},
		{"negative", -3, 9},
	}

	fmt.Println("\nTable-driven test structure:")
	for _, tt := range tests {
		result := tt.input * tt.input
		fmt.Printf("  %s: %d^2 = %d (expected %d) ✓\n",
			tt.name, tt.input, result, tt.expected)
	}
}

// StructGotchas demonstrates common pitfalls
func StructGotchas() {
	fmt.Println("\n=== COMMON GOTCHAS ===")

	// Gotcha 1: Structs are value types
	fmt.Println("\nGotcha 1: Structs are copied by value")
	p1 := Person{Name: "Alice", Age: 30}
	p2 := p1
	p2.Age = 31
	fmt.Printf("Original: %+v\n", p1)
	fmt.Printf("Copy: %+v (independent)\n", p2)

	// Use pointers to share
	p3 := &p1
	p3.Age = 32
	fmt.Printf("After pointer modification: %+v\n", p1)

	// Gotcha 2: Comparing structs with slices
	fmt.Println("\nGotcha 2: Structs with uncomparable fields")
	fmt.Println("  Cannot use == on structs containing slices/maps/functions")

	// Gotcha 3: Method receivers
	fmt.Println("\nGotcha 3: Value vs Pointer receivers")
	fmt.Println("  Value receivers: Work on copies, can't modify original")
	fmt.Println("  Pointer receivers: Can modify original, more efficient for large structs")

	point1 := Point{X: 1, Y: 1}
	point1.Distance() // Value receiver - works on copy
	point1.Scale(2)   // Pointer receiver - modifies original

	// Gotcha 4: Zero values
	fmt.Println("\nGotcha 4: Zero values can be problematic")
	var p4 Person // All fields are zero values
	fmt.Printf("Zero Person: %+v\n", p4)
	fmt.Println("  Empty strings and 0 might not be valid business values")
	fmt.Println("  Use constructor functions for validation and defaults")
}

// RunStructs runs all struct examples
func RunStructs() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("STRUCTS IN GO")
	fmt.Println(strings.Repeat("=", 60))

	StructBasics()
	StructPointers()
	StructComparison()
	StructEmbedding()
	StructMethods()
	StructTags()
	StructPatternConstructor()
	StructPatternBuilder()
	StructPatternAnonymous()
	StructGotchas()
}
