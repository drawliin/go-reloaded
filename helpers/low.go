package helper

func Lower(s string) string {
	str := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		str += string(c)
	}
	return str
}