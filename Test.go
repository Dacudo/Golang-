package prose

import (
	"strings"
	"testing"
)

//JoinWithCommas will link words
func JoinWithCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1], ",")

	result += ", and"

	result += phrases[len(phrases)-1]

	return result
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}

	if JoinWithCommas(list) != "apple and orange" {

		t.Error("Didn't match expected value")
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}

	if JoinWithCommas(list) != "apple, orange, and pear" {

		t.Error("Didn't match expected value")
	}
}
