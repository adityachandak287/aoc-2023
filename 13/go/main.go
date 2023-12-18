package main

import (
	"flag"
	"log"
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
		answer = PartA(lines)
	case "B":
		answer = PartB(lines)
	default:
		log.Panic("Invalid input for part")
	}

	log.Println("Answer", answer)
}

type Grid = [][]byte

func parseGrids(lines []string) []Grid {
	var grids []Grid

	currentGrid := Grid{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" && len(currentGrid) > 0 {
			grids = append(grids, currentGrid)
			currentGrid = Grid{}
			continue
		}

		currentGrid = append(currentGrid, []byte(line))
	}

	if len(currentGrid) > 0 {
		grids = append(grids, currentGrid)
	}

	return grids
}

func checkGridVertical(grid Grid, colLeft int, colRight int) int {
	rows := len(grid)
	cols := len(grid[0])
	diff := 0
	for cl, cr := colLeft, colRight; cl >= 0 && cr < cols; cl, cr = cl-1, cr+1 {
		for rowIdx := 0; rowIdx < rows; rowIdx++ {
			if grid[rowIdx][cl] != grid[rowIdx][cr] {
				diff += 1
			}
		}
	}
	return diff
}

func checkGridHorizontal(grid Grid, rowUp int, rowDown int) int {
	rows := len(grid)
	cols := len(grid[0])
	diff := 0
	for ru, rd := rowUp, rowDown; ru >= 0 && rd < rows; ru, rd = ru-1, rd+1 {
		for colIdx := 0; colIdx < cols; colIdx++ {
			if grid[ru][colIdx] != grid[rd][colIdx] {
				diff += 1
			}
		}
	}
	return diff
}

func calcGridScore(grid Grid, target int) int {
	for colIdx := 0; colIdx < len(grid[0])-1; colIdx++ {
		if checkGridVertical(grid, colIdx, colIdx+1) == target {
			// log.Println("Found vertical on col ", colIdx)
			return colIdx + 1
		}
	}

	for rowIdx := 0; rowIdx < len(grid)-1; rowIdx++ {
		if checkGridHorizontal(grid, rowIdx, rowIdx+1) == target {
			// log.Println("Found horizontal on row ", rowIdx)
			return (rowIdx + 1) * 100
		}
	}

	return 0
}

// func printGrids(grids []Grid) {
// 	for gridIdx, grid := range grids {
// 		log.Println("GRID", gridIdx+1)
// 		for _, row := range grid {
// 			log.Println(string(row))
// 		}
// 		log.Println("========")
// 	}
// }

func PartA(lines []string) int {
	grids := parseGrids(lines)

	// printGrids(grids)

	sum := 0
	for _, grid := range grids {
		sum += calcGridScore(grid, 0)
	}

	return sum
}

func PartB(lines []string) int {
	grids := parseGrids(lines)

	// printGrids(grids)

	sum := 0
	for _, grid := range grids {
		sum += calcGridScore(grid, 1)
	}

	return sum
}
