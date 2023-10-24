package main

import "fmt"

func productExceptSelf(nums []int) []int {
	length := len(nums)
	L, R, ans := make([]int, length), make([]int, length), make([]int, length)

	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = nums[i-1] * L[i-1]
	}
	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}

	// 对于索引 i，除 nums[i] 之外其余各元素的乘积就是左侧所有元素的乘积乘以右侧所有元素的乘积
	for i := 0; i < length; i++ {
		ans[i] = L[i] * R[i]
	}
	return ans
}

func main() {
	t := []int{1, 2, 3, 4}

	for _, v := range productExceptSelf(t) {
		fmt.Printf(" %d", v)
	}
}
