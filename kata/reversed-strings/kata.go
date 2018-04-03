package kata

func Solution(word string) string {
	var reversed []rune
	for i := range word {
		reversed = append(reversed, rune(word[len(word)-1-i]))
	}
	return string(reversed)
}
