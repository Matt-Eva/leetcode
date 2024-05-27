package main

func ProdOfArrayExceptSelf(nums []int) []int {
	length := len(nums)
    answer := make([]int, length)

    for i := 0; i < length; i++ {
      answer[i] = 1
    }

    rightValue, leftValue := 1, 1
    for leftIndex := 0; leftIndex < length; leftIndex++ {
        rightIndex := length - leftIndex - 1
        answer[leftIndex] = answer[leftIndex] * leftValue 
        leftValue *= nums[leftIndex]
        answer[rightIndex] *= rightValue
        rightValue *= nums[rightIndex]
    }

    return answer
}