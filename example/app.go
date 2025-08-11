package main

import (
	"fmt"
	"github.com/grandpaej/fmtpy"
	"github.com/grandpaej/fmtpy/color"
	"github.com/grandpaej/fmtpy/input"
)

// User represents a user profile
type User struct {
	Name     string
	Age      int
	Email    string
	Score    float64
	IsActive bool
}

func main() {
	fmt.Println(color.BoldCyan("🚀 Welcome to User Profile Manager"))
	fmt.Println(color.Green("Built with fmtpy - showcasing real-world usage"))
	
	// Create user profile
	user := createUserProfile()
	
	// Display profile
	displayUserProfile(user)
	
	// Interactive menu
	for {
		choice := showMenu()
		if !handleMenuChoice(choice, &user) {
			break
		}
	}
	
	fmt.Println(color.BoldGreen("👋 Thank you for using User Profile Manager!"))
}

func createUserProfile() User {
	fmt.Println(color.BoldYellow("\n📝 Create User Profile"))
	
	// Get user input with validation
	var name string
	for {
		name = fmtpy.Input("Enter your name: ").String()
		name = input.Trim(name)
		
		if !input.IsEmpty(name) && input.IsLetter(input.RemoveSpaces(name)) {
			name = input.Title(name) // Capitalize properly
			break
		}
		fmt.Println(color.Red("❌ Please enter a valid name (letters only)"))
	}
	
	// Get age with validation
	var age int
	for {
		age = fmtpy.Input("Enter your age (18-100): ").Int()
		if age >= 18 && age <= 100 {
			break
		}
		fmt.Println(color.Red("❌ Age must be between 18 and 100"))
	}
	
	// Get email with basic validation
	var email string
	for {
		email = fmtpy.Input("Enter your email: ").String()
		email = input.Lower(input.Trim(email))
		
		if input.Contains(email, "@") && input.Contains(email, ".") && !input.IsEmpty(email) {
			break
		}
		fmt.Println(color.Red("❌ Please enter a valid email address"))
	}
	
	// Get score
	var score float64
	for {
		score = fmtpy.Input("Enter your score (0.0-100.0): ").Float()
		if score >= 0.0 && score <= 100.0 {
			break
		}
		fmt.Println(color.Red("❌ Score must be between 0.0 and 100.0"))
	}
	
	// Get active status
	isActive := fmtpy.Input("Are you an active user? (y/n): ").Bool()
	
	return User{
		Name:     name,
		Age:      age,
		Email:    email,
		Score:    score,
		IsActive: isActive,
	}
}

func displayUserProfile(user User) {
	fmt.Println(color.BoldMagenta("\n👤 User Profile"))
	fmt.Println(color.Cyan("═══════════════════════════════════"))
	
	// Display with colors based on data
	fmt.Printf("📛 Name: %s\n", color.BoldBlue(user.Name))
	fmt.Printf("🎂 Age: %s years old\n", color.Yellow(user.Age))
	fmt.Printf("📧 Email: %s\n", color.Green(user.Email))
	
	// Color-coded score
	var scoreColor func(interface{}) string
	switch {
	case user.Score >= 90:
		scoreColor = color.BoldGreen
	case user.Score >= 70:
		scoreColor = color.Yellow
	case user.Score >= 50:
		scoreColor = color.BoldYellow
	default:
		scoreColor = color.Red
	}
	fmt.Printf("📊 Score: %s/100\n", scoreColor(user.Score))
	
	// Status with appropriate color
	status := "Inactive"
	statusColor := color.Red
	if user.IsActive {
		status = "Active"
		statusColor = color.BoldGreen
	}
	fmt.Printf("🔘 Status: %s\n", statusColor(status))
	
	// Additional info
	fmt.Printf("📏 Name length: %s characters\n", color.Cyan(input.Length(user.Name)))
	fmt.Printf("🔤 Name reversed: %s\n", color.Magenta(input.Reverse(user.Name)))
	fmt.Printf("📝 Name uppercase: %s\n", color.Blue(input.Upper(user.Name)))
}

func showMenu() int {
	fmt.Println(color.BoldCyan("\n🔧 Profile Manager Menu"))
	fmt.Println("1. " + color.Green("Update name"))
	fmt.Println("2. " + color.Yellow("Update age"))
	fmt.Println("3. " + color.Blue("Update email"))
	fmt.Println("4. " + color.Magenta("Update score"))
	fmt.Println("5. " + color.Cyan("Toggle active status"))
	fmt.Println("6. " + color.White("Show profile"))
	fmt.Println("7. " + color.BoldYellow("Profile statistics"))
	fmt.Println("8. " + color.Red("Exit"))
	
	return fmtpy.Input("Choose an option (1-8): ").Int()
}

