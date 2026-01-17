# Go fmt Package - Comprehensive Analysis

## Overview
The `fmt` package implements formatted I/O functions analogous to C's printf and scanf family. It's one of Go's most fundamental packages, providing essential formatting capabilities for all Go programs.

## Core Architecture

### 1. Main Components

**Key Interfaces:**
- `State`: Represents printer state passed to custom formatters
- `Formatter`: Implemented by types with custom formatting logic
- `Stringer`: Implemented by types with custom string representation
- `GoStringer`: Implemented by types with Go syntax representation
- `Scanner`: Implemented by types that can scan themselves from input

**Core Types:**
- `pp`: Main printer struct (reused via sync.Pool for performance)
- `fmt`: Low-level formatter with flag handling
- `buffer`: Efficient byte buffer implementation

### 2. Key Design Principles

1. **Performance-Oriented**: Uses sync.Pool to reuse printer instances
2. **Interface-Driven**: Leverages Go's interface system for extensibility
3. **Reflection-Based**: Uses reflection for complex types and method detection
4. **Safety-First**: Handles panics, nil values, and type mismatches gracefully

## Format Verbs Deep Dive

### General Verbs
- `%v`: Default format - the most commonly used verb
- `%+v`: Struct field names included
- `%#v`: Go syntax representation
- `%T`: Type information
- `%%`: Literal percent sign

### Type-Specific Verbs

**Boolean:**
- `%t`: "true" or "false"

**Integer:**
- `%d`: Decimal
- `%b`: Binary
- `%o`: Octal
- `%O`: Octal with 0o prefix
- `%x`, `%X`: Hexadecimal (lower/upper case)
- `%c`: Character from Unicode code point
- `%U`: Unicode format (U+1234)
- `%q`: Single-quoted character

**Floating Point:**
- `%f`: Decimal notation
- `%e`, `%E`: Scientific notation
- `%g`, `%G`: General format (switches between %f and %e)
- `%b`: Binary scientific notation
- `%x`, `%X`: Hexadecimal notation with exponent

**String/Slice:**
- `%s`: Direct string
- `%q`: Quoted string with escaping
- `%x`, `%X`: Hexadecimal encoding of bytes

### Flags System

**5 Key Flags:**
1. `+`: Always show sign for numbers
2. `-`: Left-justify (default is right)
3. `#`: Alternate format (0x for hex, 0 for octal, etc.)
4. ` ` (space)`: Leave space for elided sign
5. `0`: Pad with zeros instead of spaces

## Implementation Details

### 1. Printer Pool Pattern
```go
var ppFree = sync.Pool{
    New: func() any { return new(pp) },
}
```
This avoids allocations for frequent formatting operations.

### 2. Method Resolution Order
1. Check for `Formatter` interface first
2. Check for `GoStringer` (with %#v)
3. Check for `error` interface
4. Check for `Stringer` interface
5. Use reflection for other types

### 3. Buffer Management
The package uses a custom `buffer` type instead of `bytes.Buffer` to avoid large dependencies and optimize for the specific use case.

### 4. Unicode Support
- Proper UTF-8 handling throughout
- Rune counting for width calculations
- Unicode replacement character (U+FFFD) for invalid sequences

## Advanced Features

### 1. Argument Indexing
```go
fmt.Printf("%[2]d %[1]s", 11, 22) // "22 11"
```

### 2. Dynamic Width/Precision
```go
fmt.Sprintf("%*.*f", width, precision, value)
```

### 3. Error Handling in Formatting
The package gracefully handles:
- Type mismatches: `%!d(string=hello)`
- Missing arguments: `%!d(MISSING)`
- Extra arguments: `%!(EXTRA string=extra)`
- Panics in custom methods: `%!s(PANIC=error message)`

## Performance Considerations

### 1. Fast Paths
- Direct type assertions for common types (int, string, etc.)
- Avoid reflection when possible
- Pre-allocated buffers for common cases

### 2. Memory Efficiency
- Buffer reuse via sync.Pool
- Stack allocation for small buffers
- Careful buffer growth management

### 3. Best Practices
- Use `fmt.Sprintf` for complex formatting
- Use string concatenation (`+`) for simple cases
- Consider `strings.Builder` for building strings incrementally

## Interface Implementation Examples

### Custom Formatter
```go
type Temperature float64

func (t Temperature) Format(f fmt.State, verb rune) {
    switch verb {
    case 'f':
        fmt.Fprintf(f, "%.1f°C", float64(t))
    case 'v':
        fmt.Fprintf(f, "%.1f°C", float64(t))
    default:
        fmt.Fprintf(f, "%v", float64(t))
    }
}
```

### Stringer Implementation
```go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}
```

## Scanning Functions

The scanning side mirrors printing:
- `Scan`, `Scanf`, `Scanln` for stdin
- `Fscan`, `Fscanf`, `Fscanln` for io.Reader
- `Sscan`, `Sscanf`, `Sscanln` for strings

Key scanning features:
- Space handling (different from C's scanf)
- Base prefixes (0b, 0o, 0x) automatically recognized
- Width support (but no precision in scanning)
- Interface support via `Scanner` interface

## Common Patterns and Idioms

### 1. Error Handling
```go
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
```

### 2. Debugging
```go
fmt.Printf("Debug: %+v\n", complexStruct)
```

### 3. Logging
```go
log.Printf("User %s performed action on %s", userID, resource)
```

### 4. String Building
```go
// Complex formatting
msg := fmt.Sprintf("User %s (ID: %d) has %d items", name, id, count)

// Simple concatenation (faster)
simple := "Hello, " + name
```

## Comparison with Other Languages

### vs C printf
- **Similarities**: Basic format verbs, flag system
- **Differences**: No varargs (uses variadic functions), safer type system, Unicode support

### vs Python format
- **Similarities**: Type-based formatting, interface extensibility
- **Differences**: Compile-time checking, performance focus, simpler syntax

### vs Java String.format
- **Similarities**: Method naming conventions
- **Differences**: Interface-based extensibility, better performance characteristics

## Gotchas and Common Mistakes

### 1. Interface Confusion
```go
// Wrong - will print type information
fmt.Printf("%v", myInterface)

// Better - let the method handle it
fmt.Printf("%s", myStringer)
```

### 2. Format String Injection
Be careful with user input in format strings:
```go
// Dangerous - user could contain format verbs
fmt.Printf(userInput)

// Safe - treat as literal string
fmt.Printf("%s", userInput)
```

### 3. Performance Pitfalls
```go
// Inefficient in loops
for i := 0; i < 1000; i++ {
    result += fmt.Sprintf("%d,", i)
}

// Better
var buf strings.Builder
for i := 0; i < 1000; i++ {
    buf.WriteString(strconv.Itoa(i))
    buf.WriteString(",")
}
result = buf.String()
```

## Testing with fmt

The package is extensively tested with:
- Unit tests for all format verbs and combinations
- Property-based tests for edge cases
- Benchmarks for performance regression detection
- Cross-platform Unicode tests

## Conclusion

The `fmt` package exemplifies Go's design philosophy:
1. **Simplicity**: Clean, intuitive interface
2. **Safety**: Type checking and error handling
3. **Performance**: Efficient implementation with careful resource management
4. **Extensibility**: Interface-based customization
5. **Practicality**: Real-world usage patterns baked in

It's a masterclass in Go package design, balancing performance with usability while maintaining the language's core principles of clarity and correctness.