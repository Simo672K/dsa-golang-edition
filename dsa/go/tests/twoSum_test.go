package tests

import (
	"testing"

	"github.com/Simo672K/dsa/leet"
)

func TestTwoSum(t *testing.T) {
	test1_input := []int{2, 7, 11, 15}
	target := 9
	sum := 0

	test1_output := leet.TwoSum(test1_input, target)
	for _, v := range test1_output {
		sum += test1_input[v]
	}

	if target != sum {
		t.Error("test failed")
	}
}
