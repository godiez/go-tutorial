# Go Data Structures Learning Module

## Overview

This module provides a comprehensive, hands-on tutorial for understanding Go's core data structures and memory allocation mechanisms. Perfect for developers learning Go or transitioning from other languages.

## What You'll Learn

### Core Data Structures
- **Arrays**: Fixed-size collections and their value semantics
- **Slices**: Go's dynamic arrays with powerful manipulation capabilities
- **Maps**: Hash tables for key-value storage and retrieval
- **Structs**: Custom types with fields, methods, and composition

### Memory Management
- **new()**: Pointer allocation for any type
- **make()**: Initialization for slices, maps, and channels
- Understanding when to use each

### Practical Patterns
- Slice patterns: filtering, mapping, reducing
- Map patterns: counting, grouping, sets, caching
- Struct patterns: constructors, builders, embedding
- Common idioms: comma-ok, range loops, pointers

### Advanced Concepts
- Value vs reference semantics
- Backing arrays and capacity management
- Method receivers (value vs pointer)
- Struct embedding and composition
- Interface implementation (fmt.Stringer)

## Module Contents

```
datastructures/
├── main.go              # Interactive tutorial program with menu
├── arrays_slices.go     # 300+ lines of arrays/slices examples
├── maps.go              # 350+ lines of map examples
├── structs.go           # 400+ lines of struct examples
├── new_vs_make.go       # 320+ lines comparing allocation methods
├── README.md            # Full documentation
└── QUICK_REFERENCE.md   # Cheat sheet for quick lookup
```

**Total: ~1,400 lines of documented, runnable examples**

## Features

### Interactive Learning
- Menu-driven interface to explore topics at your own pace
- Run individual topics or all examples at once
- Clear output with explanations and comparisons

### Comprehensive Examples
- **Basic concepts**: Zero values, literals, initialization
- **Operations**: Insert, read, update, delete, iterate
- **Patterns**: Real-world use cases you'll encounter daily
- **Gotchas**: Common mistakes and how to avoid them
- **Best practices**: Idiomatic Go code

### Well-Documented Code
- Every function has clear explanations
- Output shows results and behavior
- Side-by-side comparisons (e.g., new() vs make())
- Comments explain the "why" not just the "what"

## Getting Started

### Prerequisites
- Docker (recommended) or Go 1.16+ installed locally
- No prior Go experience required (but basic programming knowledge helps)

### Running the Tutorial

**Option 1: Using Docker (Recommended)**
```bash
# Start the container
docker compose -f docker/docker-compose.yml up -d

# Enter the container
docker compose -f docker/docker-compose.yml exec go-learning bash

# Navigate and run
cd datastructures
go run .
```

**Option 2: Local Go Installation**
```bash
cd datastructures
go run .
```

### Interactive Menu
Once running, you'll see:
```
╔════════════════════════════════════════════════════════════╗
║          GO DATA STRUCTURES TUTORIAL                      ║
║   Arrays, Slices, Maps, Structs, new() and make()         ║
╚════════════════════════════════════════════════════════════╝

Select a topic to learn:
  1. Arrays & Slices
  2. Maps
  3. Structs
  4. new() vs make()
  5. Run ALL examples
  0. Exit
```

## Topics Covered in Detail

### 1. Arrays & Slices (arrays_slices.go)
- Array basics and fixed-size behavior
- Slice creation with literals, make(), and from arrays
- Understanding length vs capacity
- Append, copy, and slice operations
- Capacity growth and reallocation
- Filtering, mapping, and reducing patterns
- Common gotchas: backing arrays, range variables

**Key Example: Slice Growth**
```go
slice := make([]int, 0, 10)  // Pre-allocate capacity
for i := 0; i < 10; i++ {
    slice = append(slice, i)  // No reallocation!
}
```

### 2. Maps (maps.go)
- Map creation and initialization
- Insert, read, update, delete operations
- The comma-ok idiom for checking existence
- Iteration (and why order is random)
- Maps with complex value types
- Counting, grouping, set, and caching patterns
- Common gotchas: nil maps, concurrent access

**Key Example: Word Counting**
```go
counts := make(map[string]int)
for _, word := range words {
    counts[word]++  // So elegant!
}
```

### 3. Structs (structs.go)
- Struct definition and initialization
- Value vs pointer semantics
- Struct comparison and equality
- Embedding for composition
- Methods with value and pointer receivers
- Struct tags for metadata
- Constructor, builder, and anonymous struct patterns
- Common gotchas: copying, zero values

**Key Example: Struct Embedding**
```go
type Employee struct {
    Person  // Embedded - promotes fields
    ID int
}
emp := Employee{Person: Person{Name: "Alice"}, ID: 123}
fmt.Println(emp.Name)  // Promoted field access
```

### 4. new() vs make() (new_vs_make.go)
- How new() allocates zeroed memory
- How make() initializes data structures
- Return type differences (pointer vs value)
- When to use each (spoiler: rarely use new())
- Idiomatic Go patterns for initialization
- Common mistakes and how to avoid them

**Key Example: The Difference**
```go
// new() - Returns pointer, rarely used
p := new([]int)    // *[]int (pointer to nil slice)
*p = append(*p, 1) // Must dereference

// make() - Returns value, preferred
s := make([]int, 5)  // []int (ready to use)
s[0] = 1             // Can use immediately
```

## Learning Path

### Beginner (Start Here)
1. Run "Arrays & Slices" - understand the most common data structure
2. Run "Maps" - learn Go's hash tables
3. Study the QUICK_REFERENCE.md - reinforce concepts

### Intermediate
1. Run "Structs" - master custom types
2. Run "new() vs make()" - understand memory allocation
3. Review common patterns in each file

### Advanced
1. Read through all source files in detail
2. Experiment with modifying examples
3. Try implementing your own patterns
4. Study the gotchas sections carefully

## Key Takeaways

After completing this tutorial, you'll understand:

✅ **When to use arrays vs slices** (almost always slices)  
✅ **How slices grow** and how to pre-allocate for performance  
✅ **The comma-ok idiom** for safe map access  
✅ **Value vs pointer receivers** for struct methods  
✅ **Why you should use make()** not new() for slices/maps  
✅ **Struct embedding** for composition over inheritance  
✅ **Common patterns** you'll see in real Go code  
✅ **Gotchas to avoid** that trip up beginners  

## Best Practices Learned

1. **Pre-allocate** slices when you know the size
2. **Always initialize** maps with make() before use
3. **Use pointer receivers** for methods that modify
4. **Check map existence** with comma-ok idiom
5. **Use constructor functions** for validation
6. **Prefer literals** over new(): `&T{}` not `new(T)`
7. **Document nil behavior** in your APIs
8. **Use embedding** for composition

## Next Steps

After mastering this module:
- **Interfaces**: Learn Go's powerful interface system
- **Concurrency**: Goroutines and channels
- **Error Handling**: The Go way of handling errors
- **Testing**: Table-driven tests using anonymous structs
- **JSON/Encoding**: Using struct tags with encoding packages

## Additional Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Go maps in action](https://go.dev/blog/maps)

## Contributing

Found an issue or want to add more examples? Feel free to modify the files and experiment. That's the best way to learn!

---

**Happy Learning!** 🚀

Start with `go run .` in the datastructures directory and explore each topic at your own pace.
