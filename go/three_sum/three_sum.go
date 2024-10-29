package main

import (
	"fmt"
	"slices"
)

func main() {

	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(threeSum([]int{0, 1, 1}))
	// fmt.Println(threeSum([]int{0, 0, 0}))
	// fmt.Println(threeSum([]int{-1, 0, 1}))
    // fmt.Println(threeSum([]int{-1, 0, 1, -9, 0, 9, 5, 4}))
    // fmt.Println(threeSum([]int{0,0,0, 0}))
    // fmt.Println(threeSum([]int{-2,0,0,2,2}))
    fmt.Println(threeSum([]int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0}))

}

func threeSum(nums []int) [][]int {

    slices.Sort(nums)
    fmt.Println(nums)

	result := make([][]int, 0)

    for i, num := range nums {
        if (i > 0 && num == nums[i-1]){
            continue
        }

        negativeNum := num * -1
        j := i + 1
        k := len(nums) - 1

        incremented := false
        decremented := false
        for !(j >= k) {
            num2 := nums[j]
            num3 := nums[k]
            if incremented && num2 == nums[j - 1]{
                j++
            } else if decremented && num2 == nums[k + 1]{
                incremented = false
                k--
            } else if (num2 + num3) == negativeNum{
                resultArr := []int{num, num2, num3}
                result = append(result, resultArr)
                j++
                k--
                incremented = true
                decremented = true
            } else if (num2 + num3) < negativeNum{
                j++
                incremented= true
                decremented = false
            } else{
                k--
                incremented = false
                decremented = true
            }
        }
    }

	return result

}

// need to delete all other possible sums
// that means 
