package p2

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

func isXOfAKind(cardCounts map[int]int, x int, numJokers int) bool {
	for c, count := range cardCounts {
		if c == 1 {
			if count >= x {
				return true
			}
		} else {
			if count+numJokers >= x {
				return true
			}
		}
	}
	return false
}

func isFullHouse(cardCounts map[int]int, numJokers int) bool {
	hasThreeCount := false
	hasTwoCount := false
	jokerCount := numJokers
	for c, count := range cardCounts {
		if c == 1 {
			if jokerCount == 3 {
				hasThreeCount = true
				jokerCount -= 3
			} else if jokerCount == 2 {
				hasTwoCount = true
				jokerCount -= 2
			}
		} else {
			if count+jokerCount >= 3 {
				hasThreeCount = true
				jokerCount -= (3 - count)
			} else if count+jokerCount >= 2 {
				hasTwoCount = true
				jokerCount -= (2 - count)
			}
		}
	}
	return hasThreeCount && hasTwoCount
}

func isXPair(cardCounts map[int]int, x int, numJokers int) bool {
	pairCount := 0
	jokerCount := numJokers
	for c, count := range cardCounts {
		if c == 1 {
			if count == 2 {
				pairCount++
				jokerCount -= 2
			}
		} else {
			if count+jokerCount >= 2 {
				pairCount++
				jokerCount -= (2 - count)
			}
		}
	}
	return pairCount >= x
}

func getHandType(hand Hand) string {
	cardCounts := hand.cardCounts

	jokerCount := cardCounts[1]
	if isXOfAKind(cardCounts, 5, jokerCount) {
		return "five of a kind"
	}
	if isXOfAKind(cardCounts, 4, jokerCount) {
		return "four of a kind"
	}
	if isFullHouse(cardCounts, jokerCount) {
		return "full house"
	}
	if isXOfAKind(cardCounts, 3, jokerCount) {
		return "three of a kind"
	}
	if isXPair(cardCounts, 2, jokerCount) {
		return "two pair"
	}
	if isXPair(cardCounts, 1, jokerCount) {
		return "one pair"
	}
	return "high card"
}

func cardValue(card rune) int {
	switch card {
	case 'A':
		return 13
	case 'K':
		return 12
	case 'Q':
		return 11
	case 'J':
		return 1
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

type Hand struct {
	cards      []int
	cardStr    string
	bid        int
	cardCounts map[int]int
	handType   string
}

type HandList []Hand

func (h HandList) Len() int {
	return len(h)
}

func (h HandList) Less(i, j int) bool {
	handScoreI := handTypeScore[h[i].handType]
	handScoreJ := handTypeScore[h[j].handType]
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
	cardCounts := map[int]int{}
	for _, card := range cards {
		cardCounts[card]++
	}
	h := Hand{
		cards:      cards,
		bid:        bid,
		cardStr:    cardStr,
		cardCounts: cardCounts,
	}
	h.handType = getHandType(h)
	return h
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
	rankedHands := readInput("p2/problem_input.txt")
	fmt.Println("--- Part 2 ---")
	totalWinnings := 0
	for r, hand := range rankedHands {
		totalWinnings += hand.bid * (r + 1)
	}
	fmt.Println("Solution:", totalWinnings)
}
