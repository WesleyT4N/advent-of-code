package p2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var nummbersAsStrings = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func isGreaterThanZero(c int) bool {
	return int(c-'0') > 0
}

func readInput(inputFileName string) []int {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	values := []int{}
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)

		var numsAtIndexes []int
		for range text {
			numsAtIndexes = append(numsAtIndexes, 0)
		}

		for i, numString := range nummbersAsStrings {
			num := i + 1

			index := strings.Index(text, numString)
			lastIndex := strings.LastIndex(text, numString)
			if index != -1 {
				numsAtIndexes[index] = num
				numsAtIndexes[lastIndex] = num
			}
			numIndex := strings.Index(text, fmt.Sprint(num))
			numLastIndex := strings.LastIndex(text, fmt.Sprint(num))
			if numIndex != -1 {
				numsAtIndexes[numIndex] = num
				numsAtIndexes[numLastIndex] = num
			}
		}

		// get first digit & last digit
		firstDigit := 0
		lastDigit := 0
		for _, n := range numsAtIndexes {
			if n > 0 {
				if firstDigit == 0 {
					firstDigit = n
				}
				lastDigit = n
			}
		}
		fmt.Println(firstDigit, lastDigit)

		values = append(values, firstDigit*10+lastDigit)
	}
	return values
}

func Solution() {
	values := readInput("p2/problem_input.txt")
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Println("----")
	fmt.Println(sum)
}
