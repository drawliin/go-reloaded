package helper

func ValidBrackets(arg string) bool {
	var brackets = map[rune]rune{
		')': '(',
	}

	openedB := []rune{}

	for _, b := range arg {
		if b == '(' {
			openedB = append(openedB, b)
			continue
		}
		if b == ')' {
			end := len(openedB) - 1
			if end >= 0 && openedB[end] == brackets[b] {
				openedB = openedB[:end]
			} else {
				return false
			}
		}
	}
	return len(openedB) == 0
}