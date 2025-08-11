// Package fmtpy provides simplified input/output helpers for Go beginners
package fmtpy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// InputValue is a special type that can convert to different types
type InputValue string

// String returns the string value
func (iv InputValue) String() string {
	return string(iv)
}

// Int converts to integer, returns 0 if conversion fails
func (iv InputValue) Int() int {
	if i, err := strconv.Atoi(strings.TrimSpace(string(iv))); err == nil {
		return i
	}
	return 0
}

// Float converts to float64, returns 0.0 if conversion fails
func (iv InputValue) Float() float64 {
	if f, err := strconv.ParseFloat(strings.TrimSpace(string(iv)), 64); err == nil {
		return f
	}
	return 0.0
}

// Bool converts to boolean (y/yes/true/1 = true, everything else = false)
func (iv InputValue) Bool() bool {
	s := strings.ToLower(strings.TrimSpace(string(iv)))
	return s == "y" || s == "yes" || s == "true" || s == "1"
}

// Input prompts the user for input and returns an InputValue that can be converted to any type
// Usage examples:
//
//	name := Input("Enter name: ").String()
//	var age int = Input("Enter age: ").Int()
//	score := Input("Enter score: ").Float()
//	confirmed := Input("Confirm (y/n): ").Bool()
func Input(prompt string) InputValue {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return InputValue(strings.TrimSpace(text))
}

// Print formats and prints the given values.
// It supports:
// 1. Go-style formatting: Print("Hello %s", name)
// 2. Python-style f-strings: Print(f"Hello {name}")
// 3. Multiple colored strings: Print(red("Error:"), yellow("Score: {score}"))
func Print(format interface{}, args ...interface{}) {
	switch v := format.(type) {
	case string:
		if strings.HasPrefix(v, "f\"") || strings.HasPrefix(v, "f'") {
			// Python-style formatting
			v = v[2 : len(v)-1] // Remove f" and "
			v = convertPythonStyle(v)
			fmt.Printf(v, args...)
		} else if len(args) > 0 {
			if len(args) == 1 {
				// Check if it's a simple string with color
				fmt.Printf(v, args[0])
			} else {
				parts := make([]string, 0, len(args)+1)
				parts = append(parts, v)
				for _, arg := range args {
					if s, ok := arg.(string); ok {
						parts = append(parts, s)
					} else {
						parts = append(parts, fmt.Sprint(arg))
					}
				}
				fmt.Print(strings.Join(parts, " "))
			}
		} else {
			fmt.Print(v)
		}
	default:
		fmt.Print(v)
	}
	fmt.Println()
}

// convertPythonStyle converts Python-style format strings to Go format strings
// Example: "Hello {name}" -> "Hello %v"
func convertPythonStyle(s string) string {
	var result strings.Builder
	inBrace := false

	for i := 0; i < len(s); i++ {
		if s[i] == '{' {
			inBrace = true
			result.WriteString("%v")
			continue
		}
		if s[i] == '}' {
			inBrace = false
			continue
		}
		if !inBrace {
			result.WriteByte(s[i])
		}
	}

	return result.String()
}
