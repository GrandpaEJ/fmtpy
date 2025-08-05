# fmtpy

A simplified input/output helper package for Go beginners, inspired by Python's simplicity.

## Installation

```bash
go get github.com/grandpaej/fmtpy
```

## Features

### Basic Features
- Simple input function with optional newline prompts
- Print function supporting both Go-style and Python-style formatting
- Rich color support with simple API
- No performance loss compared to standard `fmt` package
- Beginner-friendly API

### Input/Output
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
}
```

<img width="684" height="561" alt="image" src="https://github.com/user-attachments/assets/ab0f4abb-0229-49b5-a1e2-8a88a4cd963f" />

### Color Support
```go
package main

import (
    "fmt"
    "github.com/grandpaej/fmtpy/color"
)

func main() {
    c := color.Get()
    
    // Basic colors - Short form
    fmt.Println(c.R("Red text"))
    fmt.Println(c.G("Green text"))
    fmt.Println(c.B("Blue text"))
    
    // Bold colors
    fmt.Println(c.BR("Bold red"))
    fmt.Println(c.BG("Bold green"))
    
    // Background colors
    fmt.Println(c.OnR("Red background"))
    fmt.Println(c.OnY("Yellow background"))
    
    // Semantic styles
    fmt.Println(c.E("This is an error"))
    fmt.Println(c.S("This is a success"))
    fmt.Println(c.I("This is info"))
    fmt.Println(c.W("This is a warning"))
    
    // With formatting
    name := "World"
    fmt.Println(c.S("Hello, {n}!", name))
}
```

### Available Color Methods

#### Short Form Colors
- `R()` - Red text
- `G()` - Green text
- `B()` - Blue text
- `Y()` - Yellow text
- `M()` - Magenta text
- `C()` - Cyan text

#### Bold Colors
- `BR()` - Bold red
- `BG()` - Bold green
- `BB()` - Bold blue
- `BY()` - Bold yellow
- `BM()` - Bold magenta
- `BC()` - Bold cyan

#### Background Colors
- `OnR()` - Red background
- `OnG()` - Green background
- `OnB()` - Blue background
- `OnY()` - Yellow background
- `OnM()` - Magenta background
- `OnC()` - Cyan background

#### Semantic Helpers
- `E()` - Error (Red)
- `S()` - Success (Green)
- `I()` - Info (Blue)
- `W()` - Warning (Yellow)

### Type-Specific Input
- `InputInt(prompt)` - Get integer input
- `InputFloat(prompt)` - Get float input
- `InputBool(prompt)` - Get yes/no input
- `InputWithDefault(prompt, default)` - Get input with default value

## Advanced Usage

### Custom Color Combinations
```go
c := color.Get()

// Mix foreground and background
fmt.Printf("%s on %s\n",
    c.BR("Bold Red"),
    c.OnY("Yellow Background"))

// Error message with custom formatting
name := "test.txt"
fmt.Println(c.E("Error: Could not open file {n}", name))
```

### Traditional Color Functions
```go
// Using traditional color functions
fmt.Println(color.RedText("Red text"))
fmt.Println(color.GreenText("Green text"))
fmt.Println(color.BoldRedText("Bold red text"))
```

## Environment Variables

- `NO_COLOR`: Set this to disable colors
- `TERM=dumb`: Colors will be disabled automatically

## Contributing

Feel free to open issues or submit pull requests!

## License

MIT License
