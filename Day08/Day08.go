package Day08

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	command string
	scale   int
}

type runCode struct {
	runLine    instruction
	hasVisited bool
}

type runTimeEnv struct {
	accumulator int
	pointer     int
	codeBase    []runCode
}

func parseSingleLineToInstruction(singleLine string) instruction {
	lineSplit := strings.Split(singleLine, " ")
	scale, _ := strconv.Atoi(lineSplit[1])
	return instruction{
		command: lineSplit[0],
		scale:   scale,
	}
}

func createInstructionList(inputLines []string) *[]instruction {
	instructionList := make([]instruction, 0)
	for _, singleLine := range inputLines {
		instructionList = append(instructionList, parseSingleLineToInstruction(singleLine))
	}
	return &instructionList
}

func solvePt1(inputLines []string) {
	instructionList := createInstructionList(inputLines)
	fmt.Printf("%v\n", instructionList)
}

func solvePt2(inputLines []string) {

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
