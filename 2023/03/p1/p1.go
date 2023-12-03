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

func isSymbol(char string) bool {
	regex := regexp.MustCompile("[^0-9.]")
	return regex.MatchString(char)
}

func hasNeighborSymbol(window [3]string, index []int) bool {
	var l int
	if index[0]-1 >= 0 {
		l = index[0] - 1
	} else {
		l = 0
	}
	var r int
	if index[1] < len(window[1]) {
		r = index[1]
	} else {
		r = len(window[1]) - 1
	}
	fmt.Println(l, r)

	for j := 0; j < 3; j++ {
		for i := l; i <= r; i++ {
			if isSymbol(string(window[j][i])) {
				fmt.Println("found symbol", string(window[j][i]))
				return true
			}
		}
	}
	return false
}

var numberRegex = regexp.MustCompile("[0-9]+")

func validPartNumbers(window [3]string) []int {
	fmt.Println("----")
	numberIndexes := numberRegex.FindAllIndex([]byte(window[1]), -1)
	fmt.Println(numberIndexes)
	validPartNumbers := []int{}
	for _, index := range numberIndexes {
		if hasNeighborSymbol(window, index) {
			partNumber, _ := strconv.Atoi(window[1][index[0]:index[1]])
			fmt.Println(partNumber)
			validPartNumbers = append(validPartNumbers, partNumber)
		}
	}
	return validPartNumbers
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

func nextWindow(window [3]string, nextLine string) [3]string {
	window[0] = window[1]
	window[1] = window[2]
	window[2] = nextLine
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

	partNumbers := []int{}
	for scanner.Scan() {
		partNumbers = append(partNumbers, validPartNumbers(window)...)
		window = nextWindow(window, scanner.Text())
	}

	partNumbers = append(partNumbers, validPartNumbers(window)...)

	emptyLine := strings.Repeat(".", len(window[1]))
	window = nextWindow(window, emptyLine)

	partNumbers = append(partNumbers, validPartNumbers(window)...)

	return partNumbers
}

func Solution() {
	partNumbers := readInput("p1/problem_input.txt")
	sum := 0
	for _, num := range partNumbers {
		fmt.Println(num)
		sum += num
	}
	fmt.Println("----")
	fmt.Println(sum)
}
