package main

// currently exceeds memory
func LargestRectangleArea(heights []int) int {

	// we need to keep track of vertical rectangles (individual column)
	// we need to keep track of horizontal rectangles
	// horizontal rectangles are created for each level within a verticle rectangle
	// if a layer of a horizontal rectangles ends ( it drops off), we calculate the area and
	// save it in an array, then remove that rectangle from the list of current rectangles
	// we need to keep track of the start and end point of a horizontal rectangle
	// the area will be calculated by the difference between the start and end point and the height

	// we can keep track of the current horizontal rectangles in a map
	// the key in the map will represent the height of the rectangle
	// each time we reach a new space, we iterate through the map to check the current heights
	// if the height is less than or equal to the rectangle in the new space, do nothing
	// if the height is greater than the rectangle in the new space, we get the length of the rectangle and multiply it by the
	// current height

	// this actually is a stack problem
	// the end value on the stack will always be the greatest
	// we just need to record the starting point when we push the rectangle onto the stack
	// it doesn't need to be a map

	// rather than appending a new rectangle for every single layer in a rectangle,
	// we simply shift the starting position of the lesser rectangle that we push onto the stack to be the position of its greater preceding rectangle.

	zeroHeight := true

	for _, height := range heights {
		if height > 0 {
			zeroHeight = false
			break
		}
	}

	if zeroHeight {
		return 0
	}

	type StackRect struct {
		Height int
		Start  int
	}

	stack := make([]StackRect, 0, len(heights))
	maxArea := 0

	for index, height := range heights {
		start := index

		for len(stack) > 0 && stack[len(stack)-1].Height > height {
			length := index - stack[len(stack)-1].Start
			area := length * stack[len(stack)-1].Height

			if area > maxArea {
				maxArea = area
			}

			start = stack[len(stack)-1].Start
			stack = stack[:len(stack)-1]
		}

		if height > 0 {
			if len(stack) == 0 || stack[len(stack)-1].Height < height {
				newRect := StackRect{
					height,
					start,
				}
				stack = append(stack, newRect)
			}
		}
	}

	heightsLen := len(heights)
	for _, rectangle := range stack {
		length := heightsLen - rectangle.Start
		area := length * rectangle.Height

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
