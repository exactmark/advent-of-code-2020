package main

import (
	"advent-of-code/Day04"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
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

	start := time.Now()
	currentDay := "04"

	//inputLines := readInputFile("./Day"+currentDay+"/test_input.txt")
	inputLines := readInputFile("./Day" + currentDay + "/input.txt")
	Day04.Solve(inputLines)
	elapsed := time.Since(start)
	fmt.Printf("solve took %s", elapsed)
}
