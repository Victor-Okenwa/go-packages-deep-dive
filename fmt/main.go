package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("=== Go fmt Package Deep Dive ===")

	// 1. Basic Printing
	name := "Victor"
	age := 22
	isStudent := false
	height := 1.75

	fmt.Println("1. Basic Printing:")
	fmt.Println("Hello, my name is", name)
	fmt.Println("I am", age, "years old. Am I a student?", isStudent)
	fmt.Println("My height is", height, "meters.")

	// 2. Printf - Formatted Printing (Most Important)
	fmt.Println("\n2. Printf - Formatted strings:")
	fmt.Printf("My name is %s\n", name)              // %s = string
	fmt.Printf("I am %d years old\n", age)           // %d = integer
	fmt.Printf("My height is %.3f meters\n", height) // %.2f = float with 2 decimal places
	fmt.Printf("Am I a student? %t\n", isStudent)    // %t = boolean

	// 3. Common Formatting Verbs
	fmt.Println("\n3. Common Formatting Verbs:")

	number := 255
	pi := math.Pi

	fmt.Printf("Decimal: %d\n", number)
	fmt.Printf("Binary: %b\n", number)
	fmt.Printf("Octal: %o\n", number)
	fmt.Printf("Hexadecimal: %x\n", number)
	fmt.Printf("Hex (uppercase): %X\n", number)
	fmt.Printf("Pi: %f\n", pi)
	fmt.Printf("Pi (scientific): %e\n", pi)
	fmt.Printf("Pi (shortest): %g\n", pi)

	// 4. Struct Formatting
	fmt.Println("\n4. Struct Formatting:")

	type Person struct {
		Name    string
		Age     int
		Balance float64
	}

	person := Person{
		Name:    "Alice",
		Age:     25,
		Balance: 2450.75,
	}

	fmt.Printf("Default: %v\n", person)         // Default format
	fmt.Printf("With fields: %+v\n", person)    // Shows field names
	fmt.Printf("Go code format: %#v\n", person) // Go syntax

	// 5. Formatting with Width and Precision
	fmt.Println("\n5. Width and Precision:")

	fmt.Printf("Price: $%8.2f\n", 199.99)  // Width 8, 2 decimals
	fmt.Printf("Price: $%-8.2f\n", 199.99) // Left aligned

	// 6. Sprintf - Format and return as string
	fmt.Println("\n6. Sprintf - Return formatted string:")

	message := fmt.Sprintf("Hello %s, you are %d years old.", name, age)
	fmt.Println("Sprintf result:", message)

	// 7. Errorf - Create formatted errors
	fmt.Println("\n7. Errorf:")

	err := fmt.Errorf("user %s not found: age must be at least %d", name, 18)
	fmt.Println("Error:", err)

	// 8. Fprintf - Write to any io.Writer
	fmt.Println("\n8. Fprintf (not commonly used in simple examples)")

	// Bonus: Println vs Printf vs Print
	fmt.Println("\n=== Println vs Printf vs Print ===")
	fmt.Println("Println adds space and newline, automatically")
	fmt.Print("Print", "does", "not", "add", "spaces\n")
	fmt.Printf("Printf %s full control\n", "gives")
}
