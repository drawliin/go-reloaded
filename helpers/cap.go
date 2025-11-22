package helper

func Capitalize(s string) string {
	index := 0
	wordStart := false
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			wordStart = true
		}
		if !(s[i] >= 'a' && s[i] <= 'z') && !(s[i] >= 'A' && s[i] <= 'Z') && wordStart {
			index = i+1
		}

	}

	str := ""
	for i, c := range s {
		if c >= 'a' && c <= 'z' && i == index {
			c -= 32
		}
		if c >= 'A' && c <= 'Z' && i != index {
			c += 32
		}
		str += string(c)
	}
	return str
}
