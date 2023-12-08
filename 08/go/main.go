package main

import (
	"flag"
	"log"
	"os"
	"regexp"
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

type Node struct {
	left  string
	right string
}

func PartA(lines []string) int {
	pattern := regexp.MustCompile(`^(?P<node>\w{3}) = \((?P<left>\w{3}), (?P<right>\w{3})\)$`)

	moves := strings.Split(lines[0], "")

	nodeMap := make(map[string]Node)

	for _, line := range lines[2:] {
		match := pattern.FindStringSubmatch(line)
		if match == nil {
			log.Panic("Line didn't match regex", line)
		}
		nodeMap[match[1]] = Node{left: match[2], right: match[3]}
	}

	current := "AAA"
	moveIdx := 0
	numMoves := 0
	for current != "ZZZ" {
		move := moves[moveIdx]
		if move == "L" {
			current = nodeMap[current].left
		} else if move == "R" {
			current = nodeMap[current].right
		} else {
			log.Panicf("Move must be L or R, found %s", move)
		}

		moveIdx = (moveIdx + 1) % len(moves)
		numMoves += 1
		// log.Printf("Current position after %d moves: %s", numMoves, current)
	}

	return numMoves
}

func PartB(lines []string) int {
	return len(lines)
}
