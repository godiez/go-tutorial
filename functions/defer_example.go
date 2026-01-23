package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("=== Defer Examples ===\n")

	// Example 1: Simple defer - executes in LIFO order
	fmt.Println("1. Simple defer - LIFO order:")
	simpleDeferExample()

	// Example 2: File operations with defer
	fmt.Println("\n2. File operations with defer:")
	fileDeferExample()

	// Example 3: Multiple defers for resource cleanup
	fmt.Println("\n3. Resource cleanup with multiple defers:")
	resourceCleanupExample()

	// Example 4: Defer with function parameters
	fmt.Println("\n4. Defer with function parameters:")
	deferWithParametersExample()

	// Example 5: Defer for recovery from panic
	fmt.Println("\n5. Panic recovery with defer:")
	panicRecoveryExample()
}

func simpleDeferExample() {
	fmt.Println("Function started")
	
	defer fmt.Println("First defer (executed last)")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer (executed first)")
	defer fmt.Println("Third defer (executed firsttttt)")
	
	fmt.Println("Function ending")
}

func fileDeferExample() {
	// Create a test file
	filename := "test.txt"
	content := "Hello, Go defer!"
	
	// Write to file
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}

	// Read file with defer for closing
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer func() {
		fmt.Println("Closing file...")
		file.Close()
	}()

	// Read file content
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("File content: %s\n", string(buffer[:n]))
}

func resourceCleanupExample() {
	fmt.Println("Setting up resources...")
	
	// Simulate database connection
	defer func() {
		fmt.Println("Closing database connection...")
	}()
	
	// Simulate file handle
	defer func() {
		fmt.Println("Closing file handle...")
	}()
	
	// Simulate network connection
	defer func() {
		fmt.Println("Closing network connection...")
	}()
	
	fmt.Println("Working with resources...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done with resources")
}

func deferWithParametersExample() {
	message := "Original message"
	
	// The parameter is evaluated immediately, but execution is deferred
	defer printMessage(message)
	
	message = "Changed message"
	
	fmt.Println("Inside function, after defer setup")
}

func printMessage(msg string) {
	fmt.Printf("Deferred message: %s\n", msg)
}

func panicRecoveryExample() {
	safeFunction := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()

		fmt.Println("About to cause a panic...")
		panic("Something went wrong!")
		// This line will never be reached
		fmt.Println("This will not print")
	}

	safeFunction()
	fmt.Println("Program continues after recovery")
}