# fmtpy

A simplified input/output helper package for Go beginners, inspired by Python's simplicity.

## Installation

```bash
go get github.com/grandpaej/fmtpy/v2@v2.0.0
```

## Features

### Basic Features
- **Smart Input function** with automatic type conversion
- Print function supporting both Go-style and Python-style formatting
- **Modular package structure** for clean imports
- **Color functions** that work seamlessly with `fmt.Print`, `fmt.Println`, `fmt.Printf`
- Comprehensive string manipulation functions
- No performance loss compared to standard `fmt` package
- Beginner-friendly API

### Package Structure

```go
import (
    "github.com/grandpaej/fmtpy/v2"        // Core input/output
    "github.com/grandpaej/fmtpy/v2/color" // Color functions
    "github.com/grandpaej/fmtpy/v2/input" // String manipulation
)
```

### Smart Input with Type Conversion

The new `Input()` function returns an `InputValue` that can be converted to any type:

```go
package main

import (
    "fmt"
    "github.com/grandpaej/fmtpy/v2"
)

func main() {
    // Direct type conversion - exactly what you wanted!
    var age int
    age = fmtpy.Input("Enter your age: ").Int()

    var score float64
    score = fmtpy.Input("Enter your score: ").Float()

    name := fmtpy.Input("Enter your name: ").String()
    confirmed := fmtpy.Input("Confirm (y/n): ").Bool()

    fmt.Printf("Name: %s, Age: %d, Score: %.2f, Confirmed: %t\n",
               name, age, score, confirmed)
}
```

### Colors That Work with fmt

```go
package main

import (
    "fmt"
    "github.com/grandpaej/fmtpy/v2/color"
)

func main() {
    // Works with any type and any fmt function!
    fmt.Println(color.Red("Error message"))
    fmt.Println(color.Green(42))           // Numbers
    fmt.Println(color.Blue(1+1))           // Expressions
    fmt.Printf("Result: %s\n", color.Yellow(123.45))

    // Bold colors
    fmt.Println(color.BoldRed("Important!"))
    fmt.Println(color.BoldGreen("Success!"))

    // Background colors
    fmt.Println(color.OnYellow("Warning"))
    fmt.Println(color.OnRed("Critical"))
}
```

<img width="684" height="561" alt="image" src="https://github.com/user-attachments/assets/ab0f4abb-0229-49b5-a1e2-8a88a4cd963f" />

### Color Support
```go
package main

import (
    "fmt"
    "github.com/grandpaej/fmtpy/v2/color"
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
- `Ask(prompt)` - Get yes/no input (returns bool)

### String Manipulation Functions

#### Case Conversion
- `Upper(s)` - Convert to uppercase
- `Lower(s)` - Convert to lowercase
- `Title(s)` - Convert to title case (first letter of each word)
- `Capitalize(s)` - Capitalize first letter, lowercase the rest
- `SwapCase(s)` - Swap case of all letters

#### String Operations
- `Reverse(s)` - Reverse a string
- `Repeat(s, n)` - Repeat string n times
- `Replace(s, old, new)` - Replace all occurrences
- `Split(s, sep)` - Split string by separator
- `Join(parts, sep)` - Join string slice with separator
- `Trim(s)` - Remove leading/trailing whitespace

#### String Analysis
- `Contains(s, substring)` - Check if string contains substring
- `StartsWith(s, prefix)` - Check if string starts with prefix
- `EndsWith(s, suffix)` - Check if string ends with suffix
- `Length(s)` - Get character count (Unicode-aware)
- `WordCount(s)` - Count words in string
- `CharCount(s)` - Count characters in string
- `Count(s, substring)` - Count occurrences of substring

#### String Extraction
- `Substring(s, start, end)` - Extract substring
- `FirstWord(s)` - Get first word
- `LastWord(s)` - Get last word
- `Lines(s)` - Split into lines
- `Words(s)` - Split into words

#### String Filtering
- `OnlyLetters(s)` - Keep only alphabetic characters
- `OnlyNumbers(s)` - Keep only numeric characters
- `RemoveSpaces(s)` - Remove all spaces

#### String Validation
- `IsEmpty(s)` - Check if empty or whitespace only
- `IsNumber(s)` - Check if string is a valid number
- `IsLetter(s)` - Check if string contains only letters

#### Type Conversion
- `ToInt(s)` - Convert to int (returns 0 if invalid)
- `ToFloat(s)` - Convert to float64 (returns 0.0 if invalid)
- `ToString(v)` - Convert any value to string

#### String Formatting
- `Center(s, width)` - Center string within width
- `LeftPad(s, width, char)` - Pad left with character
- `RightPad(s, width, char)` - Pad right with character

### String Manipulation Examples
```go
package main

import (
    "fmt"
    "github.com/grandpaej/fmtpy/v2"
    "github.com/grandpaej/fmtpy/v2/input"
)

