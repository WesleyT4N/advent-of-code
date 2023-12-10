package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func allZero(sequence []int) bool {
	for _, num := range sequence {
		if num != 0 {
			return false
		}
	}
	return true
}

func getNextSequence(sequence []int) (nextSequence []int) {
	for i := 1; i < len(sequence); i++ {
		nextSequence = append(nextSequence, sequence[i]-sequence[i-1])
	}
	return nextSequence
}

func getSequenceDifferences(sequence []int) [][]int {
	sequences := [][]int{sequence}
	currentSequence := sequence
	for !allZero(currentSequence) {
		currentSequence = getNextSequence(currentSequence)
		sequences = append(sequences, currentSequence)
	}
	return sequences
}

func extrapolate(sequence []int) int {
	sequences := getSequenceDifferences(sequence)

	zeroSequence := sequences[len(sequences)-1]
	zeroSequence = append(zeroSequence, 0)

	for i := len(sequences) - 2; i >= 0; i-- {
		difference := sequences[i+1][len(sequences[i+1])-1]
		extrapolation := sequences[i][len(sequences[i])-1] + difference
		sequences[i] = append(sequences[i], extrapolation)
	}
	return sequences[0][len(sequences[0])-1]
}

func readInput(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sequences := [][]int{}
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), " ")
		sequence := []int{}
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			sequence = append(sequence, n)
		}
		sequences = append(sequences, sequence)
	}
	sum := 0
	for _, sequence := range sequences {
		sum += extrapolate(sequence)
	}
	return sum
}

func Solution() {
	sum := readInput("p1/problem_input.txt")
	fmt.Println(sum)
}
