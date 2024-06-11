package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	// so, we need to push characters onto the stack
	// operation characters are '+' '-' '*' and '/'
	// Whenever we encounter an operation character, we want to
	// use it to operate on the preceding two characters, which
	// will be numbers
	// We will want to pop the preceding two characters off the stack
	// then push the result onto the stack.

	stack := make([]int, 0, len(tokens))
	operatorMap := map[string]func(int, int) int {
		"+": func(a, b int) int {return a + b},
		"-": func(a, b int) int {return a - b},
		"*": func(a, b int) int {return a * b},
		"/": func(a, b int) int {return a / b},
	}

	for _, token := range tokens{
		if operatorMap[token] == nil {
			intToken, _ := strconv.Atoi(token)
			stack = append(stack, intToken)
			continue
		}

		length := len(stack)
		prevInts := stack[length - 2:]
		stack = stack[:length-2]
		newInt := operatorMap[token](prevInts[0], prevInts[1])
		stack = append(stack, newInt)
	}

   return stack[0] 
}