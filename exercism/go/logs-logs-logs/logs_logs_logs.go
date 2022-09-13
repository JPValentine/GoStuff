package logs

import(
	"fmt"
	"strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _,char := range log{
		if fmt.Sprintf("%U",char) == "U+2757"{
			return "recommendation"
		}
		if fmt.Sprintf("%U",char) == "U+1F50D"{
			return "search"
		}
		if fmt.Sprintf("%U",char) == "U+2600"{
			return "weather"
		}
	}
	return "default" 
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newStr:=fmt.Sprintf("%c",newRune)
	oldStr:=fmt.Sprintf("%c",oldRune)
	return strings.Replace(log,oldStr,newStr,-1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log))<=limit
}
