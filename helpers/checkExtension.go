package helper

import (
	"strings"
)

func CheckExtension(s string) bool {
	invalidExtensions := []string{"go", "mod", "sum"}
	
	fields := strings.Split(s, ".")
	
	extension := fields[len(fields)-1]
	
	for _, ext := range invalidExtensions {
		if ext == extension {
			return false
		}
	}
	
	return true
}
