package main

import (
	"testing"
)

func TestProdOfArrayExceptSelf(t * testing.T) {
	firstCase := []int{1, 2, 3, 4}
	firstAnswer := []int{24, 12, 8, 6}
	result := ProdOfArrayExceptSelf(firstCase)
	for index, val := range result {
		if val != firstAnswer[index]{
			t.Fatalf(`Non matching entries at %d: answer value %d does not equal case value %d`, index, val, firstAnswer[index])
		}
	}
}