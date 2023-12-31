package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Mapper struct {
	dest   int
	origin int
	length int
	finish int
}

type RangeOfSeeds struct {
	start int
	end   int
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

	allMaps := make([]map[int]Mapper, 7)
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
	/*
		lowestLocation, err := strconv.Atoi(seeds[0])
		if err != nil {
			log.Fatalf("error converting seed to int: %s", err)
		}

		LargestFinalSoil := findLargestMappedFinal(toSoil)
		LargestFinalFertilizer := findLargestMappedFinal(toFertilizer)
		LargestFinalWater := findLargestMappedFinal(toWater)
		LargestFinalLight := findLargestMappedFinal(toLight)
		LargestFinalTemperature := findLargestMappedFinal(toTemperature)
		LargestFinalHumidity := findLargestMappedFinal(toHumidity)
		LargestFinalLocation := findLargestMappedFinal(toLocation)
		fmt.Println("Calculated largest mapped final for each mapper")
	*/
	seedRanges := make([]RangeOfSeeds, 10)

	allMaps[0] = toSoil
	allMaps[1] = toFertilizer
	allMaps[2] = toWater
	allMaps[3] = toLight
	allMaps[4] = toTemperature
	allMaps[5] = toHumidity
	allMaps[6] = toLocation

	//create starting ranges
	ct := 0
	for i := 0; i < len(seeds); i += 2 {

		snum, err := strconv.Atoi(seeds[i])
		if err != nil {
			log.Fatalf("error converting string: %s", err)
		}
		slength, err := strconv.Atoi(seeds[i+1])
		seedRanges[ct] = RangeOfSeeds{snum, (snum + slength)}
		ct++
	}

	sort.Slice(seedRanges, func(i, j int) bool {
		return seedRanges[i].start < seedRanges[j].start
	})

	//iterate over maps to generate the ranges after each mapping

	for numMpr, mpr := range allMaps {

		for seedRangePos := 0; seedRangePos < len(seedRanges); seedRangePos++ {
			fmt.Println(numMpr)

			for _, m := range mpr {

				//check for mapper function
				if m.origin <= seedRanges[seedRangePos].start && m.finish >= seedRanges[seedRangePos].start {
					//found correct mapper for start of range

					//check if current range falls completely within mapper
					if seedRanges[seedRangePos].end <= m.finish {
						//range falls fully within mapper function no extra ranges needed

						seedRanges[seedRangePos].start = m.dest + seedRanges[seedRangePos].start - m.origin
						seedRanges[seedRangePos].end = m.dest + seedRanges[seedRangePos].end - m.origin
						break
					} else {
						seedRanges = append(seedRanges, RangeOfSeeds{(m.finish + 1), seedRanges[seedRangePos].end})
						seedRanges[seedRangePos].start = m.dest + seedRanges[seedRangePos].start - m.origin
						seedRanges[seedRangePos].end = m.dest + m.length
						break
					}
				} else if m.origin <= seedRanges[seedRangePos].end && m.finish >= seedRanges[seedRangePos].end {
					//this handles the case that the range doesnt start in a mapper but ends in it
					seedRanges = append(seedRanges, RangeOfSeeds{seedRanges[seedRangePos].start, m.origin - 1})
					seedRanges[seedRangePos].start = m.dest
					seedRanges[seedRangePos].end = m.dest + seedRanges[seedRangePos].end - m.origin

					break

				} else if seedRanges[seedRangePos].start < m.origin && seedRanges[seedRangePos].end > m.finish {

					seedRanges = append(seedRanges, RangeOfSeeds{seedRanges[seedRangePos].start, m.origin - 1})
					seedRanges = append(seedRanges, RangeOfSeeds{m.finish + 1, seedRanges[seedRangePos].end})
					seedRanges[seedRangePos].start = m.dest
					seedRanges[seedRangePos].end = m.dest + m.length
					break
				} else {
					continue
				}
			}
		}
	}

	sort.Slice(seedRanges, func(i, j int) bool {
		return seedRanges[i].start < seedRanges[j].start
	})

	fmt.Println(allMaps)
	fmt.Println("The lowest location is: ", seedRanges[0].start)
}

func addToMapperList(list map[int]Mapper, line string) map[int]Mapper {

	if line == "" {
		return list
	}
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
				num = diff + cur.dest
				return num
			}
		}

	}
	//fmt.Println("We did not find any matches")
	return num
}

func genSeeds(seedStart string, seedRange string) []int {

	seedFinish := make([]int, 0)

	s, err := strconv.Atoi(seedStart)
	if err != nil {
		log.Fatalf("error converting seed start: %s", err)
	}
	r, err := strconv.Atoi(seedRange)
	if err != nil {
		log.Fatalf("error converting seed start range: %s", err)
	}

	newSeeds := make([]int, r)
	nSCounter := 0
	for j := s; j < (s + r); j++ {
		newSeeds[nSCounter] = j
		nSCounter++
	}
	seedFinish = append(seedFinish, newSeeds...)

	fmt.Println("boo")

	return seedFinish
}

func findLargestMappedFinal(list map[int]Mapper) int {

	largestVal := 0

	for _, mpr := range list {
		if mpr.finish > largestVal {
			largestVal = mpr.finish
		}
	}

	return largestVal
}

//create all intervals for seeds

//iterate for each mapper function

//for each interval calculate intervals for next function
//adding new intervals to the end?

//return lowest location
