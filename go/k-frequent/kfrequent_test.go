package main

import (
	"fmt"
	"testing"
)

func TestKFrequent(t *testing.T) {
	type Argument struct {
		Nums      []int
		Frequency int
		Result    []int
	}

	cases := []Argument{
		{
			[]int{1, 1, 2, 2, 3},
			2,
			[]int{1, 2},
		},
		{
			[]int{3, 4, 5, 5},
			1,
			[]int{5},
		},
		{
			[]int{2, 2, 2, 8, 8, -1, -1, -1, -1},
			2,
			[]int{2, -1},
		},
	}

	for _, singleCase := range cases {
		res := topKFrequent(singleCase.Nums, singleCase.Frequency)

		for _, num := range singleCase.Result {
			match := contains(res, num)

			if !match {
				t.Errorf("result does not match")
			}
		}
	}
}

func contains(result []int, value int) bool {
	for _, integer := range result {
		if integer == value {
			return true
		}
	}
	return false
}
