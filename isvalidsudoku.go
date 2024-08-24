package bootcamp

func ValidSudoku(n [9][9]int) bool {
	m := make(map[int][][]int)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if n[i][j] == '.' {
				continue
			}
			locs := m[n[i][j]]
			for k := range locs {
				if locs[k][0] == i {
					return false
				}
				if locs[k][1] == j {
					return false
				}
				if locs[k][0]/3 == i/3 && locs[k][1]/3 == j/3 {
					return false
				}
			}
			m[n[i][j]] = append(m[n[i][j]], []int{i, j})
		}
	}

	return true
}
