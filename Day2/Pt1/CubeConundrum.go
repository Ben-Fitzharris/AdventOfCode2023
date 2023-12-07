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

	redMax := 12
	greenMax := 13
	blueMax := 14
	gameTotal := 0

	for _, game := range games {

		gamePossible := true
		prefix := strings.Split(game, ":")
		game := strings.Split(prefix[0], " ")
		gameNumber := game[1]
		pulls := strings.Split(prefix[1], ";")

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
					if number > redMax {
						gamePossible = false
					}

				case "green":
					if number > greenMax {
						gamePossible = false
					}

				case "blue":
					if number > blueMax {
						gamePossible = false
					}

				default:
					log.Fatalf("There was an error with checking game possibility")
				}
				if !gamePossible {
					break
				}

			}

			if !gamePossible {
				break
			}
		}

		if gamePossible {
			gN, err := strconv.Atoi(gameNumber)
			if err != nil {
				log.Fatalf("There was an error converting game number to int: %s", err)
			}
			gameTotal += gN
		}
	}

	fmt.Println("The sum of all valid game numbers is: ", gameTotal)

}
