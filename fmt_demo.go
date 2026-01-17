package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Example types to demonstrate fmt interfaces
type Person struct {
	Name string
	Age  int
}

// String method implements fmt.Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("Nameee:%s (ageeeee %d)", p.Name, p.Age)
}

// Custom type implementing fmt.Formatter
type Temperature float64

func (t Temperature) Format(f fmt.State, verb rune) {
	// Custom formatting for temperature
	switch verb {
	case 'f':
		fmt.Fprintf(f, "%.1f°C", float64(t))
	case 'v':
		if f.Flag('#') {
			fmt.Fprintf(f, "Temperature(%v°C)", float64(t))
		} else {
			fmt.Fprintf(f, "%.1f°C", float64(t))
		}
	default:
		// Fall back to default behavior
		fmt.Fprintf(f, "%v", float64(t))
	}
}

func main() {
	fmt.Println("=== Go fmt Package Deep Dive ====\n")

	// 1. Basic printing functions
	fmt.Println("1. BASIC PRINTING FUNCTIONS:")
	fmt.Print("Hello, ")
	fmt.Print("Go!\n")
	fmt.Println("Hello, Go!")
	name := "Alice"
	age := 30
	active := true
	fmt.Printf("Name: %s, Age: %d, Active: %t\n", name, age, active)

	// 2. Format verbs demonstration
	fmt.Println("\n2. FORMAT VERBS:")
	
	// General verbs
	value := 42
	fmt.Printf("%%v (default): %v\n", value)
	fmt.Printf("%%+v (with field names): %+v\n", Person{Name: "Bob", Age: 25})
	fmt.Printf("%%#v (Go syntax): %#v\n", value)
	fmt.Printf("%%T (type): %T\n", value)
	
	// Boolean
	flag := true
	fmt.Printf("%%t (boolean): %t\n", flag)
	
	// Integer formatting
	num := 255
	fmt.Printf("%%d (decimal): %d\n", num)
	fmt.Printf("%%b (binary): %b\n", num)
	fmt.Printf("%%o (octal): %o\n", num)
	fmt.Printf("%%x (hex): %x\n", num)
	fmt.Printf("%%X (HEX): %X\n", num)
	fmt.Printf("%%#x (hex with 0x): %#x\n", num)
	fmt.Printf("%%U (Unicode): %U\n", num)
	
	// Floating point
	pi := 3.14159
	fmt.Printf("%%f (decimal): %f\n", pi)
	fmt.Printf("%%.2f (precision): %.2f\n", pi)
	fmt.Printf("%%e (scientific): %e\n", pi)
	fmt.Printf("%%g (general): %g\n", pi)
	
	// String formatting
	text := "Hello \"World\""
	fmt.Printf("%%s (string): %s\n", text)
	fmt.Printf("%%q (quoted): %q\n", text)
	fmt.Printf("%%x (hex bytes): %x\n", text)
	
	// Pointer
	slice := []int{1, 2, 3}
	fmt.Printf("%%p (pointer): %p\n", slice)

	// 3. Width and precision
	fmt.Println("\n3. WIDTH AND PRECISION:")
	fmt.Printf("|%10s|%10d|%10.2f|\n", "hello", 42, 3.14159)
	fmt.Printf("|%-10s|%-10d|%-10.2f|\n", "hello", 42, 3.14159) // left-justified
	fmt.Printf("%08d\n", 42) // zero padding
	fmt.Printf("%8.2f\n", 3.1) // width with precision

	// 4. Flags
	fmt.Println("\n4. FLAGS:")
	fmt.Printf("Default: %d\n", 42)
	fmt.Printf("Plus flag: %+d\n", 42)
	fmt.Printf("Space flag: % d\n", 42)
	fmt.Printf("Sharp flag: %#o, %#x\n", 42, 42)
	fmt.Printf("Minus flag: %-10d|\n", 42)
	fmt.Printf("Zero flag: %010d\n", 42)

	// 5. Interface implementations
	fmt.Println("\n5. INTERFACE IMPLEMENTATIONS:")
	
	// Stringer interface
	person := Person{Name: "Charlie", Age: 35}
	fmt.Printf("Stringer: %v\n", person)
	fmt.Printf("Stringer with %%s: %s\n", person)
	
	// Formatter interface
	temp := Temperature(23.5)
	fmt.Printf("Formatter (default): %v\n", temp)
	fmt.Printf("Formatter (float): %f\n", temp)
	fmt.Printf("Formatter (sharp): %#v\n", temp)

	// 6. Error handling in formatting
	fmt.Println("\n6. ERROR HANDLING:")
	fmt.Printf("Wrong type: %d\n", "hello") // type mismatch
	fmt.Printf("Missing arg: %d %s\n", 42) // missing argument
	fmt.Printf("Extra arg: %d\n", 42, "extra") // extra argument

	// 7. Complex types
	fmt.Println("\n7. COMPLEX TYPES:")
	
	// Slices and arrays
	fmt.Printf("Slice: %v\n", []int{1, 2, 3, 4})
	fmt.Printf("Slice with %%+v: %+v\n", []Person{{Name: "A", Age: 1}, {Name: "B", Age: 2}})
	
	// Maps
	m := map[string]int{"apple": 5, "banana": 3}
	fmt.Printf("Map: %v\n", m)
	fmt.Printf("Map with %%+v: %+v\n", m)
	
	// Channels
	ch := make(chan int)
	fmt.Printf("Channel: %v\n", ch)
	
	// Functions
	funcVar := func(x int) int { return x * 2 }
	fmt.Printf("Function: %v\n", funcVar)

	// 8. Scanning functions
	fmt.Println("\n8. SCANNING FUNCTIONS:")
	
	// Simulate input for scanning
	input := "Alice 25 3.14"
	var name2 string
	var age2 int
	var score float64
	n, err := fmt.Sscanf(input, "%s %d %f", &name2, &age2, &score)
	fmt.Printf("Scanned: name=%s, age=%d, score=%.2f (n=%d, err=%v)\n", 
		name2, age2, score, n, err)

	// 9. Using with io.Writer
	fmt.Println("\n9. CUSTOM WRITERS:")
	
	// Write to string buffer
	var buf strings.Builder
	n, err = fmt.Fprintf(&buf, "Formatted output: %s scored %.1f points", "Bob", 87.5)
	fmt.Printf("Wrote %d bytes to buffer: %s\n", n, buf.String())
	
	// 10. Advanced formatting tricks
	fmt.Println("\n10. ADVANCED TRICKS:")
	
	// Argument indexing
	fmt.Printf("Reordering: %[2]d %[1]s %[3]f\n", "hello", 42, 3.14)
	
	// Using * for width and precision from arguments
	width, precision := 8, 3
	value2 := 3.14159
	fmt.Printf("Dynamic width/precision: %*.*f\n", width, precision, value2)
	
	// Format string reconstruction
	fmt.Printf("Format string reconstruction: %s\n", 
	fmt.Sprintf("%[2]*.[1]*[3]f", precision, width, value2))

	// 11. Performance considerations
	fmt.Println("\n11. PERFORMANCE:")
	
	// Reusing format strings vs building strings
	simple := fmt.Sprintf("Name: %s, Age: %d", "David", 40)
	fmt.Printf("Sprintf result: %s\n", simple)
	
	// For simple concatenation, + operator is often faster
	fast := "Name: " + "David" + ", Age: " + fmt.Sprintf("%d", 40)
	fmt.Printf("Concatenation result: %s\n", fast)

	// 12. Type reflection with fmt
	fmt.Println("\n12. TYPE REFLECTION:")
	
	var x interface{} = 42
	fmt.Printf("Type of x: %T\n", x)
	fmt.Printf("Value of x: %v\n", x)
	
	x = "hello"
	fmt.Printf("Type of x: %T\n", x)
	fmt.Printf("Value of x: %v\n", x)
	
	x = []int{1, 2, 3}
	fmt.Printf("Type of x: %T\n", x)
	fmt.Printf("Value of x: %v\n", x)

	// Demonstrate reflect.Value formatting
	val := reflect.ValueOf([]int{10, 20, 30})
	fmt.Printf("reflect.Value: %v\n", val)

	fmt.Println("\n=== End of fmt Package Demo ====")
}