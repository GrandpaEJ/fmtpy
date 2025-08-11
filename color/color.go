package color

import (
	"fmt"
	"os"
	"strings"
)

// Attributes for text styling
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	Inverse
	Concealed
	CrossedOut
)

// Foreground colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Background colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Bright foreground colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Bright background colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// Attribute defines a single SGR Code
type Attribute int

// NoColor defines if the output is colorized or not
var NoColor = os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb"

// Color represents a text color
type Color struct {
	params []Attribute
}

// New creates a new Color instance
func New(attrs ...Attribute) *Color {
	c := &Color{params: make([]Attribute, 0)}
	c.Add(attrs...)
	return c
}

// Add adds attributes to the color
func (c *Color) Add(attrs ...Attribute) *Color {
	c.params = append(c.params, attrs...)
	return c
}

// Sprint formats using the default formats for its operands and returns the resulting string
func (c *Color) Sprint(args ...interface{}) string {
	return c.wrap(fmt.Sprint(args...))
}

// Sprintf formats according to a format specifier and returns the resulting string
func (c *Color) Sprintf(format string, args ...interface{}) string {
	return c.wrap(fmt.Sprintf(format, args...))
}

// formatString replaces {placeholders} with arguments
func formatString(template string, args ...interface{}) string {
	result := template
	for _, arg := range args {
		// Replace the first occurrence of {word} with the arg string
		start := strings.Index(result, "{")
		end := strings.Index(result, "}")
		if start != -1 && end != -1 && end > start {
			result = result[:start] + fmt.Sprint(arg) + result[end+1:]
		}
	}
	return result
}

// Format using Python-style {placeholder} formatting
func (c *Color) Format(template string, args ...interface{}) string {
	return c.wrap(formatString(template, args...))
}

// wrap wraps the text with color escape codes
func (c *Color) wrap(text string) string {
	if NoColor {
		return text
	}

	format := make([]string, len(c.params))
	for i, attr := range c.params {
		format[i] = fmt.Sprintf("\033[%dm", attr)
	}

	return strings.Join(format, "") + text + "\033[0m"
}

// Simple color functions that work with any type - perfect for fmt.Print, fmt.Println, fmt.Printf
func Red(v interface{}) string {
	return New(FgRed).Sprint(v)
}

func Green(v interface{}) string {
	return New(FgGreen).Sprint(v)
}

func Blue(v interface{}) string {
	return New(FgBlue).Sprint(v)
}

func Yellow(v interface{}) string {
	return New(FgYellow).Sprint(v)
}

func Magenta(v interface{}) string {
	return New(FgMagenta).Sprint(v)
}

func Cyan(v interface{}) string {
	return New(FgCyan).Sprint(v)
}

func White(v interface{}) string {
	return New(FgWhite).Sprint(v)
}

func Black(v interface{}) string {
	return New(FgBlack).Sprint(v)
}

// Bold color functions
func BoldRed(v interface{}) string {
	return New(Bold, FgRed).Sprint(v)
}

func BoldGreen(v interface{}) string {
	return New(Bold, FgGreen).Sprint(v)
}

func BoldBlue(v interface{}) string {
	return New(Bold, FgBlue).Sprint(v)
}

func BoldYellow(v interface{}) string {
	return New(Bold, FgYellow).Sprint(v)
}

func BoldMagenta(v interface{}) string {
	return New(Bold, FgMagenta).Sprint(v)
}

func BoldCyan(v interface{}) string {
	return New(Bold, FgCyan).Sprint(v)
}

// Background color functions
func OnRed(v interface{}) string {
	return New(BgRed).Sprint(v)
}

func OnGreen(v interface{}) string {
	return New(BgGreen).Sprint(v)
}

func OnBlue(v interface{}) string {
	return New(BgBlue).Sprint(v)
}

func OnYellow(v interface{}) string {
	return New(BgYellow).Sprint(v)
}

