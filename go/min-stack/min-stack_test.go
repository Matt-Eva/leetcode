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
	var stackStruct StackStruct

	if val <= this.Min {
		currentMinStruct := this.MinStruct
		stackStruct = StackStruct{
			Value:   val,
			PrevMin: currentMinStruct,
		}
	} else {
		stackStruct = StackStruct{
			Value:   val,
			PrevMin: nil,
		}
	}

	if len(this.Stack) == 0 || val <= this.Min{
		this.Min = val
		this.MinStruct = &stackStruct
	}
	
	this.Stack = append(this.Stack, stackStruct)
}

func (this *MinStack) Pop() {
	stackLength := len(this.Stack)
	
	if stackLength == 1 {
		this.Min = 0
		this.MinStruct = nil
		this.Stack = this.Stack[:0]
		return
	}
	
	topStruct := &this.Stack[stackLength-1]

	if topStruct.Value == this.Min {
		previousMinStruct := topStruct.PrevMin

		this.Min = previousMinStruct.Value
		this.MinStruct = previousMinStruct
	}

	this.Stack = this.Stack[:stackLength-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1].Value
}

func (this *MinStack) GetMin() int {
	return this.Min
}

// Very cool solution that uses two lists of equal length below
// It just adds a new item to the list of minimums for every new 
// item added to the stack
// The new item added to the list of minimums will be guaranteed to be 
// the current minimum.

// type MinStack struct {
//     s []int
//     min []int
//     len int
// }


// func Constructor() MinStack {
//     return MinStack{}
// }


// func (this *MinStack) Push(val int)  {
//     this.s = append(this.s, val)
//     curMin := val
//     if this.len > 0 {
//         curMin = getMin(this.min[this.len-1], val)
//     }
//     this.min = append(this.min, curMin) 
//     this.len ++
// }


// func (this *MinStack) Pop()  {
//     this.s = this.s[0:this.len-1]
//     this.min = this.min[0:this.len-1]
//     this.len --
// }


// func (this *MinStack) Top() int {
//     return this.s[this.len-1]
// }


// func (this *MinStack) GetMin() int {
//     return this.min[this.len-1]
// }

// func getMin(a int, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }
