package main

import (
	"flag"
	"log"
	"math"
	"os"
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
		answer = PartAB(lines, 2) // expand by 2
	case "B":
		answer = PartAB(lines, 1e6) // expand by 1 million
	default:
		log.Panic("Invalid input for part")
	}

	log.Println("Answer", answer)
}

type Position struct {
	x int
	y int
}

func PartAB(lines []string, expansionFactor int) int {
	emptyRows := make([]bool, len(lines))
	emptyCols := make([]bool, len(lines[0]))

	for row := range emptyRows {
		emptyRows[row] = true
	}
	for col := range emptyCols {
		emptyCols[col] = true
	}

	var galaxies []Position

	for rowIdx, row := range lines {
		for colIdx, col := range strings.Split(row, "") {
			emptyRows[rowIdx] = emptyRows[rowIdx] && col == "."
			emptyCols[colIdx] = emptyCols[colIdx] && col == "."

			if col == "#" {
				galaxies = append(galaxies, Position{x: colIdx, y: rowIdx})
			}
		}
	}

	rowOffset := make([]int, len(emptyRows))
	colOffset := make([]int, len(emptyCols))
	for row := range emptyRows {
		prev := 0
		if row > 0 {
			prev = rowOffset[row-1]
		}

		curr := 0
		if emptyRows[row] {
			curr = expansionFactor - 1
		}
		rowOffset[row] = prev + curr
	}

	for col := range emptyCols {
		prev := 0
		if col > 0 {
			prev = colOffset[col-1]
		}

		curr := 0
		if emptyCols[col] {
			curr = expansionFactor - 1
		}
		colOffset[col] = prev + curr
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := 0; j < i; j++ {
			g1 := galaxies[i]
			g2 := galaxies[j]

			// Manhattan Distance |x1-x2| + |y1-y2|
			dist := int(math.Abs(float64(g1.x+colOffset[g1.x])-float64(g2.x+colOffset[g2.x]))) + int(math.Abs(float64(g1.y+rowOffset[g1.y])-float64(g2.y+rowOffset[g2.y])))
			// log.Println("dist", g1, g2, dist)
			sum += dist
		}
	}

	return sum
}
