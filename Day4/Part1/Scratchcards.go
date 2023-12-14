package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type card struct {
	points int
	count  int
}

func main() {

	file, err := os.ReadFile("scratchcardNumbers.txt")
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	strFile := string(file)
	lines := strings.Split(strFile, "\n")

	numberReg := regexp.MustCompile(`\d+`)

	sum := 0
	stackOfCards := make([]card, len(lines))

	for cNum, line := range lines {

		removeCN := strings.Split(line, ":")
		nums := strings.Split(removeCN[1], "|")
		winning := numberReg.FindAllString(nums[0], -1)
		scratch := numberReg.FindAllString(nums[1], -1)

		winningCount := 0
		linePoints := 0

		for _, n := range scratch {
			for _, wn := range winning {
				if n == wn {
					winningCount++
				}
			}
		}
		//add 1 to count as the original
		stackOfCards[cNum].count += 1
		//add copies to later cards
		for i := 0; i < winningCount; i++ {
			if i == 0 {
				linePoints = 1
			} else {
				linePoints = linePoints * 2
			}

			stackOfCards[cNum+1+i].count += stackOfCards[cNum].count
		}
		stackOfCards[cNum].points = linePoints

		sum += stackOfCards[cNum].count

	}

	fmt.Println("The sum of the winning cards is: ", sum)
}
