package p2

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

func calcWaysToBeatRecord(totalTime int, recordDist int) int {
	waysToBeatRecord := 0
	for releaseTime := 0; releaseTime < totalTime; releaseTime++ {
		speed := releaseTime
		if canBeatRecord(speed, totalTime-releaseTime, recordDist) {
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
	timeStr := ""
	for _, num := range numsRgex.FindAllString(scanner.Text(), -1) {
		timeStr += num
	}
	time, _ := strconv.Atoi(timeStr)
	times = append(times, time)

	scanner.Scan()
	recordDistStr := ""
	for _, num := range numsRgex.FindAllString(scanner.Text(), -1) {
		recordDistStr += num
	}
	recordDist, _ := strconv.Atoi(recordDistStr)
	recordDists = append(recordDists, recordDist)

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
	waysToBeatRecord := readInput("p2/problem_input.txt")
	solution := 1
	for _, way := range waysToBeatRecord {
		solution *= way
	}
	fmt.Println("Solution:", solution)
}
