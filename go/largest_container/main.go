package main

func main() {}

func maxArea(height []int) int {

	volume := 0

	for i := 0; i < len(height); i++ {
		startHeight := height[i]

		for j := len(height) - 1; j > i; j-- {
			endHeight := height[j]
			width := j - i

			if endHeight >= startHeight{
				newVolume := startHeight * width

				if newVolume > volume {
					volume = newVolume
				}

				break
			}

			newVolume := endHeight * width

			if newVolume > volume {
				volume = newVolume
			}
		}
	}

	return volume
}

// use two pointers
// set start at first pointer
// move the second pointer forward until it reaches a height greater than the height of the first pointer or the end of the container
// use the minimum height - either the start or the end - to set the height of the container
// calculate the width of the container by subtracting the first pointer index from the second pointer index
// multiply the width by the height to caluclate the volume
// save this volume in a variable
// move the first pointer forward by one, and repeat the process
// if a new volume is greater than the volume stored in the variable update the variable to contain the new volume
// ^^^^^^ Incomplete

// bug - we need to keep track of preceding volume containers until we reach a height that is greater than the start or the end of the array
// Why?
// In the case of reaching

// Question - can containers pass through the heights of other containers?
// Apparently yes.
// Is this just an n^2 problem?
// we can solve it as such first, then consider if there are optimizations to be made.
// n^2 works, but exceeds time limit.

// start the second pointer at the end of the array
// if the end of the array height is greater than or equal to the start height,
// then we know that's the max rectangle size for that starting height.
// compare against the volume and continue.
// if it's not greater than or equal to, iterate backwards until we reach a height that
// is greater than or equal to the starting height.
// Calculate volume for each step. If it exceeds the previous volume, reset the volume.
// Once we iterate backward to the point where we reach a height that is greater than
// or equal to the starting height, we can calculate that volume.
// If it's greater than the current volume, reset the volume and break.
// If not, break.
// ^^ This solution was accepted, but it's still slow - n^2 time complexity.

// Better solution (not mine)
// We iterate both pointers for an n^2 solution
// If one height is less than another, we move that pointer either to the right or left
// depending on where it started.
// We calculate the volume at each step and update the volume accordingly.