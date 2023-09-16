package sudoku

import "fmt"

func Solve(arr [9][9]int, row int, col int) bool {
	// defer func() { arr[row][col] = 0 }()
	if row == 9 {
		return true
	}

	if arr[row][col] != 0 {
		nextRow , nextCol := NextBlock(row, col)
		// fmt.Println("if")
		return Solve(arr, nextRow, nextCol)
		
		// fmt.Println(arr[row][col])
	} else {
		for i := 1; i <= 9; i++ {
			if IsValid(arr, i, row, col) {
				arr[row][col] = i
				PrintSudoku(arr)
				nextRow , nextCol := NextBlock(row, col)
				if Solve(arr, nextRow, nextCol) {
					return true
				}
				arr[row][col] = 0
				
			}
		}
		fmt.Println("else for out")
		return false
	}
	fmt.Println("out")
	return false
}