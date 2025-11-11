package helper

import (
	"regexp"
	"strings"
)

func ParseString(s string) string{
	arrS := strings.Split(s, " ")
	for i := 0; i < len(arrS); i++ {
		if arrS[i] == "(hex)" {
			arrS[i-1] = Hex(arrS[i-1])
			continue
		} else if arrS[i] == "(bin)" {
			arrS[i-1] = Bin(arrS[i-1])
			continue
		}
	}
	return Join(arrS)
}

func Join(slice []string) string {
	output := ""
	for i,str := range slice {
		if isMod(str) {
			continue
		}
		output += str
		if i != len(slice)-1 {
			output += " "
		}
	}
	return output
}

func isMod(s string) bool {
	re1 := regexp.MustCompile(`^\((low|cap|up)(,\s*\d+)?\)$`)
	re2 := regexp.MustCompile(`^\(hex|bin\)$`)
	return re1.MatchString(s) || re2.MatchString(s)
}