package strings

import (
	"reflect"
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
			t.Errorf("Revert (%d): expected %d, actual %d", string(tt.value), string(tt.expected), string(res))
		}
	}
}

func TestRevertInline(t *testing.T) {
	var removeTests = []struct {
		value    []rune
		expected []rune
	}{
		{[]rune("first second third"), []rune("third second first")},
		{[]rune("first second"), []rune("second first")},
	}
	for _, tt := range removeTests {
		reversedValue := make([]rune, len(tt.value))
		copy(reversedValue, tt.value)
		reverseInline(reversedValue)

		if !reflect.DeepEqual(reversedValue, tt.expected) {
			t.Errorf("Revert (%d): expected %d, actual %d", string(tt.value), string(tt.expected), string(reversedValue))
		}
	}
}

func TestRevertString(t *testing.T) {

	var removeTests = []struct {
		value    []rune
		expected []rune
	}{
		{[]rune("abc"), []rune("cba")},
	}
	for _, tt := range removeTests {
		reversedValue := make([]rune, len(tt.value))
		copy(reversedValue, tt.value)
		reverseString(reversedValue)

		if !reflect.DeepEqual(reversedValue, tt.expected) {
			t.Errorf("Revert (%d): expected %d, actual %d", string(tt.value), string(tt.expected), string(reversedValue))
		}
	}
}
