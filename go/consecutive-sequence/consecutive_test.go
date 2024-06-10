package main

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	case1 := []int{100, 4, 200, 1, 3, 2}
	answer1 := 4
	testResult1 := longestConsecutive(case1)
	fmt.Println(testResult1)
	if answer1 != testResult1 {
		t.Fatal("first results did not match!")
	}
	case2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	answer2 := 10
	testResult2 := longestConsecutive(case2)
	if answer2 != testResult2 {
		t.Fatal("Second results did not match!")
	}
}
