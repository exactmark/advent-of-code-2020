package Day15

import (
	"fmt"
	"strconv"
	"strings"
)

type MemGame struct {
	recall  map[int]int
	lastNum int
	nextNum int
	turn    int
}

func (self *MemGame) seedRecall(startingNumbers []int) {
	//turn:=0
	for turn, num := range startingNumbers {
		self.turn = turn + 1
		self.recall[num] = turn + 1
		self.lastNum = num
	}
	self.turn=self.turn+1
	self.recall[0]=self.turn
	self.lastNum = 0
}

func (self *MemGame) seedRecallFromString(startingNumbersString string) {
	stringSplit := strings.Split(startingNumbersString, ",")
	startingNumbers := make([]int, len(stringSplit))
	for x, numString := range stringSplit {
		startingNumbers[x], _ = strconv.Atoi(numString)
	}
	self.seedRecall(startingNumbers)
}

func (self *MemGame) processTurn() {
	self.turn = self.turn + 1
	v, ok := self.recall[self.lastNum]
	if !ok {
		last0Turn,_:=self.recall[0]
		self.nextNum=
	}
}

func solvePt1(inputLines []string) {
	memState := MemGame{
		recall:  make(map[int]int),
		lastNum: 0,
		turn:    0,
	}
	memState.seedRecallFromString(inputLines[0])
	memState.processTurn()
	fmt.Printf("%v\n", memState)
}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
