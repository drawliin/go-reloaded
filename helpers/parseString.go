package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ParseString(s string) string {
	arrS := Split(s)
	fmt.Printf("%q\n", arrS) ///////////////
	stack := []string{}
	for i := 0; i < len(arrS); i++ {
		if i == 0 && isMod(arrS[i]) {
			continue
		}

		if ValidBrackets(arrS[i]) {
			if strings.HasPrefix(arrS[i], "(") && strings.HasSuffix(arrS[i], ")") {
				arrS[i] = "(" + ParseString(arrS[i][1:len(arrS[i])-1]) + ")"
			}
		} else if strings.HasPrefix(arrS[i], "(") {
			arrS[i] = "(" + ParseString(arrS[i][1:])
		}

		if !isMod(arrS[i]) {
			stack = append(stack, arrS[i])
		}
		hex, bin, cap, low, up := checkMod(arrS[i])
		if hex {
			end := len(stack) - 1
			stack[end] = Hex(stack[end])
		} else if bin {
			end := len(stack) - 1
			stack[end] = Bin(stack[end])
		} else if cap {
			ApplyMod(stack, parseMod(arrS[i]), Capitalize)
		} else if low {
			ApplyMod(stack, parseMod(arrS[i]), Lower)
		} else if up {
			ApplyMod(stack, parseMod(arrS[i]), Upper)
		}
	}
	fmt.Printf("%q\n", stack) //////////////////
	return Join(stack)
}

func parseMod(s string) int {
	num := ""
	for _, c := range s {
		if (c >= '0' && c <= '9') || c == '-' {
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
	lowMatch := regexp.MustCompile(`^\(low(,\s*\-?\d+)?\)$`)
	upMatch := regexp.MustCompile(`^\(up(,\s*\-?\d+)?\)$`)
	capMatch := regexp.MustCompile(`^\(cap(,\s*\-?\d+)?\)$`)
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
	re1 := regexp.MustCompile(`^\((low|cap|up)(,\s*\-?\d+)?\)$`)
	re2 := regexp.MustCompile(`^\((hex|bin)\)$`)
	return re1.MatchString(s) || re2.MatchString(s)
}

func ApplyMod(stack []string, count int, fn func(string) string) {
	if count > len(stack) {
		count = len(stack)
	}
	if count < 0 {
		count = 0
	}
	end := len(stack) - 1
	applied := 0
	pos := end
	for pos >= 0 && applied < count {
		if strings.Contains(stack[pos], "\n") {
			break
		}
		if hasLetter(stack[pos]) {
			stack[pos] = fn(stack[pos])
			applied++
		}
		pos--
	}
}

func Join(slice []string) string {
	out := []string{}
	result := ""

	for i := 0; i < len(slice); i++ {

		if i != len(slice)-1 && Lower(slice[i]) == "a" && startWithVowel(Lower(slice[i+1])) {
			slice[i] = slice[i] + "n"
		}
		out = append(out, slice[i])

	}

	foundSingleQuote := 0
	for i, str := range out {

		if (isPunctuation(str) &&
			puncAlone(str) &&
			i > 0 &&
			!isPunctuation(out[i-1])) ||
			(i > 0 &&
				strings.ContainsAny(out[i-1], "!?") &&
				containsOnlyExclOrInterr(str)) {
			result = result[:len(result)-1]
		}
		
		if strings.Contains(str, "'") {
			foundSingleQuote++
		}

		if foundSingleQuote == 2 {
			str = handleSecondQuote(str)
			foundSingleQuote = 0
		}
		result += str

		if i < len(out)-1 && !strings.Contains(str, "\n") && !strings.Contains(out[i+1], "\n") {
			result += " "
		}

		if foundSingleQuote == 1 {
			if str == "'" {
				result = result[:len(result)-1]
				continue
			}

			if i < len(out)-1 && strings.HasPrefix(out[i+1], "'") {
				result = result[:len(result)-1]
			}
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
			if s[wordStart:i] != "" {
				arr = append(arr, s[wordStart:i])
				wordStart = i
			}
			nest := 1
			for i+1 < len(s) {
				i++
				if s[i] == '(' {
					nest++
				}
				if s[i] == ')' {
					nest--
				}
				if nest == 0 {
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
		} else if s[i] == '\'' {
			if s[wordStart:i] != "" {
				arr = append(arr, s[wordStart:i])
				wordStart = i
			} else {
				i++
			}
		} else if s[i] == '\n' {
			if s[wordStart:i] != "" {
				arr = append(arr, s[wordStart:i])
				wordStart = i
			}

			if i < len(s)-1 {
				for i+1 < len(s) {
					i++
					if s[i] != '\n' {
						arr = append(arr, s[wordStart:i])
						wordStart = i
						break
					}
				}
			} else {
				i++
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

func startWithVowel(s string) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'h'}
	for _, c := range vowels {
		if c == s[0] {
			return true
		}
	}
	return false
}

func hasLetter(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func containsOnlyExclOrInterr(s string) bool {
	if s[0] == '!' || s[0] == '?' {
		return true
	}
	return false
}

func handleSecondQuote(s string) string {
	result := ""
	for _, c := range s {
		result += string(c)
		if c == '\'' && len(s) > 1 {
			result += " "
		}
	}
	return result
}
