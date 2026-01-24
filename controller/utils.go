package controller

import "regexp"

// IsValidEmail checks if the provided email address is valid.
func IsValidEmail(email string) bool {
	// Define a regular expression for validating email addresses
	// This regex covers most common email formats
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// Compile the regular expression
	re := regexp.MustCompile(emailRegex)

	// Return whether the email address matches the regular expression
	return re.MatchString(email)
}

func IsValidPhone(phone string) bool {
	// Indian phone numbers start with 6-9 and are 10 digits long
	phoneRegex := `^[6-9]\d{9}$`
	matched, _ := regexp.MatchString(phoneRegex, phone)
	if !matched {
		return false
	}
	return true
}
