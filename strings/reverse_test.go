package strings

import (
	"testing"
)

func TestRevert(t *testing.T) {
	var removeTests = []struct {
		value    []rune
		expected string
	}{
		{[]rune("first second third"), "third second first"},
		{[]rune("first second"), "second first"},
	}
	for _, tt := range removeTests {
		res := reverse(tt.value)

		if res != tt.expected {
			t.Errorf("Revert (%d): expected %d, actual %d", tt.value, tt.expected, res)
		}
	}
}
