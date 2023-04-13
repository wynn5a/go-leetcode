package solution36

// isValidSudoku checks if a given 9x9 Sudoku board is valid by ensuring that each row, column, and 3x3 box contains only unique numbers from 1 to 9.
// It returns true if the board is valid, and false otherwise.
func isValidSudoku(board [][]byte) bool {
	// initialize maps to keep track of numbers in rows, columns, and boxes
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	// iterate through each cell in the board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != '.' {
				// check if the number already exists in the row, column, or box
				if rows[i][num] || cols[j][num] || boxes[(i/3)*3+j/3][num] {
					return false
				}
				// if the number doesn't exist in the row, column, or box, add it to the maps
				rows[i][num] = true
				cols[j][num] = true
				boxes[(i/3)*3+j/3][num] = true
			}
		}
	}
	return true
}
