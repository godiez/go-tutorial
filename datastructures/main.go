package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║          GO DATA STRUCTURES TUTORIAL                      ║")
	fmt.Println("║   Arrays, Slices, Maps, Structs, new() and make()         ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n" + strings.Repeat("─", 60))
		fmt.Println("Select a topic to learn:")
		fmt.Println("  1. Arrays & Slices")
		fmt.Println("  2. Maps")
		fmt.Println("  3. Structs")
		fmt.Println("  4. new() vs make()")
		fmt.Println("  5. Run ALL examples")
		fmt.Println("  0. Exit")
		fmt.Print("\nYour choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			RunArraysSlices()
		case "2":
			RunMaps()
		case "3":
			RunStructs()
		case "4":
			RunNewVsMake()
		case "5":
			RunAll()
		case "0":
			fmt.Println("\nHappy coding! 🚀")
			return
		default:
			fmt.Println("\n❌ Invalid choice. Please enter 0-5.")
		}

		fmt.Println("\n" + strings.Repeat("─", 60))
		fmt.Print("Press ENTER to continue...")
		reader.ReadString('\n')
	}
}

// RunAll executes all examples in sequence
func RunAll() {
	RunArraysSlices()
	RunMaps()
	RunStructs()
	RunNewVsMake()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("ALL EXAMPLES COMPLETED!")
	fmt.Println(strings.Repeat("=", 60))
}
