package tests

import (
	"testing"

	"github.com/Simo672K/dsa/leet"
)

func TestContainsDuplicate(t *testing.T) {
	test1 := []int{1, 2, 3}
	test2 := []int{1, 2, 2}

	got := leet.ContainsDuplicate(test1)
	got2 := leet.ContainsDuplicate(test2)

	expected := false
	expected2 := true

	if got != expected {
		t.Errorf("for %v , expected %t got %t", test1, expected, got)
		return
	}

	if got2 != expected2 {
		t.Errorf("for %v , expected %t got %t", test2, expected2, got2)
		return
	}
}
