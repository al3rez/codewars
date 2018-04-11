package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCentury(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		assert.Equal(t, 20, century(int(2000)))
		assert.Equal(t, 18, century(int(1705)))
	})
}
