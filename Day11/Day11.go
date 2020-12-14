package Day11

import (
	"fmt"
	"strings"
	"sync"
)

type space struct {
	isSeat    bool
	occupied  bool
	nextState bool
	neighbors []*space
}

func (self *space) countOccupiedNeighbors() int {
	tally := 0
	for _, neighbor := range self.neighbors {
		if neighbor.occupied {
			tally++
		}
	}
	return tally
}

//If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
//If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
//Otherwise, the seat's state does not change.
func (self *space) populateNextState() {
	occupiedNeighbors := self.countOccupiedNeighbors()
	if self.occupied {
		if occupiedNeighbors >= 4 {
			self.nextState = false
		}
	} else {
		if occupiedNeighbors == 0 {
			self.nextState = true
		}
	}
}

func (self *space) concPopulateNextState(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	self.populateNextState()
}

type board struct {
	gameSpace [][]*space
}

func (self *board) getBoardString() string {
	returnArray := make([]string, len(self.gameSpace))
	for j, singleRow := range self.gameSpace {
		thisRow := make([]rune, len(singleRow))
		for i, singleSpace := range singleRow {
			if !singleSpace.isSeat {
				thisRow[i] = '.'
			} else if singleSpace.occupied {
				thisRow[i] = '#'
			} else {
				thisRow[i] = 'L'
			}
		}
		returnArray[j] = string(thisRow)
	}
	return strings.Join(returnArray, "\n")
}

func (self *board) processTimeStep() {
	var waitGroup sync.WaitGroup

	for _, singleRow := range self.gameSpace {
		for _, singleSpace := range singleRow {
			if singleSpace.isSeat {
				waitGroup.Add(1)
				go singleSpace.concPopulateNextState(&waitGroup)
			}
		}
	}

	waitGroup.Wait()
	for _, singleRow := range self.gameSpace {
		for _, singleSpace := range singleRow {
			singleSpace.occupied = singleSpace.nextState
		}
	}
}

func (self *board) countOccupiedSeats() int {
	tally := 0
	for _, singleRow := range self.gameSpace {
		for _, singleSpace := range singleRow {
			if singleSpace.occupied {
				tally++
			}
		}
	}
	return tally
}

func createBoard(inputLines []string) board {

	gameSpace := make([][]*space, 0)

	for _, singleLine := range inputLines {
		singleRow := make([]*space, 0)
		for _, singleRune := range singleLine {
			//
			newSpace := space{
				isSeat:    false,
				occupied:  false,
				nextState: false,
				neighbors: nil,
			}
			if singleRune == 'L' {
				newSpace.isSeat = true
			}
			singleRow = append(singleRow, &newSpace)
		}
		gameSpace = append(gameSpace, singleRow)
	}

	for y := 0; y < len(gameSpace); y++ {
		for x := 0; x < len(gameSpace[y]); x++ {
			if gameSpace[y][x].isSeat {
				for yOffset := -1; yOffset < 2; yOffset++ {
					for xOffset := -1; xOffset < 2; xOffset++ {
						neighborY := y + yOffset
						neighborX := x + xOffset
						if neighborY >= 0 && neighborY < len(gameSpace) {
							if neighborX >= 0 && neighborX < len(gameSpace[y]) {
								if neighborX == x && neighborY == y {
									//	don't add yourself as neighbor
								} else if gameSpace[neighborY][neighborX].isSeat {
									gameSpace[y][x].neighbors = append(gameSpace[y][x].neighbors, gameSpace[neighborY][neighborX])
								}
							}
						}
					}
				}
			}
		}
	}

	return board{gameSpace: gameSpace}

}

func solvePt1(inputLines []string) {
	thisBoard := createBoard(inputLines)
	lastState := thisBoard.getBoardString()
	thisBoard.processTimeStep()
	currentState := thisBoard.getBoardString()
	for lastState != currentState {
		lastState = currentState
		thisBoard.processTimeStep()
		currentState = thisBoard.getBoardString()
	}

	fmt.Println(thisBoard.getBoardString())
	fmt.Printf("%v occupied seats\n", thisBoard.countOccupiedSeats())
}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
