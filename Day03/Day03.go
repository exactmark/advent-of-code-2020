package Day03

import "fmt"

func solveGeneralSlope(inputLines []string, xMove int, yMove int) int {
	currentX := 0
	numTrees := 0
	repeatLength := len(inputLines[0])

	for y := 0; y < len(inputLines); y += yMove {
		single_line := inputLines[y]
		if single_line[currentX] == '#' {
			numTrees++
		}
		currentX += xMove
		if currentX >= repeatLength {
			currentX -= repeatLength
		}
	}
	return numTrees
}

func solvePt1(inputLines []string) {
	numTrees := solveGeneralSlope(inputLines, 3, 1)
	fmt.Printf("Encountered %v trees.", numTrees)
}

func solvePt2(inputLines []string) {
	possibleSlopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	productNumTrees := 1
	for _, slopeVals := range possibleSlopes {
		thisNumTrees := solveGeneralSlope(inputLines, slopeVals[0], slopeVals[1])
		fmt.Printf("For slope %v, %v found %v trees.\n", slopeVals[0], slopeVals[1], thisNumTrees)
		if thisNumTrees > 0 {
			productNumTrees *= thisNumTrees
		}
	}

	fmt.Printf("Solution is %v trees.", productNumTrees)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
