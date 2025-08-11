# fmtpy Examples

This directory contains comprehensive examples demonstrating all fmtpy features.

## 📁 Example Files

### 1. `basic_usage.go`
**Perfect for beginners** - Shows fundamental features:
- Smart input with type conversion
- Basic color usage
- Simple string manipulation
- Boolean input handling

**Run it:**
```bash
go run basic_usage.go
```

### 2. `color_showcase.go`
**Complete color demonstration** - Shows all color features:
- Basic colors (Red, Green, Blue, etc.)
- Bold colors
- Background colors
- Colors with different data types
- Integration with fmt functions
- Practical examples (progress bars, log levels, tables)

**Run it:**
```bash
go run color_showcase.go
```

### 3. `string_manipulation.go`
**String processing powerhouse** - Demonstrates:
- Case transformations
- String analysis and validation
- Text filtering and searching
- String formatting and alignment
- Substring operations
- Type conversions

**Run it:**
```bash
go run string_manipulation.go
```

### 4. `advanced_features.go`
**Advanced techniques** - Shows:
- Text processing pipelines
- Input validation patterns
- Interactive menu systems
- Complex string operations
- Real-world usage patterns

**Run it:**
```bash
go run advanced_features.go
```

### 5. `real_world_app.go`
**Complete application** - A full user profile manager featuring:
- Data validation
- Interactive menus
- Profile statistics
- Error handling
- Professional UI with colors

**Run it:**
```bash
go run real_world_app.go
```

### 6. `main.go`
**Feature overview** - Comprehensive demonstration of all features

**Run it:**
```bash
go run main.go
```

## 🚀 Quick Start

1. **Clone and navigate:**
   ```bash
   git clone https://github.com/grandpaej/fmtpy
   cd fmtpy/example
   ```

2. **Run any example:**
   ```bash
   go run basic_usage.go
   ```

3. **Try the interactive examples:**
   ```bash
   go run advanced_features.go
   go run real_world_app.go
   ```

## 💡 Key Features Demonstrated

### Smart Input with Type Conversion
```go
// The new way - exactly what you wanted!
var age int = fmtpy.Input("Enter age: ").Int()
var score float64 = fmtpy.Input("Enter score: ").Float()
name := fmtpy.Input("Enter name: ").String()
confirmed := fmtpy.Input("Confirm (y/n): ").Bool()
```

### Colors That Work with fmt
```go
// Works with any fmt function and any data type
fmt.Println(color.Red("Error message"))
fmt.Println(color.Green(42))
fmt.Printf("Result: %s\n", color.Blue(1+1))
```

### Comprehensive String Manipulation
```go
// From the input package
text := "Hello World"
fmt.Println(input.Upper(text))      // "HELLO WORLD"
fmt.Println(input.Reverse(text))    // "dlroW olleH"
fmt.Println(input.WordCount(text))  // 2
```

## 🎯 Learning Path

1. **Start with `basic_usage.go`** - Learn the fundamentals
2. **Explore `color_showcase.go`** - Master colors
3. **Try `string_manipulation.go`** - Understand text processing
4. **Study `advanced_features.go`** - Learn advanced patterns
5. **Build with `real_world_app.go`** - See complete application

## 🔧 Customization

All examples are designed to be:
- **Modular** - Copy and adapt code snippets
- **Educational** - Well-commented and explained
- **Practical** - Real-world usage patterns
- **Progressive** - From simple to complex

## 📚 Documentation

- **README.md** - Complete API reference
- **EXAMPLES.md** - This file
- **Code comments** - Detailed explanations in each example

## 🎉 Have Fun!

These examples show the power and simplicity of fmtpy. Use them as:
- Learning materials
- Code templates
- Inspiration for your projects
- Testing ground for new ideas

Happy coding with fmtpy! 🚀
