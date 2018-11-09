package main

import (
	"fmt"
)

// check row for an element
func checkRow(sudoku *[9][9]int, row, elem int) bool {
	for i := 0; i < 9; i++ {
		if sudoku[row][i] == elem {
			return false
		}
	}
	return true
}

// check column for an element
func checkCol(sudoku *[9][9]int, col, elem int) bool {
	for i := 0; i < 9; i++ {
		if sudoku[i][col] == elem {
			return false
		}
	}
	return true
}

// check 3x3 square for an element
func checkSquare(sudoku *[9][9]int, row, col, elem int) bool {
	row = (row / 3) * 3
	col = (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku[row+i][col+j] == elem {
				return false
			}
		}
	}
	return true
}

func solve_sudoku(sudoku *[9][9]int, row, col int) {
	for k := 0; k < 9; k++ {
		if checkRow(sudoku, row, k) &&
			checkCol(sudoku, col, k) &&
			checkSquare(sudoku, row, col, k) {
			// set value
			sudoku[row][col] = k
		}
	}
}

func main() {
	sudoku := [9][9]int{
		{4, 0, 5, 0, 1, 0, 2, 0, 9},
		{7, 0, 0, 0, 9, 0, 0, 8, 0},
		{8, 0, 0, 6, 4, 3, 1, 0, 5},

		{0, 0, 0, 0, 0, 4, 0, 0, 0},
		{2, 0, 0, 5, 7, 8, 0, 0, 6},
		{0, 0, 0, 9, 0, 0, 0, 0, 0},

		{9, 0, 3, 7, 5, 6, 0, 0, 1},
		{0, 8, 0, 0, 3, 0, 0, 0, 7},
		{1, 0, 2, 0, 8, 0, 5, 0, 3},
	}

	fmt.Println(sudoku)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			solve_sudoku(&sudoku, i, j)
		}
	}
	fmt.Println(sudoku)
}
