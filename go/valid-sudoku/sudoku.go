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
			} else if val != "."{
                rowMap[val] = true
            }
		}

		columnMap := make(map[string]bool)
		for n := 0; n < 9; n++ {
			row := board[n]
			val := row[i]
			if columnMap[val] {
				return false
			} else if val != "."{
                columnMap[val] = true
            }
		}
	}

	for i := 0; i < 9; i+=3 {
        box1 := make(map[string]bool)
		for n := 0; n < 3; n++{
            row := board[n]
            for b:= i; b < i + 3; b++{
                val := row[b]
                if box1[val]{
                    return false
                } else if val != "."{
                    box1[val] = true
                }
            }
        }
        box2 := make(map[string]bool)
        for n:= 3; n < 6; n++{
            row := board[n]
            for b:= i; b < i + 3; b++{
                val := row[b]
                if box2[val]{
                    return false
                } else if val != "."{
                    box2[val] = true
                }
            }
        }
        box3 := make(map[string]bool)
        for n:= 6; n < 9; n++{
            row := board[n]
            for b:= i; b < i + 3; b++{
                val := row[b]
                if box3[val]{
                    return false
                } else if val != "."{
                    box3[val] = true
                }
            }
        }
	}

	// iterate through columns, rows, and boxes 3 at a time

	return true
}

func validateByThree(i int, board [][]string) bool {
	rowMap := make(map[string]bool)
	columnMap := make(map[string]bool)

	for n := i; n < i+3; n++ {
		row := board[i]
		for _, val := range row {
			if rowMap[val] {
				return false
			}
			rowMap[val] = true
		}
		for b := 0; b < 9; b++ {
			column := board[b]
			val := column[i]
			if columnMap[val] {
				return false
			}
			columnMap[val] = true
		}
	}
	return true
}
