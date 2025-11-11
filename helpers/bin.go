package helper

import "strconv"

func binary(c rune) int {
	switch c {
	case '0':
		return 0
	case '1':
		return 1
	default:
		panic("Invalid Binary")
	}
}
func Bin(s string) string {
	num := 0
	for _, c := range s {
		num = num*2 + binary(c)
	}
	return strconv.Itoa(num)
}
