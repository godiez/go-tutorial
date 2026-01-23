# Go Data Structures Quick Reference

## Arrays vs Slices vs Maps

| Feature | Array | Slice | Map |
|---------|-------|-------|-----|
| **Size** | Fixed | Dynamic | Dynamic |
| **Type** | Value | Reference | Reference |
| **Declaration** | `[5]int` | `[]int` | `map[string]int` |
| **Zero value** | Array of zeros | `nil` | `nil` |
| **Creation** | `[3]int{1,2,3}` | `make([]int, 5)` | `make(map[K]V)` |
| **Access** | `arr[i]` | `slice[i]` | `m[key]` |
| **Length** | `len(arr)` | `len(slice)` | `len(m)` |
| **Capacity** | Fixed | `cap(slice)` | N/A |

## Memory Allocation: new() vs make()

```go
// new(T) - Returns *T, zeroed memory, works with ANY type
intPtr := new(int)              // *int, value = 0
structPtr := new(Person)        // *Person, fields zeroed

// make(T, args) - Returns T, initialized, ONLY slices/maps/channels
slice := make([]int, 5)         // []int, len=5, cap=5
sliceCap := make([]int, 5, 10)  // []int, len=5, cap=10
m := make(map[string]int)       // map[string]int, ready to use
```

### When to use what?

- **Slices**: `make([]T, len, cap)` or `[]T{1, 2, 3}`
- **Maps**: `make(map[K]V)` or `map[K]V{k: v}`
- **Structs**: `T{}` or `&T{field: value}`
- **new()**: Rarely needed (returns pointer to zero value)

## Slices

```go
// Creation
var s1 []int                    // nil slice
s2 := []int{}                   // empty slice literal
s3 := []int{1, 2, 3}           // initialized
s4 := make([]int, 5)           // len=5, cap=5, zeros
s5 := make([]int, 3, 10)       // len=3, cap=10

// Operations
s = append(s, 1, 2, 3)         // Add elements
s = append(s, other...)        // Append slice
copy(dest, src)                // Copy elements
sub := s[1:3]                  // Slicing [low:high]
sub := s[:3]                   // From start
sub := s[2:]                   // To end

// Patterns
// Filter
filtered := s[:0]
for _, v := range s {
    if v > 5 {
        filtered = append(filtered, v)
    }
}

// Map
mapped := make([]int, len(s))
for i, v := range s {
    mapped[i] = v * 2
}

// Reduce
sum := 0
for _, v := range s {
    sum += v
}
```

## Maps

```go
// Creation
var m1 map[string]int           // nil map (can't insert!)
m2 := make(map[string]int)      // empty map, ready to use
m3 := map[string]int{           // initialized
    "one": 1,
    "two": 2,
}

// Operations
m["key"] = value                // Insert/Update
value := m["key"]               // Read (zero if missing)
delete(m, "key")                // Delete
len(m)                          // Number of entries

// Check existence (comma-ok idiom)
if value, exists := m["key"]; exists {
    // key exists
}

// Iteration (order is random!)
for key, value := range m {
    fmt.Println(key, value)
}

// Set pattern (map[T]bool or map[T]struct{})
set := make(map[string]bool)
set["item"] = true
if set["item"] {
    // item in set
}
```

## Structs

```go
// Definition
type Person struct {
    Name string
    Age  int
    City string
}

// Creation
var p1 Person                   // Zero value
p2 := Person{                   // With field names (preferred)
    Name: "Alice",
    Age:  30,
}
p3 := Person{"Bob", 25, "NYC"}  // Without names (not recommended)
p4 := &Person{Name: "Charlie"}  // Pointer to struct
p5 := new(Person)               // Pointer, zeroed

// Methods
// Value receiver (works on copy)
func (p Person) GetName() string {
    return p.Name
}

// Pointer receiver (can modify)
func (p *Person) SetAge(age int) {
    p.Age = age
}

// Embedding (composition)
type Employee struct {
    Person              // Embedded (promoted fields)
    EmployeeID int
}
emp := Employee{
    Person: Person{Name: "Alice"},
    EmployeeID: 123,
}
fmt.Println(emp.Name)  // Promoted field
```

## Common Patterns

### Constructor Function
```go
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
        City: "Unknown",  // Default
    }
}
```

### Functional Options (Builder)
```go
type Option func(*Config)

func WithHost(host string) Option {
    return func(c *Config) { c.Host = host }
}

func NewConfig(opts ...Option) *Config {
    c := &Config{Host: "localhost"}  // Defaults
    for _, opt := range opts {
        opt(c)
    }
    return c
}

cfg := NewConfig(WithHost("example.com"))
```

### Counting with Maps
```go
counts := make(map[string]int)
for _, word := range words {
    counts[word]++
}
```

### Grouping with Maps
```go
groups := make(map[string][]Item)
for _, item := range items {
    key := item.Category
    groups[key] = append(groups[key], item)
}
```

## Common Gotchas

### 1. Nil Map Assignment
```go
var m map[string]int   // nil
m["key"] = 1           // PANIC!

m = make(map[string]int)  // Fix
m["key"] = 1              // OK
```

### 2. Slice Backing Array
```go
s := []int{1, 2, 3, 4, 5}
sub := s[0:2]          // {1, 2}
sub = append(sub, 99)  // Modifies s!

// Fix: limit capacity
sub := s[0:2:2]        // len=2, cap=2
sub = append(sub, 99)  // New array
```

### 3. Range Variable Reuse
```go
var ptrs []*int
for _, v := range []int{1, 2, 3} {
    ptrs = append(ptrs, &v)  // BUG: all point to same v
}

// Fix: create new variable
for _, v := range []int{1, 2, 3} {
    v := v  // New variable
    ptrs = append(ptrs, &v)
}
```

### 4. Struct Copying
```go
p1 := Person{Name: "Alice"}
p2 := p1           // Copied by value
p2.Name = "Bob"    // Doesn't affect p1

// Use pointers to share
p3 := &p1
p3.Name = "Charlie"  // Modifies p1
```

### 5. Zero Values vs Missing Keys
```go
m := map[string]int{"zero": 0}
v1 := m["zero"]     // 0 (exists)
v2 := m["missing"]  // 0 (doesn't exist!)

// Use comma-ok to distinguish
if v, exists := m["zero"]; exists {
    // Key exists, v is 0
}
```

## Best Practices

1. **Pre-allocate** when size is known: `make([]T, 0, capacity)`
2. **Initialize maps** with `make()` before use
3. **Use pointer receivers** for methods that modify or for large structs
4. **Use comma-ok idiom** to check map existence
5. **Use constructor functions** for validation and defaults
6. **Prefer struct literals** over `new()`: `&T{}` not `new(T)`
7. **Use embedding** for composition, not inheritance
8. **Document** whether functions expect nil or initialized values

## Performance Tips

1. Pre-allocate slices when size is known
2. Use `map[T]struct{}` instead of `map[T]bool` for sets (saves memory)
3. Avoid appending in tight loops (pre-allocate instead)
4. Use pointer receivers for large structs to avoid copying
5. Be aware of slice capacity growth (doubles until 1024, then 25% increments)