func OnMagenta(v interface{}) string {
	return New(BgMagenta).Sprint(v)
}

func OnCyan(v interface{}) string {
	return New(BgCyan).Sprint(v)
}

// Text color functions (legacy - for backward compatibility)
func RedText(template string, args ...interface{}) string {
	return New(FgRed).Format(template, args...)
}

func GreenText(template string, args ...interface{}) string {
	return New(FgGreen).Format(template, args...)
}

func BlueText(template string, args ...interface{}) string {
	return New(FgBlue).Format(template, args...)
}

func YellowText(template string, args ...interface{}) string {
	return New(FgYellow).Format(template, args...)
}

func MagentaText(template string, args ...interface{}) string {
	return New(FgMagenta).Format(template, args...)
}

func CyanText(template string, args ...interface{}) string {
	return New(FgCyan).Format(template, args...)
}

// Style combinations
func BoldRedText(template string, args ...interface{}) string {
	return New(Bold, FgRed).Format(template, args...)
}

func ItalicGreenText(template string, args ...interface{}) string {
	return New(Italic, FgGreen).Format(template, args...)
}

func HighlightedText(template string, args ...interface{}) string {
	return New(Bold, FgBlack, BgYellow).Format(template, args...)
}

func ErrorText(template string, args ...interface{}) string {
	return New(Bold, FgRed, BgHiWhite).Format(template, args...)
}

func SuccessText(template string, args ...interface{}) string {
	return New(Bold, FgGreen).Format(template, args...)
}

func InfoText(template string, args ...interface{}) string {
	return New(FgCyan).Format(template, args...)
}

func WarningText(template string, args ...interface{}) string {
	return New(Bold, FgYellow).Format(template, args...)
}

// Get returns a new Color instance
func Get() *Color {
	return New()
}

// Short form methods
func (c *Color) R(s string) string { return New(FgRed).Sprint(s) }
func (c *Color) G(s string) string { return New(FgGreen).Sprint(s) }
func (c *Color) B(s string) string { return New(FgBlue).Sprint(s) }
func (c *Color) Y(s string) string { return New(FgYellow).Sprint(s) }
func (c *Color) M(s string) string { return New(FgMagenta).Sprint(s) }
func (c *Color) C(s string) string { return New(FgCyan).Sprint(s) }

// Bold versions
func (c *Color) BR(s string) string { return New(Bold, FgRed).Sprint(s) }
func (c *Color) BG(s string) string { return New(Bold, FgGreen).Sprint(s) }
func (c *Color) BB(s string) string { return New(Bold, FgBlue).Sprint(s) }
func (c *Color) BY(s string) string { return New(Bold, FgYellow).Sprint(s) }
func (c *Color) BM(s string) string { return New(Bold, FgMagenta).Sprint(s) }
func (c *Color) BC(s string) string { return New(Bold, FgCyan).Sprint(s) }

// Background versions
func (c *Color) OnR(s string) string { return New(BgRed).Sprint(s) }
func (c *Color) OnG(s string) string { return New(BgGreen).Sprint(s) }
func (c *Color) OnB(s string) string { return New(BgBlue).Sprint(s) }
func (c *Color) OnY(s string) string { return New(BgYellow).Sprint(s) }
func (c *Color) OnM(s string) string { return New(BgMagenta).Sprint(s) }
func (c *Color) OnC(s string) string { return New(BgCyan).Sprint(s) }

// Semantic helpers
func (c *Color) E(s string, args ...interface{}) string { return New(FgRed).Format(s, args...) }    // Error
func (c *Color) S(s string, args ...interface{}) string { return New(FgGreen).Format(s, args...) }  // Success
func (c *Color) I(s string, args ...interface{}) string { return New(FgBlue).Format(s, args...) }   // Info
func (c *Color) W(s string, args ...interface{}) string { return New(FgYellow).Format(s, args...) } // Warning
