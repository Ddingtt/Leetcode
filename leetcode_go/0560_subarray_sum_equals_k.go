package main

func subarraySum(nums []int, k int) int {
	i, j := 0, 0
	res := 0
	for i <= j {
		sum := 0
		if i == j {
			sum = nums[i]
		} else {
			sum = nums[i] + nums[j]
		}

		if sum == k {
			res++
			i++
			j++
		} else if sum < k {
			j++
		} else if sum > k {
			i++
		}
		if j > len(nums)-1 {
			break
		}
	}
	return res
}

func main() {
	println(subarraySum([]int{1, 1, 1}, 2))
}
