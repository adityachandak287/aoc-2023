package main

import (
	"flag"
	"log"
	"math"
	"os"
	"regexp"
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

func PartA(lines []string) int {
	linePattern := regexp.MustCompile(`:\s+| \| `)

	score := 0

	for _, line := range lines {
		parts := linePattern.Split(line, 3)
		winnerStr := parts[1]
		numberStr := parts[2]

		var winners []int
		var numbers []int

		for _, winnerStr := range strings.Fields(winnerStr) {
			num, err := strconv.Atoi(winnerStr)
			if err != nil {
				panic(err)
			}
			winners = append(winners, num)
		}
		for _, numberStr := range strings.Fields(numberStr) {
			num, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, num)
		}

		matches := 0
		for _, winner := range winners {
			for _, number := range numbers {
				if winner == number {
					matches++
				}
			}
		}

		if matches != 0 {
			score += int(math.Pow(2, float64(matches-1)))
		}

	}

	return score
}

func PartB(lines []string) int {
	linePattern := regexp.MustCompile(`:\s+| \| `)

	type CardScore struct {
		card  string
		score int
	}

	var scores []CardScore

	for _, line := range lines {
		parts := linePattern.Split(line, 3)
		cardId := parts[0]
		winnerStr := parts[1]
		numberStr := parts[2]

		var winners []int
		var numbers []int

		for _, winnerStr := range strings.Fields(winnerStr) {
			num, err := strconv.Atoi(winnerStr)
			if err != nil {
				panic(err)
			}
			winners = append(winners, num)
		}
		for _, numberStr := range strings.Fields(numberStr) {
			num, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, num)
		}

		matches := 0
		for _, winner := range winners {
			for _, number := range numbers {
				if winner == number {
					matches++
				}
			}
		}

		scores = append(scores, CardScore{card: cardId, score: matches})
	}

	cardCount := make(map[string]int)

	for cardIdx, score := range scores {
		numCards := 1 + cardCount[score.card]

		for idx := cardIdx + 1; idx < cardIdx+1+score.score; idx++ {
			cardCount[scores[idx].card] = cardCount[scores[idx].card] + numCards
		}

		cardCount[score.card] = numCards
	}

	score := 0

	for _, numCards := range cardCount {
		score += numCards
	}

	return score
}