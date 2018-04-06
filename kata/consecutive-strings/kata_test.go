package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsec(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		assert.Equal(t, "abigailtheta", LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
		assert.Equal(t, "oocccffuucccjjjkkkjyyyeehh", LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1))
		assert.Equal(t, "", LongestConsec([]string{}, 3))
		assert.Equal(t, "wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck", LongestConsec([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2))
	})
}
