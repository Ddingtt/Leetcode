package leetcode_go

import (
	"sort"
	"strings"
)

func findAnagrams(s string, p string) []int {
	l := 0
	r := len(p) - 1
	var res []int

	for r < len(s) {
		if isAna(s[l:r+1], p) {
			res = append(res, l)
		}
		l++
		r++
	}

	return res
}

func isAna(s string, p string) bool {
	ss, pp := []rune(s), []rune(p)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})

	sort.Slice(pp, func(i, j int) bool {
		return pp[i] < pp[j]
	})
	if strings.Compare(string(ss), string(pp)) == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	s, p := "cbaebabacd", "abc"
	for i := range s[:len(s)-len(p)] {
		println(i)
		println(s[i])
	}
}
