# Go Data Structures Tutorial

A comprehensive guide to data structures, memory allocation, and common patterns in Go.

## Topics Covered

### 1. Arrays & Slices (`arrays_slices.go`)
- **Arrays**: Fixed-size, value types
- **Slices**: Dynamic, reference types backed by arrays
- **Key Operations**:
  - Creating slices with `make()`
  - Appending and copying
  - Slicing operations
  - Capacity and growth behavior
- **Common Patterns**:
  - Filtering
  - Mapping
  - Reducing
- **Gotchas**: Shared backing arrays, range variable reuse

### 2. Maps (`maps.go`)
- **Basics**: Key-value pairs (hash tables)
- **Key Operations**:
  - Creating with `make()`
  - Insert, read, delete
  - Checking existence (comma-ok idiom)
  - Iteration (unordered)
- **Common Patterns**:
  - Grouping data
  - Counting occurrences
  - Set implementation
  - Caching/memoization
- **Gotchas**: Nil maps, concurrent access, uncomparable types

### 3. Structs (`structs.go`)
- **Basics**: Typed collections of fields
- **Key Concepts**:
  - Value vs pointer semantics
  - Struct embedding (composition)
  - Methods (value vs pointer receivers)
  - Struct tags
- **Common Patterns**:
  - Constructor functions
  - Builder pattern (functional options)
  - Anonymous structs
  - Table-driven tests
- **Gotchas**: Copying by value, zero values, method receivers

### 4. new() vs make() (`new_vs_make.go`)
- **new(T)**:
  - Allocates zeroed memory
  - Returns pointer `*T`
  - Works with any type
  - Rarely used in practice
- **make(T)**:
  - Initializes and returns `T`
  - Only for slices, maps, channels
  - Returns ready-to-use value
  - Preferred for these types
- **Practical Guidance**: When to use each, idiomatic patterns

## Running the Examples

### Option 1: Using Docker (Recommended)

1. Start the container:
   ```bash
   docker compose -f docker/docker-compose.yml up -d
   ```

2. Enter the container:
   ```bash
   docker compose -f docker/docker-compose.yml exec go-learning bash
   ```

3. Navigate to the datastructures directory:
   ```bash
   cd datastructures
   ```

4. Run the interactive tutorial:
   ```bash
   go run .
   ```

### Option 2: Local Go Installation

If you have Go installed locally:

```bash
cd datastructures
go run .
```

## Interactive Menu

The main program presents an interactive menu:

```
Select a topic to learn:
  1. Arrays & Slices
  2. Maps
  3. Structs
  4. new() vs make()
  5. Run ALL examples
  0. Exit
```

Choose a number to run specific examples or see all of them at once.

## File Structure

```
datastructures/
├── main.go           # Interactive menu program
├── arrays_slices.go  # Arrays and slices examples
├── maps.go           # Map examples
├── structs.go        # Struct examples
├── new_vs_make.go    # Memory allocation comparison
└── README.md         # This file
```

## Key Takeaways

### Memory Allocation
- **Slices, Maps**: Always use `make()` or literals
- **Structs**: Use literals `T{}` or `&T{}`
- **new()**: Rarely needed; returns pointer to zero value

### Value vs Reference
- **Value types** (copied): Arrays, structs, basic types
- **Reference types** (share data): Slices, maps, channels

### Best Practices
1. **Pre-allocate** when size is known: `make([]int, 0, capacity)`
2. **Use the comma-ok idiom** to check map keys
3. **Prefer pointer receivers** for methods that modify or for large structs
4. **Use constructor functions** for validation and defaults
5. **Avoid nil maps/slices** - initialize with `make()` or literals

### Common Pitfalls
1. Appending to slices can modify shared backing arrays
2. Nil maps panic on assignment (must initialize first)
3. Maps are not safe for concurrent access
4. Range loop variables are reused (be careful with pointers)
5. Structs are copied by value when passed to functions

## Next Steps

After completing this tutorial, you should understand:
- How to choose the right data structure
- When to use `new()` vs `make()`
- Common patterns for working with collections
- How to avoid common pitfalls

## Additional Resources

- [Effective Go - Data](https://go.dev/doc/effective_go#data)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Go maps in action](https://go.dev/blog/maps)
- [Go by Example](https://gobyexample.com/)
