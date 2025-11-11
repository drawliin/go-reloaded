package helper

func Capitalize(s string) string {
	str := ""
	for i, c := range s {
		if c >= 'a' && c <= 'z' && i == 0{
			c -=  32
		}
		if c >= 'A' && c <= 'Z' && i != 0 {
			c += 32
		}
		str += string(c)
	}
	return str
}