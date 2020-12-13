package Day09

import (
	"advent-of-code/Day01"
	"fmt"
	"sort"
	"strconv"
)

func convertInputToIntArray(inputLines []string)[]int{
	returnArray:=make([]int,len(inputLines))
	for i,singleLine := range inputLines{
		returnArray[i],_=strconv.Atoi(singleLine)
	}
	return returnArray
}

func solvePt1(inputLines []string){
	intArray := convertInputToIntArray(inputLines)
	preambleSize :=25
	for x:=preambleSize;x<len(intArray);x++{
		found,_,_:=Day01.FindMatchingPairSafe(intArray[x],intArray[x-preambleSize:x],intArray[x-preambleSize:x])
		if !found{
			fmt.Printf("Could not find %v\n",intArray[x])
			break
		}
	}
}


func solvePt2(inputLines []string){
	intArray := convertInputToIntArray(inputLines)
	preambleSize :=25
	targetVal:=0
	for x:=preambleSize;x<len(intArray);x++{
		found,_,_:=Day01.FindMatchingPairSafe(intArray[x],intArray[x-preambleSize:x],intArray[x-preambleSize:x])
		if !found{
			targetVal=intArray[x]
			break
		}
	}
	fmt.Printf("Could not find %v\n",targetVal)
	sliceSize:=0
	sliceStart:=0
	for x:=0;x<len(intArray);x++{
		sliceSum:=0
		sliceSize=0
		for sliceSum<targetVal{
						sliceSum+=intArray[x+sliceSize]
			sliceSize++
		}
		if sliceSum==targetVal{
						sliceStart=x
			break
		}
	}
	subSlice := make([]int,sliceSize)
	copy(subSlice,intArray[sliceStart:sliceStart+sliceSize])
	sort.Ints(subSlice)
	fmt.Printf("%v + %v = %v\n",subSlice[0],subSlice[len(subSlice)-1],subSlice[0]+subSlice[len(subSlice)-1])



}


func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
