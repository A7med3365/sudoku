package sudoku

// this function check if the number is vaild for the given block
func IsValid(arr [9][9]int, nbr int, row int, col int) bool {
	if IsValidBlock(arr, nbr, row, col) && IsValidCol(arr, nbr, col) && IsValidRow(arr, nbr, row) {
		return true
	} else {
		return false
	}
}