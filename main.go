package main

import (
	"advent-of-code/Day09"
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

	currentDay := "09"

	//inputLines := readInputFile("./Day" + currentDay + "/test_input.txt")
	inputLines := readInputFile("./Day" + currentDay + "/input.txt")

	start := time.Now()

	Day09.Solve(inputLines)

	elapsed := time.Since(start)
	fmt.Printf("solve took %s\n", elapsed)
}
