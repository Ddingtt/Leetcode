package leetcode_go

func groupAnagrams(strs []string) [][]string {
	var res [][]string

	for _, str := range strs {
		if len(res) == 0 {
			res = append(res, []string{str})
			continue
		}

		for _, strIn := range res {
			if isEquel(strIn[0], str) {
				strIn = append(strIn, str)
			} else {
				res = append(res, []string{str})
			}
		}
	}

	return res
}

func isEquel(strKey string, str string) bool {
	if len(strKey) != len(str) {
		return false
	}

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(strKey); j++ {
			if str[i:i+1] != strKey[j:j+1] {
				return false
			}
		}
	}
	return true
}
