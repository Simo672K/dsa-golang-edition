package leet

// import "fmt"

func ProductExceptSelf(nums []int) []int {
	l := len(nums) - 1
	prefixProd := make([]int, len(nums))
	sufixProd := make([]int, len(nums))
	answer := make([]int, len(nums))

	prefixProd[0] = 1
	sufixProd[l] = 1

	pp := 1
	sp := 1

	for i := 1; i <= l; i++ {
		pp *= nums[i-1]
		sp *= nums[l+1-i]

		prefixProd[i] = pp
		sufixProd[l-i] = sp
	}

	// fmt.Println()
	for i := range answer {
		answer[i] = prefixProd[i] * sufixProd[i]
	}

	return answer
}

/*

	[ 1 , 2 , 3 , 4 ]
	  ^   ^   ^   ^

*/
