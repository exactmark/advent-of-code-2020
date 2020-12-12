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
}

type canHoldMap struct {
	targetColor  string
	canHoldIndex map[string]canHoldLink
}

type canHoldLink struct {
	canHoldColor bool
	evaluated    bool
}

func makeSingleBag(descriptor string) bag {
	descriptorSplit := strings.Split(descriptor, " bags contain ")
	bagColor := descriptorSplit[0]
	containsString := descriptorSplit[1]
	containsString = containsString[:len(containsString)-1]
	contains := make(map[string]int)

	if containsString != "no other bags" {
		for _, singleContained := range strings.Split(containsString, ", ") {
			spaceSplit := strings.Split(singleContained, " ")
			containedNumber, _ := strconv.Atoi(spaceSplit[0])
			containedColor := strings.Join(spaceSplit[1:len(spaceSplit)-1], " ")
			contains[containedColor] = containedNumber
		}
	}

	return bag{
		color:       bagColor,
		contains:    contains,
		containedBy: make([]string, 0),
	}
}

func processInputToBagMap(inputLines []string) *map[string]bag {
	returnBagMap := make(map[string]bag)

	for _, singleLine := range inputLines {
		thisBag := makeSingleBag(singleLine)
		returnBagMap[thisBag.color]=thisBag
	}

	return &returnBagMap
}

func populateTrackBack(bagMapPtr *map[string]bag) {
	bagMapActual := *bagMapPtr
	for _,containingBag := range bagMapActual{
		for containedBag,_ :=range containingBag.contains{
			fmt.Printf("%v\n",containedBag)
			bagMapActual[containedBag].containedBy=append(bagMapActual[containedBag].containedBy,containingBag.color)
		}
	}
	fmt.Printf("%v\n", bagMapActual)
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
