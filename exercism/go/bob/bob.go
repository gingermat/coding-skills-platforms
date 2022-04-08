package bob

import (
	"strings"
)

func isEmpty(remark string) bool {
	return remark == ""
}
func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isUpperCase(remark string) bool {
	return remark != strings.ToLower(remark) && remark == strings.ToUpper(remark)
}

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case isUpperCase(remark) && isQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case isUpperCase(remark):
		return "Whoa, chill out!"
	case isQuestion(remark):
		return "Sure."
	case isEmpty(remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
