package main

import (
	"fmt"

	"github.com/Simo672K/dsa/leet"
)

// main entry function
func main() {
	// dsa.InitSLL()
	// leet.ContainsDuplicate([]int{1, 2, 2})
	out := leet.TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(out)
}
