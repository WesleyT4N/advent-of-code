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

func isGear(char string) bool {
	return char == "*"
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isNeighboringNumber(numberStart, numberEnd, gearInd int) bool {
	return Abs(gearInd-numberEnd) <= 1 || Abs(numberStart-gearInd) <= 1 || (numberStart <= gearInd && gearInd <= numberEnd)
}

func getGearRatio(window [3]string, gearIndex []int) int {
	numberRegex := regexp.MustCompile("[0-9]+")
	neighboringNumbers := []int{}
	for j := 0; j < 3; j++ {
		numberIndexes := numberRegex.FindAllIndex([]byte(window[j]), -1)
		for _, numIndex := range numberIndexes {
			if isNeighboringNumber(numIndex[0], numIndex[1]-1, gearIndex[0]) {
				partNumber, _ := strconv.Atoi(window[j][numIndex[0]:numIndex[1]])
				neighboringNumbers = append(neighboringNumbers, partNumber)
			}
		}
		if len(neighboringNumbers) > 2 {
			return 0
		}
	}

	fmt.Printf("neighbors: %v\n", neighboringNumbers)
	gearRatio := 0
	if len(neighboringNumbers) == 2 {
		gearRatio = neighboringNumbers[0] * neighboringNumbers[1]
	}
	return gearRatio
}

var gearRegex = regexp.MustCompile("[*]")

func getGearRatios(window [3]string) []int {
	fmt.Println("----")
	gearRatios := []int{}
	gearIndexes := gearRegex.FindAllIndex([]byte(window[1]), -1)
	fmt.Printf("gearIndexes: %v\n", gearIndexes)
	for _, index := range gearIndexes {
		gearRatios = append(gearRatios, getGearRatio(window, index))
	}

	fmt.Printf("gearRatios: %v\n", gearRatios)
	return gearRatios
}

func nextWindow(window [3]string, nextLine string) [3]string {
	window[0] = window[1]
	window[1] = window[2]
	window[2] = nextLine
	return window
}

func initWindow(scanner *bufio.Scanner) [3]string {
	window := [3]string{}
	scanner.Scan()
	window[1] = scanner.Text()
	scanner.Scan()
	window[2] = scanner.Text()
	emptyLine := strings.Repeat(".", len(window[1]))
	window[0] = emptyLine
	return window
}

func readInput(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	window := initWindow(scanner)

	gearRatios := []int{}
	for scanner.Scan() {
		gearRatios = append(gearRatios, getGearRatios(window)...)
		window = nextWindow(window, scanner.Text())
	}

	gearRatios = append(gearRatios, getGearRatios(window)...)

	emptyLine := strings.Repeat(".", len(window[1]))
	window = nextWindow(window, emptyLine)

	gearRatios = append(gearRatios, getGearRatios(window)...)

	return gearRatios
}

func Solution() {
	gearRatios := readInput("p2/problem_input.txt")
	sum := 0
	for _, num := range gearRatios {
		sum += num
	}
	fmt.Println("----")
	fmt.Println(sum)
}
