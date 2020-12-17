package Day13

import (
	"fmt"
	"strconv"
	"strings"
)

func solvePt1(inputLines []string) {

	startTime, _ := strconv.Atoi(inputLines[0])
	possibleIds := strings.Split(inputLines[1], ",")

	busIdArray := make([]int, 0)
	for _, busId := range possibleIds {
		busNum, err := strconv.Atoi(busId)
		if err == nil {
			busIdArray = append(busIdArray, busNum)
		}
	}
	fmt.Printf("%v\n%v\n", startTime, busIdArray)
	closestId := 0
	closestWaiting := startTime
	for _, busId := range busIdArray {
		thisWait := busId - (startTime % busId)
		if thisWait < closestWaiting {
			closestId = busId
			closestWaiting = thisWait
		}
	}
	fmt.Printf("%v*%v=%v\n", closestId, closestWaiting, closestId*closestWaiting)

}

func findFirstCyclePoint(cycle1 int64, c1start int64, cycle2 int64, c2start int64, delay int64) (int64, int64) {
	// firstDay % c1 = firstDay % c2 + 1
	// t = firstDay%c1
	// t = firstDay%c2 +1

	firstDay := c1start

	for (firstDay-c1start)%cycle1 != 0 || ((firstDay+delay)%cycle2) != 0 {
		firstDay += cycle1
		//fmt.Printf("%v\n",firstDay)
	}
	var cycleLength int64 = cycle1 * cycle2

	return cycleLength, firstDay
}

func findPt2TimeStamp(idList string) int64 {

	//"17,x,13,19"

	idListSplit := strings.Split(idList, ",")
	index := 0
	var cycle1 int64
	var cycle2 int64
	var delay int64

	cycle1 = 0
	cycle2 = 0
	delay = 0
	//panic("fuu")
	for cycle1 == 0 {
		maybe_cycle, err := strconv.Atoi(idListSplit[index])
		if err == nil {
			cycle1 = int64(maybe_cycle)
		}
		index++
		delay++
	}

	for cycle2 == 0 {
		maybe_cycle, err := strconv.Atoi(idListSplit[index])
		if err == nil {
			cycle2 = int64(maybe_cycle)
		} else {
			index++
			delay++
		}
	}

	cycleLength, startPoint := findFirstCyclePoint(cycle1, cycle1, cycle2, cycle2, delay)
	index++
	delay++
	for index<len(idListSplit)  {
		maybe_cycle, err := strconv.Atoi(idListSplit[index])
		if err == nil {
			cycle2 = int64(maybe_cycle)
			cycleLength, startPoint = findFirstCyclePoint(cycleLength, startPoint, cycle2, cycle2, delay)
		}
		delay++
		index++
	}

	return int64(startPoint)
}

func findPt2TimeStampProofOfConcept(idList string) int64 {

	//67,7,59,61
	cycleLength, startPoint := findFirstCyclePoint(67, 67, 7, 7, 1)
	fmt.Printf("%v,%v\n", cycleLength, startPoint)

	cycleLength, startPoint = findFirstCyclePoint(cycleLength, startPoint, 59, 59, 2)
	fmt.Printf("%v,%v\n", cycleLength, startPoint)

	cycleLength, startPoint = findFirstCyclePoint(cycleLength, startPoint, 61, 61, 3)
	fmt.Printf("%v,%v\n", cycleLength, startPoint)
	fmt.Printf("%v\n", (startPoint)%67)
	fmt.Printf("%v\n", (startPoint+1)%7)
	fmt.Printf("%v\n", (startPoint+2)%59)
	fmt.Printf("%v\n", (startPoint+3)%61)

	//"17,x,13,19"

	fmt.Printf("%v\n", "17,x,13,19")
	cycleLength, startPoint = findFirstCyclePoint(17, 17, 13, 13, 2)
	fmt.Printf("%v,%v\n", cycleLength, startPoint)

	cycleLength, startPoint = findFirstCyclePoint(cycleLength, startPoint, 19, 19, 3)
	fmt.Printf("%v,%v\n", cycleLength, startPoint)

	fmt.Printf("%v\n", startPoint%17)

	fmt.Printf("%v\n", (startPoint+2)%13)
	fmt.Printf("%v\n", (startPoint+3)%19)

	return int64(startPoint)
}

func solvePt2(inputLines []string) {
	fmt.Printf("%v\n", findPt2TimeStamp(inputLines[1]))
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
