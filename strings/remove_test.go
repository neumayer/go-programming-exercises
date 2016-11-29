package strings

import (
	"reflect"
	"testing"
)

func TestLookup(t *testing.T) {
	var lookupTests = []struct {
		value    string
		expected []int
	}{
		{"abc", []int{97, 98, 99}},
	}
	for _, tt := range lookupTests {
		res := getLookup(tt.value)
		expected := make([]bool, 128)

		for _, i := range tt.expected {
			expected[i] = true
		}

		if !reflect.DeepEqual(res, expected) {
			t.Errorf("Lookup (%d): expected %d, actual %d", tt.value, expected, res)
		}
	}
}

func TestRemove(t *testing.T) {
	var removeTests = []struct {
		value    []rune
		remove   string
		expected string
	}{
		{[]rune("clarionhotel"), "acl", "rionhote"},
		{[]rune("clarionhotel"), "c", "larionhotel"},
		{[]rune("clarionhotel"), "e", "clarionhotl"},
	}
	for _, tt := range removeTests {
		res := remove(tt.value, tt.remove)

		if res != tt.expected {
			t.Errorf("Remove (%d, %d): expected %d, actual %d", tt.value, tt.remove, tt.expected, res)
		}
	}
}
