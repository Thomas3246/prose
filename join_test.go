package prose

import "testing"

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	if JoinWinthCommas(list) != "apple and orange" {
		t.Error("didn't match expected value")
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	if JoinWinthCommas(list) != "apple, orange, and pear" {
		t.Error("didn't match expected value")
	}
}
