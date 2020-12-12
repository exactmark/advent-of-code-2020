package main

import (
	"advent-of-code/Day07"
	"bufio"
	"log"
	"os"
)

func readInputFile(filename string) []string {
	var returnStrings []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		returnStrings = append(returnStrings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return returnStrings
}

func main() {

	currentDay := "07"

	inputLines := readInputFile("./Day"+currentDay+"/test_input.txt")
	//inputLines := readInputFile("./Day" + currentDay + "/input.txt")
	Day07.Solve(inputLines)

}
