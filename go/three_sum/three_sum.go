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

	sumMap := make(map[int][]map[string]int)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			keyNum := nums[i] + nums[j]
			sumMap[keyNum] = []map[string]int{
				{
					"index": i,
					"num":   nums[i],
				},
				{
					"index": j,
					"num":   nums[j],
				},
			}
		}
	}

	fmt.Println("populatedSumMap", sumMap)

	result := make([][]int, 0)

	for i, num := range nums {
		negativeNum := num * -1
		if len(sumMap[negativeNum]) == 2 {
			duplicateIndex := false
			for _, numMap := range sumMap[negativeNum] {
				if numMap["index"] == i {
					duplicateIndex = true
				}
			}

			if !duplicateIndex {
				resultArr := []int{num, sumMap[negativeNum][0]["num"], sumMap[negativeNum][1]["num"]}
                result = append(result, resultArr)
			}
		}
	}

	// fmt.Println("end sumMap", sumMap)

	return result

}
