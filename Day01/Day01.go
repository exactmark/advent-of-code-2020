package Day01

import (
	"fmt"
	"sort"
	"strconv"
)

func FindMatchingPairSafe(target int, lowList []int, highList []int) (bool, int, int) {
	lowListCopy:=make([]int,len(lowList))
	highListCopy:=make([]int,0)
	copy(lowListCopy,lowList)
	if (&lowList)!=(&highList){
		highListCopy=make([]int,len(highList))
		copy(highListCopy,highList)

	}else {
		highListCopy=lowListCopy
	}
	return findMatchingPair(target,lowListCopy,highListCopy)
}

func findMatchingPair(target int, lowList []int, highList []int) (bool, int, int) {
	sort.Ints(lowList)
	sort.Ints(highList)

	lowPtr := 0
	highPtr := len(highList) - 1
	found := false
	for !found {
		this_sum := lowList[lowPtr] + highList[highPtr]
		if this_sum == target {
			found = true
		} else {
			if this_sum > target {
				highPtr -= 1
			} else {
				lowPtr += 1
			}
		}
		if lowPtr == len(lowList) || highPtr == 0 {
			break
		}
	}

	if found {
		return found, lowList[lowPtr], highList[highPtr]
	} else
	{
		return found, 0, 0
	}
}

func solvePt1(inputLines []string) {
	var numList = make([]int, 0)
	for _, singleLine := range inputLines {
		thisNumber, _ := strconv.Atoi(singleLine)
		numList = append(numList, thisNumber)
	}
	_, firstVal, secondVal := findMatchingPair(2020, numList, numList)
	fmt.Printf("%v,%v\n", firstVal, secondVal)
	fmt.Printf("Product is %v\n", firstVal*secondVal)

}

func solvePt2(inputLines []string) {
	var singlesList = make([]int, 0)
	for _, singleLine := range inputLines {
		thisNumber, _ := strconv.Atoi(singleLine)
		singlesList = append(singlesList, thisNumber)
	}
	sort.Ints(singlesList)
	var doublesList = make([]int, 0)
	for x := 0; x < len(singlesList); x++ {
		for y := 0; y < len(singlesList); y++ {
			doublesSum := singlesList[x] + singlesList[y]
			if doublesSum < 2020 {
				doublesList = append(doublesList, doublesSum)
			}else{
				break
			}
		}
	}

	_, firstVal, doublesVal := findMatchingPair(2020, singlesList, doublesList)
	fmt.Printf("%v,%v\n", firstVal, doublesVal)
	_,secondVal,thirdVal:=findMatchingPair(doublesVal,singlesList,singlesList)
	fmt.Printf("%v,%v\n", secondVal,thirdVal)
	fmt.Printf("Product is %v\n", firstVal*secondVal*thirdVal)

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