func handleMenuChoice(choice int, user *User) bool {
	switch choice {
	case 1:
		updateName(user)
	case 2:
		updateAge(user)
	case 3:
		updateEmail(user)
	case 4:
		updateScore(user)
	case 5:
		toggleActiveStatus(user)
	case 6:
		displayUserProfile(*user)
	case 7:
		showProfileStatistics(*user)
	case 8:
		return false
	default:
		fmt.Println(color.Red("❌ Invalid choice. Please try again."))
	}
	return true
}

func updateName(user *User) {
	fmt.Printf("Current name: %s\n", color.Blue(user.Name))
	newName := fmtpy.Input("Enter new name: ").String()
	newName = input.Title(input.Trim(newName))
	
	if !input.IsEmpty(newName) {
		user.Name = newName
		fmt.Printf("✅ Name updated to: %s\n", color.Green(user.Name))
	} else {
		fmt.Println(color.Red("❌ Invalid name. No changes made."))
	}
}

func updateAge(user *User) {
	fmt.Printf("Current age: %s\n", color.Yellow(user.Age))
	newAge := fmtpy.Input("Enter new age: ").Int()
	
	if newAge >= 18 && newAge <= 100 {
		user.Age = newAge
		fmt.Printf("✅ Age updated to: %s\n", color.Green(user.Age))
	} else {
		fmt.Println(color.Red("❌ Invalid age. Must be between 18 and 100."))
	}
}

func updateEmail(user *User) {
	fmt.Printf("Current email: %s\n", color.Green(user.Email))
	newEmail := fmtpy.Input("Enter new email: ").String()
	newEmail = input.Lower(input.Trim(newEmail))
	
	if input.Contains(newEmail, "@") && input.Contains(newEmail, ".") {
		user.Email = newEmail
		fmt.Printf("✅ Email updated to: %s\n", color.Green(user.Email))
	} else {
		fmt.Println(color.Red("❌ Invalid email format."))
	}
}

func updateScore(user *User) {
	fmt.Printf("Current score: %s\n", color.Yellow(user.Score))
	newScore := fmtpy.Input("Enter new score: ").Float()
	
	if newScore >= 0.0 && newScore <= 100.0 {
		user.Score = newScore
		fmt.Printf("✅ Score updated to: %s\n", color.Green(user.Score))
	} else {
		fmt.Println(color.Red("❌ Invalid score. Must be between 0.0 and 100.0."))
	}
}

func toggleActiveStatus(user *User) {
	user.IsActive = !user.IsActive
	status := "Inactive"
	statusColor := color.Red
	if user.IsActive {
		status = "Active"
		statusColor = color.Green
	}
	fmt.Printf("✅ Status changed to: %s\n", statusColor(status))
}

func showProfileStatistics(user User) {
	fmt.Println(color.BoldCyan("\n📊 Profile Statistics"))
	fmt.Println(color.Cyan("═══════════════════════════════"))
	
	// Name analysis
	nameStats := analyzeText(user.Name)
	fmt.Printf("📛 Name Analysis:\n")
	fmt.Printf("   Length: %s characters\n", color.Yellow(nameStats.Length))
	fmt.Printf("   Words: %s\n", color.Green(nameStats.WordCount))
	fmt.Printf("   First word: %s\n", color.Blue(nameStats.FirstWord))
	fmt.Printf("   Uppercase: %s\n", color.Red(nameStats.Uppercase))
	fmt.Printf("   Lowercase: %s\n", color.Magenta(nameStats.Lowercase))
	
	// Email analysis
	emailStats := analyzeText(user.Email)
	fmt.Printf("\n📧 Email Analysis:\n")
	fmt.Printf("   Length: %s characters\n", color.Yellow(emailStats.Length))
	fmt.Printf("   Domain: %s\n", color.Green(getDomain(user.Email)))
	fmt.Printf("   Contains numbers: %s\n", color.Blue(input.Contains(user.Email, input.OnlyNumbers(user.Email))))
	
	// Score analysis
	fmt.Printf("\n📊 Score Analysis:\n")
	grade := getGrade(user.Score)
	fmt.Printf("   Grade: %s\n", getGradeColor(grade)(grade))
	fmt.Printf("   Percentage: %s%%\n", color.Yellow(user.Score))
	fmt.Printf("   Points to 100: %s\n", color.Cyan(100.0-user.Score))
}

type TextStats struct {
	Length    int
	WordCount int
	FirstWord string
	Uppercase string
	Lowercase string
}

func analyzeText(text string) TextStats {
	return TextStats{
		Length:    input.Length(text),
		WordCount: input.WordCount(text),
		FirstWord: input.FirstWord(text),
		Uppercase: input.Upper(text),
		Lowercase: input.Lower(text),
	}
}

func getDomain(email string) string {
	parts := input.Split(email, "@")
	if len(parts) > 1 {
		return parts[1]
	}
	return "unknown"
}

func getGrade(score float64) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func getGradeColor(grade string) func(interface{}) string {
	switch grade {
	case "A":
		return color.BoldGreen
	case "B":
		return color.Green
	case "C":
		return color.Yellow
	case "D":
		return color.BoldYellow
	default:
		return color.Red
	}
}
