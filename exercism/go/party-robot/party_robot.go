package partyrobot

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, table int, neighbour, direction string, distance float64) string {
	positionText := fmt.Sprintf("You have been assigned to table %X. Your table is %s, exactly %.1f meters from here.",
		table, direction, distance)
	neighbourText := fmt.Sprintf("You will be sitting next to %s.", neighbour)

	return Welcome(name) + "\n" + positionText + "\n" + neighbourText

}
