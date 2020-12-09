package Day05

import (
	"fmt"
	"sort"
	"strconv"
)

type Ticket struct{
	ticketString string
	rowNumber int
	seatNumber int
	seatID int
}

func convTicketPartToInt(inputString string) int {
	workingBinary := ""
	for _, single_rune := range inputString {
		if single_rune == 'B' || single_rune == 'R' {
			workingBinary += "1"
		} else
		{
			workingBinary += "0"
		}
	}
	i, err := strconv.ParseInt(workingBinary, 2, 64)
	if err != nil {
		fmt.Println(err)
		panic("incorrect binary parse")
	}
	return int(i)
}

func createTicket(ticketString string)Ticket{
	rowLetters := ticketString[0:7]
	seatLetters := ticketString[7:10]
	rowNumber := convTicketPartToInt(rowLetters)
	seatNumber:=convTicketPartToInt(seatLetters)
	seatID := rowNumber*8 + seatNumber
	return Ticket{
		ticketString: ticketString,
		rowNumber:    rowNumber,
		seatNumber:   seatNumber,
		seatID:       seatID,
	}

}

//make a sort!
type BySeatId []Ticket
func (a BySeatId) Len() int           { return len(a) }
func (a BySeatId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySeatId) Less(i, j int) bool { return a[i].seatID < a[j].seatID }


func solvePt1(inputLines []string) {

	ticketList:=make(map[string]Ticket,0)
	highestId:=0
	for _,singleTicket := range inputLines{
		thisTicket:=createTicket(singleTicket)
		if thisTicket.seatID>highestId{
			highestId=thisTicket.seatID
		}
		ticketList[singleTicket]=thisTicket
	}

	fmt.Printf("Highest ticket id is %v\n",highestId)
}

func solvePt2(inputLines[]string){
	ticketList:=make([]Ticket,0)
	for _,singleTicket := range inputLines{
		thisTicket:=createTicket(singleTicket)
		ticketList= append(ticketList,thisTicket)
	}
	sort.Sort(BySeatId(ticketList))
	currentTicket := ticketList[0].seatID
	for _,singleTicket := range ticketList{
		if singleTicket.seatID == currentTicket{
			currentTicket++
		}else {
			break
		}
	}
	fmt.Printf("My ticket is %v\n",currentTicket)

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
