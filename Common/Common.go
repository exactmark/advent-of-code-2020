package Common

import "strconv"

func ConvertInputToIntArray(inputLines []string)[]int{
	returnArray:=make([]int,len(inputLines))
	for i,singleLine := range inputLines{
		returnArray[i],_=strconv.Atoi(singleLine)
	}
	return returnArray
}
