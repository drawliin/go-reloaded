package helper

import (
	"slices"
	"strings"
)

func CheckExtension(s string) bool {
	invalidExtensions := []string{"go", "mod", "sum"}

	fields := strings.Split(s, ".")

	extension := fields[len(fields)-1]

	if slices.Contains(invalidExtensions, extension) {
		return false
	}

	return true
}
