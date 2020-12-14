package Day10

import (
	"advent-of-code/Common"
	"fmt"
	"sort"
)

func solvePt2(inputLines []string){
	joltArray:=Common.ConvertInputToIntArray(inputLines)
	//joltArray=[]int{16,10,15,5,1,11,7,19,6,12,4}
	joltArray=append(joltArray,0)
	sort.Ints(joltArray)
	joltArray=append(joltArray,joltArray[len(joltArray)-1]+3)
	trackbacker:=make([]int,len(joltArray))
	trackbacker[0]=1
	for x:=1; x<len(joltArray);  x++{
		thisSum:=0
		for y:=1;y<4;y++{
			if x-y<0{
				break
			}
			if joltArray[x]-joltArray[x-y]<=3{
				thisSum+=trackbacker[x-y]
			}else{
				break
			}
		}
		trackbacker[x]=thisSum
	}
fmt.Printf("There were %v possibilities\n",trackbacker[len(trackbacker)-1])
}

func solvePt1(inputLines []string) {

	joltArray := Common.ConvertInputToIntArray(inputLines)
	// add entry for wall
	joltArray=append(joltArray,0)
	sort.Ints(joltArray)
	// add entry for wall charger
	joltArray=append(joltArray,joltArray[len(joltArray)-1]+3)

	numOnes := 0
	// the last one
	numThrees := 0
	fmt.Printf("%v\n",joltArray)
	for x := 1; x < len(joltArray); x++ {
		if joltArray[x]-joltArray[x-1] == 1 {
			numOnes++
		} else if joltArray[x]-joltArray[x-1] == 3 {
			numThrees++
		} else {
			panic("incorrect joltage")
		}
	}
	fmt.Printf("There were %v 1's, and %v 3's\n", numOnes, numThrees)
	fmt.Printf("The product is %v\n",numOnes*numThrees)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
