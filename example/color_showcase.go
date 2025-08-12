package main

import (
	"fmt"

	"github.com/grandpaej/fmtpy/v2/color"
)

func main() {
	fmt.Println(color.BoldCyan("=== fmtpy Color Showcase ==="))

	// Basic Colors
	fmt.Println(color.Green("\n🎨 Basic Colors:"))
	fmt.Printf("Red: %s\n", color.Red("Hello World"))
	fmt.Printf("Green: %s\n", color.Green("Hello World"))
	fmt.Printf("Blue: %s\n", color.Blue("Hello World"))
	fmt.Printf("Yellow: %s\n", color.Yellow("Hello World"))
	fmt.Printf("Magenta: %s\n", color.Magenta("Hello World"))
	fmt.Printf("Cyan: %s\n", color.Cyan("Hello World"))
	fmt.Printf("White: %s\n", color.White("Hello World"))
	fmt.Printf("Black: %s\n", color.Black("Hello World"))

	// Bold Colors
	fmt.Println(color.Green("\n💪 Bold Colors:"))
	fmt.Printf("Bold Red: %s\n", color.BoldRed("Important!"))
	fmt.Printf("Bold Green: %s\n", color.BoldGreen("Success!"))
	fmt.Printf("Bold Blue: %s\n", color.BoldBlue("Information"))
	fmt.Printf("Bold Yellow: %s\n", color.BoldYellow("Warning!"))
	fmt.Printf("Bold Magenta: %s\n", color.BoldMagenta("Special"))
	fmt.Printf("Bold Cyan: %s\n", color.BoldCyan("Highlight"))

	// Background Colors
	fmt.Println(color.Green("\n🌈 Background Colors:"))
	fmt.Printf("On Red: %s\n", color.OnRed("  ERROR  "))
	fmt.Printf("On Green: %s\n", color.OnGreen("  SUCCESS  "))
	fmt.Printf("On Blue: %s\n", color.OnBlue("  INFO  "))
	fmt.Printf("On Yellow: %s\n", color.OnYellow("  WARNING  "))
	fmt.Printf("On Magenta: %s\n", color.OnMagenta("  SPECIAL  "))
	fmt.Printf("On Cyan: %s\n", color.OnCyan("  NOTICE  "))

	// Colors with Different Data Types
	fmt.Println(color.Green("\n🔢 Colors with Different Types:"))
	fmt.Printf("String: %s\n", color.Red("text"))
	fmt.Printf("Integer: %s\n", color.Green(42))
	fmt.Printf("Float: %s\n", color.Blue(3.14159))
	fmt.Printf("Boolean: %s\n", color.Yellow(true))
	fmt.Printf("Expression: %s\n", color.Magenta(10*5))
	fmt.Printf("Complex: %s\n", color.Cyan([]int{1, 2, 3}))

	// Using with fmt functions
	fmt.Println(color.Green("\n📝 Using with fmt Functions:"))

	// fmt.Print
	fmt.Print("fmt.Print: ")
	fmt.Print(color.Red("Red "))
	fmt.Print(color.Green("Green "))
	fmt.Print(color.Blue("Blue"))
	fmt.Println()

	// fmt.Println
	fmt.Println("fmt.Println:", color.Yellow("Yellow text"))

	// fmt.Printf
	name := "Alice"
	age := 25
	fmt.Printf("fmt.Printf: Hello %s, you are %s years old!\n",
		color.Cyan(name), color.Magenta(age))

	// Practical Examples
	fmt.Println(color.Green("\n🛠️  Practical Examples:"))

	// Status messages
	fmt.Printf("✅ %s\n", color.BoldGreen("Operation completed successfully"))
	fmt.Printf("⚠️  %s\n", color.BoldYellow("Warning: Check your input"))
	fmt.Printf("❌ %s\n", color.BoldRed("Error: Something went wrong"))
	fmt.Printf("ℹ️  %s\n", color.BoldBlue("Info: Process started"))

	// Progress indicators
	fmt.Println(color.Green("\n📊 Progress Indicators:"))
	fmt.Printf("[%s%s%s] 33%%\n",
		color.Green("███"),
		color.White("      "),
		color.White(""))
	fmt.Printf("[%s%s%s] 66%%\n",
		color.Green("██████"),
		color.White("   "),
		color.White(""))
	fmt.Printf("[%s] 100%%\n", color.Green("█████████"))

	// Log levels
	fmt.Println(color.Green("\n📋 Log Levels:"))
	fmt.Printf("%s This is a debug message\n", color.Cyan("[DEBUG]"))
	fmt.Printf("%s This is an info message\n", color.Blue("[INFO]"))
	fmt.Printf("%s This is a warning message\n", color.Yellow("[WARN]"))
	fmt.Printf("%s This is an error message\n", color.Red("[ERROR]"))
	fmt.Printf("%s This is a fatal message\n", color.BoldRed("[FATAL]"))

	// Highlighting code
	fmt.Println(color.Green("\n💻 Code Highlighting:"))
	fmt.Printf("%s %s%s%s {\n",
		color.Blue("func"),
		color.Yellow("main"),
		color.White("()"),
		color.White(""))
	fmt.Printf("    %s.%s(%s)\n",
		color.Green("fmt"),
		color.Yellow("Println"),
		color.Red("\"Hello, World!\""))
	fmt.Printf("}\n")

	// Table-like output
	fmt.Println(color.Green("\n📊 Table Output:"))
	fmt.Printf("| %-15s | %-10s | %-8s |\n",
		color.BoldCyan("Name"),
		color.BoldCyan("Age"),
		color.BoldCyan("Status"))
	fmt.Printf("| %-15s | %-10s | %-8s |\n",
		color.White("Alice"),
		color.Yellow("25"),
		color.Green("Active"))
	fmt.Printf("| %-15s | %-10s | %-8s |\n",
		color.White("Bob"),
		color.Yellow("30"),
		color.Red("Inactive"))

	fmt.Println(color.BoldMagenta("\n🎉 That's the power of fmtpy colors!"))
}
