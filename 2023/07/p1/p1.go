package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var handTypeScore = map[string]int{
	"five of a kind":  7,
	"four of a kind":  6,
	"full house":      5,
	"three of a kind": 4,
	"two pair":        3,
	"one pair":        2,
	"high card":       1,
	"":                0,
}

func isXOfAKind(cardCounts map[int]int, x int) bool {
	for _, count := range cardCounts {
		if count == x {
			return true
		}
	}
	return false
}

func isFullHouse(cardCounts map[int]int) bool {
	hasThreeCount := false
	hasTwoCount := false
	for _, count := range cardCounts {
		if count == 3 {
			hasThreeCount = true
		}
		if count == 2 {
			hasTwoCount = true
		}
	}
	return hasThreeCount && hasTwoCount
}

func isXPair(cardCounts map[int]int, x int) bool {
	pairCount := 0
	for _, count := range cardCounts {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount == x
}

func getHandType(hand Hand) string {
	cardCounts := map[int]int{}
	for _, card := range hand.cards {
		cardCounts[card]++
	}

	if isXOfAKind(cardCounts, 5) {
		return "five of a kind"
	}
	if isXOfAKind(cardCounts, 4) {
		return "four of a kind"
	}
	if isFullHouse(cardCounts) {
		return "full house"
	}
	if isXOfAKind(cardCounts, 3) {
		return "three of a kind"
	}
	if isXPair(cardCounts, 2) {
		return "two pair"
	}
	if isXPair(cardCounts, 1) {
		return "one pair"
	}
	return "high card"
}

func cardValue(card rune) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

type Hand struct {
	cards   []int
	cardStr string
	bid     int
}

type HandList []Hand

func (h HandList) Len() int {
	return len(h)
}

func (h HandList) Less(i, j int) bool {
	handTypeI := getHandType(h[i])
	handTypeJ := getHandType(h[j])

	handScoreI := handTypeScore[handTypeI]
	handScoreJ := handTypeScore[handTypeJ]
	if handScoreI != handScoreJ {
		return handScoreI < handScoreJ
	}

	for k := 0; k < len(h[i].cards); k++ {
		if h[i].cards[k] != h[j].cards[k] {
			return h[i].cards[k] < h[j].cards[k]
		}
	}
	return false
}

func (h HandList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func parseHand(hand string) Hand {
	s := strings.Fields(hand)
	cardStr := s[0]
	bid, _ := strconv.Atoi(s[1])

	cards := []int{}
	for _, c := range cardStr {
		cards = append(cards, cardValue(c))
	}
	return Hand{
		cards:   cards,
		bid:     bid,
		cardStr: cardStr,
	}
}

func readInput(filePath string) HandList {
	// Scan lines from file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := HandList{}
	for scanner.Scan() {
		hand := parseHand(scanner.Text())
		hands = append(hands, hand)
	}
	sort.Sort(hands)
	return hands
}

func Solution() {
	rankedHands := readInput("p1/problem_input.txt")
	fmt.Println("--- Part 1 ---")
	totalWinnings := 0
	for r, hand := range rankedHands {
		totalWinnings += hand.bid * (r + 1)
	}
	fmt.Println("Solution:", totalWinnings)
}
