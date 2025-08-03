# fmt2

# fmt2

A simplified input/output helper package for Go beginners, inspired by Python's simplicity.

## Installation

```bash
go get github.com/GrandpaEJ/fmt2
```

## Usage

```go
package main

import "github.com/GrandpaEJ/fmt2"

func main() {
    // Get input from user
    name := fmt2.Input("Enter your name: ")
    age := fmt2.Input("Enter your age: 
")  // With newline

    // Print using Go style formatting
    fmt2.Print("Your name is %s and age is %s", name, age)

    // Print using Python-style formatting
    fmt2.Print(f"Your name is {name} and age is {age}")

    // Simple printing
    fmt2.Print("Hello World!")
    fmt2.Print(42)
}
```

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
