package main


// // problem : https://leetcode.com/problems/generate-parentheses/description/

import (
	"strings"
)

func generateParenthesis(n int) []string{
	stack := make([]string, 0, 2*n)
	result := []string {}

	backtrack(0, 0)
	
	return result
}

func backtrack (stack, result *[]string, open, closed, n int) {
	if open == closed && open == n {
		joined := strings.Join((*stack), "")
		(*result) = append((*result), joined)
		return
	}

	// initially we add all possible openings to start.
	// then, as function calls resolve, we begin introducing closing
	// parentheses as interruptions.
	// This creates a new chain of cascading function calls, which repeat the process
	if open < n {
		(*stack) = append((*stack), "(")
		backtrack((*stack), result, open + 1, closed, n)
		// ^^ blocking function call
		// once resolved, moves on to next if statement
		(*stack) = (*stack)[:len((*stack)) - 1]
	}

	// as the function calls from the initial pass resolve
	// they move on to this if statement, but retain the value for "open"
	// that they had when called.
	if closed < open {
		(*stack) = append((*stack), ")")
		backtrack((*stack), result, open, closed + 1, n)
		// ^^ blocking function call
		(*stack) = (*stack)[:len((*stack)) - 1]
	}
}



// func generateParenthesis(n int) []string {
// 	// we can push a certain number of starting parentheses onto the stack
// 	// we can start by pushing all of them, then decreasing the number pushed
// 	// for each decrease, we then need to push all the various iterations of that
// 	// number itself - i.e. the various combinations following the starting combination.

// 	// for n starting parentheses pushed onto the stack, we will need
// 	// to push a factorial of n ending parentheses onto the stack
// 	// we need the total number of each pushed onto the stack.
// 	var result []string

// 	stack := make([]string, 0, n*2)
// 	startPushed := 0
// 	endPushed := 0

// 	// generates starting cases for starting parentheses pushed
// 	// startPushed ultimately needs to equal n
// 	// this is an n^2 solution :(
// 	for i := n; i > 0; i-- {
// 		appendStart(&stack, i, n, &startPushed, &endPushed)
// 		parenString := strings.Join(stack, "")
// 		result = append(result, parenString)
// 		stack = stack[:0]
// 		startPushed = 0
// 		endPushed = 0
// 	}

// 	// for number of starting parentheses pushed, need to be factorial of number of end parentheses pushed, starting with 1.

// 	// following number of end parentheses pushed, there will be a new set of starting parentheses pushed

// 	// these operations happen recursively until all end parentheses and start parentheses have been added.

// 	return result
// }

// func appendStart(stack *[]string, pushNum, n int, startPushed, endPushed *int){

// 	for i := 0; i < pushNum; i++ {
// 		(*stack) = append((*stack), "(")
// 		(*startPushed) += 1
// 	}
// 	appendEnd(stack, pushNum, n, startPushed, endPushed)
// }

// func appendEnd(stack *[]string, pushNum, n int, startPushed, endPushed *int) {
// 	for i := 0; i < pushNum; i++{

// 	}
// }

// // new approach
// // we take starting and ending parentheses and assign them an index value
// // we then go into the next layer and give them an index value
// // and proceed so on.

// // whenever pushed becomes equal to n, we know that we have exited the loop
// // we reset pushed to 0 and empty the stack.
// // or do we need to pop off the stack?
// // maybe we create a struct for each pair with the corresponding indexes?