package leetcode_go

func moveZeroes(nums []int) {
	for i, num := range nums {
		if num != 0 {
			continue
		}
		if num == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i] = nums[j]
					nums[j] = 0
					break
				}
			}
		}
	}
}
