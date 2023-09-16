package sudoku

// import "fmt"

// this function taks the sudoku array ( arr ) and check if the number passed ( nbr ) is valid for the the block its in
// based on the givin row and column ( row, col ), 
// and returns boolean value for the result
func IsValidBlock(arr [9][9]int, nbr int, row int, col int) bool {
	// fmt.Printf("block %d %d", row/3, col/3)
	blockX := col/3
	blockY := row/3

	for i := blockY*3; i < blockY*3+3; i++ {
		for j := blockX*3; j < blockX*3+3; j++ {
				if arr[i][j] == nbr {
					return false
				}
			}		
	}
	return true
}