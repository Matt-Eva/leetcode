package main

import (
	"fmt"
)

func main() {

}

// fastest solution - not mine - just for study
// brilliant
// registered slower than previous solution
// func dailyTemperatures(temperatures []int) []int {
//     result := make([]int, len(temperatures))
    
// 	// iterating backward
//     for i := len(temperatures) - 1; i >= 0; i-- {
// 		// but incrementing j forward
//         j := i + 1
        
// 		// while j is not out of bounds and the index that j represents
// 		// is less than the current temperature
//         for j < len(temperatures) && temperatures[j] <= temperatures[i] {
// 			// if the result slice has a value of zero, we break - there
// 			// are no future values that are greater than the value at j
//             if result[j] <= 0 {
//                 break
//             }
			
// 			// otherwise, we will add the number of days until it is warmer
// 			// than the day represented by j to j.
// 			// this is the minimum number of days we would need to move 
// 			// forward by, since temp[j] <= temp[i] and this would be the first
// 			// occurence where something is > temp[j]
//             j += result[j]
//         } 
        
// 		// now that we have value of j
// 		// check if temperatures[j] is > temperatures i
// 		// if so, add that info to the result
//         if j < len(temperatures) && temperatures[j] > temperatures[i] {
//             result[i] = j - i
//         }
        
//     }
//     return result
// }


// working solution with stack
func dailyTemperatures(temperatures []int) []int {
	type Day struct {
		Temp  int
		Index int
	}

	result := make([]int, len(temperatures))
	stack := make([]Day, 0, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > stack[len(stack)-1].Temp {
			index := stack[len(stack)-1].Index
			stack = stack[:len(stack)-1]
			result[index] = i - index
		}

		newDay := Day{
			temperatures[i],
			i,
		}
		stack = append(stack, newDay)
	}

	return result
}

// slow map solution - passes
// func dailyTemperatures(temperatures []int) []int {
// 	length := len(temperatures)
// 	result := make([]int, length, length)

// 	tempMap := make(map[int][]int)
// 	min := temperatures[0]

// 	for index, val := range temperatures {
// 		_, ok := tempMap[val]

// 		if ok {
// 			tempMap[val] = append(tempMap[val], index)
// 		} else {
// 			tempMap[val] = []int{index}
// 		}

// 		if val <= min{
// 			min = val
// 		} else {
// 			for key, slice := range tempMap {
// 				if val > key{
// 					for _, i := range slice {
// 						result[i] = index - i
// 					}
// 					delete(tempMap, key)
// 				}
// 			}
// 		}
// 	}

// 	return result
// }

//simple O n^2 solution
// exceeds time limit
// func dailyTemperatures(temperatures []int) []int {
// 	result := make([]int, 0, len(temperatures))

//     for index, val := range temperatures{
// 		count := 0
// 		for i := index + 1; i < len(temperatures); i++ {
// 			if temperatures[i] > val {
// 				count += i - index
// 				break
// 			}
// 		}
// 		result = append(result, count)
// 	}

// 	return result
// }
