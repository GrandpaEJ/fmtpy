package main

import (
	"fmt"

	"github.com/grandpaej/fmtpy/v2"
	"github.com/grandpaej/fmtpy/v2/color"
	"github.com/grandpaej/fmtpy/v2/input"
)

func main() {
	fmt.Println(color.BoldMagenta("=== Advanced fmtpy Features ==="))

	// Advanced Example 1: Text Processing Pipeline
	fmt.Println(color.Green("\n1. Text Processing Pipeline:"))

	sentence := fmtpy.Input("Enter a sentence: ").String()

	// Process the text step by step
	cleaned := input.Trim(sentence)
	wordCount := input.WordCount(cleaned)
	charCount := input.CharCount(cleaned)
	firstWord := input.FirstWord(cleaned)
	lastWord := input.LastWord(cleaned)

	fmt.Printf("📝 Original: %s\n", color.Blue(sentence))
	fmt.Printf("🧹 Cleaned: %s\n", color.Green(cleaned))
	fmt.Printf("📊 Stats: %s words, %s characters\n",
		color.Cyan(wordCount), color.Yellow(charCount))
	fmt.Printf("🔤 First word: %s | Last word: %s\n",
		color.Red(firstWord), color.Magenta(lastWord))

	// Advanced Example 2: Data Validation
	fmt.Println(color.Green("\n2. Input Validation:"))

	email := fmtpy.Input("Enter your email: ").String()

	if input.Contains(email, "@") && input.Contains(email, ".") {
		fmt.Println(color.BoldGreen("✅ Email format looks valid"))
	} else {
		fmt.Println(color.BoldRed("❌ Invalid email format"))
	}

	// Advanced Example 3: String Filtering and Extraction
	fmt.Println(color.Green("\n3. String Filtering:"))

	mixed_input := fmtpy.Input("Enter text with numbers and letters: ").String()

	letters_only := input.OnlyLetters(mixed_input)
	numbers_only := input.OnlyNumbers(mixed_input)
	no_spaces := input.RemoveSpaces(mixed_input)

	fmt.Printf("🔤 Letters only: %s\n", color.Blue(letters_only))
	fmt.Printf("🔢 Numbers only: %s\n", color.Green(numbers_only))
	fmt.Printf("🚫 No spaces: %s\n", color.Yellow(no_spaces))

	// Advanced Example 4: Text Formatting and Alignment
	fmt.Println(color.Green("\n4. Text Formatting:"))

	title := fmtpy.Input("Enter a title: ").String()
	width := fmtpy.Input("Enter display width: ").Int()

	centered := input.Center(title, width)
	left_padded := input.LeftPad(title, width, "-")
	right_padded := input.RightPad(title, width, "*")

	fmt.Printf("📍 Centered: '%s'\n", color.Cyan(centered))
	fmt.Printf("⬅️  Left padded: '%s'\n", color.Yellow(left_padded))
	fmt.Printf("➡️  Right padded: '%s'\n", color.Magenta(right_padded))

	// Advanced Example 5: Case Transformations
	fmt.Println(color.Green("\n5. Case Transformations:"))

	text := fmtpy.Input("Enter mixed case text: ").String()

	fmt.Printf("🔼 UPPER: %s\n", color.Red(input.Upper(text)))
	fmt.Printf("🔽 lower: %s\n", color.Blue(input.Lower(text)))
	fmt.Printf("📖 Title Case: %s\n", color.Green(input.Title(text)))
	fmt.Printf("🎯 Capitalized: %s\n", color.Yellow(input.Capitalize(text)))
	fmt.Printf("🔄 SwapCase: %s\n", color.Magenta(input.SwapCase(text)))

	// Advanced Example 6: Interactive Menu System
	fmt.Println(color.Green("\n6. Interactive Menu:"))

	for {
		fmt.Println(color.BoldCyan("\n--- Text Processor Menu ---"))
		fmt.Println("1. Reverse text")
		fmt.Println("2. Count characters")
		fmt.Println("3. Extract numbers")
		fmt.Println("4. Exit")

		choice := fmtpy.Input("Choose an option (1-4): ").Int()

		switch choice {
		case 1:
			text := fmtpy.Input("Enter text to reverse: ").String()
			fmt.Printf("Reversed: %s\n", color.Green(input.Reverse(text)))
		case 2:
			text := fmtpy.Input("Enter text to count: ").String()
			fmt.Printf("Character count: %s\n", color.Cyan(input.CharCount(text)))
		case 3:
			text := fmtpy.Input("Enter text with numbers: ").String()
			fmt.Printf("Numbers found: %s\n", color.Yellow(input.OnlyNumbers(text)))
		case 4:
			fmt.Println(color.BoldGreen("👋 Goodbye!"))
			return
		default:
			fmt.Println(color.BoldRed("❌ Invalid choice. Please try again."))
		}
	}
}
