package Day08

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type instruction struct {
	command string
	scale   int
}

type runCode struct {
	runLine    instruction
	orderExec  int
	hasVisited bool
}

type runTimeEnv struct {
	accumulator int
	pointer     int
	steps       int
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
	thisRunTime.steps++
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
	//fmt.Printf("%v %v\n", nextInstruction.runLine.command, nextInstruction.runLine.scale)
	nextInstruction.orderExec = thisRunTime.steps
	thisRunTime.pointer += step
	if thisRunTime.pointer >= len(thisRunTime.codeBase) {
		//thisRunTime.pointer-=step
		thisRunTime.accumulator += scale
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
		fmt.Printf("Accumulator is %v with bailresult of '%v' at %v\n", thisRunTime.accumulator, bailResult, thisRunTime.pointer)
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

func (thisRunTime *runTimeEnv) startProcessingPt2() {
	bailResult := ""
	for {
		bailResult = thisRunTime.processCurrentInstruction()
		if bailResult != "" {
			break
		}
	}
	if bailResult == "hit EOF" {
		fmt.Printf("Accumulator is %v with bailresult of '%v' at %v\n", thisRunTime.accumulator, bailResult, thisRunTime.pointer)
	} else {
		//fmt.Printf("Accumulator is %v with bailresult of '%v' at %v\n", thisRunTime.accumulator, bailResult, thisRunTime.pointer)

	}
}

func swapInstructionAndRun(waitgroup *sync.WaitGroup, swapPtr int, instructionList []instruction) {
	defer waitgroup.Done()

	targetInstruction := instructionList[swapPtr]
	if targetInstruction.command == "nop" || targetInstruction.command == "jmp" {
		instructionCopy := make([]instruction, len(instructionList))
		copy(instructionCopy, instructionList)
		newCommand := "nop"
		if targetInstruction.command == "nop" {
			newCommand = "jmp"
		}
		instructionInsert := instruction{
			command: newCommand,
			scale:   targetInstruction.scale,
		}
		//fmt.Printf("%v\n%v\n",targetInstruction,instructionInsert)
		//fmt.Printf("%v\n",instructionList)
		instructionCopy[swapPtr] = instructionInsert
		//fmt.Printf("%v\n",instructionCopy)
		thisRunTime := createRunTimeEnv(&instructionCopy)
		thisRunTime.startProcessingPt2()
	}
}

func solvePt2(inputLines []string) {
	var waitgroup sync.WaitGroup

	instructionList := createInstructionList(inputLines)

	for i, _ := range *instructionList {
		waitgroup.Add(1)
		go swapInstructionAndRun(&waitgroup, i, *instructionList)
	}
	waitgroup.Wait()

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
