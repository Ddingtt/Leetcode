package leetcode_go

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, v := range nums {
		numSet[v] = true
	}

	longestStreak := 0
	for num := range numSet {
		// 如果当前数字的前一个数字不在哈希集合中，则当前数字是连续序列的开头
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			// 继续递增当前数字，直到连续序列结束
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			// 更新最长连续序列的长度
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}

func main() {
	println(longestConsecutive([]int{}))
}
