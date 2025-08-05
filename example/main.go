package main

import "fmtpy"

func main() {
	// Test the Input function
	name := fmtpy.Input("What's your name? ")
	
	// Test different Print styles
	fmtpy.Print("Hello %s!", name)                    // Go style
	fmtpy.Print(`f"Nice to meet you, {name}"`, name) // Python style
	fmtpy.Print(42)                                  // Simple print number
	fmtpy.Print("Here's a string")                   // Simple print string
	
	// Get more input with newline
	hobby := fmtpy.Input("What's your favorite hobby?\n")
	fmtpy.Print(`f"So you like {hobby}, that's cool!"`, hobby)
}
