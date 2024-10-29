package main

import "fmt"

func main() {

	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(threeSum([]int{0, 1, 1}))
	// fmt.Println(threeSum([]int{0, 0, 0}))
	// fmt.Println(threeSum([]int{-1, 0, 1}))
    fmt.Println(threeSum([]int{-1, 0, 1, -9, 0, 9, 5, 4}))

}

func threeSum(nums []int) [][]int {


	result := make([][]int, 0)

    for i, num := range nums {
        
        numMap := make(map[int][]int)
        for j := i + 1; j < len(nums); j++{
            numMap[nums[j] + num] = []int{nums[j], num}
        }
        for j := i + 1; j < len(nums); j++ {
            if len(numMap[nums[j] * -1]) == 2{
                resultArr := []int{nums[j], numMap[nums[j] * -1][0], numMap[nums[j] * -1][1]}
                result = append(result, resultArr)
            }
        }
    }

	return result

}

// need to delete all other possible sums
// that means 
