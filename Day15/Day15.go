package Day15

import (
	"fmt"
	"strconv"
	"strings"
)

type MemGame struct {
	recall  []int
	hasSeen map[int]bool
	turn    int
}

func (self *MemGame) seedRecall(startingNumbers []int) {
	//turn:=0
	for turn, num := range startingNumbers {
		self.recall[turn] = num
		self.turn = turn + 1
	}
}

func (self *MemGame) seedRecallFromString(startingNumbersString string) {
	stringSplit := strings.Split(startingNumbersString, ",")
	startingNumbers := make([]int, len(stringSplit))
	for x, numString := range stringSplit {
		startingNumbers[x], _ = strconv.Atoi(numString)
	}
	self.seedRecall(startingNumbers)
}

func (self *MemGame) getLastUtterance(num int) int {
	if _, ok := self.hasSeen[num]; !ok {
		return 0
	}
	endPtr := self.turn - 2
	for endPtr >= 0 {
		if self.recall[endPtr] == num {
			return endPtr + 1
		}
		endPtr--
	}
	return 0
}

func (self *MemGame) processTurn() {

	lastNum := self.recall[self.turn-1]
	lastSpoke := self.getLastUtterance(lastNum)
	//fmt.Printf("%v\n",lastSpoke)
	var nextNum int
	if lastSpoke == 0 {
		nextNum = 0
	} else {
		nextNum = self.turn - lastSpoke
	}
	self.hasSeen[nextNum] = true
	self.recall[self.turn] = nextNum
	if self.turn%10000 == 0 {
		fmt.Printf("%v\n", self.turn)
	}
	self.turn++
}

//func solvePt1(inputLine string) {
//	memState := MemGame{
//		recall: make([]int, 0),
//		turn:   0,
//	}
//	memState.seedRecallFromString(inputLine)
//	for memState.turn < 2020 {
//		memState.processTurn()
//	}
//	fmt.Printf("%v\n", memState.recall[len(memState.recall)-1])
//}

func solvePt1(inputLine string, turnLimit int) {
	memState := MemGame{
		recall:  make([]int, turnLimit),
		hasSeen: make(map[int]bool),
		turn:    0,
	}
	memState.seedRecallFromString(inputLine)
	for memState.turn < turnLimit {
		memState.processTurn()
	}
	//fmt.Printf("%v\n", memState.recall)
	fmt.Printf("%v\n", memState.recall[memState.turn-1])
}

func Solve(inputLines []string) {
	//for _, singleLine := range inputLines {
	//	solvePt1(singleLine)
	//}
	for _, singleLine := range inputLines {
		solvePt1(singleLine, 2020)
		//solvePt1(singleLine, 30000000)
	}
}
