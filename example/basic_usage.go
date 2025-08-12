package main

import (
	"fmt"

	"github.com/grandpaej/fmtpy/v2"
	"github.com/grandpaej/fmtpy/v2/color"
	"github.com/grandpaej/fmtpy/v2/input"
)

func main() {
	fmt.Println(color.BoldBlue("=== Basic fmtpy Usage Examples ==="))

	// Example 1: Simple input and type conversion
	fmt.Println(color.Green("\n1. Basic Input with Type Conversion:"))

	name := fmtpy.Input("Enter your name: ").String()
	fmt.Printf("Hello, %s!\n", color.Blue(name))

	var age int
	age = fmtpy.Input("Enter your age: ").Int()
	fmt.Printf("You are %s years old\n", color.Cyan(age))

	var height float64
	height = fmtpy.Input("Enter your height (in meters): ").Float()
	fmt.Printf("Your height is %s meters\n", color.Yellow(height))

	// Example 2: Boolean input
	fmt.Println(color.Green("\n2. Boolean Input:"))

	likes_programming := fmtpy.Input("Do you like programming? (y/n): ").Bool()
	if likes_programming {
		fmt.Println(color.BoldGreen("🎉 Awesome! Programming is fun!"))
	} else {
		fmt.Println(color.BoldYellow("🤔 Maybe you'll discover it later!"))
	}

	// Example 3: String manipulation
	fmt.Println(color.Green("\n3. String Manipulation:"))

	text := fmtpy.Input("Enter some text: ").String()
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Uppercase: %s\n", color.Red(input.Upper(text)))
	fmt.Printf("Lowercase: %s\n", color.Blue(input.Lower(text)))
	fmt.Printf("Reversed: %s\n", color.Magenta(input.Reverse(text)))
	fmt.Printf("Word count: %s\n", color.Cyan(input.WordCount(text)))

	// Example 4: Colors with different data types
	fmt.Println(color.Green("\n4. Colors with Different Types:"))

	fmt.Printf("String: %s\n", color.Red("Hello World"))
	fmt.Printf("Integer: %s\n", color.Green(42))
	fmt.Printf("Float: %s\n", color.Blue(3.14159))
	fmt.Printf("Boolean: %s\n", color.Yellow(true))
	fmt.Printf("Expression: %s\n", color.Magenta(10+5))

	// Example 5: Background colors and formatting
	fmt.Println(color.Green("\n5. Background Colors and Formatting:"))

	fmt.Println(color.OnRed("  ERROR  "))
	fmt.Println(color.OnGreen("  SUCCESS  "))
	fmt.Println(color.OnYellow("  WARNING  "))
	fmt.Println(color.OnBlue("  INFO  "))

	// Example 6: Combining everything
	fmt.Println(color.Green("\n6. Combining Features:"))

	username := fmtpy.Input("Create a username: ").String()
	processed_username := input.Lower(input.Trim(username))

	fmt.Printf("Original: %s\n", color.Red(username))
	fmt.Printf("Processed: %s\n", color.Green(processed_username))
	fmt.Printf("Length: %s characters\n", color.Cyan(input.Length(processed_username)))

	if input.Length(processed_username) >= 3 {
		fmt.Println(color.BoldGreen("✅ Username is valid!"))
	} else {
		fmt.Println(color.BoldRed("❌ Username too short!"))
	}

	fmt.Println(color.BoldBlue("\n🎯 That's the power of fmtpy!"))
}
