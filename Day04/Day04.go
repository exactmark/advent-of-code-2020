package Day04

import (
	"fmt"
	"strings"
)

func readPassport(inputLines []string) map[string]string {
	returnMap := make(map[string]string)
	inputString := strings.Join(inputLines, " ")
	for _, inputVal := range strings.Split(inputString, " ") {
		keyValPieces := strings.Split(inputVal, ":")
		_, exists := returnMap[keyValPieces[0]]
		if exists {
			panic("val already exists")
		} else {
			returnMap[keyValPieces[0]] = keyValPieces[1]
		}
	}
	return returnMap
}

func getPassportListFromInput(inputLines []string) []map[string]string {
	returnArray := make([]map[string]string, 0)
	currentString := make([]string, 0)
	for _, singleLine := range inputLines {
		if singleLine == "" {
			returnArray = append(returnArray, readPassport(currentString))
			currentString = make([]string, 0)
		} else {
			currentString = append(currentString, singleLine)
		}
	}
	if len(currentString) > 0 {
		returnArray = append(returnArray, readPassport(currentString))
	}
	return returnArray
}

func passportMapIsValid(passport map[string]string) bool {
	//Let's just be explicit here
	if len(passport) == 8 {
		return true
	} else if len(passport) == 7 {
		_, exists := passport["cid"]
		if exists {
			//if cid is here, something important is missing.
			return false
		} else {
			return true
		}
	}

	return false
}

func solvePt1(inputLines []string) {
	allPassports := getPassportListFromInput(inputLines)
	validPassports := 0
	for _, singlePassPort := range allPassports {
		if passportMapIsValid(singlePassPort) {
			validPassports++
		}
	}
	fmt.Printf("Found %v valid passports\n", validPassports)
}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
