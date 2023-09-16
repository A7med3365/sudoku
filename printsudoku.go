package sudoku

import (
	"fmt"
	"time"
)

func printSudoku(s [9][9]int) {
	fmt.Print("\033[2J") // clear the screen
    fmt.Print("\033[H")  // move the cursor to the top-left corner
	fmt.Println("╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗")
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if j == 0 {
                fmt.Print("║")
            }
            if s[i][j] == 0 {
                fmt.Print("   ")
            } else {
                fmt.Printf(" %d ", s[i][j])
            }
            if (j+1)%3 == 0 {
                fmt.Print("║")
            } else {
                fmt.Print("│")
            }
        }
        fmt.Println()
        if (i+1)%3 == 0 && i != 8 {
            fmt.Println("╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣")
        } else if i != 8 {
            fmt.Println("╟───┼───┼───╫───┼───┼───╫───┼───┼───╢")
        }
    }
    fmt.Println("╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝")
	time.Sleep(20 * time.Millisecond) // wait for 500 milliseconds
}