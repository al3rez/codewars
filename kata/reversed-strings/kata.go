package kata

import "strings"

func Solution(word string) string {
	var reversed []string
	for i := range word {
		reversed = append(reversed, string(word[len(word)-1-i]))
	}
	return strings.Join(reversed, "")
}
