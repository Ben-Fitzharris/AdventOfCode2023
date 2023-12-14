package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Gear struct {
	lineIndex int
	idx       int
}

type numLocation struct {
	value     int
	lineIndex int
	start     int
	end       int
}

func main() {

	file, err := os.ReadFile("../Part1/partNumbers.txt")
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	sumParts := 0

	lines := bytes.Split(file, []byte("\n"))
	numberOfLines := len(lines)

	//gears := make([][]Gear, len(lines))

	numberReg := regexp.MustCompile(`\d+`)
	gearReg := regexp.MustCompile(`\*`)

	allNumbers := make([]numLocation, 0)
	allGears := make([]Gear, 0)

	//gather all numbers
	for lineIdx, line := range lines {
		strLine := string(line)
		numMatches := numberReg.FindAllStringIndex(strLine, -1)
		gearMatches := gearReg.FindAllStringIndex(strLine, -1)

		for _, match := range numMatches {
			s := string(line[match[0]:match[1]])
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("error converting string to number: %s", err)
			}
			t := numLocation{
				value:     v,
				lineIndex: lineIdx,
				start:     match[0],
				end:       match[1],
			}
			allNumbers = append(allNumbers, t)
		}

		for _, match := range gearMatches {
			t := Gear{
				lineIndex: lineIdx,
				idx:       match[0],
			}
			allGears = append(allGears, t)
		}

	}

	for _, gear := range allGears {

		touchedNumbers := make([]int, 0)
		for _, n := range allNumbers {

			if checkNumberTouch(gear, n, numberOfLines) {
				touchedNumbers = append(touchedNumbers, n.value)
			}

		}
		if len(touchedNumbers) == 2 {
			sumParts += touchedNumbers[0] * touchedNumbers[1]
		}

	}

	fmt.Println("The total of all parts is: ", sumParts)
}

func checkNumberTouch(gear Gear, num numLocation, lenLines int) bool {
	gearLeft := gear.idx - 1
	gearRight := gear.idx + 1
	aboveGear := gear.lineIndex - 1
	belowGear := gear.lineIndex + 1
	//assumption: num.end is the location of the char after the end

	//check if line index even worth considering
	if aboveGear <= num.lineIndex && num.lineIndex <= belowGear {

		//check row above
		if gear.lineIndex > 0 {
			if num.lineIndex == gear.lineIndex-1 {
				if (num.start <= gearRight && num.start >= gearLeft) || (num.end-1 >= gearLeft && num.end-1 <= gearRight) {
					return true
				}
			}

		}

		//check left and right of gear
		if num.lineIndex == gear.lineIndex {
			if num.end == gear.idx || num.start == gearRight {
				return true
			}
		}

		//check row below

		if gear.lineIndex < lenLines-1 {
			if num.lineIndex == gear.lineIndex+1 {
				if (num.start <= gearRight && num.start >= gearLeft) || (num.end-1 >= gearLeft && num.end-1 <= gearRight) {
					return true
				}
			}
		}
	}

	return false
}
