package fmtpy

import (
	"testing"
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
