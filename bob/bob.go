package bob

import (
	"strings"
)

func Hey(remark string) string {
	remark = strings.ReplaceAll(remark, " ", "")
	remark = strings.ReplaceAll(remark, "\r", "")
	remark = strings.ReplaceAll(remark, "\t", "")
	remark = strings.ReplaceAll(remark, "\n", "")
	qn4 := ""
	if remark == qn4 {
		return "Fine. Be that way!"
	} else if remark[len(remark)-1:] == "?" {
		if strings.ContainsAny(remark, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			if remark == strings.ToUpper(remark) {
				return "Calm down, I know what I'm doing!"
			}
		}
		return "Sure."
	} else if remark == strings.ToUpper(remark) {
		if strings.ContainsAny(remark, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			return "Whoa, chill out!"
		}
	}
	return "Whatever."
}
