package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

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
		firstDigit := 0
		lastDigit := 0
		fmt.Println(text)
		for _, c := range text {
			if unicode.IsDigit(c) {
				if firstDigit == 0 {
					firstDigit = int(c - '0')
				}
				lastDigit = int(c - '0')
			}
		}
		fmt.Println(firstDigit, lastDigit)

		values = append(values, firstDigit*10+lastDigit)
	}
	return values
}

func Solution() {
	values := readInput("p1/problem_input.txt")
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Println("----")
	fmt.Println(sum)
}
