package helper

import (
	"regexp"
	"strconv"
)

func ParseString(s string) string {
	arrS := Split(s)
	for i := 0; i < len(arrS); i++ {
		if i == 0 || isMod(arrS[i-1]) {
			continue
		}
		hex, bin, cap, low, up := checkMod(arrS[i])
		if hex {
			arrS[i-1] = Hex(arrS[i-1])
		} else if bin {
			arrS[i-1] = Bin(arrS[i-1])
		} else if cap {
			ApplyMod(arrS, i, parseMod(arrS[i]), Capitalize)
		} else if low {
			ApplyMod(arrS, i, parseMod(arrS[i]), Lower)
		} else if up {
			ApplyMod(arrS, i, parseMod(arrS[i]), Upper)
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
	re2 := regexp.MustCompile(`^\((hex|bin)\)$`)
	return re1.MatchString(s) || re2.MatchString(s)
}

func ApplyMod(arrS []string, i, count int, fn func(string) string) {
	if count > i {
		count = i
	}
	for j := 1; j <= count; j++ {
		arrS[i-j] = fn(arrS[i-j])
	}
}

func Join(slice []string) string {
	out := []string{}
	result := ""
	for _, str := range slice {
		if !isMod(str) {
			out = append(out, str)
		}
	}
	for i, str := range out {
		if isPunctuation(str) && puncAlone(str) {
			result = result[:len(result)-1]
		}
		result += str
		if i != len(out)-1 && str != "'" && containQuote(out[i+1]){
			result += " "
		}
	}
	return result
}

func Split(s string) []string {
	arr := []string{}
	wordStart := 0
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			if s[wordStart:i] != "" {
				arr = append(arr, s[wordStart:i])
			}
			wordStart = i + 1
			i = wordStart
		} else if s[i] == '(' {
			for i+1 < len(s) {
				i++
				if s[i] == ')' {
					arr = append(arr, s[wordStart:i+1])
					wordStart = i + 1
					i = wordStart
					break
				}
			}
		} else if s[i] == ',' || s[i] == ';' || s[i] == ':' {
			arr = append(arr, s[wordStart:i+1])
			wordStart = i + 1
			i = wordStart
		} else if s[i] == '.' && i < len(s)-1 {
			for i+1 < len(s) {
				i++
				if s[i] != '.' {
					arr = append(arr, s[wordStart:i])
					wordStart = i
					break
				}
			}
		} else if (s[i] == '!' || s[i] == '?') && i < len(s)-1 {
			for i+1 < len(s) {
				i++
				if s[i] != '!' && s[i] != '?' {
					arr = append(arr, s[wordStart:i])
					wordStart = i
					break
				}
			}
		} else {
			i++
		}
	}
	if wordStart < len(s) {
		arr = append(arr, s[wordStart:])
	}
	return arr
}

func isPunctuation(s string) bool {
	switch s[len(s)-1] {
	case ',', '.', ':', ';', '!', '?':
		return true
	default:
		return false
	}

}

func puncAlone(s string) bool {
	for _, c := range s {
		if !isPunctuation(string(c)) {
			return false
		}
	}
	return true
}

func containQuote(s string) bool {
	for _,c := range s {
		if c == '\'' {
			return true
		}
	}
		return false
}