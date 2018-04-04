package kata

func HasUniqueChar(s string) bool {
	h := map[rune]int{}

	for _, r := range s {
		if h[r] += 1; h[r] == 2 {
			return false
		}
	}

	return true
}
