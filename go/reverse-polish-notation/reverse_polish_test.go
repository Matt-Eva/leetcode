package main

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T){
	type CaseStruct struct {
		Case []string
		Result int
	}
	cases := []CaseStruct{
		{
			[]string{"1", "2", "+"},
			3,
		},
		{
			[]string{"4", "3", "/"},
			1,
		},
		{
			[]string{"2", "3", "*", "5", "/", "4", "-"},
			-3,
		},
	}

	for _, val := range cases {
		if val.Result != evalRPN(val.Case){
			t.Fatal("Results do not match")
			fmt.Printf(": %d != %d", val.Result, evalRPN(val.Case))
		}
	}

}