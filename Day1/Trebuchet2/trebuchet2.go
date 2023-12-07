package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "TrebuchetInput.txt"
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file: ", filename)
	}

	finalTotal := trebuchet2(input)
	fmt.Println("The final total is: ", finalTotal)
}

func trebuchet2(input []byte) (total int) {

	lines := bytes.Split(input, []byte("\n"))

	for _, l := range lines {
		linestring := string(l)
		fw, sw := checkForNumberWords(linestring)
		numbersFromLine := make([]int, 0)
		for index, b := range l {
			if checkIfNumber(b) {
				numbersFromLine = append(numbersFromLine, index)
			}
		}
		number1 := ""
		number2 := ""
		//compare byte first and last index and word first and last index
		if numbersFromLine[0] < fw[0] {
			number1 = string(l[numbersFromLine[0]])
		} else {
			number1 = fmt.Sprint(fw[1])
		}

		if numbersFromLine[len(numbersFromLine)-1] > sw[0] {
			number2 = string(l[numbersFromLine[len(numbersFromLine)-1]])
		} else {
			number2 = fmt.Sprint(sw[1])
		}
		lineNumber := number1 + number2
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

func checkForNumberWords(lineIn string) (firstWord [2]int, secondWord [2]int) {
	//start with first index of each
	numbers := make([]int, 9)
	numbers[0] = strings.Index(lineIn, "one")
	numbers[1] = strings.Index(lineIn, "two")
	numbers[2] = strings.Index(lineIn, "three")
	numbers[3] = strings.Index(lineIn, "four")
	numbers[4] = strings.Index(lineIn, "five")
	numbers[5] = strings.Index(lineIn, "six")
	numbers[6] = strings.Index(lineIn, "seven")
	numbers[7] = strings.Index(lineIn, "eight")
	numbers[8] = strings.Index(lineIn, "nine")

	lowestIndex := 1000
	firstValue := 0

	for idx, val := range numbers {
		if val == -1 {
			continue
		}

		if val < lowestIndex {
			lowestIndex = val
			firstValue = idx + 1
		}
	}

	numbers[0] = strings.LastIndex(lineIn, "one")
	numbers[1] = strings.LastIndex(lineIn, "two")
	numbers[2] = strings.LastIndex(lineIn, "three")
	numbers[3] = strings.LastIndex(lineIn, "four")
	numbers[4] = strings.LastIndex(lineIn, "five")
	numbers[5] = strings.LastIndex(lineIn, "six")
	numbers[6] = strings.LastIndex(lineIn, "seven")
	numbers[7] = strings.LastIndex(lineIn, "eight")
	numbers[8] = strings.LastIndex(lineIn, "nine")

	highestIndex := -1
	lastValue := 0

	for idx, val := range numbers {
		if val == -1 {
			continue
		}
		if val > highestIndex {
			highestIndex = val
			lastValue = idx + 1
		}

	}

	firstWord = [2]int{lowestIndex, firstValue}
	secondWord = [2]int{highestIndex, lastValue}

	return firstWord, secondWord
}
