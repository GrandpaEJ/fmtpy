package main

import "fmt2"

func main() {
	// Test the Input function
	name := fmt2.Input("What's your name? ")
	
	// Test different Print styles
	fmt2.Print("Hello %s!", name)                    // Go style
	fmt2.Print(`f"Nice to meet you, {name}"`, name) // Python style
	fmt2.Print(42)                                  // Simple print number
	fmt2.Print("Here's a string")                   // Simple print string
	
	// Get more input with newline
	hobby := fmt2.Input("What's your favorite hobby?\n")
	fmt2.Print(`f"So you like {hobby}, that's cool!"`, hobby)
}
