package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	type CaseStruct struct {
		Case   string
		Result bool
	}

	cases := []CaseStruct{
		{
			"aba",
			true,
		},
		{
			"abb",
			false,
		},
		{
			"abba",
			true,
		},
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			" ",
			true,
		},
	}

	for _, testCase := range cases {
		result := isPalindrome(testCase.Case)
		if result != testCase.Result {
			t.Errorf("Result '%t' does not match test case result '%t'\n for case '%s'", result, testCase.Result, testCase.Case)
		}
	}
}
