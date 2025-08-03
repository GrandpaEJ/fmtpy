// Package fmt2 provides simplified input/output helpers for Go beginners
package fmt2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Input prompts the user for input and returns the entered string
// If the prompt ends with \n, it will be printed on a new line
func Input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
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
