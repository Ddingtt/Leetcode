package leetcode_go

func setZeroes(matrix [][]int) {
	var tmpX []int
	var tmpY []int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				tmpX = append(tmpX, i)
				tmpY = append(tmpY, j)
			}
		}
	}

	for _, x := range tmpX {
		for i2 := range matrix[x] {
			matrix[x][i2] = 0
		}
	}

	for _, y := range tmpY {
		for _, m := range matrix {
			m[y] = 0
		}
	}
}
