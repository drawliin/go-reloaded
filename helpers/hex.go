package helper

import "strconv"

func hexToDecimal(c rune) int {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0')
	case c >= 'a' && c <= 'f':
		return int(c - 'a' + 10)
	case c >= 'A' && c <= 'F':
		return int(c - 'A' + 10)
	default:
		panic("Invalid hex caracter")
	}
}

func Hex(s string) string {
	num := 0
	for _, c := range s {
		num = num*16 + hexToDecimal(c)
	}
	return strconv.Itoa(num)
}
