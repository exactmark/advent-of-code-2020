package Day07

import (
	"fmt"
	"strconv"
	"strings"
)

type bag struct {
	color         string
	contains      map[string]int
	containedBy   []string
	bagsContained int
}

type canHoldStruct struct {
	colorToFind string
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
		color:         bagColor,
		contains:      contains,
		containedBy:   make([]string, 0),
		bagsContained: -1,
	}
}

func makeBagMap(inputLines []string) *map[string]*bag {
	bagMap := *(processInputToBagMap(inputLines))
	populateTrackBack(&bagMap)
	populateBagsContained(&bagMap)
	return &bagMap
}

func processInputToBagMap(inputLines []string) *map[string]*bag {
	returnBagMap := make(map[string]*bag)

	for _, singleLine := range inputLines {
		thisBag := makeSingleBag(singleLine)
		returnBagMap[thisBag.color] = &thisBag
	}

	return &returnBagMap
}

func populateTrackBack(bagMapPtr *map[string]*bag) {
	bagMapActual := *bagMapPtr
	for _, containingBag := range bagMapActual {
		for containedBag, _ := range containingBag.contains {
			//fmt.Printf("%v\n", containedBag)
			bagMapActual[containedBag].containedBy = append(bagMapActual[containedBag].containedBy, containingBag.color)
		}
	}
	//fmt.Printf("%v\n", bagMapActual)
}

func populateBagsContained(bagMapPtr *map[string]*bag) {
	bagMapActual := *bagMapPtr
	for _, singleBag := range bagMapActual {
		populateBagsContainedLogic(bagMapPtr, singleBag.color)
	}
}

//we're going to pretend there are no circular loops
func populateBagsContainedLogic(bagMapPtr *map[string]*bag, colorToCheck string) int {
	bagMapActual := *bagMapPtr
	bagToCheck, _ := bagMapActual[colorToCheck]
	if bagToCheck.bagsContained>=0 {
		return bagToCheck.bagsContained
	}else{
		containedDirectly:=0
		secondDegreeContained:=0
		for containedColor,numContainedBag:=range bagToCheck.contains{
			containedDirectly+=numContainedBag
			secondDegreeContained+=numContainedBag*populateBagsContainedLogic(bagMapPtr,containedColor)
		}
		bagToCheck.bagsContained=containedDirectly+secondDegreeContained
		return bagToCheck.bagsContained
	}

}

func determineHowManyCanHold(target string, bagMapPtr *map[string]*bag) {
	workingHoldStruct := canHoldStruct{
		colorToFind: target,
		canHoldMap:  make(map[string]*canHoldLink),
	}

	for _, singleBag := range *bagMapPtr {
		populateCanHold(&workingHoldStruct, bagMapPtr, singleBag.color)
		//	for each bag call populateCanHold
	}

	//fmt.Printf("%v\n", workingHoldStruct)
	sum := 0
	for _, singleCanHoldLink := range workingHoldStruct.canHoldMap {
		if singleCanHoldLink.canHoldColor {
			sum++
		}
	}
	fmt.Printf("%v bags can hold %v\n", sum, target)
	//	go through workingHoldStruct, canHoldMap. Count each .canContain = true

}

//this is doing a depth-first search with trimming based on "evaluated"
//there is probably a better way to do this.
func populateCanHold(holdStruct *canHoldStruct, bagMapPtr *map[string]*bag, colorToCheck string) bool {
	bagMapActual := *bagMapPtr
	workingHoldStruct := *holdStruct
	colorToFind := workingHoldStruct.colorToFind

	val, exists := workingHoldStruct.canHoldMap[colorToCheck]
	if exists {
		if val.evaluated {
			return val.canHoldColor
		} else {
			//check if any of the contained bags are the colorToFind
			for containedBag, _ := range bagMapActual[colorToCheck].contains {
				if containedBag == colorToFind {
					val.canHoldColor = true
					val.evaluated = true
					return true
				}
			}
			//if none of the contained bags are the colorToFind, then check whether each child bag can hold colorToFind
			for containedBag, _ := range bagMapActual[colorToCheck].contains {
				found := populateCanHold(holdStruct, bagMapPtr, containedBag)
				if found {
					val.canHoldColor = true
					val.evaluated = false
					return true
				}
			}
			val.canHoldColor = false
			val.evaluated = true
			return false
		}
	} else {
		//      make an entry
		//            call populateCanHold with each contained color
		//                  if any return true, set as true and set as evaluated.
		//                  else set as false and set as evaluated
		//	          return true or false up the stack.
		val = &canHoldLink{
			canHoldColor: false,
			evaluated:    true,
		}
		workingHoldStruct.canHoldMap[colorToCheck] = val
		//check if any of the contained bags are the colorToFind
		for containedBag, _ := range bagMapActual[colorToCheck].contains {
			if containedBag == colorToFind {
				val.canHoldColor = true
				val.evaluated = true
				return true
			}
		}
		//if none of the contained bags are the colorToFind, then check whether each child bag can hold colorToFind
		for containedBag, _ := range bagMapActual[colorToCheck].contains {
			found := populateCanHold(holdStruct, bagMapPtr, containedBag)
			if found {
				val.canHoldColor = true
				val.evaluated = false
				return true
			}
		}
		val.canHoldColor = false
		val.evaluated = true
		return false

	}
	//you should never get here
	val = &canHoldLink{
		canHoldColor: false,
		evaluated:    true,
	}

	return false
}

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
	determineHowManyCanHold("shiny gold", &bagMap)
}

func solvePt2(inputLines []string) {
	bagMap := *(makeBagMap(inputLines))
	targetColor:="shiny gold"
	fmt.Printf("%v can hold %v bags\n",targetColor,bagMap[targetColor].bagsContained)

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
