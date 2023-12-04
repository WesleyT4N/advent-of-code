package p1

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

func calcPoints(winningNumbers []int, myNumbers []int) int {
	winnNumberSet := make(map[int]bool)
	for _, num := range winningNumbers {
		winnNumberSet[num] = true
	}

	num_matches := 0
	for _, myNum := range myNumbers {
		if winnNumberSet[myNum] {
			num_matches += 1
		}
	}

	if num_matches > 0 {
		points := 1
		for i := 1; i < num_matches; i++ {
			points *= 2
		}
		return points
	}
	return 0
}

func readInput(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := []int{}
	for scanner.Scan() {
		winningNums, mynums := getNumbers(scanner.Text())
		points = append(points, calcPoints(winningNums, mynums))
	}
	return points
}

func Solution() {
	points := readInput("p1/problem_input.txt")
	sum := 0
	for _, num := range points {
		sum += num
	}
	fmt.Println("----")
	fmt.Println(sum)
}
