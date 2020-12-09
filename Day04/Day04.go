package Day04

import (
	"fmt"
	"regexp"
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

func pt2PassportIsValid(passport map[string]string) bool {
	//byr (Birth Year) - four digits; at least 1920 and at most 2002.
	//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	//hgt (Height) - a number followed by either cm or in:
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	//pid (Passport ID) - a nine-digit number, including leading zeroes.
	//cid (Country ID) - ignored, missing or not.

	keyFmts := [][]string{{"byr", "^(19[2-9][0-9]|200[0-2])$"},
		{"iyr", "^(2020|201[0-9])$"},
		{"eyr", "^(2030|202[0-9])$"},
		{"hgt", "^(1([5-8][0-9]|9[0-3])cm|(59|6[0-9]|7[0-6])in)$"},
		{"hcl", "^#([0-9]|[a-f]){6}$"},
		{"ecl", "^(amb|blu|brn|gry|grn|hzl|oth)$"},
		{"pid", "^[0-9]{9}$"}}

	for _, rule := range keyFmts {
		val, exists := passport[rule[0]]
		if !exists {
			return false
		}
		isValid, _ := regexp.MatchString(rule[1], val)
		if !isValid {
			return false
		}
	}

	return true
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

func solvePt2(inputLines []string) {
	allPassports := getPassportListFromInput(inputLines)
	validPassports := 0
	for _, singlePassPort := range allPassports {
		if pt2PassportIsValid(singlePassPort) {
			validPassports++
		}
	}
	fmt.Printf("Found %v valid passports\n", validPassports)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
