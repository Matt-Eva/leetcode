package main

import (
	"slices"
)

// we need to find the slowest car that will reach the target first
// that will determine the first fleet
// We can then repeat this operation for subsequent fleets

// for each car, we check and see if there is another car ahead
// that is slower and will reach the end at the same time or later
// than the current car.
// this is a STACK problem
// we can push cars onto the stack
// whenever a car is pushed onto the stack that is slower and
// will reach the target later than the current car on the stack
// we pop the stack.
// we only push onto the stack if the current car is slower and / or arrives later
// than the top of the stack

// criteria to pop
// behind current car
// faster than current car
// will reach end at same time or sooner than current car

// what makes a fleet?

// if there is a car that is at the same position or forward in position
// that will reach the target after or at the same time as the current car
// the current car joins that fleet

// conditions:
// if currentPosition == fleetPosition {
//		join fleet
//		set fleet complete time to slowest car
//	}
// else if currentPosition < fleetPosition && currentComplete <= fleetComplete {
//		join fleet
//	}

// currentPosition <= fleetPosition
// 		if the currentCar is ahead of the fleetPosition, it will be unable to join that
//		fleet and will become its own fleet

// currentComplete <= fleetComplete
// 		if the currentCar is behind or at the same position as the fleet, but will reach the target in the same number of
// 		rounds or fewer than the fleet, it will catch up to the fleet by the time the fleet
//		reaches the target

//	else if currentPosition > fleetPosition && currentComplete >= fleetComplete {
// 		fleet joins car (essentially the same thing as car joining fleet)
// 		set fleetComplete to new current car
// 		set fleetPosition to new current car
// 	}

// else {
//	create new fleet
//}

// nlog(n) solution
func carFleet(target int, position []int, speed []int) int {
	type Car struct {
		Position int
		Complete float64
	}

	carSortedByPosition := make([]*Car, 0, len(position))
	
	for i, p := range position{
		complete := calculateComplete(target,p, speed[i])
		newCar := &Car{
			p, complete,
		}
		carSortedByPosition = append(carSortedByPosition, newCar)
	}

	slices.SortFunc(carSortedByPosition, func(a, b *Car) int {
		return b.Position - a.Position
	})

	fleetStack := make([]*Car, 0, len(position))
	fleetStack = append(fleetStack, carSortedByPosition[0])

	for i:= 1; i < len(carSortedByPosition); i++{
		car := carSortedByPosition[i]
		tail := fleetStack[len(fleetStack) - 1]
		if car.Complete <= tail.Complete {
			continue
		}
		fleetStack = append(fleetStack, car)
	}

	return len(fleetStack)
}

func calculateComplete(target, position, speed int) float64 {
	completeTime := float64((target - position)) / float64(speed)
	return completeTime
}

// n^2 solution
// func carFleet(target int, position []int, speed []int) int {
// 	type Fleet struct {
// 		Position int
// 		Complete float64
// 	}
// 	fleetSlice := make([]*Fleet, 0, len(position))

// 	for i, pos := range position {
// 		completeTime := calculateComplete(target, pos, speed[i])
// 		joinedFleet := false

// 		for _, fleet := range fleetSlice {
// 			if pos == fleet.Position {
// 				maxTime := max(completeTime, fleet.Complete)
// 				fleet.Complete = maxTime
// 				joinedFleet = true
// 			} else if pos < fleet.Position && completeTime <= fleet.Complete {
// 				joinedFleet = true
// 			} else if pos > fleet.Position && completeTime >= fleet.Complete {
// 				fleet.Position = pos
// 				fleet.Complete = completeTime
// 				joinedFleet = true
// 			}
// 		}

// 		if !joinedFleet {
// 			newFleet := &Fleet{pos, completeTime}
// 			fleetSlice = append(fleetSlice, newFleet)
// 		}
// 	}

// 	fleetMap := make(map[int]bool)

// 	for _, fleet := range fleetSlice{
// 		if fleetMap[fleet.Position] {
// 			continue
// 		}

// 		fleetMap[fleet.Position] = true
// 	}

// 	return len(fleetMap)
// }

// necessary because integer division rounds down in go.
// Ex: 3 goes into 4 once, so 4/3 == 1
// Whereas if a car was at position 6 and had a speed of 3
// it would only reach the target on its 2nd turn, not its first.

// O(n^2) solution - does not work
// func carFleet(target int, position []int, speed []int) int {
// 	type Car struct {
// 		Speed        int
// 		Position     int
// 		CompleteTime int
// 		// we can keep this at an int b/c if a faster car will hit
// 		// the same point as a slower car on the same turn
// 		// it will match the slower car.
// 	}

// 	stack := make([]Car, 0, len(position))

// 	for index, carSpeed := range speed {
// 		carPosition := position[index]
// 		multiple := (target - carPosition) / carSpeed
// 		completeTime := multiple

// 		if (multiple * carSpeed) != (target - carPosition){
// 			completeTime = multiple + 1
// 		}

// 		for len(stack) != 0 && stack[len(stack)-1].Position <= carPosition && stack[len(stack)-1].Speed > carSpeed && stack[len(stack)-1].CompleteTime <= completeTime {
// 			stack = stack[:len(stack) - 1]
// 			fmt.Println(stack)
// 		}

// 		newCar := Car{
// 			carSpeed, carPosition, completeTime,
// 		}

// 		if len(stack) == 0{
// 			stack = append(stack, newCar)
// 			continue
// 		}

// 		newFleet := true

// 		for _, car := range stack {
// 			if car.Position >= newCar.Position && car.Speed < newCar.Speed && car.CompleteTime >= newCar.CompleteTime{
// 				newFleet = false
// 			}
// 		}

// 		if newFleet{
// 			stack = append(stack, newCar)
// 		}

// 	}

// 	return len(stack)
// }
