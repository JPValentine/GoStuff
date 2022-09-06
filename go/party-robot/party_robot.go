package partyrobot

import(
	"strconv"
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	a := strconv.Itoa(age)
	return "Happy birthday "+ name +"! You are now "+ a + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	x := Welcome(name)
	// Welcome to my party, Christiane!
	t := fmt.Sprintf("%03d", table)
	//t := strconv.Itoa(table)
	s := strconv.FormatFloat(distance, 'f',1,64)
	y := "\nYou have been assigned to table "+ t + ". Your table is "+direction+", exactly "+s+" meters from here.\n"
	// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
	z:= "You will be sitting next to "+ neighbor+"."
	// You will be sitting next to Frank.

	return x + y + z

}
