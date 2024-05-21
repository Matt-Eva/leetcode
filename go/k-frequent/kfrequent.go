package main

import (
	"container/heap"
	"fmt"
)

func main() {
	result := topKFrequent([]int{1, 1, 2, 2, 3}, 2)
	fmt.Println(result)
}

func topKFrequent(nums []int, k int) []int {
	// m := make(map[int]int)

	// for _, v := range nums {
	// 	_, ok := m[v]
	// 	if ok {
	// 		m[v] += 1
	// 	} else {
	// 		m[v] = 1
	// 	}
	// }

	// fmt.Println(m)
	// max := 1
	// for _, v := range m {
	// 	if v > max {
	// 		max = v
	// 	}
	// }

	// s := make([][]int, max+1)
	// for key, v := range m {
	// 	fmt.Println(s[v])
	// 	if s[v] == nil {
	// 		s[v] = []int{key}
	// 	} else {
	// 		s[v] = append(s[v], key)
	// 	}
	// }

	// solution := make([]int, 0, k)
	// for kCount, i := k, len(s)-1; kCount > 0 && i > 0; i-- {
	// 	if s[i] != nil {
	// 		for _, val := range s[i] {
	// 			solution = append(solution, val)
	// 			kCount--

	// 			if kCount == 0 {
	// 				return solution
	// 			}
	// 		}
	// 	}
	// }

	// return solution

	// ^^ My solution - bad. Need to look into using Heaps.
	// https://en.wikipedia.org/wiki/Heap_(data_structure)
	// https://pkg.go.dev/container/heap

	myMap := make(map[int]int)

	for _, num := range nums {
		myMap[num]++
	}

	h := &StructHeap{}

	heap.Init(h)
	for num, frequency := range myMap {
		fs := &FrequencyStruct{
			Number:    num,
			Frequency: frequency,
		}
		heap.Push(h, fs)
	}

	result := make([]int, 0, k)
	minLen := h.Len() - k
	for h.Len() > minLen  {
		popped := heap.Pop(h).(*FrequencyStruct)
		number := popped.Number
		fmt.Println(number)

		result = append(result, number)
	}

	return result
}

type FrequencyStruct struct {
	Number    int
	Frequency int
}

type StructHeap []*FrequencyStruct

func (h StructHeap) Len() int { return len(h) }

// setting Less to a greater than operation turns it into a Max Heap instead of a Min Heap
func (h StructHeap) Less(i, j int) bool { return h[i].Frequency > h[j].Frequency }
func (h StructHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StructHeap) Push(x any) {
	item := x.(*FrequencyStruct)
	*h = append(*h, item)
}

func (h *StructHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil // need to dereference struct to prevent memory leak
	*h = old[0 : n-1]
	return x
}

// Heap Notes from Claude AI:
// In a binary heap implemented using an array, the leaf nodes are guaranteed to be in the range n/2 + 1 to n (inclusive), where n is the length of the array. This property arises from the way the heap is represented in the array.
// In a binary heap, the root node is stored at index 0 in the array. For any node at index i, its left child is at index 2*i + 1, and its right child is at index 2*i + 2. This relationship is based on the array indexing, where the first element is at index 0.
// Now, let's consider the last node in the array, which is at index n-1. Since this node has no children (it's a leaf node), its children's indices would be 2*(n-1) + 1 = 2n - 1 and 2*(n-1) + 2 = 2n, which are both out of bounds for an array of length n.
// Therefore, the last node that can have children in an array of length n is the node at index n/2 - 1. This is because the children of this node would be at indices 2*(n/2 - 1) + 1 = n - 1 and 2*(n/2 - 1) + 2 = n, which are the last valid indices in the array.
// This means that all the nodes from index n/2 onwards are leaf nodes, as they have no children within the bounds of the array. Hence, the leaf nodes in a binary heap implemented using an array of length n are guaranteed to be at indices n/2 + 1, n/2 + 2, ..., n-1, n.
// In summary, the leaf nodes in a binary heap array are guaranteed to be in the range n/2 + 1 to n because of the way the heap is represented in the array, with the children of a node at index i being located at indices 2*i + 1 and 2*i + 2.

// Standford paper on Heaps: https://web.stanford.edu/class/archive/cs/cs161/cs161.1168/lecture4.pdf

// A note on the Go Heap implementation included in the standard library:

// The Pop IntHeap function that we see for the min IntHeap method was initially confusing.
// It seemed to be removing the final element from the end of the slice, rather than the first one.
// However, this makes sense - we only know that the first element is the minimum.
// We want to make the underlying slice implementing the heap shorter, so that we can remove the initial element
// However, removing the first element would mean that we would need to possibly need to re-index the new slice?
// There is no spec on this, but perhaps its true?
// Therefore, we can remove the final element, but return it from the function to save it.
// Then the heap implementation can try and
