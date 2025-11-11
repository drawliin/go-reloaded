package helper

func Upper(s string) string {
	str := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c -= 32
		}
		str += string(c)
	}
	return str
}