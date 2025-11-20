package helper

import (
	"fmt"
	"math"
	"strconv"
)

func hexToDecimal(c rune) (int, bool) {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0'), false
	case c >= 'a' && c <= 'f':
		return int(c - 'a' + 10), false
	case c >= 'A' && c <= 'F':
		return int(c - 'A' + 10), false
	default:
		return 0, true
	}
}

func Hex(s string) string {
	var num uint64 = 0
	sign := ""
	if s[0] == '-' {
		s = s[1:]
		sign = "-"
	}
	for _, c := range s {
		d, err := hexToDecimal(c)
		if err {
			return s
		}
		if num > (math.MaxUint64-uint64(d))/16 {
			fmt.Printf("Hexadecimal Overflowing...\n")
			return s
		}
		num = num*16 + uint64(d)
	}
	return sign + strconv.FormatUint(num, 10)
}
