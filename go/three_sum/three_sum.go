package main

import "fmt"

func main(){

    threeSum([]int{-1,0,1,2,-1,-4})
    threeSum([]int{0,1,1})
    threeSum([]int{0,0,0})
    
}

func threeSum(nums []int) {
    m := make(map[int][]int)

    for i, num1 := range nums {
        for j, num2 := range nums {
            if j <= i {
                continue
            }
            sum := num1 + num2
            sumArr := []int{num1, num2}
            m[sum] = sumArr
        }
    }

    // fmt.Println(m)

    result := make([][]int, 0)
    for _, num := range nums {
        if len(m[num]) == 2 {
            resultArr := []int{num, m[num][0], m[num][1]}
            result = append(result, resultArr)
        }
    }

    fmt.Println(result)


}
