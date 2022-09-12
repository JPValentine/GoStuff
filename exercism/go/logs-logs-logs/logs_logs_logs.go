package logs

import(
	"fmt"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _,char := range log{
		x:=fmt.Sprintf("%U",char)
		if x == "U+2757"{
			return "recommendation"
		}
		if x == "U+1F50D"{
			return "search"
		}
		if x == "U+2600"{
			return "weather"
		}
	}
	return "default" 
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	y:=fmt.Sprintf("%U",oldRune)
	myStr:=""
	for i,c := range log{
		x:=fmt.Sprintf("%U", c)
		if x == y{
			log[i] = newRune
		}
	}
	return log
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	panic("Please implement the WithinLimit() function")
}