func main() {
    text := "Hello World"

    // Case conversion
    fmt.Println(input.Upper(text))      // "HELLO WORLD"
    fmt.Println(input.Lower(text))      // "hello world"
    fmt.Println(input.Title(text))      // "Hello World"
    fmt.Println(input.SwapCase(text))   // "hELLO wORLD"

    // String operations
    fmt.Println(input.Reverse(text))    // "dlroW olleH"
    fmt.Println(input.Repeat("Hi", 3))  // "HiHiHi"

    // Analysis
    fmt.Println(input.WordCount(text))  // 2
    fmt.Println(input.Length(text))     // 11
    fmt.Println(input.FirstWord(text))  // "Hello"

    // Filtering
    mixed := "Hello123World"
    fmt.Println(input.OnlyLetters(mixed))  // "HelloWorld"
    fmt.Println(input.OnlyNumbers(mixed))  // "123"

    // Formatting
    fmt.Println(input.Center("Hi", 10))           // "    Hi    "
    fmt.Println(input.LeftPad("Hi", 5, "*"))     // "***Hi"
    fmt.Println(input.RightPad("Hi", 5, "-"))    // "Hi---"

    // Smart input with type conversion
    var age int = fmtpy.Input("Enter your age: ").Int()
    var score float64 = fmtpy.Input("Enter your score: ").Float()

    // Yes/no questions
    if fmtpy.Input("Do you like Go? (y/n): ").Bool() {
        fmt.Println("Great choice!")
    }
}
```

### Complete Example

```go
package main

import (
    "fmt"
    "github.com/grandpaej/fmtpy/v2"
    "github.com/grandpaej/fmtpy/v2/color"
    "github.com/grandpaej/fmtpy/v2/input"
)

func main() {
    // Get user input with automatic type conversion
    name := fmtpy.Input("What's your name? ").String()
    var age int = fmtpy.Input("How old are you? ").Int()
    var score float64 = fmtpy.Input("Enter your score: ").Float()

    // Use colors with fmt functions
    fmt.Println(color.Green("=== User Profile ==="))
    fmt.Printf("Name: %s\n", color.Blue(name))
    fmt.Printf("Age: %s\n", color.Cyan(age))
    fmt.Printf("Score: %s\n", color.Yellow(score))

    // String manipulation
    fmt.Printf("Name in uppercase: %s\n", color.BoldRed(input.Upper(name)))
    fmt.Printf("Reversed name: %s\n", color.Magenta(input.Reverse(name)))

    // Combine everything
    if fmtpy.Input("Save profile? (y/n): ").Bool() {
        fmt.Println(color.BoldGreen("✅ Profile saved successfully!"))
    } else {
        fmt.Println(color.BoldYellow("⚠️  Profile not saved"))
    }
}
```

## 📚 Examples

The `example/` directory contains comprehensive examples:

- **`basic_usage.go`** - Simple input, colors, and string manipulation
- **`advanced_features.go`** - Text processing, validation, and interactive menus
- **`color_showcase.go`** - Complete color demonstration with fmt functions
- **`string_manipulation.go`** - All string manipulation features
- **`main.go`** - Complete feature demonstration

Run any example:
```bash
cd example
go run basic_usage.go
go run color_showcase.go
```

## 🎯 Quick Reference

### Input with Type Conversion
```go
// The new way - exactly what you wanted!
var age int = fmtpy.Input("Age: ").Int()
var score float64 = fmtpy.Input("Score: ").Float()
name := fmtpy.Input("Name: ").String()
confirmed := fmtpy.Input("Confirm (y/n): ").Bool()
```

### Colors (work with any fmt function)
```go
fmt.Println(color.Red("Error"))           // Basic colors
fmt.Println(color.BoldGreen("Success"))   // Bold colors
fmt.Println(color.OnYellow("Warning"))    // Background colors
fmt.Printf("Number: %s\n", color.Cyan(42)) // Any data type
```

### String Manipulation
```go
input.Upper("hello")           // "HELLO"
input.Reverse("hello")         // "olleh"
input.WordCount("hello world") // 2
input.OnlyNumbers("abc123")    // "123"
input.Center("hi", 10)         // "    hi    "
```

## 📖 Complete API Reference

### Core Package (`fmtpy`)

#### Input Functions
- `Input(prompt) InputValue` - Smart input with type conversion
- `Print(format, args...)` - Enhanced print with Python f-string support

#### InputValue Methods
- `.String()` - Convert to string
- `.Int()` - Convert to int (returns 0 if invalid)
- `.Float()` - Convert to float64 (returns 0.0 if invalid)
- `.Bool()` - Convert to bool (y/yes/true/1 = true)

### Color Package (`fmtpy/color`)

#### Basic Colors
- `Red(v)`, `Green(v)`, `Blue(v)`, `Yellow(v)`, `Magenta(v)`, `Cyan(v)`, `White(v)`, `Black(v)`

#### Bold Colors
- `BoldRed(v)`, `BoldGreen(v)`, `BoldBlue(v)`, `BoldYellow(v)`, `BoldMagenta(v)`, `BoldCyan(v)`

#### Background Colors
- `OnRed(v)`, `OnGreen(v)`, `OnBlue(v)`, `OnYellow(v)`, `OnMagenta(v)`, `OnCyan(v)`

*All color functions accept any data type and work with `fmt.Print`, `fmt.Println`, `fmt.Printf`*

### Input Package (`fmtpy/input`)

#### Case Conversion
- `Upper(s)` - Convert to UPPERCASE
- `Lower(s)` - Convert to lowercase
- `Title(s)` - Convert To Title Case
- `Capitalize(s)` - Capitalize first letter only
- `SwapCase(s)` - sWAP cASE

#### String Operations
- `Reverse(s)` - Reverse string
- `Repeat(s, n)` - Repeat string n times
- `Replace(s, old, new)` - Replace all occurrences
- `Split(s, sep)` - Split by separator
- `Join(parts, sep)` - Join with separator
- `Trim(s)` - Remove leading/trailing whitespace

#### String Analysis
- `Contains(s, sub)` - Check if contains substring
- `StartsWith(s, prefix)` - Check if starts with prefix
- `EndsWith(s, suffix)` - Check if ends with suffix
- `Length(s)` - Character count (Unicode-aware)
- `WordCount(s)` - Count words
- `CharCount(s)` - Count characters
- `Count(s, sub)` - Count occurrences of substring

#### String Extraction
- `Substring(s, start, end)` - Extract substring
- `FirstWord(s)` - Get first word
- `LastWord(s)` - Get last word
- `Lines(s)` - Split into lines
- `Words(s)` - Split into words

#### String Filtering
- `OnlyLetters(s)` - Keep only alphabetic characters
- `OnlyNumbers(s)` - Keep only numeric characters
- `RemoveSpaces(s)` - Remove all spaces

#### String Validation
- `IsEmpty(s)` - Check if empty or whitespace only
- `IsNumber(s)` - Check if valid number
- `IsLetter(s)` - Check if contains only letters

#### Type Conversion
- `ToInt(s)` - Convert to int (returns 0 if invalid)
- `ToFloat(s)` - Convert to float64 (returns 0.0 if invalid)
- `ToString(v)` - Convert any value to string

#### String Formatting
- `Center(s, width)` - Center string within width
- `LeftPad(s, width, char)` - Pad left with character
- `RightPad(s, width, char)` - Pad right with character

## 🚀 Migration Guide

### From Old fmtpy
```go
// Old way
name := fmtpy.Input("Name: ")
age := fmtpy.InputInt("Age: ")

