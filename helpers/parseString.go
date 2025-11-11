package helper

import (
	"regexp"
	"strconv"
	"strings"
)

func ParseString(s string) string {
	arrS := strings.Split(s, " ")
	for i := 0; i < len(arrS); i++ {
		hex, bin, cap, low, up := checkMod(arrS[i])
		if hex {
			arrS[i-1] = Hex(arrS[i-1])
		} else if bin {
			arrS[i-1] = Bin(arrS[i-1])
		} else if cap {
			count := parseMod(arrS[i])
			for j := 1; j <= count; j++ {
				arrS[i-j] = Capitalize(arrS[i-j])
			}
		} else if low {
			count := parseMod(arrS[i])
			for j := 1; j <= count; j++ {
				arrS[i-j] = Lower(arrS[i-j])
			}
		} else if up {
			count := parseMod(arrS[i])
			for j := 1; j <= count; j++ {
				arrS[i-j] = Upper(arrS[i-j])
			}
		}
	}
	return Join(arrS)
}

func parseMod(s string) int {
	num := ""
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num += string(c)
		}
	}
	if num == "" {
		return 1
	}
	count, _ := strconv.Atoi(num)
	return count
}

func checkMod(s string) (bool, bool, bool, bool, bool) {
	hex, bin, cap, low, up := false, false, false, false, false
	lowMatch := regexp.MustCompile(`^\(low(,\s*\d+)?\)$`)
	upMatch := regexp.MustCompile(`^\(up(,\s*\d+)?\)$`)
	capMatch := regexp.MustCompile(`^\(cap(,\s*\d+)?\)$`)
	if s == "(hex)" {
		hex = true
	} else if s == "(bin)" {
		bin = true
	} else if lowMatch.MatchString(s) {
		low = true
	} else if upMatch.MatchString(s) {
		up = true
	} else if capMatch.MatchString(s) {
		cap = true
	}
	return hex, bin, cap, low, up
}

func isMod(s string) bool {
	re1 := regexp.MustCompile(`^\((low|cap|up)(,\s*\d+)?\)$`)
	re2 := regexp.MustCompile(`^\(hex|bin\)$`)
	return re1.MatchString(s) || re2.MatchString(s)
}

func Join(slice []string) string {
	output := ""
	for i, str := range slice {
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
