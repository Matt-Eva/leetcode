package main

import "testing"

func TestCarFleet(t *testing.T) {
	carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})
	carFleet(10, []int{3}, []int{3})
	carFleet(10, []int{6, 8}, []int{3, 2})
	carFleet(13, []int{10, 2, 5, 7, 4, 6, 11}, []int{7, 5, 10, 5, 9, 4, 1})
}
