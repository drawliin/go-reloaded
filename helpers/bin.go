package helper

import "strconv"

func binary(c rune) (int, bool) {
	switch c {
	case '0':
		return 0, false
	case '1':
		return 1, false
	default:
		return 0, true
	}
}

func Bin(s string) string {
	num := 0
	sign := ""
	if s[0] == '-' {
		s = s[1:]
		sign = "-"
	}
	for _, c := range s {
		d, err := binary(c)
		if err {
			return s
		}
		num = num*2 + d
	}
	return sign+strconv.Itoa(num)
}
