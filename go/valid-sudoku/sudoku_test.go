package main

import (
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	firstCase := [][]string{[]string{"5","3",".",".","7",".",".",".","."}, []string{"6",".",".","1","9","5",".",".","."}, []string{".","9","8",".",".",".",".","6","."}, []string{"8",".",".",".","6",".",".",".","3"}, []string{"4",".",".","8",".","3",".",".","1"}, []string{"7",".",".",".","2",".",".",".","6"}, []string{".","6",".",".",".",".","2","8","."}, []string{".",".",".","4","1","9",".",".","5"}, []string{".",".",".",".","8",".",".","7","9"}}
	firstCaseResult := true
	firstTestResult := IsValidSudoku(firstCase)
	if firstCaseResult != firstTestResult{
		t.Fatalf("First test results do not match")
	}

	secondCase := [][]string{[]string{"5","3",".",".","7",".",".",".","."}, []string{"6",".","3","1","9","5",".",".","."}, []string{".","9","8",".",".",".",".","6","."}, []string{"8",".",".",".","6",".",".",".","3"}, []string{"4",".",".","8",".","3",".",".","1"}, []string{"7",".",".",".","2",".",".",".","6"}, []string{".","6",".",".",".",".","2","8","."}, []string{".",".",".","4","1","9",".",".","5"}, []string{".",".",".",".","8",".",".","7","9"}}
	secondCaseResult := false
	secondTestResult := IsValidSudoku(secondCase)
	if secondCaseResult != secondTestResult{
		t.Fatalf("Second test results do not match")
	}
}
