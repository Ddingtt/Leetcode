package leetcode_go

func maxArea(height []int) int {
	res := 0
	for i, j := 0, len(height)-1; i < j; {
		res = max(res, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max1(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	println(maxArea([]int{1, 2}))
}
