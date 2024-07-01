package main

func twoSum(numbers []int, target int) []int {
	// we have to pointers
	// one is at the start
	// the other is at the end

	// while the sum is less than the indexes at the two pointers, move the start forward
		// we know that if the current start plus the current end is greater than the target
		// the correct start index must be greater than the current start index
		// Why?
		// The current end is at its current position, because all indexes after it have resulted
		// in something greater than the target.
		// This held true for all indexes preceding the current starting index
		// We know that the values preceding the current starting index are less than
		// the current starting index.
		// Therefore, we cannot move the end forward, because we are guaranteed that the sum
		// will be greater than the target.

	// while the sum is greater than the indexes at the two pointers, move the end backward


	length := len(numbers)
	start := 0
	end := length - 1

	for {
		if numbers[start]+numbers[end] == target {
			return []int{start + 1, end + 1}
		} else if numbers[start]+numbers[end] < target {
			start += 1
		} else if numbers[start]+numbers[end] > target {
			end -= 1
		}
	}

}

