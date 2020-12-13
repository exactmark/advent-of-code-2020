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
	codeBase    []*runCode
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

func createRunTimeEnv(instructionList *[]instruction) *runTimeEnv {
	runCodeArr := make([]*runCode, len(*instructionList))
	for i, singleInstruct := range *instructionList {
		runCodeArr[i] = &runCode{
			runLine:    singleInstruct,
			hasVisited: false,
		}
	}
	return &runTimeEnv{
		accumulator: 0,
		pointer:     0,
		codeBase:    runCodeArr,
	}
}

func (thisRunTime *runTimeEnv) processCurrentInstruction() string {
	nextInstruction := thisRunTime.codeBase[thisRunTime.pointer]
	step := 0
	scale := 0
	switch nextCommand := nextInstruction.runLine.command; nextCommand {

	case "nop":
		step = 1
	case "acc":
		step = 1
		scale = nextInstruction.runLine.scale
	case "jmp":
		step = nextInstruction.runLine.scale
	}

	thisRunTime.pointer += step
	if thisRunTime.pointer >= len(thisRunTime.codeBase) {
		//thisRunTime.pointer-=step
		return "hit EOF"
	} else if thisRunTime.codeBase[thisRunTime.pointer].hasVisited {
		//thisRunTime.pointer-=step
		return "hit loop"
	} else {
		thisRunTime.codeBase[thisRunTime.pointer].hasVisited = true
		thisRunTime.accumulator += scale
	}
	return ""

}

func (thisRunTime *runTimeEnv) startProcessing() {
	bailResult := ""
	for {
		bailResult = thisRunTime.processCurrentInstruction()
		if bailResult != "" {
			break
		}
	}
	if bailResult != "" {
		fmt.Printf("Accumulator is %v with bailresult of '%v' at %v\n", thisRunTime.accumulator, bailResult,thisRunTime.pointer)
	} else {
		fmt.Printf("Accumulator is %v\n", thisRunTime.accumulator)
	}
}

func (thisRunTime *runTimeEnv) runCodeFromStart() {
	thisRunTime.accumulator = 0
	thisRunTime.pointer = 0

}

func solvePt1(inputLines []string) {
	instructionList := createInstructionList(inputLines)
	//fmt.Printf("%v\n", instructionList)
	thisRunTime := createRunTimeEnv(instructionList)
	thisRunTime.startProcessing()
}

func solvePt2(inputLines []string) {

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
