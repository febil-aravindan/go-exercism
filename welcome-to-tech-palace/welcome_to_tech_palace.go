package techpalace

import (
	s "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace," + " " + s.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	str := s.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + s.Repeat("*", numStarsPerLine)
	return str
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	starCleanedString := s.ReplaceAll(oldMsg, "*", "")
	trimedString := s.TrimSpace(starCleanedString)
	return trimedString
}
