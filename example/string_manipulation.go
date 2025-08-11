package main

import (
	"fmt"
	"github.com/grandpaej/fmtpy"
	"github.com/grandpaej/fmtpy/color"
	"github.com/grandpaej/fmtpy/input"
)

func main() {
	fmt.Println(color.BoldGreen("=== String Manipulation Showcase ==="))

	// Get sample text
	text := fmtpy.Input("Enter some text to manipulate: ").String()
	
	// Case Transformations
	fmt.Println(color.Cyan("\n🔤 Case Transformations:"))
	fmt.Printf("Original: %s\n", color.White(text))
	fmt.Printf("Upper: %s\n", color.Red(input.Upper(text)))
	fmt.Printf("Lower: %s\n", color.Blue(input.Lower(text)))
	fmt.Printf("Title: %s\n", color.Green(input.Title(text)))
	fmt.Printf("Capitalize: %s\n", color.Yellow(input.Capitalize(text)))
	fmt.Printf("SwapCase: %s\n", color.Magenta(input.SwapCase(text)))

	// String Operations
	fmt.Println(color.Cyan("\n🔧 String Operations:"))
	fmt.Printf("Reverse: %s\n", color.Red(input.Reverse(text)))
	fmt.Printf("Repeat 3x: %s\n", color.Green(input.Repeat(text, 3)))
	fmt.Printf("No spaces: %s\n", color.Blue(input.RemoveSpaces(text)))
	
	// String Analysis
	fmt.Println(color.Cyan("\n📊 String Analysis:"))
	fmt.Printf("Length: %s characters\n", color.Yellow(input.Length(text)))
	fmt.Printf("Word count: %s words\n", color.Green(input.WordCount(text)))
	fmt.Printf("Character count: %s chars\n", color.Blue(input.CharCount(text)))
	
	if input.WordCount(text) > 0 {
		fmt.Printf("First word: %s\n", color.Red(input.FirstWord(text)))
		fmt.Printf("Last word: %s\n", color.Magenta(input.LastWord(text)))
	}

	// String Filtering
	fmt.Println(color.Cyan("\n🔍 String Filtering:"))
	fmt.Printf("Only letters: %s\n", color.Green(input.OnlyLetters(text)))
	fmt.Printf("Only numbers: %s\n", color.Yellow(input.OnlyNumbers(text)))
	
	// String Validation
	fmt.Println(color.Cyan("\n✅ String Validation:"))
	fmt.Printf("Is empty: %s\n", color.Blue(input.IsEmpty(text)))
	fmt.Printf("Is number: %s\n", color.Green(input.IsNumber(text)))
	fmt.Printf("Is letters only: %s\n", color.Yellow(input.IsLetter(text)))
	
	// String Searching
	fmt.Println(color.Cyan("\n🔎 String Searching:"))
	search_term := fmtpy.Input("Enter a word to search for: ").String()
	
	contains := input.Contains(text, search_term)
	count := input.Count(text, search_term)
	starts_with := input.StartsWith(text, search_term)
	ends_with := input.EndsWith(text, search_term)
	
	fmt.Printf("Contains '%s': %s\n", search_term, color.Green(contains))
	fmt.Printf("Count of '%s': %s\n", search_term, color.Yellow(count))
	fmt.Printf("Starts with '%s': %s\n", search_term, color.Blue(starts_with))
	fmt.Printf("Ends with '%s': %s\n", search_term, color.Magenta(ends_with))

	// String Formatting
	fmt.Println(color.Cyan("\n📐 String Formatting:"))
	width := fmtpy.Input("Enter formatting width: ").Int()
	
	if width > 0 {
		fmt.Printf("Centered: '%s'\n", color.Green(input.Center(text, width)))
		fmt.Printf("Left padded with '-': '%s'\n", color.Yellow(input.LeftPad(text, width, "-")))
		fmt.Printf("Right padded with '*': '%s'\n", color.Blue(input.RightPad(text, width, "*")))
	}

	// String Splitting and Joining
	fmt.Println(color.Cyan("\n✂️  String Splitting:"))
	
	if input.Contains(text, " ") {
		words := input.Split(text, " ")
		fmt.Printf("Split by spaces: %s\n", color.Green(fmt.Sprintf("%v", words)))
		
		joined := input.Join(words, " | ")
		fmt.Printf("Joined with ' | ': %s\n", color.Yellow(joined))
	}
	
	lines := input.Lines(text)
	if len(lines) > 1 {
		fmt.Printf("Lines: %s\n", color.Blue(fmt.Sprintf("%v", lines)))
	}
	
	words := input.Words(text)
	fmt.Printf("Words: %s\n", color.Magenta(fmt.Sprintf("%v", words)))

	// String Replacement
	fmt.Println(color.Cyan("\n🔄 String Replacement:"))
	old_text := fmtpy.Input("Enter text to replace: ").String()
	new_text := fmtpy.Input("Enter replacement text: ").String()
	
	if input.Contains(text, old_text) {
		replaced := input.Replace(text, old_text, new_text)
		fmt.Printf("Original: %s\n", color.Red(text))
		fmt.Printf("Replaced: %s\n", color.Green(replaced))
	} else {
		fmt.Printf("'%s' not found in the text\n", color.Yellow(old_text))
	}

	// Substring Operations
	fmt.Println(color.Cyan("\n✂️  Substring Operations:"))
	if input.Length(text) > 3 {
		start := fmtpy.Input("Enter start position: ").Int()
		end := fmtpy.Input("Enter end position: ").Int()
		
		if start >= 0 && end <= input.Length(text) && start < end {
			substring := input.Substring(text, start, end)
			fmt.Printf("Substring [%d:%d]: %s\n", start, end, color.Green(substring))
		} else {
			fmt.Println(color.Red("Invalid positions"))
		}
	}

	// Type Conversion Examples
	fmt.Println(color.Cyan("\n🔢 Type Conversion:"))
	
	number_str := fmtpy.Input("Enter a number: ").String()
	as_int := input.ToInt(number_str)
	as_float := input.ToFloat(number_str)
	
	fmt.Printf("As string: %s\n", color.White(number_str))
	fmt.Printf("As int: %s\n", color.Green(as_int))
	fmt.Printf("As float: %s\n", color.Blue(as_float))
	
	// Convert other types to string
	fmt.Printf("Boolean to string: %s\n", color.Yellow(input.ToString(true)))
	fmt.Printf("Number to string: %s\n", color.Magenta(input.ToString(42)))

	fmt.Println(color.BoldGreen("\n🎉 String manipulation complete!"))
}
