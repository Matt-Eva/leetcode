package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
    // boxes are read from top-bottom, left to right
    // box two is left most in middle

    for i := 0; i < 9; i++{
        row := board[i]
        rowMap := make(map[byte]boolean)
        for _, val:= range row {
            if rowMap[val]{
                return false
            }
            rowMap[val] = true
        }

        columnMap := make(map[byte]boolean)
        for n := 0; n < 9; n++{
            row := board[n]
            val := row[i]
            if columnMap[val]{
                return false
            }
            columnMap[val] = true
        }
    }

    for i := 0; i < 9; i+=3{
        for n := i; n < i + 3; n++{
            box := make(map[byte]boolean)
        }
    }

    return true
}