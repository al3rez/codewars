package kata

import (
	"strings"
)

func duplicateCount(s string) (c int) {
	h := map[rune]int{}

	for _, r := range strings.ToLower(s) {
		if h[r] += 1; h[r] == 2 {
			c++
		}
	}

	return
}
