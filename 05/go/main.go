package main

import (
	"flag"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := flag.String("input", "input", "Input file relative path")
	part := flag.String("part", "A", "Implementation of part A or B of the problem")

	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		panic(err)
	}

	input := string(data)

	lines := strings.Split(input, "\n")

	var answer int

	switch *part {
	case "A":
		answer = PartA(lines)
	case "B":
		answer = PartB(lines)
	default:
		log.Panic("Invalid input for part")
	}

	log.Println("Answer", answer)
}

func parseInts(row string) []int {
	var ints []int
	numStrings := strings.Fields(row)

	for _, ns := range numStrings {
		num, err := strconv.Atoi(ns)
		if err != nil {
			log.Panic(err)
		}
		ints = append(ints, num)
	}

	return ints
}

type MapRange struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

type CategoryMap struct {
	name   string
	ranges []MapRange
}

func (c *CategoryMap) findDestination(src int) int {
	for _, r := range c.ranges {
		if src >= r.sourceStart && src <= (r.sourceStart + r.rangeLen - 1) {
			return r.destStart + (src - r.sourceStart)
		}
	}
	// If no mapping found, dest = src
	return src
}

func PartA(lines []string) int {
	seedsRow := strings.Split(lines[0], "seeds: ")[1]
	log.Println(seedsRow)

	seeds := parseInts(seedsRow)

	log.Println("Seeds", seeds)

	var categoryMaps []CategoryMap

	currentMap := CategoryMap{}

	for _, line := range lines[2:] {
		if line == "" {
			categoryMaps = append(categoryMaps, currentMap)
			currentMap = CategoryMap{}
			continue
		}

		if strings.HasSuffix(line, " map:") {
			currentMap.name = line
		} else {
			rangeRow := parseInts(line)
			mapRange := MapRange{destStart: rangeRow[0], sourceStart: rangeRow[1], rangeLen: rangeRow[2]}
			currentMap.ranges = append(currentMap.ranges, mapRange)
		}
	}
	categoryMaps = append(categoryMaps, currentMap)

	minLocation := math.MaxInt
	for _, seed := range seeds {
		src := seed
		
		for _, catMap := range categoryMaps {
			dest := catMap.findDestination(src)
			src = dest
		}

		log.Println("Location for seed", seed, src)

		if src < minLocation {
			minLocation = src
		}
	}


	return minLocation
}

func PartB(lines []string) int {
	return len(lines)
}
