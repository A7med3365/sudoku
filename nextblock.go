package sudoku

func NextBlock(row int, col int) (int, int) {
	if col == 8 {
		return row+1, 0
	} else {
		return row, col+1
	}
}