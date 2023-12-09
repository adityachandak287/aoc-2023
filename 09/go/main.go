package main

import (
	"flag"
	"log"
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

func PartA(lines []string) int {
	sum := 0
	for _, line := range lines {
		nums := parseInts(line)

		var series [][]int
		series = append(series, nums)
		found := false

		seriesIdx := 0

		for !found {
			s := series[seriesIdx]
			if len(s) < 2 {
				log.Panicf("Series must have atleast 2 elements, found %v", s)
			}

			diff := 0
			zeroDiffs := 0
			newSeries := make([]int, len(s)-1)

			for idx := 1; idx < len(s); idx++ {
				diff = s[idx] - s[idx-1]
				newSeries[idx-1] = diff
				if diff == 0 {
					zeroDiffs += 1
				}
			}

			series = append(series, newSeries)

			if zeroDiffs == len(newSeries) {
				found = true
			}

			seriesIdx += 1
		}

		series[len(series)-1] = append(series[len(series)-1], 0)
		for sIdx := len(series) - 2; sIdx >= 0; sIdx-- {
			currSeries := series[sIdx]
			nextSeries := series[sIdx+1]
			next := currSeries[len(currSeries)-1] + nextSeries[len(nextSeries)-1]
			series[sIdx] = append(series[sIdx], next)
		}

		sum += series[0][len(series[0])-1]
	}
	return sum
}

func PartB(lines []string) int {
	sum := 0
	for _, line := range lines {
		nums := parseInts(line)

		var series [][]int
		series = append(series, nums)
		found := false

		seriesIdx := 0

		for !found {
			s := series[seriesIdx]
			if len(s) < 2 {
				log.Panicf("Series must have atleast 2 elements, found %v", s)
			}

			diff := 0
			zeroDiffs := 0
			newSeries := make([]int, len(s)-1)

			for idx := 1; idx < len(s); idx++ {
				diff = s[idx] - s[idx-1]
				newSeries[idx-1] = diff
				if diff == 0 {
					zeroDiffs += 1
				}
			}

			series = append(series, newSeries)

			if zeroDiffs == len(newSeries) {
				found = true
			}

			seriesIdx += 1
		}

		series[len(series)-1] = append([]int{0}, series[len(series)-1]...)
		for sIdx := len(series) - 2; sIdx >= 0; sIdx-- {
			currSeries := series[sIdx]
			nextSeries := series[sIdx+1]
			prev := currSeries[0] - nextSeries[0]
			series[sIdx] = append([]int{prev}, series[sIdx]...)
		}

		sum += series[0][0]
	}
	return sum
}
