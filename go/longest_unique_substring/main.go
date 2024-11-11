package main

import (
	"fmt"
	"log"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func main() {

	cases := []string{
		"abcabcbb",
		"dddddd",
		"dabbad",
		"cajhdiccoin",
		"",
		"a",
		"tmmzuxt",
	}

	caseAnswers := []int{
		3,
		1,
		3,
		6,
		0,
		1,
		5,
	}

	for i, c := range cases {
		result := lengthOfLongestSubstring(c)
		answer := caseAnswers[i]
		if result != answer {
			log.Fatalf("result %d does not equal answer %d for case %s \n", result, answer, c)
		} else {
			fmt.Printf("success - result %d equals answer %d for case %s \n", result, answer, c)
		}
	}

}

func lengthOfLongestSubstring(s string) int {
	bytes := []byte(s)

	byteMap := make(map[byte]int)
	maxLength := 0
	leftSide := 0
	for i, b := range bytes {
		_, ok := byteMap[b]
		if !ok {
			byteMap[b] = i
			maxLength = max(maxLength, i-leftSide+1)
		} else {
			leftSide = max(byteMap[b]+1, leftSide)
			byteMap[b] = i
			maxLength = max(maxLength, i-leftSide+1)
		}
	}

	return maxLength
}

/// === Below solution works, but unecessarily deletes from hash map >>

// use two pointers
// iterate forward using second pointer
// store characters in a hash map
// keep track of substring length in a sublength variable
// as long as there are no duplicates, increment sub-length variable
// have a maxlength variable that stores the final answer
// if the sublength ever exceeds the max length, update the maxlength

// if a duplicate occurs, increment the first pointer to the character after the first duplicate.
// while the first pointer is incrementing forward, remove each character from the hashmap.
// set the new duplicate character into the hashmap.
// update substring length by subtracting current position from new starting position.

// func lengthOfLongestSubstring(s string) int {
//     bytes := []byte(s)

// 	byteMap := make(map[byte]bool)
// 	subLength := 0
// 	maxLength := 0
// 	i := 0
// 	for j := 0; j < len(bytes); j++{
// 		if !byteMap[bytes[j]] {
// 			byteMap[bytes[j]] = true
// 			subLength += 1
// 			maxLength = max(maxLength, subLength)
// 		} else {
// 			for i < len(bytes) {
// 				if bytes[i] == bytes[j]{
// 					i = i + 1
// 					subLength = j - i + 1
// 					break
// 				} else {
// 					delete(byteMap, bytes[i])
// 					i++
// 				}
// 			}
// 		}
// 	}

// 	return maxLength
// }
