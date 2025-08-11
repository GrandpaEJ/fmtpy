// Package input provides simplified input helpers for Go beginners
package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var reader = bufio.NewReader(os.Stdin)

// Input prompts the user for input and returns the entered string
// If the prompt ends with \n, it will be printed on a new line
func Input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// InputInt prompts for input and converts to integer
func InputInt(prompt string) int {
	return ToInt(Input(prompt))
}

// InputFloat prompts for input and converts to float
func InputFloat(prompt string) float64 {
	return ToFloat(Input(prompt))
}

// Ask prompts for yes/no input and returns boolean
func Ask(prompt string) bool {
	response := Lower(Trim(Input(prompt + " (y/n): ")))
	return response == "y" || response == "yes" || response == "true" || response == "1"
}

// String manipulation functions for easy text processing

// Upper converts a string to uppercase
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Lower converts a string to lowercase
func Lower(s string) string {
	return strings.ToLower(s)
}

// Title converts a string to title case (first letter of each word capitalized)
func Title(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}

// Capitalize converts the first character to uppercase and the rest to lowercase
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Repeat repeats a string n times
func Repeat(s string, n int) string {
	return strings.Repeat(s, n)
}

// Replace replaces all occurrences of old with new in string s
func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Split splits a string by separator
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Join joins a slice of strings with separator
func Join(parts []string, sep string) string {
	return strings.Join(parts, sep)
}

// Trim removes leading and trailing whitespace
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// Contains checks if string s contains substring
func Contains(s, substring string) bool {
	return strings.Contains(s, substring)
}

// StartsWith checks if string s starts with prefix
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// EndsWith checks if string s ends with suffix
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Length returns the length of a string (number of characters, not bytes)
func Length(s string) int {
	return len([]rune(s))
}

// Substring returns a substring from start to end (exclusive)
func Substring(s string, start, end int) string {
	runes := []rune(s)
	if start < 0 {
		start = 0
	}
	if end > len(runes) {
		end = len(runes)
	}
	if start >= end {
		return ""
	}
	return string(runes[start:end])
}

// IsEmpty checks if a string is empty or contains only whitespace
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// ToInt converts a string to integer, returns 0 if conversion fails
func ToInt(s string) int {
	if i, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
		return i
	}
	return 0
}

// ToFloat converts a string to float64, returns 0.0 if conversion fails
func ToFloat(s string) float64 {
	if f, err := strconv.ParseFloat(strings.TrimSpace(s), 64); err == nil {
		return f
	}
	return 0.0
}

// ToString converts any value to string
func ToString(v interface{}) string {
	return fmt.Sprint(v)
}

// Center centers a string within a given width, padding with spaces
func Center(s string, width int) string {
	if len(s) >= width {
		return s
	}
	padding := width - len(s)
	leftPad := padding / 2
	rightPad := padding - leftPad
	return strings.Repeat(" ", leftPad) + s + strings.Repeat(" ", rightPad)
}

// LeftPad pads a string on the left to reach the specified width
func LeftPad(s string, width int, padChar string) string {
	if len(padChar) == 0 {
		padChar = " "
	}
	if len(s) >= width {
		return s
	}
	padding := width - len(s)
	return strings.Repeat(padChar, padding) + s
}

// RightPad pads a string on the right to reach the specified width
func RightPad(s string, width int, padChar string) string {
	if len(padChar) == 0 {
		padChar = " "
	}
	if len(s) >= width {
		return s
	}
	padding := width - len(s)
	return s + strings.Repeat(padChar, padding)
}

// Additional utility functions

// Lines splits text into lines
func Lines(s string) []string {
	return strings.Split(s, "\n")
}

// Words splits text into words
func Words(s string) []string {
	return strings.Fields(s)
}

// WordCount returns the number of words in a string
func WordCount(s string) int {
	return len(Words(s))
}

// CharCount returns the number of characters in a string
func CharCount(s string) int {
	return Length(s)
}

// RemoveSpaces removes all spaces from a string
func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// OnlyLetters returns only alphabetic characters from a string
func OnlyLetters(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// OnlyNumbers returns only numeric characters from a string
func OnlyNumbers(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// IsNumber checks if a string represents a valid number
func IsNumber(s string) bool {
	_, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	return err == nil
}

// IsLetter checks if a string contains only letters
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return len(s) > 0
}

// SwapCase swaps the case of all letters in a string
func SwapCase(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsUpper(r) {
			result.WriteRune(unicode.ToLower(r))
		} else if unicode.IsLower(r) {
			result.WriteRune(unicode.ToUpper(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// Count counts occurrences of substring in string
func Count(s, substring string) int {
	return strings.Count(s, substring)
}

// FirstWord returns the first word from a string
func FirstWord(s string) string {
	words := Words(s)
	if len(words) > 0 {
		return words[0]
	}
	return ""
}

// LastWord returns the last word from a string
func LastWord(s string) string {
	words := Words(s)
	if len(words) > 0 {
		return words[len(words)-1]
	}
	return ""
}
