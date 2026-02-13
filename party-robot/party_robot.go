package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	firstLine := fmt.Sprintf("Welcome to my party, %s!\n", name)
	secondLine := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, ", table, direction)
	thirdLine := fmt.Sprintf("exactly %.1f meters from here.\nYou will be sitting next to %s.", distance, neighbor)

	return firstLine + secondLine + thirdLine
}
