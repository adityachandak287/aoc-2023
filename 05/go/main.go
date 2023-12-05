package main

import (
	"flag"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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
	case "BV2":
		answer = PartBV2(lines)
	case "BV3":
		answer = PartBV3(lines)
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
		if src >= r.sourceStart && src <= (r.sourceStart+r.rangeLen-1) {
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

	for idx := 0; idx < len(seeds); idx += 2 {
		start := time.Now()

		startRange := seeds[idx]
		rangeLen := seeds[idx+1]
		endRange := seeds[idx] + rangeLen

		log.Printf("Seeds [%d, %d) starting", startRange, endRange)
		for seed := startRange; seed < endRange; seed += 1 {
			src := seed
			for _, catMap := range categoryMaps {
				dest := catMap.findDestination(src)
				src = dest
			}

			// log.Println("Location for seed", seed, src)

			if src < minLocation {
				minLocation = src
			}
		}
		log.Printf("Finished %d seeds in %s: min location %d", rangeLen, time.Since(start), minLocation)
	}

	return minLocation
}

// Added concurrency
func PartBV2(lines []string) int {
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

	var wg sync.WaitGroup

	numSeedPairs := len(seeds) / 2

	results := make(chan int, numSeedPairs)

	for idx := 0; idx < len(seeds); idx += 2 {
		wg.Add(1)

		startRange := seeds[idx]
		endRange := seeds[idx] + seeds[idx+1]

		go func(startSeed int, endSeed int, result chan int) {
			defer wg.Done()
			localMin := math.MaxInt
			start := time.Now()
			log.Printf("Seeds [%d, %d) starting", startSeed, endSeed)
			for seed := startSeed; seed < endSeed; seed += 1 {
				src := seed
				for _, catMap := range categoryMaps {
					dest := catMap.findDestination(src)
					src = dest
				}

				// log.Println("Location for seed", seed, src)

				if src < localMin {
					localMin = src
				}
			}
			log.Printf("Finished %d seeds in %s: min location %d", endSeed-startSeed, time.Since(start), localMin)
			result <- localMin
		}(startRange, endRange, results)

	}

	wg.Wait()
	close(results)

	minLocation := math.MaxInt

	for result := range results {
		if result < minLocation {
			minLocation = result
		}
	}

	return minLocation
}

// Bounded concurrency - limit max number of go routines
func PartBV3(lines []string) int {
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

	var wg sync.WaitGroup

	numSeedPairs := len(seeds) / 2

	results := make(chan int, numSeedPairs)

	const maxRoutines = 6
	guard := make(chan int, maxRoutines)

	for idx := 0; idx < len(seeds); idx += 2 {
		guard <- idx
		wg.Add(1)

		startRange := seeds[idx]
		endRange := seeds[idx] + seeds[idx+1]

		go func(startSeed int, endSeed int, result chan int) {
			defer wg.Done()
			localMin := math.MaxInt
			start := time.Now()
			log.Printf("Seeds [%d, %d) starting", startSeed, endSeed)
			for seed := startSeed; seed < endSeed; seed += 1 {
				src := seed
				for _, catMap := range categoryMaps {
					dest := catMap.findDestination(src)
					src = dest
				}

				// log.Println("Location for seed", seed, src)

				if src < localMin {
					localMin = src
				}
			}
			log.Printf("Finished %d seeds in %s: min location %d", endSeed-startSeed, time.Since(start), localMin)
			result <- localMin
			<-guard
		}(startRange, endRange, results)

	}

	wg.Wait()
	close(results)

	minLocation := math.MaxInt

	for result := range results {
		if result < minLocation {
			minLocation = result
		}
	}

	return minLocation
}
