package fmtpy

import (
	"testing"

	"github.com/grandpaej/fmtpy/v2/input"
)

func TestConvertPythonStyle(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "Hello {name}",
			expected: "Hello %v",
		},
		{
			input:    "Your name is {name} and age is {age}",
			expected: "Your name is %v and age is %v",
		},
		{
			input:    "No brackets here",
			expected: "No brackets here",
		},
	}

	for _, test := range tests {
		result := convertPythonStyle(test.input)
		if result != test.expected {
			t.Errorf("convertPythonStyle(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestStringManipulation(t *testing.T) {
	// Test Upper
	if input.Upper("hello") != "HELLO" {
		t.Errorf("Upper failed")
	}

	// Test Lower
	if input.Lower("HELLO") != "hello" {
		t.Errorf("Lower failed")
	}

	// Test Title
	if input.Title("hello world") != "Hello World" {
		t.Errorf("Title failed")
	}

	// Test Capitalize
	if input.Capitalize("hELLO wORLD") != "Hello world" {
		t.Errorf("Capitalize failed")
	}

	// Test Reverse
	if input.Reverse("hello") != "olleh" {
		t.Errorf("Reverse failed")
	}

	// Test Repeat
	if input.Repeat("hi", 3) != "hihihi" {
		t.Errorf("Repeat failed")
	}

	// Test Replace
	if input.Replace("hello world", "world", "Go") != "hello Go" {
		t.Errorf("Replace failed")
	}

	// Test Split and Join
	parts := input.Split("a,b,c", ",")
	if len(parts) != 3 || parts[0] != "a" {
		t.Errorf("Split failed")
	}
	if input.Join(parts, "-") != "a-b-c" {
		t.Errorf("Join failed")
	}

	// Test Trim
	if input.Trim("  hello  ") != "hello" {
		t.Errorf("Trim failed")
	}

	// Test Contains
	if !input.Contains("hello world", "world") {
		t.Errorf("Contains failed")
	}

	// Test StartsWith and EndsWith
	if !input.StartsWith("hello", "he") {
		t.Errorf("StartsWith failed")
	}
	if !input.EndsWith("hello", "lo") {
		t.Errorf("EndsWith failed")
	}

	// Test Length
	if input.Length("hello") != 5 {
		t.Errorf("Length failed")
	}

	// Test Substring
	if input.Substring("hello", 1, 4) != "ell" {
		t.Errorf("Substring failed")
	}

	// Test IsEmpty
	if !input.IsEmpty("   ") {
		t.Errorf("IsEmpty failed")
	}

	// Test ToInt
	if input.ToInt("123") != 123 {
		t.Errorf("ToInt failed")
	}
	if input.ToInt("abc") != 0 {
		t.Errorf("ToInt should return 0 for invalid input")
	}

	// Test ToFloat
	if input.ToFloat("123.45") != 123.45 {
		t.Errorf("ToFloat failed")
	}

	// Test ToString
	if input.ToString(123) != "123" {
		t.Errorf("ToString failed")
	}

	// Test Center
	if input.Center("hi", 6) != "  hi  " {
		t.Errorf("Center failed")
	}

	// Test LeftPad and RightPad
	if input.LeftPad("hi", 5, "0") != "000hi" {
		t.Errorf("LeftPad failed")
	}
	if input.RightPad("hi", 5, "0") != "hi000" {
		t.Errorf("RightPad failed")
	}
}

func TestUtilityFunctions(t *testing.T) {
	// Test Lines
	lines := input.Lines("line1\nline2\nline3")
	if len(lines) != 3 || lines[0] != "line1" {
		t.Errorf("Lines failed")
	}

	// Test Words
	words := input.Words("hello world test")
	if len(words) != 3 || words[1] != "world" {
		t.Errorf("Words failed")
	}

	// Test WordCount
	if input.WordCount("hello world") != 2 {
		t.Errorf("WordCount failed")
	}

	// Test CharCount
	if input.CharCount("hello") != 5 {
		t.Errorf("CharCount failed")
	}

	// Test RemoveSpaces
	if input.RemoveSpaces("hello world") != "helloworld" {
		t.Errorf("RemoveSpaces failed")
	}

	// Test OnlyLetters
	if input.OnlyLetters("hello123world") != "helloworld" {
		t.Errorf("OnlyLetters failed")
	}

	// Test OnlyNumbers
	if input.OnlyNumbers("hello123world456") != "123456" {
		t.Errorf("OnlyNumbers failed")
	}

	// Test IsNumber
	if !input.IsNumber("123.45") {
		t.Errorf("IsNumber failed for valid number")
	}
	if input.IsNumber("abc") {
		t.Errorf("IsNumber should return false for invalid number")
	}

	// Test IsLetter
	if !input.IsLetter("hello") {
		t.Errorf("IsLetter failed for valid letters")
	}
	if input.IsLetter("hello123") {
		t.Errorf("IsLetter should return false for mixed content")
	}

	// Test SwapCase
	if input.SwapCase("Hello World") != "hELLO wORLD" {
		t.Errorf("SwapCase failed")
	}

	// Test Count
	if input.Count("hello hello", "hello") != 2 {
		t.Errorf("Count failed")
	}

	// Test FirstWord and LastWord
	if input.FirstWord("hello world test") != "hello" {
		t.Errorf("FirstWord failed")
	}
	if input.LastWord("hello world test") != "test" {
		t.Errorf("LastWord failed")
	}
}

func TestInputValue(t *testing.T) {
	// Test InputValue type conversions
	iv := InputValue("123")
	if iv.Int() != 123 {
		t.Errorf("InputValue.Int() failed")
	}

	iv = InputValue("123.45")
	if iv.Float() != 123.45 {
		t.Errorf("InputValue.Float() failed")
	}

	iv = InputValue("hello")
	if iv.String() != "hello" {
		t.Errorf("InputValue.String() failed")
	}

	iv = InputValue("y")
	if !iv.Bool() {
		t.Errorf("InputValue.Bool() failed for 'y'")
	}

	iv = InputValue("no")
	if iv.Bool() {
		t.Errorf("InputValue.Bool() should return false for 'no'")
	}
}
