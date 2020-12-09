package Day06

import (
	"fmt"
	"strings"
)

func getCharCountFromString(inputLines []string) map[rune]int {

	inputLine := strings.Join(inputLines, "")

	charMap := make(map[rune]int)
	for _, thisRune := range inputLine {
		curVal, exists := charMap[thisRune]
		if !exists {
			charMap[thisRune] = 1
		} else {
			charMap[thisRune] = curVal + 1
		}
	}
	return charMap
}

func createAnswerMapArray(inputLines []string) ([]map[rune]int, []int) {

	returnSurveyCounts := make([]map[rune]int, 0)
	returnParticipantCounts := make([]int, 0)
	currentString := make([]string, 0)
	for _, singleLine := range inputLines {
		if singleLine == "" {
			returnSurveyCounts = append(returnSurveyCounts, getCharCountFromString(currentString))
			returnParticipantCounts=append(returnParticipantCounts,len(currentString))
			currentString = make([]string, 0)
		} else {
			currentString = append(currentString, singleLine)
		}
	}
	if len(currentString) > 0 {
		returnSurveyCounts = append(returnSurveyCounts, getCharCountFromString(currentString))
		returnParticipantCounts=append(returnParticipantCounts,len(currentString))
	}
	return returnSurveyCounts, returnParticipantCounts
}

func solvePt2(inputLines []string) {

	surveyArray, countArray := createAnswerMapArray(inputLines)

	yesSum := 0

	for x:=0;x<len(surveyArray) ;x++  {
		//thisCount:=0
		thisPartCount:=countArray[x]
		for _,val :=range surveyArray[x]  {
			if val == thisPartCount{
				yesSum++
			}
		}
	}


	fmt.Printf("Questions with all Y: %v\n", yesSum)
}

func solvePt1(inputLines []string) {

	surveyArray, _ := createAnswerMapArray(inputLines)

	yesSum := 0
	for _, singleSurvey := range surveyArray {
		yesSum += len(singleSurvey)
	}
	fmt.Printf("Questions with Y: %v\n", yesSum)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
