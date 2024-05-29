package main

//  "fmt"

func IsValidSudoku(board [][]string) bool {
	// boxes are read from top-bottom, left to right
	// box two is left most in middle

	for i := 0; i < 9; i++ {
		row := board[i]
		rowMap := make(map[string]bool)
		for _, val := range row {
			if rowMap[val] {
				return false
			} else if val != "." {
				rowMap[val] = true
			}
		}

		columnMap := make(map[string]bool)
		for n := 0; n < 9; n++ {
			row := board[n]
			val := row[i]
			if columnMap[val] {
				return false
			} else if val != "." {
				columnMap[val] = true
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for n := 0; n < 9; n += 3 {
			isValid := validateBox(i, n, board)
			if !isValid {
				return false
			}
		}
	}

	// iterate through columns, rows, and boxes 3 at a time

	return true
}

func validateBox(i, n int, board [][]string) bool {
	box := make(map[string]bool)
	for b := n; b < n+3; b++ {
		row := board[b]
		for a := i; a < i+3; a++ {
			val := row[a]
			if box[val] {
				return false
			} else if val != "." {
				box[val] = true
			}
		}
	}
	return true
}
