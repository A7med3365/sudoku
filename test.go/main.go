package main 

import "sudoku"
import "fmt"
import "os"

func main() {
	args := os.Args[1:]

	if len(args) != 9 {
		fmt.Println("Error")
		return
	}
	for i := 0; i <= 8; i++ {
		if len(args[i]) != 9 {
			fmt.Println("Error")
			return
		}
	}

	arr := [9][9]int{}

	for i, str := range args {
		for j, v := range str {
			if v == '.' {
				arr[i][j] = 0
			} else {
				nbr := v - '0'
				arr[i][j] = int(nbr)
			}
		}
	}

	sudoku.PrintSudoku(arr)

	if sudoku.Solve(arr, 0, 0) {
		fmt.Println("solution found")
	} else {
		fmt.Println("no solution found")
	}

}