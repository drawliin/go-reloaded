package helper

func ToUpper(s string) string {
	out := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c = c-32
		}
		out += string(c)
	}
	return out
}
