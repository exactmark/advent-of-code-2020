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

type canHoldStruct struct {
	targetColor string
	canHoldMap  map[string]*canHoldLink
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

func makeBagMap(inputLines []string) *map[string]*bag {
	bagMap := *(processInputToBagMap(inputLines))
	populateTrackBack(&bagMap)
	return &bagMap
}

func processInputToBagMap(inputLines []string) *map[string]*bag {
	returnBagMap := make(map[string]*bag)

	for _, singleLine := range inputLines {
		thisBag := makeSingleBag(singleLine)
		returnBagMap[thisBag.color]=&thisBag
	}

	return &returnBagMap
}

func populateTrackBack(bagMapPtr *map[string]*bag) {
	bagMapActual := *bagMapPtr
	for _,containingBag := range bagMapActual{
		for containedBag,_ :=range containingBag.contains{
			fmt.Printf("%v\n",containedBag)
			bagMapActual[containedBag].containedBy=append(bagMapActual[containedBag].containedBy,containingBag.color)
		}
	}
	fmt.Printf("%v\n", bagMapActual)
}

func determineHowManyCanHold(target string,bagMapPtr *map[string]*bag) {
	workingHoldStruct:= canHoldStruct{
		targetColor: target,
		canHoldMap:  make(map[string]*canHoldLink),
	}

	for _,singleBag := range *bagMapPtr{
	//	for each bag call populateCanHold
	}

//	go through workingHoldStruct, canHoldMap. Count each .canContain = true

}

func populateCanHold()

// populateCanHold will take ptr to canHoldStruct,bagMapPtr,colorToEvaluate{
// check if colorToEvaluate is in canHoldStruct.
//   if yes,
//     check if it's been evaluated.
//         if evaluated yes,
//		      return canHoldValue
//         if not evaluated,
//            with each contained color
//                 check if target color
//                       set as true and set as evaluated.
//                 else
//                	call populateCanHold
//                  if any return true, set as true and set as evaluated.
//                  else set as false and set as evaluated
//	          return true or false up the stack.
//   if not in CanHold
//      make an entry
//            call populateCanHold with each contained color
//                  if any return true, set as true and set as evaluated.
//                  else set as false and set as evaluated
//	          return true or false up the stack.
//

func solvePt1(inputLines []string) {
	bagMap := *(makeBagMap(inputLines))
	determineHowManyCanHold("shiny gold",&bagMap)
}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
