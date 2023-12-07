package main

import (
	"flag"
	"log"
	"os"
	"sort"
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

const (
	HighCard = iota + 1
	SinglePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	cards    []byte
	bid      int
	strength uint8
}

func PartA(lines []string) int {
	allCards := []byte("23456789TJQKA")

	allCardsMap := make(map[byte]byte)

	for i, c := range allCards {
		allCardsMap[c] = byte(i)
	}

	hands := make([]Hand, len(lines))

	for lineIdx, line := range lines {
		parts := strings.Fields(line)
		cards := parts[0]
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Panic(err)
		}

		hand := Hand{cards: []byte(cards), bid: bid, strength: HighCard}

		counts := make([]int, len(allCards))
		hasFive := false
		hasFour := false
		hasThree := false
		hasPairA := false
		hasPairB := false
		var pairA byte
		var pairB byte

		for _, card := range hand.cards {
			counts[allCardsMap[card]] += 1

			count := counts[allCardsMap[card]]

			if count == 5 {
				hasFive = true
			} else if count == 4 {
				hasFour = true
			} else if count == 3 {
				hasThree = true
			} else if count == 2 {
				if hasPairA {
					if card == pairA {
						log.Panic("Should be three of a kind!")
					}
					if !hasPairB {
						hasPairB = true
						pairB = card
					} else {
						log.Panic("3 pairs cannot exist!", line, string(card))
					}
				} else {
					hasPairA = true
					pairA = card
				}
			}
		}

		if hasFive {
			hand.strength = FiveOfAKind
		} else if hasFour {
			hand.strength = FourOfAKind
		} else if hasThree && hasPairA && hasPairB && pairA != pairB {
			hand.strength = FullHouse
		} else if hasThree && hasPairA && !hasPairB {
			hand.strength = ThreeOfAKind
		} else if hasPairA && hasPairB && pairA != pairB {
			hand.strength = TwoPair
		} else if hasPairA && !hasPairB {
			hand.strength = SinglePair
		} else {
			if hasFive || hasFour || hasThree || hasPairA || hasPairB {
				log.Panic("Hand is not high card!", cards)
			}
			hand.strength = HighCard
		}

		hands[lineIdx] = hand
	}

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].strength < hands[j].strength {
			return true
		} else if hands[i].strength > hands[j].strength {
			return false
		} else {
			for idx := 0; idx < len(hands[i].cards); idx++ {
				cardA := allCardsMap[hands[i].cards[idx]]
				cardB := allCardsMap[hands[j].cards[idx]]
				if cardA < cardB {
					return true
				} else if cardA > cardB {
					return false
				}
			}
			log.Panicf("Both hands are same! [%s] [%s]", string(hands[i].cards), string(hands[j].cards))
			return true
		}
	})

	// for _, h := range hands {
	// 	log.Println(string(h.cards), h.bid, h.strength)
	// }

	answer := 0
	for idx, hand := range hands {
		rank := idx + 1
		answer += rank * hand.bid
	}

	return answer
}

func PartB(lines []string) int {
	allCards := []byte("J23456789TQKA")

	allCardsMap := make(map[byte]byte)

	for i, c := range allCards {
		allCardsMap[c] = byte(i)
	}

	hands := make([]Hand, len(lines))

	for lineIdx, line := range lines {
		parts := strings.Fields(line)
		cards := parts[0]
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Panic(err)
		}

		hand := Hand{cards: []byte(cards), bid: bid, strength: HighCard}

		counts := make([]int, len(allCards))
		hasFive := false
		hasFour := false
		hasThree := false
		hasPairA := false
		hasPairB := false
		var pairA byte
		var pairB byte
		jokers := 0

		for _, card := range hand.cards {
			counts[allCardsMap[card]] += 1

			if card == 'J' {
				jokers += 1
			}

			count := counts[allCardsMap[card]]

			if count == 5 {
				hasFive = true
			} else if count == 4 {
				hasFour = true
			} else if count == 3 {
				hasThree = true
			} else if count == 2 {
				if hasPairA {
					if card == pairA {
						log.Panic("Should be three of a kind!")
					}
					if !hasPairB {
						hasPairB = true
						pairB = card
					} else {
						log.Panic("3 pairs cannot exist!", line, string(card))
					}
				} else {
					hasPairA = true
					pairA = card
				}
			}
		}

		if hasFive {
			hand.strength = FiveOfAKind
		} else if hasFour {
			hand.strength = FourOfAKind
		} else if hasThree && hasPairA && hasPairB && pairA != pairB {
			hand.strength = FullHouse
		} else if hasThree && hasPairA && !hasPairB {
			hand.strength = ThreeOfAKind
		} else if hasPairA && hasPairB && pairA != pairB {
			hand.strength = TwoPair
		} else if hasPairA && !hasPairB {
			hand.strength = SinglePair
		} else {
			if hasFive || hasFour || hasThree || hasPairA || hasPairB {
				log.Panic("Hand is not high card!", cards)
			}
			hand.strength = HighCard
		}

		if jokers > 0 {
			switch hand.strength {
			case FiveOfAKind:
				// already best hand
			case FourOfAKind:
				// AAAAK, joker can be A or K
				if jokers == 1 || jokers == 4 {
					hand.strength = FiveOfAKind
				} else {
					log.Panic("FourOfAKind can only have 1 or 4 jokers, found ", jokers)
				}
			case FullHouse:
				// AAAKK, joker can be A or K
				if jokers == 2 || jokers == 3 {
					hand.strength = FiveOfAKind
				} else {
					log.Panic("FullHouse can only have 2 or 3 jokers, found ", jokers)
				}
			case ThreeOfAKind:
				// 777AK, joker can be 7, A or K
				if jokers == 1 || jokers == 3 {
					hand.strength = FourOfAKind
				} else {
					log.Panic("ThreeOfAKind can only have 1 joker, found ", jokers)
				}
			case TwoPair:
				// AAKKT, joker can be A,K,T
				if jokers == 2 {
					hand.strength = FourOfAKind
				} else if jokers == 1 {
					hand.strength = FullHouse
				} else {
					log.Panic("TwoPair can only have 1 or 2 jokers, found ", jokers)
				}
			case SinglePair:
				// AA234, joker can be A, 2, 3, 4
				if jokers == 1 || jokers == 2 {
					hand.strength = ThreeOfAKind
				} else {
					log.Panic("SinglePair can only have 1 or 2 jokers, found ", jokers)
				}
			case HighCard:
				// A2345, joker can be any of the cards
				if jokers == 1 {
					hand.strength = SinglePair
				} else {
					log.Panic("HighCard can only have 1 joker, found ", jokers)
				}
			}
		}

		hands[lineIdx] = hand
	}

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].strength < hands[j].strength {
			return true
		} else if hands[i].strength > hands[j].strength {
			return false
		} else {
			for idx := 0; idx < len(hands[i].cards); idx++ {
				cardA := allCardsMap[hands[i].cards[idx]]
				cardB := allCardsMap[hands[j].cards[idx]]
				if cardA < cardB {
					return true
				} else if cardA > cardB {
					return false
				}
			}
			log.Panicf("Both hands are same! [%s] [%s]", string(hands[i].cards), string(hands[j].cards))
			return true
		}
	})

	// for idx, h := range hands {
	// 	log.Println(string(h.cards), h.bid, h.strength, idx+1)
	// }

	answer := 0
	for idx, hand := range hands {
		rank := idx + 1
		answer += rank * hand.bid
	}

	return answer
}
