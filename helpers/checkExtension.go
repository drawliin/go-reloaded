package helper

import (
	"strings"
)

func CheckExtension(s string) bool {
	fields := strings.Split(s, ".")

	extension := fields[len(fields)-1]

	if extension == "txt" {
		return true
	}

	return false
}
