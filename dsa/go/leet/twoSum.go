package leet

// problem #1 in leet code -> https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {
	prev_diffs := make(map[int]int)

	for i, val := range nums {
		diff := target - val
		if v, exists := prev_diffs[val]; exists {
			return []int{v, i}
		}
		prev_diffs[diff] = i
	}
	return []int{}
}
