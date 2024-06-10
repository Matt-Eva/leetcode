package main

//  "fmt"

func isValidSudoku(board [][]byte) bool {
    period := "."
    periodByte := period[0]

	for i := 0; i < 9; i++ {
		row := board[i]
		rowMap := make(map[byte]bool)
		for _, val := range row {
			if rowMap[val] {
				return false
			} else if val != periodByte {
				rowMap[val] = true
			}
		}

		columnMap := make(map[byte]bool)
		for n := 0; n < 9; n++ {
			row := board[n]
			val := row[i]
			if columnMap[val] {
				return false
			} else if val != periodByte {
				columnMap[val] = true
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for n := 0; n < 9; n += 3 {
			isValid := validateBox(i, n, board, periodByte)
			if !isValid {
				return false
			}
		}
	}

	return true
}

func validateBox(i, n int, board [][]byte, periodByte byte) bool {
	box := make(map[byte]bool)
	for b := n; b < n+3; b++ {
		row := board[b]
		for a := i; a < i+3; a++ {
			val := row[a]
			if box[val] {
				return false
			} else if val != periodByte {
				box[val] = true
			}
		}
	}
	return true
}
