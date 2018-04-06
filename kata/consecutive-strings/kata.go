package kata

import (
	"strings"
)

func LongestConsec(strarr []string, k int) string {
	lng := ""
	for i := 0; k > 0 && i <= len(strarr)-k; i++ {
		str := strings.Join(strarr[i:i+k], "")
		if len(str) > len(lng) {
			lng = str
		}
	}

	return lng
}
