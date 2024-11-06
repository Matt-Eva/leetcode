package main

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	trap([]int{4,2,0,3,2,5})
	trap([]int{6, 4, 2, 0, 3, 2, 0, 3, 1, 4, 5, 3, 2, 7, 5, 3, 0, 1, 2, 1, 3, 4, 6, 8, 1, 3})

	// fmt.Println(1 + 3 + 1 + 3 + 2 + 2 + 3 + 2 + 4 + 7 + 6 + 5 + 6 + 4 + 1 + 2)
}

func trap(height []int) int {
	volume := 0

	leftMax := 0
	rightMax := 0
	i := 0
	j := len(height) - 1

	for i < j{
		if height[i] < height[j]{
			 leftMax = max(leftMax, height[i])
			 volume += leftMax - height[i]
			i++
		} else {
			rightMax = max(rightMax, height[j])
			volume += rightMax - height[j]
			j--
		}
	}

	return volume
}

// func trap(height []int) int {
	
// 	volume := 0
	
// 	leftMax := 0
// 	for i := 0; i < len(height); i++{
// 		if height[i] > leftMax {
// 			leftMax = height[i]
// 		} else {
// 			volume += leftMax - height[i]
// 		}
// 	}

// 	rightMax := 0
// 	for i := len(height) - 1; i > 0; i-- {
// 		if height[i] > rightMax{
// 			rightMax = height[i]
// 		}
		
// 		if leftMax > height[i]{
// 			volume -= leftMax - rightMax
// 		}
// 	}

// 	return volume
// }

// func trap(height []int) int {

// 	l := len(height)
// 	if l < 2 {
// 		return 0
// 	}

// 	volume := 0
	
// 	leftSide := height[0]
// 	leftSides := make([]int, l)
// 	for i, h := range height {
// 		if h < leftSide {
// 			leftSides[i] = leftSide
// 		} else {
// 			leftSide = h
// 		}
// 	}

// 	rightSide := height[l-1]
// 	for j := l - 1; j > 0; j-- {
// 		if height[j] < rightSide && height[j] < leftSides[j] {
// 			minHeight := min(rightSide, leftSides[j])
// 			volume += minHeight - height[j]
// 		} else {
// 			rightSide = height[j]
// 		}
// 	}

// 	return volume
// }

// func trap(height []int) int {

// 	volume := 0
// 	peaks := make([][]int, 0)
// 	// peaks is slice of slices - internal slices have index as first element,
// 	// height as second element

// 	for i, h := range height {
// 		if len(peaks) == 0 && i < len(height) - 1 && h > height[i + 1]{
// 			peaks = append(peaks, []int{i, h})
// 			continue
// 		}
// 		// accounts for empty slice where bowl begins.

// 		if i > 0 && h > height[i - 1]{
// 			for len(peaks) > 1 && peaks[len(peaks) - 1][1] <= peaks[len(peaks) - 2][1] && peaks[len(peaks) - 1][1] <= h {
// 				peaks = peaks[:len(peaks) - 1]
// 			}
// 			peaks = append(peaks, []int{i, h})
// 			continue
// 		}
// 		// removes peaks that occur between the start and end of a bowl - adds new peak
// 		// as end of bowl.

// 		if i > 0 && i < len(height) - 1 && h > height[i - 1] && h < height[i + 1]{
// 			peaks = append(peaks, []int{i, h})
// 			continue
// 		}
// 		// adds peaks that may occur within a bowl, or in general.

// 		if len(peaks) > 0 && i < len(height) - 1 && h >= peaks[len(peaks) - 1][1] && h > height[i + 1]{
// 			peaks = append(peaks, []int{i, h})
// 		}
// 		// accounts for a plateau that may occur.
// 	}

// 	fmt.Println(peaks)

// 	if len(peaks) < 2 {
// 		return 0
// 	}

// 	for i := 0; i < len(peaks) - 1; i++{
// 		startPeak := peaks[i]
// 		endPeak := peaks[i + 1]
// 		minHeight := min(startPeak[1], endPeak[1])
// 		fmt.Println(minHeight, startPeak, endPeak)

// 		for j := startPeak[0]; j < endPeak[0]; j++{
// 			unitHeight := height[j]
// 			storedVolume := minHeight - unitHeight
// 			fmt.Println("minHeight", minHeight)
// 			fmt.Println("unitHeight", unitHeight)
// 			fmt.Println("storedVolume", storedVolume)
// 			if storedVolume > 0{
// 				volume += storedVolume
// 			}
// 			fmt.Println("volume", volume)
// 		}
// 	}

// 	fmt.Println(volume)

// 	return volume
// }

// func trap(height []int) int {

// 	// for normal bucket:
// 	volume := 0
// 	for i := 0; i < len(height)-1; {
// 		startingHeight := height[i]
// 		maxBucketVolume := 0
// 		closedBowl := true
// 		stack := make([]map[string]int, 0, len(height))
// 		for j := i + 1; j < len(height); j++ {

// 			if i > 0 && height[i] > height[i-1] {
// 				if len(stack) > 0 {
// 					for len(stack) > 0 && stack[len(stack)-1]["height"] < height[i] {
// 						stack = stack[:len(stack)-1]
// 					}
// 				}
// 				stack = append(stack, map[string]int{
// 					"index":  i,
// 					"height": height[i],
// 				})
// 			} else if i < len(height) - 1 && height[i] > height[i + 1] && height[i] == stack[len(stack) - 1]["height"]{
// 				// case accounts for plateau
// 				stack = append(stack, map[string]int{
// 					"index":  i,
// 					"height": height[i],
// 				})
// 			}

