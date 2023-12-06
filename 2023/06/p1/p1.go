package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var numsRgex = regexp.MustCompile(`\d+`)

func canBeatRecord(speed, timeRemaining, recordDist int) bool {
	return speed*timeRemaining > recordDist
}

func calcWaysToBeatRecord(time int, recordDist int) int {
	waysToBeatRecord := 0
	for releaseTime := 0; releaseTime < time; releaseTime++ {
		speed := releaseTime
		if canBeatRecord(speed, time-releaseTime, recordDist) {
			waysToBeatRecord++
		}
	}
	return waysToBeatRecord
}

func readInput(filePath string) []int {
	// Scan lines from file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	times := []int{}
	recordDists := []int{}

	scanner.Scan()
	for _, num := range numsRgex.FindAllString(scanner.Text(), -1) {
		time, _ := strconv.Atoi(num)
		times = append(times, time)
	}

	scanner.Scan()
	for _, num := range numsRgex.FindAllString(scanner.Text(), -1) {
		dist, _ := strconv.Atoi(num)
		recordDists = append(recordDists, dist)
	}
	fmt.Println("times", times)
	fmt.Println("recordDists", recordDists)

	waysToBeatRecord := []int{}
	for i := 0; i < len(times); i++ {
		time := times[i]
		recordDist := recordDists[i]
		waysToBeatRecord = append(waysToBeatRecord, calcWaysToBeatRecord(time, recordDist))
	}

	return waysToBeatRecord
}

func Solution() {
	waysToBeatRecord := readInput("p1/problem_input.txt")
	solution := 1
	for _, way := range waysToBeatRecord {
		solution *= way
	}
	fmt.Println("Solution:", solution)
}
