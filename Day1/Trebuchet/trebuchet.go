package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := "TrebuchetInput.txt"
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file: ", filename)
	}

	finalTotal := trebuchet(input)
	fmt.Println("The final total is: ", finalTotal)
}

func trebuchet(input []byte) (total int) {

	lines := bytes.Split(input, []byte("\n"))

	for _, l := range lines {
		numbersFromLine := make([]int, 0)
		for index, b := range l {
			if checkIfNumber(b) {
				numbersFromLine = append(numbersFromLine, index)
			}
		}
		number1 := string(l[numbersFromLine[0]])
		number2 := string(l[numbersFromLine[len(numbersFromLine)-1]])
		var lineNumber string
		lineNumber = number1 + number2
		num, err := strconv.Atoi(lineNumber)
		if err != nil {
			fmt.Println("conversion error")
		}
		total += num
	}

	return total

}

func checkIfNumber(in byte) (isNumber bool) {
	if int(in) <= 57 && int(in) >= 48 {
		isNumber = true
	} else {
		isNumber = false
	}

	return isNumber
}
