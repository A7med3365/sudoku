package sudoku

// import "fmt"


//this function taks the sudoku array ( arr ) and check if the number passed ( nbr ) is valid for the givin row ( row ), 
//and returns boolean value for the result
func IsValidRow(arr [9][9]int, nbr int, row int) bool {
	rowToCheck := arr[row][:]

	for _, n := range rowToCheck {
		if n == nbr {
			return false
		}
	}
	return true
}