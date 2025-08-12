package main

import (
	"fmt"

	"github.com/grandpaej/fmtpy/v2"
	"github.com/grandpaej/fmtpy/v2/color"
	"github.com/grandpaej/fmtpy/v2/input"
)

func main() {
	// Test the new Input function with type conversion
	name := fmtpy.Input("What's your name? ").String()

	// Test different Print styles
	fmtpy.Print("Hello %s!", name)                   // Go style
	fmtpy.Print(`f"Nice to meet you, {name}"`, name) // Python style

	// Demonstrate new color functions that work with fmt
	fmt.Println("\n=== Color Demo ===")
	fmt.Println(color.Red("This is red text"))
	fmt.Println(color.Green("This is green text"))
	fmt.Println(color.Blue("This is blue text"))
	fmt.Println(color.BoldYellow("This is bold yellow"))

	// Colors work with any type!
	fmt.Println(color.Cyan(42))                     // Numbers
	fmt.Println(color.Magenta(1 + 1))               // Expressions
	fmt.Printf("Result: %s\n", color.Green(123.45)) // With Printf

	// Demonstrate string manipulation functions from input package
	fmt.Println("\n=== String Manipulation Demo ===")

	// Case conversion
	fmt.Printf("Uppercase: %s\n", input.Upper(name))
	fmt.Printf("Lowercase: %s\n", input.Lower(name))
	fmt.Printf("Title Case: %s\n", input.Title(name))
	fmt.Printf("Capitalized: %s\n", input.Capitalize(name))

	// Fun with strings
	fmt.Printf("Reversed: %s\n", input.Reverse(name))
	fmt.Printf("Repeated 3x: %s\n", input.Repeat(name, 3))
	fmt.Printf("Swap Case: %s\n", input.SwapCase(name))

	// Get a sentence for more demos
	sentence := fmtpy.Input("\nEnter a sentence: ").String()

	// Text analysis
	fmt.Printf("Word count: %d\n", input.WordCount(sentence))
	fmt.Printf("Character count: %d\n", input.CharCount(sentence))
	fmt.Printf("First word: %s\n", input.FirstWord(sentence))
	fmt.Printf("Last word: %s\n", input.LastWord(sentence))

	// String operations
	fmt.Printf("Only letters: %s\n", input.OnlyLetters(sentence))
	fmt.Printf("Only numbers: %s\n", input.OnlyNumbers(sentence))

	// Demonstrate the new type conversion
	var age int
	age = fmtpy.Input("How old are you? ").Int() // Direct assignment!
	fmt.Printf("Next year you'll be %d\n", age+1)

	// Get a float score
	var score float64
	score = fmtpy.Input("Enter your score: ").Float()
	fmt.Printf("Your score: %.2f\n", score)

	// Ask a yes/no question
	confirmed := fmtpy.Input("Do you like programming? (y/n): ").Bool()
	if confirmed {
		fmt.Println(color.Green("That's awesome! Keep coding!"))
	} else {
		fmt.Println(color.Yellow("Maybe you'll discover the joy of coding soon!"))
	}

	// Formatting demo
	fmt.Println("\n=== Formatting Demo ===")
	text := "Hello"
	fmt.Printf("Centered (10): '%s'\n", input.Center(text, 10))
	fmt.Printf("Left padded: '%s'\n", input.LeftPad(text, 10, "*"))
	fmt.Printf("Right padded: '%s'\n", input.RightPad(text, 10, "-"))

	// Show the power of combining colors with formatting
	fmt.Println("\n=== Color + Formatting Demo ===")
	fmt.Println(color.BoldRed(input.Center("ERROR", 20)))
	fmt.Println(color.BoldGreen(input.Center("SUCCESS", 20)))
	fmt.Println(color.OnYellow(input.Center("WARNING", 20)))
}
