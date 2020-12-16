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
		if err==nil {
			busIdArray = append(busIdArray,busNum )
		}
	}
	fmt.Printf("%v\n%v\n",startTime,busIdArray)
	closestId:=0
	closestWaiting:=startTime
	for _,busId := range busIdArray{
		thisWait :=busId-( startTime%busId)
		if thisWait<closestWaiting{
			closestId = busId
			closestWaiting=thisWait
		}
	}
	fmt.Printf("%v*%v=%v\n",closestId,closestWaiting,closestId*closestWaiting)

}

func findPt2TimeStamp(idList string) int64 {


	return 0
}

func solvePt2(inputLines []string) {
	fmt.Printf("%v\n",findPt2TimeStamp(inputLines[1]))
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
