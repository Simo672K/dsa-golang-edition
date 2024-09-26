package leet

// problem #1 in leet code -> https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {
	prev_diffs := make(map[int]int)

	for i, val := range nums {
		diff := val - target
		if v, exists := prev_diffs[diff]; exists {
			return []int{i, v}
		}
	}
	return []int{}
}
