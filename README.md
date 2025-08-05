# fmtpy

A simplified input/output helper package for Go beginners, inspired by Python's simplicity.

## Installation

```bash
go get github.com/grandpaej/fmtpy
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/grandpaej/fmtpy"
)

func main() {
	// Get input from user
	name := fmtpy.Input("Enter your name: ")
	age := fmtpy.Input("Enter your age: \n") // With newline

	// Print using Python-style formatting
	fmtpy.Print("f\"Your name is {name} and age is {age}\"", name, age)

	// Simple printing
	fmtpy.Print("Hello World!")
	fmtpy.Print(42)

	// Input type
	fmt.Printf("Type of name: %T\n", name) // Output: Type of name: string
	fmt.Printf("Type of age: %T\n", age)   // Output: Type of age: string
}

```
<img width="684" height="561" alt="image" src="https://github.com/user-attachments/assets/ab0f4abb-0229-49b5-a1e2-8a88a4cd963f" />

## Features

### Basic Features
- Simple input function with optional newline prompts
- Print function supporting both Go-style and Python-style formatting
- No performance loss compared to standard `fmt` package
- Beginner-friendly API

### Type-Specific Input
- `InputInt(prompt)` - Get integer input
- `InputFloat(prompt)` - Get float input
- `InputBool(prompt)` - Get yes/no input
- `InputWithDefault(prompt, default)` - Get input with default value


## Contributing

Feel free to open issues or submit pull requests!

## License

MIT License