// 			if height[j] >= startingHeight {
// 				i = j
// 				stack = stack[:0]
// 				break
// 			} else if j == len(height)-1 {
// 				i = j
// 				closedBowl = false
// 				break
// 			} else {
// 				unitVolume := startingHeight - height[j]
// 				maxBucketVolume += unitVolume
// 			}
// 		}

// 		if closedBowl {
// 			volume += maxBucketVolume
// 		}
// 	}

// 	return volume
// }

// stick with two pointer approach, but use stack to keep track of relative maximums
// whenever there is a peak within a bowl, check it against the stack
// if it is higher than the previous peak, pop the stack
// continue until it reaches a peak that it is less than or equal to
// push it onto the stack
// If we reach the end of the array without closing the bowl,
// Iterate from the last peak to the preceding peak, calculating volumes,
// then pop the stack.
// Do this until we get back to i.

// we need to account for plateaus. I.e. we need to check and see if the current value
// matches the height of the previous peak but only descends to the left.
// We don't need to check subsequent values to create a peak - we can create one whenever something is ascending.

// function handles case where end of array does not close bucket
// we start at starting point, then iterate forward until we find a descending

// we don't need this - we can just try to find the relative maximum peak.
// if we reach the end of the array, we use the relative maximum peaks


// can do this in O(n) using two pointers
// we will calculate the size of individual "buckets"
// iterate second pointer forward until we reach a point where
// the height is greater than or equal to the starting point.
// As we iterate, keep track of two values - one for handling a case
// where we are caculating the volume of a normal bucket, one for
// a case where we reach the end of the array without finding a
// height greater than or equal to our starting height.
// Use these two values to calculate a "max volume" or a "min volume"
// for the bucket - if we close the bucket, add the "max volume" to the
// total volume. If we reach the end of the array and don't close the
// bucket, add the "min volume" to the total volume.
// Once we have found a closing point, set i = j and jump to the next
// starting point.

// case one - normal bucket
// since a bucket's height cannot exceed it's starting height,
// the max volume it can have can be calculated by subtracting
// the current height from the starting height while iterating.
// For example, if the starting height is 3, and the current height is 1,
// that would store 2 units.

// case two - bucket without a closing side
// in this case, we need to actually calculate the volume
// of individual units.

// bug in current approach
// say we start with a height of 7, we drop to 0, go to 3, then up back to
// 5.
// We need to keep track of the peaks and valleys within incomplete buckets.
// But, we only need to do this for handling unclosed buckets.

// So, perhaps we could write our normal formula, then write a special function
// handling our edge case.
// Solving for the normal bucket is easy, so let's just do that first.

// Breaking down end of array problem

// If we iterate through

// we can keep track of when a bowl starts to descend.

// what if we create new bowls each time? Using a stack?
// Instead of calculating vertical units, we calculate horizontal units
// we move forward - as long as we're descending, we push a height and an index
// onto the stack.
// Once we reach a height that is ascending, we check it against the end of the stack
// If the end of the stack height is less than the current height, we subtract
// the end of stack height from the current height, subtract the current index from
// the end of stack index, and calculate the volume
// We then pop the stack and repeat this process until the end of the stack is greater
// than the current value.
// If the two are equal, we still pop the end of the stack.
// Then we push the current value onto the end of the stack

// keep track of the starting height - push onto stack
// if there are descending values, push onto stack.
// if values are increasing, check them against start height.
// if they are greater than start height, use start height.
// if they are less than start height, use current value.
// subtract end of stack height from selected value.
// subtract end of stack index from current index.
// multiply the two values.
// pop the stack.
// Repeat until current value is descending, or stack is empty.
// Push onto the stack.

// func trap(height []int) int {

// 	volume := 0
// 	stack := make([]map[string]int, 0, len(height))

// 	for i := 0; i < len(height); i++ {
// 		heightMap := map[string]int{
// 			"index": i,
// 			"height": height[i],
// 		}

// 		if len(stack) == 0 {
// 			stack = append(stack, heightMap)
// 			continue
// 		}

// 		end := stack[len(stack) - 1]
// 		startingHeight := stack[0]["height"]

// 		if end["height"] >= height[i] {
// 			stack = append(stack, heightMap)
// 			continue
// 		}

// 		maxHeight := 0

// 		if height[i] > startingHeight {
// 			maxHeight = startingHeight
// 		} else {
// 			maxHeight = height[i]
// 		}

// 		for len(stack) > 0 && height[i] > stack[len(stack) - 1]["height"]{
// 			width := i - stack[len(stack) - 1]["index"]
// 			unitHeight := maxHeight - stack[len(stack) - 1]["height"]
// 			unitVolume := width * unitHeight
// 			fmt.Println("stack before pop", stack)
// 			fmt.Println("index", i)
// 			fmt.Println("current height", height[i])
// 			fmt.Println("unitVolume", unitVolume)
// 			volume += unitVolume
// 			fmt.Println("volume", volume)
// 			stack = stack[:len(stack) - 1]
// 			fmt.Println("stack after pop", stack)
// 		}

// 		stack = append(stack, heightMap)
// 	}

// 	fmt.Println(volume)

// 	return volume
// }

// ^^ this is wrong
