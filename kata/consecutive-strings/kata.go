package kata

import (
	"strings"
)

func LongestConsec(strarr []string, k int) (longest string) {
	for i := 0; k > 0 && i <= len(strarr)-k; i++ {
		str := strings.Join(strarr[i:i+k], "")
		if len(str) > len(longest) {
			longest = str
		}
	}
	return
}
