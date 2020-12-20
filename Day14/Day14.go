package Day14

import (
	"fmt"
	"strconv"
	"strings"
)

type memState struct {
	currentMask string
	memory      map[int]int64
}

func getMemAddressFromInputLine(singleLine string) int {
	memAddress := strings.Split(strings.Split(singleLine, "[")[1], "]")[0]
	memInt, _ := strconv.Atoi(memAddress)
	return memInt
}

func padBinString(binString string, desiredLength int) string {
	runeRep := make([]rune, desiredLength)
	for x := 1; x <= 36; x++ {
		if x <= len(binString) {
			runeRep[desiredLength-x] = rune(binString[len(binString)-x])
		} else {
			runeRep[desiredLength-x] = '0'
		}
	}
	return string(runeRep)
}

func (self *memState) applyMask(number int64) int64 {
	var binRep string
	binRep = strconv.FormatInt(number, 2)
	padBinRep := padBinString(binRep, 36)
	runeRep := []rune(padBinRep)

	for x := 1; x <= 36; x++ {
		if self.currentMask[len(self.currentMask)-x] == '0' {
			runeRep[len(padBinRep)-x] = '0'
		} else if self.currentMask[len(self.currentMask)-x] == '1' {
			runeRep[len(padBinRep)-x] = '1'
		}
	}
	//fmt.Printf("%v\n", string(runeRep))
	newVal, err := strconv.ParseInt(string(runeRep), 2, 64)
	if err != nil {
		fmt.Printf("%v\n", self.currentMask)
		fmt.Printf("%v\n", number)
		fmt.Printf("%v\n", binRep)
		fmt.Printf("%v\n", padBinRep)
		fmt.Printf("%v\n", string(runeRep))
		panic("incorrect conversion")
	}
	if newVal == 0 {
		fmt.Printf("%v gave 0\n", number)
	}
	return newVal

}

func (self *memState) applyAssignment(singleLine string) {
	memAddress := getMemAddressFromInputLine(singleLine)
	var newVal int64
	inputVal, err := strconv.Atoi(strings.Split(singleLine, " = ")[1])
	if err != nil {
		panic("Bad string conversion")
	}
	newVal = self.applyMask(int64(inputVal))
	self.memory[memAddress] = newVal
}

func (self *memState) sumAllMemory() int64 {
	var sum int64
	for _, val := range self.memory {
		sum += val
	}
	return sum
}

func solvePt1(inputLines []string) {

	memory := memState{memory: make(map[int]int64)}
	for _, singleLine := range inputLines {
		if singleLine[:4] == "mask" {
			memory.currentMask = singleLine[7:]
		} else {
			memory.applyAssignment(singleLine)
		}
	}
	total := memory.sumAllMemory()
	fmt.Printf("memTotal is %v\n", total)

}

func solvePt2(inputLines []string) {

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
