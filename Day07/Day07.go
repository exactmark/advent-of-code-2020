package Day07

import (
	"fmt"
	"strconv"
	"strings"
)

type bag struct {
	color       string
	contains    map[string]int
	containedBy []string
	canHoldGold bool
}

func makeSingleBag(descriptor string) bag {
	descriptorSplit := strings.Split(descriptor, " bags contain ")
	bagColor:=descriptorSplit[0]
	containsString := descriptorSplit[1]
	containsString = containsString[:len(containsString)-1]
	contains := make(map[string]int)

	if containsString != "no other bags" {
		for _,singleContained := range strings.Split(containsString,", "){
			spaceSplit:=strings.Split(singleContained," ")
			containedNumber,_:=strconv.Atoi( spaceSplit[0])
			containedColor:=strings.Join(spaceSplit[1:len(spaceSplit)-1]," ")
			contains[containedColor]=containedNumber
		}
	}

	return bag{
		color:       bagColor,
		contains:    contains,
		containedBy: make([]string,0),
		canHoldGold: false,
	}
}

func processInputToBagMap(inputLines []string) *map[string]bag {
	returnBagMap := make(map[string]bag)

	for _, singleLine := range inputLines {
		thisBag := makeSingleBag(singleLine)
		fmt.Printf("%v\n", thisBag)
	}

	return &returnBagMap
}

func populateTrackBack(bagMapPtr *map[string]bag){
	bagMap:=*bagMapPtr
}

func solvePt1(inputLines []string) {
	bagMap := *(processInputToBagMap(inputLines))
	populateTrackBack(&bagMap)
	fmt.Printf("%v\n", bagMap)
}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
