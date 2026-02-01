package fundamentals

import "fmt"

// AgeChecker evaluates a person's life stage based on their age.
// It uses a switch statement to categorize the age value into three groups:
// - Child: for ages less than 13
// - Teen: for ages between 13 and 19 (inclusive)
// - Adult: for ages 20 and above
// The function currently uses a hardcoded age value of 18 for demonstration purposes.
func AgeChecker() {
	age := 18

	switch true {
	case age < 13:
		fmt.Println("Child")
	case age >= 13 && age <= 19:
		fmt.Println("Teen")
	default:
		fmt.Println("Adult")
	}
}
