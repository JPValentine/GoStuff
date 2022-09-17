package parsinglogfiles

import(
	"regexp"
	//"fmt"
)
func IsValidLine(text string) bool {
	myText := text
	if len(myText)>5{
		myText = myText[0:5]
	}
	switch{
	case myText == "[TRC]":
		return true
	case myText=="[DBG]":
		return true
	case myText=="[INF]":
		return true
	case myText=="[WRN]":
		return true
	case myText=="[ERR]":
		return true
	case myText=="[FTL]":
		return true
	default:
		return false
	}//switch
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[\*\~\=\-]*>`)
	return re.Split(text,-1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"\D*(?i)\b(password)\D*"`)
	count:=0
	for _,val := range lines{
		if(re.MatchString(val)){ count++}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`\b(end-of-line)*[\d]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	usr := regexpMustCompile(`(?i)\b(user)\s*(\S+)`)
	
}
