package main

import (
	"flag"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func main() {
	inputFile := flag.String("input", "input", "Input file relative path")
	part := flag.String("part", "A", "Implementation of part A or B of the problem")
	checkpoint := flag.Int("checkpoint", int(1e6), "Checkpoint for reporting part B solution")
	known := flag.Bool("known", false, "Part B will run with known progress length (by first solving using PartBV2)")

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
		answer = PartB(lines, *checkpoint, *known)
	case "BV2":
		answer = PartBV2(lines)
	default:
		log.Panic("Invalid input for part")
	}

	log.Println("Answer", answer)
}

type Node struct {
	left  string
	right string
}

var pattern = regexp.MustCompile(`^(?P<node>\w{3}) = \((?P<left>\w{3}), (?P<right>\w{3})\)$`)

func PartA(lines []string) int {
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

// NOTE: Will take ~800 hours, lol
func PartB(lines []string, checkpoint int, known bool) int {
	moves := strings.Split(lines[0], "")

	nodeMap := make(map[string]Node)

	var positions []string

	for _, line := range lines[2:] {
		match := pattern.FindStringSubmatch(line)
		if match == nil {
			log.Panic("Line didn't match regex", line)
		}
		nodeMap[match[1]] = Node{left: match[2], right: match[3]}
		if strings.HasSuffix(match[1], "A") {
			positions = append(positions, match[1])
		}
	}

	moveIdx := 0
	numMoves := 0
	numZ := 0

	var progress *progressbar.ProgressBar
	if !known {
		progress = progressbar.Default(-1)
	} else {
		progress = progressbar.Default(int64(PartBV2(lines)))
	}
	progress.Describe("Calculating moves")

	for numZ != len(positions) {
		move := moves[moveIdx]
		moveIdx = (moveIdx + 1) % len(moves)

		numZ = 0
		for posIdx := range positions {
			if move == "L" {
				positions[posIdx] = nodeMap[positions[posIdx]].left
			} else if move == "R" {
				positions[posIdx] = nodeMap[positions[posIdx]].right
			} else {
				log.Panicf("Move must be L or R, found %s", move)
			}

			if strings.HasSuffix(positions[posIdx], "Z") {
				numZ += 1
			}
		}

		numMoves += 1

		if numMoves%checkpoint == 0 {
			progress.Add(checkpoint)
		}
	}

	return numMoves
}

func PartBV2(lines []string) int {
	moves := strings.Split(lines[0], "")

	nodeMap := make(map[string]Node)

	var positions []string

	for _, line := range lines[2:] {
		match := pattern.FindStringSubmatch(line)
		if match == nil {
			log.Panic("Line didn't match regex", line)
		}
		nodeMap[match[1]] = Node{left: match[2], right: match[3]}
		if strings.HasSuffix(match[1], "A") {
			positions = append(positions, match[1])
		}
	}

	movesRequired := make([]int, len(positions))

	for idx, pos := range positions {
		current := pos
		moveIdx := 0
		numMoves := 0
		for !strings.HasSuffix(current, "Z") {
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
		}
		// log.Printf("%s takes %d moves", pos, numMoves)
		movesRequired[idx] = numMoves
	}

	numMoves := arrayLcm(movesRequired)

	return numMoves
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func arrayLcm(integers []int) int {
	if len(integers) < 2 {
		log.Panic("Not enough starting positions")
		return 0
	}

	result := integers[0]

	for _, num := range integers {
		result = LCM(result, num)
	}

	return result
}
