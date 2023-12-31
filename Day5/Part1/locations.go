package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Mapper struct {
	destination int
	origin      int
	length      int
	finish      int
}

func main() {

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	strFile := string(file)
	lines := strings.Split(strFile, "\n")

	seeds := make([]string, 0)
	toSoil := make(map[int]Mapper, 0)
	toFertilizer := make(map[int]Mapper, 0)
	toWater := make(map[int]Mapper, 0)
	toLight := make(map[int]Mapper, 0)
	toTemperature := make(map[int]Mapper, 0)
	toHumidity := make(map[int]Mapper, 0)
	toLocation := make(map[int]Mapper, 0)

	for lineNum, line := range lines {

		trimmed := strings.Trim(line, "\r")

		if lineNum == 2 {
			seeds = strings.Split(trimmed, " ")
		} else if 4 < lineNum && lineNum < 32 {
			toSoil = addToMapperList(toSoil, trimmed)
		} else if 33 < lineNum && lineNum < 54 {
			toFertilizer = addToMapperList(toFertilizer, trimmed)
		} else if 55 < lineNum && lineNum < 104 {
			toWater = addToMapperList(toWater, trimmed)
		} else if 105 < lineNum && lineNum < 148 {
			toLight = addToMapperList(toLight, trimmed)
		} else if 149 < lineNum && lineNum < 174 {
			toTemperature = addToMapperList(toTemperature, trimmed)
		} else if 175 < lineNum && lineNum < 201 {
			toHumidity = addToMapperList(toHumidity, trimmed)
		} else if 203 < lineNum && lineNum < 241 {
			toLocation = addToMapperList(toLocation, trimmed)
		}

	}

	lowestLocation, err := strconv.Atoi(seeds[0])
	if err != nil {
		log.Fatalf("error converting seed to int: %s", err)
	}

	for _, seed := range seeds {

		workingSToL, err := strconv.Atoi(seed)
		if err != nil {
			log.Fatalf("There was an error converting a seed: %s", err)
		}
		workingSToL = convertIt(workingSToL, toSoil)
		workingSToL = convertIt(workingSToL, toFertilizer)
		workingSToL = convertIt(workingSToL, toWater)
		workingSToL = convertIt(workingSToL, toLight)
		workingSToL = convertIt(workingSToL, toTemperature)
		workingSToL = convertIt(workingSToL, toHumidity)
		workingSToL = convertIt(workingSToL, toLocation)

		if workingSToL < lowestLocation {
			lowestLocation = workingSToL
		}

	}

	fmt.Println(seeds[0])
	fmt.Println(toFertilizer[0], toHumidity[0], toLight[0], toLocation[0], toSoil[0], toTemperature[0], toWater[0])
	fmt.Println("The lowest location is: ", lowestLocation)
}

func addToMapperList(list map[int]Mapper, line string) map[int]Mapper {

	s := strings.Split(line, " ")
	dest, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatalf("There was an error converting destination: %s", err)
	}
	orig, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatalf("There was an error converting origin: %s", err)
	}
	r, err := strconv.Atoi(s[2])
	if err != nil {
		log.Fatalf("There was an error converting range: %s", err)
	}

	fini := orig + r
	temp := Mapper{dest, orig, r, fini}

	list[orig] = temp

	return list
}

func convertIt(num int, mpr map[int]Mapper) int {

	for _, cur := range mpr {

		if num >= cur.origin {
			if num <= cur.finish {
				diff := num - cur.origin
				num = diff + cur.destination
				return num
			}
		}

	}
	fmt.Println("We did not find any matches")
	return num
}
