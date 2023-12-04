package p2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var winningNumbersRegex = regexp.MustCompile(`(\d+(?:\s+\d+)*)\s*\|`)
var myNumbersRegex = regexp.MustCompile(`\|\s*(\d+(?:\s+\d+)*)\s*`)

func getNumbers(line string) (winningNumbers []int, myNumbers []int) {
	winningMatches := winningNumbersRegex.FindStringSubmatch(line)
	winningNumbersStr := strings.Fields(winningMatches[1])
	for _, numStr := range winningNumbersStr {
		num, _ := strconv.Atoi(numStr)
		winningNumbers = append(winningNumbers, num)
	}

	myMatches := myNumbersRegex.FindStringSubmatch(line)
	myNumbersStr := strings.Fields(myMatches[1])
	for _, numStr := range myNumbersStr {
		num, _ := strconv.Atoi(numStr)
		myNumbers = append(myNumbers, num)
	}
	return winningNumbers, myNumbers
}

func calcMatches(winningNumbers []int, myNumbers []int) int {
	winnNumberSet := make(map[int]bool)
	for _, num := range winningNumbers {
		winnNumberSet[num] = true
	}

	numMatches := 0
	for _, myNum := range myNumbers {
		if winnNumberSet[myNum] {
			numMatches += 1
		}
	}
	return numMatches
}

func calcCardsWon(lastCardNum int, cardMatches map[int]int) int {
	cardCopies := map[int]int{}
	// start with 1 copy of each card
	for i := 1; i <= lastCardNum; i++ {
		cardCopies[i] = 1
	}

	// calcualte total copies bottom up
	totalCards := 0
	for i := 1; i <= lastCardNum; i++ {
		numMatches := cardMatches[i]
		currentCardCopies := cardCopies[i]
		for k := i + 1; k < i+1+numMatches; k++ {
			cardCopies[k] += currentCardCopies
		}
		totalCards += cardCopies[i]
	}
	return totalCards
}

func readInput(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cardMatches := map[int]int{}
	c := 1
	for scanner.Scan() {
		winningNums, mynums := getNumbers(scanner.Text())
		cardMatches[c] = calcMatches(winningNums, mynums)
		c += 1
	}

	cardsWon := calcCardsWon(c-1, cardMatches)
	return cardsWon
}

func Solution() {
	cardsWon := readInput("p2/problem_input.txt")
	fmt.Println("----")
	fmt.Println(cardsWon)
}
