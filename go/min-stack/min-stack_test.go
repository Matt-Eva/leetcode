package main

// In order to make getting the minimum constant time
// We will actually be storing the minimum separately
// We will also be storing the differences between subsequent minimums
// on the stack itself.
// Or, alternatively, we could store a pointer to the previous min
// Then also store the min separately, along with a pointer

type StackStruct struct {
	Value   int
	PrevMin *StackStruct
}

type MinStack struct {
	Min       int
	MinStruct *StackStruct
	Stack     []StackStruct
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if val < this.Min && len(this.Stack) == 0 {
		stackStruct := StackStruct{
			Value:   val,
			PrevMin: nil,
		}

		this.Min = val
		this.MinStruct = &stackStruct
		this.Stack = append(this.Stack, stackStruct)
	} else if val < this.Min {
		currentMinStruct := this.MinStruct
		stackStruct := StackStruct{
			Value:   val,
			PrevMin: currentMinStruct,
		}

		this.Min = val
		this.MinStruct = &stackStruct
		this.Stack = append(this.Stack, stackStruct)
	} else {
		stackStruct := StackStruct{
			Value:   val,
			PrevMin: nil,
		}

		this.Stack = append(this.Stack, stackStruct)
	}
}

func (this *MinStack) Pop() {

}

func (this *MinStack) Top() int {
	return 0
}

func (this *MinStack) GetMin() int {
	return 0
}
