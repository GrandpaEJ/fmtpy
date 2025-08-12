package main

import (
	"fmt"

	"github.com/grandpaej/fmtpy"
	"github.com/grandpaej/fmtpy/color"
	"github.com/grandpaej/fmtpy/input"
)

func main() {
	fmt.Println("=== Testing New fmtpy Features ===")

	// Test 1: InputValue type conversion
	fmt.Println("\n1. Testing InputValue type conversion:")

	// Simulate input for testing
	testInput := fmtpy.InputValue("123")
	var age int = testInput.Int()
	fmt.Printf("   String '123' -> Int: %d\n", age)

	testInput = fmtpy.InputValue("123.45")
	var score float64 = testInput.Float()
	fmt.Printf("   String '123.45' -> Float: %.2f\n", score)

	testInput = fmtpy.InputValue("y")
	var confirmed bool = testInput.Bool()
	fmt.Printf("   String 'y' -> Bool: %t\n", confirmed)

	// Test 2: Color functions with fmt
	fmt.Println("\n2. Testing color functions with fmt:")
	fmt.Println("   ", color.Red("Red text"))
	fmt.Println("   ", color.Green("Green text"))
	fmt.Println("   ", color.Blue("Blue text"))
	fmt.Println("   ", color.BoldYellow("Bold Yellow"))

	// Test colors with different types
	fmt.Printf("   Number: %s\n", color.Cyan(42))
	fmt.Printf("   Expression: %s\n", color.Magenta(1+1))
	fmt.Printf("   Float: %s\n", color.Green(123.45))

	// Test 3: String manipulation from input package
	fmt.Println("\n3. Testing string manipulation:")
	testStr := "Hello World"
	fmt.Printf("   Original: %s\n", testStr)
	fmt.Printf("   Upper: %s\n", input.Upper(testStr))
	fmt.Printf("   Lower: %s\n", input.Lower(testStr))
	fmt.Printf("   Reversed: %s\n", input.Reverse(testStr))
	fmt.Printf("   Word count: %d\n", input.WordCount(testStr))

	// Test 4: Combining colors with string manipulation
	fmt.Println("\n4. Testing color + string manipulation:")
	fmt.Println("   ", color.BoldRed(input.Upper("error")))
	fmt.Println("   ", color.BoldGreen(input.Upper("success")))
	fmt.Println("   ", color.OnYellow(input.Center("WARNING", 15)))

	fmt.Println("\n✅ All tests completed successfully!")
}