// New way (recommended)
name := fmtpy.Input("Name: ").String()
var age int = fmtpy.Input("Age: ").Int()
```

### From Standard Go
```go
// Standard Go
var age int
fmt.Print("Enter age: ")
fmt.Scanf("%d", &age)

// fmtpy way
var age int = fmtpy.Input("Enter age: ").Int()
```

## 🎨 Advanced Color Usage

### Custom Color Combinations
```go
// Legacy color functions (still supported)
fmt.Println(color.RedText("Red text with {placeholder}", "value"))
fmt.Println(color.BoldRedText("Bold red text"))

// New simple functions (recommended)
fmt.Println(color.Red("Simple red text"))
fmt.Println(color.BoldRed("Bold red text"))
```

### Color with Complex Data
```go
// Colors work with any data type
users := []string{"Alice", "Bob", "Charlie"}
fmt.Printf("Users: %s\n", color.Green(users))

result := map[string]int{"score": 95, "level": 5}
fmt.Printf("Result: %s\n", color.Blue(result))
```

## 🔧 Performance Notes

- **Zero speed abuse**: All functions are optimized for performance
- **Memory efficient**: No unnecessary allocations
- **Unicode support**: Proper handling of international characters
- **Safe conversions**: Type conversions never panic, return sensible defaults

## 🌍 Environment Variables

- `NO_COLOR`: Set this to disable colors
- `TERM=dumb`: Colors will be disabled automatically

## 🤝 Contributing

Feel free to open issues or submit pull requests!

### Development Setup
```bash
git clone https://github.com/grandpaej/fmtpy
cd fmtpy
go test ./...
```

## 📄 License

MIT License

## 🎯 Why fmtpy?

- **Beginner-friendly**: Simple, intuitive API
- **Type-safe**: Smart input conversion without panics
- **Modular**: Import only what you need
- **Compatible**: Works with standard `fmt` functions
- **Colorful**: Rich color support for better UX
- **Fast**: Zero performance overhead
- **Unicode**: Proper international character support

Perfect for:
- 🎓 Learning Go programming
- 🛠️ Building CLI tools
- 📊 Creating interactive programs
- 🎨 Adding colors to terminal output
- ⚡ Rapid prototyping
