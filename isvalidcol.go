package sudoku

// import "fmt"


//this function taks the sudoku array ( arr ) and check if the number passed ( nbr ) is valid for the givin column ( col ), 
//and returns boolean value for the result
func IsValidCol(arr [9][9]int, nbr int, col int) bool {

	for i := 0; i < 9; i++ {
		if arr[i][col] == nbr {
			return false
		}
	}
	return true
}