package Day12

import (
	"fmt"
	"math"
	"strconv"
)

type ship struct {
	heading        int
	compassHeading rune
	xPos           int
	yPos           int
	xVelocity      int
	yVelocity      int
}

func parseInputToMove(inputLine string) (rune, int) {
	command := rune(inputLine[0])
	scale, _ := strconv.Atoi(inputLine[1:])
	return command, scale
}

func (self *ship) normalizeHeading() {
	for self.heading < 0 {
		self.heading += 360
	}
	for self.heading >= 360 {
		self.heading -= 360
	}
}

func (self *ship) setCompassHeading() {
	switch self.heading {
	case 0:
		self.compassHeading = 'N'
	case 90:
		self.compassHeading = 'E'
	case 180:
		self.compassHeading = 'S'
	case 270:
		self.compassHeading = 'W'
	default:
		panic("non-compass heading")
	}
}

func (self *ship) changeHeading(headingDelta int) {
	self.heading = self.heading + headingDelta
	self.normalizeHeading()
	self.setCompassHeading()
}

func (self *ship) processPt1Move(inputLine string) {
	fmt.Printf("Processing %v\n", inputLine)
	command, scale := parseInputToMove(inputLine)

	xDelta := 0
	yDelta := 0
	headingDelta := 0

	switch command {

	case 'N':
		yDelta = scale
	case 'S':
		yDelta = -scale
	case 'E':
		xDelta = scale
	case 'W':
		xDelta = -scale
	case 'L':
		headingDelta = -scale
	case 'R':
		headingDelta = scale
	case 'F':
		fCommand := string(self.compassHeading) + strconv.Itoa(scale)
		//fmt.Printf("Processing %v\n",fCommand)
		self.processPt1Move(fCommand)
	}

	if headingDelta != 0 {
		self.changeHeading(headingDelta)
	}

	self.xPos += xDelta
	self.yPos += yDelta

}

func solvePt1(inputLines []string) {

	myShip := ship{
		heading:        90,
		compassHeading: 'E',
		xPos:           0,
		yPos:           0,
	}
	for _, singleCommand := range inputLines {
		myShip.processPt1Move(singleCommand)
		fmt.Printf("ship is at %v,%v\n", myShip.xPos, myShip.yPos)
	}

	fmt.Printf("Ship has moved %v\n", math.Abs(float64(myShip.yPos))+math.Abs(float64(myShip.xPos)))
}

func solvePt2(inputLines []string) {

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
