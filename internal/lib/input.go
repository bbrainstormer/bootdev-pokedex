package lib

import "unicode"

func CleanInput(input string) []string {
	buf := ""
	words := []string{}
	for _, r := range input {
		if unicode.IsSpace(r) {
			if buf != "" {
				words = append(words, buf)
				buf = ""
			}
		} else {
			buf += string(unicode.ToLower(r))
		}
	}
	if buf != "" {
		words = append(words, buf)
	}
	return words
}
