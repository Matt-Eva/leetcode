package main

// O(N) struct example - final solution
// Ironically, the solution that just uses slices.Sort runs the fastest
// and uses the least memory

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	type NumStruct struct {
		Visited bool
	}

	set := make(map[int]*NumStruct)
	max := 1

	for _, val := range nums {
		set[val] = &NumStruct{
			false,
		}
	}

	for _, num := range nums {
		if set[num].Visited {
			continue
		}

		count := -1
		for current := num; set[current] != nil; current++ {
			count++
			set[current].Visited = true
		}

		for current := num; set[current] != nil; current-- {
			count++
			set[current].Visited = true
		}

		if count > max {
			max = count
		}
	}

	return max
}

// MAP SOLUTION

// func longestConsecutive(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	set := make(map[int]bool)
// 	max := 1

// 	for _, val := range nums {
// 		set[val] = true
// 	}

// 	for _, num := range nums {
// 		if set[num-1] {
// 			continue
// 		}

// 		count := 0
// 		for current := num; set[current]; current++ {
// 			count++
// 		}

// 		if count > max {
// 			max = count
// 		}
// 	}

// 	return max
// }

//     return newMax
// }

// OVERLY COMPLICATED HEAP SOLUTION

// func main() {
// 	longestConsecutive([]int{100,4,200,1,3,2})
// }

// func longestConsecutive(nums []int) int {
// 	numMap := make(map[int]bool)
// 	numHeap := &IntHeap{}
// 	heap.Init(numHeap)

// 	for _, num := range nums{
// 		if !numMap[num]{
// 			numMap[num] = true
// 			heap.Push(numHeap, num)
// 		}
// 	}

// 	max := 0
// 	for numHeap.Len() > 0 {
// 		newMax := 0
// 		for minimum := (*numHeap)[0]; ;minimum++ {
// 			if numMap[minimum]{
// 				newMax += 1
// 				heap.Pop(numHeap)

// 				if newMax > max{
// 					max = newMax
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 	}

// 	return max
// }

// type IntHeap []int

// func (h IntHeap) Len() int {return len(h)}
// func (h IntHeap) Less(i, j int) bool {return h[i] < h[j]}
// func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}

// func (h *IntHeap) Pop() any {
// 	old := *h
// 	length := len(old)
// 	leaf := old[length - 1]
// 	*h = old[0 : length - 1]
// 	return leaf
// }

// func (h *IntHeap) Push(x any) {
// 	if integer, ok := x.(int); ok {
// 		*h = append(*h, integer)
// 	}
// }
