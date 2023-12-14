package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	file, err := os.ReadFile("partNumbers.txt")
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	ast := make([]byte, 1)
	ast = []byte("&")
	//ast := "&"
	fmt.Printf("Is a recognised as a unicode symbol: %t \n", containsSymbol(ast))

	sumParts := 0

	lines := bytes.Split(file, []byte("\n"))

	for lineNumber, l := range lines {
		//extract numbers and remove symbols
		lineIndexToBeAdded := 0
		lineTotal := 0
		numbers := bytes.Split(l, []byte("."))
		//numbersAndIndexes := make([][]int, len(numbers))
		splitNumbers := make([][]byte, len(numbers)+12)
		splitNumbersIdx := 0
		strLine := string(l)

		//sometimes need to split numbers up, add all numbers to new slice.
		for _, n := range numbers {
			if len(n) > 4 {
				newNums := bytes.Split(n, []byte("*"))
				//numList.PushBack(newNums[0])
				if len(newNums) > 1 {
					for _, new := range newNums {

						splitNumbers[splitNumbersIdx] = new
						splitNumbersIdx++
					}
				}
			} else {
				splitNumbers[splitNumbersIdx] = n
				splitNumbersIdx++
			}
		}

		for _, num := range splitNumbers {

			//nb := bytes.Runes(num)
			trimmedNum := make([]byte, 5)
			if bytes.ContainsFunc(num, unicode.IsNumber) && containsSymbol(num) {
				trimmedNum = bytes.TrimFunc(num, unicode.IsSymbol)
				if containsSymbol(trimmedNum) {
					trimmedNum = bytes.Trim(trimmedNum, "*&\r/-!@#$%^()_=+")
				}

			} else if bytes.ContainsFunc(num, unicode.IsNumber) {

				trimmedNum = num

			} else {
				continue
			}
			strNum := string(trimmedNum)
			start := (strings.Index(strLine, strNum) + lineIndexToBeAdded)
			numLength := len(strNum)
			end := start + numLength - 1
			t := strings.SplitAfterN(strLine, strNum, 2)
			lineIndexToBeAdded += len(t[0])
			if len(t) > 1 {
				strLine = t[1]
			}

			//addNumTo3WideArr(numbersAndIndexes, idx, trimmedNum, start, end)

			symbolPresent := checkForSymbol(lines, lineNumber, start, end)
			if symbolPresent {
				iNum, err := strconv.Atoi(strNum)
				fmt.Println(strNum)
				if err != nil {
					log.Fatalf("error converting from string to int: %s", err)
				}
				sumParts += iNum
				lineTotal += iNum
			}
		}
		//fmt.Printf("%d\n", lineTotal)
	}

	fmt.Println("The final total of all part numbers is: ", sumParts)
}

func addNumTo3WideArr(m [][]int, index int, number []byte, start int, end int) {
	numberConverted, err := strconv.Atoi(string(number))
	if err != nil {
		log.Fatalf("error converting number to int: %s", err)
	}
	m[index][0] = numberConverted
	m[index][1] = start
	m[index][2] = end

	return
}

func checkForSymbol(lines [][]byte, lineNumber int, start int, end int) bool {

	symStart := start - 1
	symStop := end + 2

	//do not look over edges
	if symStart < 0 {
		symStart = 0
	}
	if symStop > len(lines[lineNumber])-1 {
		symStop = len(lines[lineNumber]) - 1
	}
	//check line above
	if lineNumber-1 > -1 {
		PartialTopLine := lines[lineNumber-1][symStart:symStop]
		if containsSymbol(PartialTopLine) {
			return true
		}
	}
	//check left and right of number
	left := []byte{lines[lineNumber][symStart]}
	right := []byte{lines[lineNumber][symStop-1]}
	if containsSymbol(left) || containsSymbol(right) {
		return true
	}
	//check line below
	if lineNumber+1 < len(lines) {
		PartialLowerLine := lines[lineNumber+1][symStart:symStop]
		if containsSymbol(PartialLowerLine) {
			return true
		}
	}

	return false
}

func containsSymbol(input []byte) bool {

	for _, b := range input {
		if !unicode.IsDigit(rune(b)) && string(b) != "." {
			return true
		}

	}
	return false
}
