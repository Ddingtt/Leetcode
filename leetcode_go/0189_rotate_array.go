package leetcode_go

func rotate(nums []int, k int) {
	ans := make([]int, len(nums))

	for i, v := range nums {
		ans[(i+k)%len(nums)] = v
	}

	copy(nums, ans)
}
