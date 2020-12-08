package Day02

import (
	"fmt"
	"strconv"
	"strings"
)

type PasswordEntry struct {
	password   string
	lowCount   int
	highCount  int
	targetRune rune
	matches    bool
	pt2Matches bool
}

func (thisEntry *PasswordEntry) populateMatches() {
	matchCount := 0
	passwordRunes := []rune(thisEntry.password)
	for _, singleRune := range passwordRunes {
		if thisEntry.targetRune == singleRune {
			matchCount++
		}
	}
	if matchCount <= thisEntry.highCount && matchCount >= thisEntry.lowCount {
		thisEntry.matches = true
	} else {
		thisEntry.matches = false
	}
}

func (thisEntry *PasswordEntry) populatePt2Matches() {
	passwordRunes := []rune(thisEntry.password)
	//Use != as XOR
	if (passwordRunes[thisEntry.lowCount-1] == thisEntry.targetRune) != (passwordRunes[thisEntry.highCount-1] == thisEntry.targetRune) {
		thisEntry.pt2Matches = true
	} else {
		thisEntry.pt2Matches = false
	}
}

func makePasswordEntry(test_string string) PasswordEntry {
	// example test_string
	// 1-3 a: abcde
	stringParts := strings.Split(test_string, ": ")
	thisPassword := PasswordEntry{}
	thisPassword.password = stringParts[1]
	rules := strings.Split(stringParts[0], " ")
	//I'm sure there's a better way to get one rune.
	thisPassword.targetRune = []rune(rules[1])[0]
	thisPassword.lowCount, _ = strconv.Atoi(strings.Split(rules[0], "-")[0])
	thisPassword.highCount, _ = strconv.Atoi(strings.Split(rules[0], "-")[1])
	thisPassword.populateMatches()
	thisPassword.populatePt2Matches()
	return thisPassword
}

func Solve(inputLines []string){
	numberValid := 0
	pt2NumberValid := 0
	for _, singleLine := range inputLines {
		singlePasswordEntry := makePasswordEntry(singleLine)

		if singlePasswordEntry.matches {
			numberValid++
		}
		if singlePasswordEntry.pt2Matches {
			pt2NumberValid++
		}

	}
	fmt.Printf("There were %v valid part 1 matches.\n", numberValid)

	fmt.Printf("There were %v valid part 2 matches.\n", pt2NumberValid)
}