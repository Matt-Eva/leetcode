package main

import "testing"

func TestLargestRectangleArea(t *testing.T) {

	type TestCase struct {
		Case   []int
		Result int
	}

	testCases := []TestCase{
		{
			[]int{2, 1, 2},
			3,
		},
		{
			[]int{2, 1, 5, 6, 2, 3},
			10,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{0, 0},
			0,
		},
	}

	for _, testCase := range testCases {
		result := LargestRectangleArea(testCase.Case)
		if testCase.Result != result {
			t.Errorf("result '%d' does not match test case result '%d'", result, testCase.Result)
		}
	}
}
