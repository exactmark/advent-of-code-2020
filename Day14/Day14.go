package Day14

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type memState struct {
	currentMask string
	maskList    []string
	memory      map[int]int64
}

func getMemAddressFromInputLine(singleLine string) int {
	memAddress := strings.Split(strings.Split(singleLine, "[")[1], "]")[0]
	memInt, _ := strconv.Atoi(memAddress)
	return memInt
}

func padBinString(binString string, desiredLength int) string {
	runeRep := make([]rune, desiredLength)
	for x := 1; x <= desiredLength; x++ {
		if x <= len(binString) {
			runeRep[desiredLength-x] = rune(binString[len(binString)-x])
		} else {
			runeRep[desiredLength-x] = '0'
		}
	}
	return string(runeRep)
}

func getBinStringOfNumber(number int64, desiredLength int) string {
	var binRep string
	binRep = strconv.FormatInt(number, 2)
	padBinRep := padBinString(binRep, desiredLength)
	return padBinRep
}

func applyGenericMask(padBinRep string, mask string) int64 {
	runeRep := []rune(padBinRep)

	for x := 1; x <= len(mask); x++ {
		if mask[len(mask)-x] == '0' {
			runeRep[len(padBinRep)-x] = '0'
		} else if mask[len(mask)-x] == '1' {
			runeRep[len(padBinRep)-x] = '1'
		}
	}
	fmt.Printf("%v\n",string(runeRep))

	newVal, err := strconv.ParseInt(string(runeRep), 2, 64)
	if err != nil {
		fmt.Printf("%v\n", mask)
		//fmt.Printf("%v\n", number)
		//fmt.Printf("%v\n", binRep)
		fmt.Printf("%v\n", padBinRep)
		fmt.Printf("%v\n", string(runeRep))
		panic("incorrect conversion")
	}
	if newVal == 0 {
		fmt.Printf("%v gave 0\n", padBinRep)
	}
	return newVal

}

func (self *memState) applyMask(number int64) int64 {
	padBinRep := getBinStringOfNumber(number, 36)

	newVal := applyGenericMask(padBinRep, self.currentMask)

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

func (self *memState) applyPt2Assignment(singleLine string) {
	baseMemAddress := getMemAddressFromInputLine(singleLine)
	baseMemAddressPadBin := getBinStringOfNumber(int64(baseMemAddress), 36)
	//var newVal int64
	inputVal, err := strconv.Atoi(strings.Split(singleLine, " = ")[1])
	if err != nil {
		panic("Bad string conversion")
	}
	for _, subMask := range self.maskList {
		fmt.Printf("%v\n%v\n",baseMemAddressPadBin, subMask)
		memAddress := applyGenericMask(baseMemAddressPadBin, subMask)
		fmt.Printf("%v\n", memAddress)
		self.memory[int(memAddress)] = int64(inputVal)
	}
}

func (self *memState) sumAllMemory() int64 {
	var sum int64
	for _, val := range self.memory {
		sum += val
	}
	return sum
}

func (self *memState) setCurrentMask(mask string) {
	self.currentMask = mask
	self.makeMaskList()
}

func (self *memState) makeMaskList() {
	self.maskList = make([]string, 0)
	baseMask := []rune(self.currentMask)
	xPtrs := make([]int, 0)
	// count number of x's
	for i, singleRune := range self.currentMask {
		if singleRune == 'X' {
			xPtrs = append(xPtrs, i)
		}
	}
	xCount := len(xPtrs)
	// make for loop of 2^x numbers
	for x := 0; x < int(math.Pow(2, float64(xCount))); x++ {
		// for each x, convert to binary string.
		xBinString := getBinStringOfNumber(int64(x), xCount)

		newMask := make([]rune, len(baseMask))
		copy(newMask, baseMask)
		fmt.Printf("%v\n",xBinString)
		fmt.Printf("%v\n",string(newMask))

		for i, xRune := range xBinString {
			newMask[xPtrs[i]] = xRune
		}
		fmt.Printf("%v\n",string(newMask))

		// convert to pt1 format
		pt1NewMask := make([]rune, len(baseMask))
		for i, baseRune := range newMask {
			if baseRune == '1' {
				pt1NewMask[i] = '1'
			} else {
				pt1NewMask[i] = 'X'
			}
		}
		fmt.Printf("%v\n",string(pt1NewMask))

		self.maskList = append(self.maskList, string(pt1NewMask))
	}
}

func solvePt1(inputLines []string) {

	memory := memState{memory: make(map[int]int64)}
	for _, singleLine := range inputLines {
		if singleLine[:4] == "mask" {
			memory.setCurrentMask(singleLine[7:])
		} else {
			memory.applyAssignment(singleLine)
		}
	}
	total := memory.sumAllMemory()
	fmt.Printf("memTotal is %v\n", total)

}

func solvePt2(inputLines []string) {
	memory := memState{memory: make(map[int]int64)}
	for _, singleLine := range inputLines {
		if singleLine[:4] == "mask" {
			memory.setCurrentMask(singleLine[7:])
		} else {
			memory.applyPt2Assignment(singleLine)
		}
	}
	total := memory.sumAllMemory()
	fmt.Printf("memTotal is %v\n", total)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
