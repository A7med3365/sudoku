package sudoku

func ArgsCheck(arr [9][9]int) bool {
	// Check if each row contains unique values.
	for i := 0; i < 9; i++ {
		seen := map[int]bool{}
		for j := 0; j < 9; j++ {
			if arr[i][j] == 0 {
				continue
			}
			if _, ok := seen[arr[i][j]]; ok {
				return false
			}
			seen[arr[i][j]] = true
		}
	}

	// Check if each column contains unique values.
	for j := 0; j < 9; j++ {
		seen := map[int]bool{}
		for i := 0; i < 9; i++ {
			if arr[i][j] == 0 {
				continue
			}
			if _, ok := seen[arr[i][j]]; ok {
				return false
			}
			seen[arr[i][j]] = true
		}
	}

	// Check if each 3x3 block contains unique values.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			seen := map[int]bool{}
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if arr[3*i+k][3*j+l] == 0 {
						continue
					}
					if _, ok := seen[arr[3*i+k][3*j+l]]; ok {
						return false
					}
					seen[arr[3*i+k][3*j+l]] = true
				}
			}
		}
	}

	return true
}
