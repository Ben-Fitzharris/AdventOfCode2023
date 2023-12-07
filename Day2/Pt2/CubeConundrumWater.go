package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("CubeConundrumData.txt")
	if err != nil {
		log.Fatalf("There was an error reading the file: %s", err)
	}

	sFile := string(file)

	games := strings.Split(sFile, "\n")

	sum := 0

	for _, game := range games {

		prefix := strings.Split(game, ":")
		pulls := strings.Split(prefix[1], ";")

		redMinimum := 0
		greenMinimum := 0
		blueMinimum := 0

		for _, p := range pulls {
			pTrimmed := strings.Trim(p, "\r")
			cubes := strings.Split(pTrimmed, ",")

			for _, c := range cubes {
				a := strings.Split(c, " ")
				number, err := strconv.Atoi(a[1])
				if err != nil {
					log.Fatalf("There was an error converting the number of cubes: %s", err)
				}
				colour := a[2]

				switch colour {
				case "red":
					if number > redMinimum {
						redMinimum = number
					}

				case "green":
					if number > greenMinimum {
						greenMinimum = number
					}

				case "blue":
					if number > blueMinimum {
						blueMinimum = number
					}

				default:
					log.Fatalf("There was an error with checking game possibility")
				}

			}

		}
		sum += (redMinimum * greenMinimum * blueMinimum)

	}

	fmt.Println("The sum of the power of the games is: ", sum)

}
