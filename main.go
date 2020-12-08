package main

import (
	"advent-of-code/Day02"
	"bufio"
	"log"
	"os"
)

func readInputFile(filename string) []string {
	var returnStrings [] string
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

	inputLines := readInputFile("input.txt")
	Day02.Solve(inputLines)
}
