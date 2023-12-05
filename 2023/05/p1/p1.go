package p1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var seedsRegex = regexp.MustCompile(`\d+`)

func getSeeds(seedText string) []int {
	seedsStr := seedsRegex.FindAllString(seedText, -1)
	seeds := []int{}
	for _, seedStr := range seedsStr {
		seed, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, seed)
	}
	return seeds
}

var entityMapping = map[string]string{
	"seed":        "soil",
	"soil":        "fertilizer",
	"fertilizer":  "water",
	"water":       "light",
	"light":       "temperature",
	"temperature": "humidity",
	"humidity":    "location",
	"location":    "",
}

type RangeMapping map[Range]Range

func setupMaps() map[string]RangeMapping {
	mapping := make(map[string]RangeMapping)
	for key := range entityMapping {
		mapping[key] = make(RangeMapping)
	}
	return mapping
}

type Range struct {
	start int
	end   int
}

func updateMapping(line string, mapping RangeMapping) {
	splitLine := strings.Fields(line)
	srcStart, _ := strconv.Atoi(splitLine[1])
	dstStart, _ := strconv.Atoi(splitLine[0])
	rangeLen, _ := strconv.Atoi(splitLine[2])
	startRange := Range{srcStart, srcStart + rangeLen - 1}
	dstRange := Range{dstStart, dstStart + rangeLen - 1}
	mapping[startRange] = dstRange
}

var titleRegex = regexp.MustCompile(`^(\w+)-`)

func getMapTitle(rawMapping []string) string {
	if len(rawMapping) == 0 {
		return ""
	}

	title := titleRegex.FindStringSubmatch(rawMapping[0])
	return title[1]
}

func populateMapping(scanner *bufio.Scanner, mapping map[string]RangeMapping) {
	rawMapping := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			title := getMapTitle(rawMapping)
			if mapping[title] != nil {
				for _, line := range rawMapping[1:] {
					updateMapping(line, mapping[title])
				}
			}
			rawMapping = []string{}
		} else {
			rawMapping = append(rawMapping, line)
		}
	}
	title := getMapTitle(rawMapping)
	if mapping[title] != nil {
		for _, line := range rawMapping[1:] {
			updateMapping(line, mapping[title])
		}
	}
}

func getNextId(id int, mapping RangeMapping) int {
	for srcRange, dstRange := range mapping {
		if id >= srcRange.start && id <= srcRange.end {
			distFromStart := id - srcRange.start
			return dstRange.start + distFromStart
		}
	}
	return id
}

func getLocation(seed int, mapping map[string]RangeMapping) int {
	key := "seed"
	currId := seed
	for key != "location" {
		currId = getNextId(currId, mapping[key])
		key = entityMapping[key]
	}

	return currId
}

func readInput(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	seeds := getSeeds(scanner.Text())
	scanner.Scan()
	maps := setupMaps()

	populateMapping(scanner, maps)
	locationNumbers := []int{}
	for _, seed := range seeds {
		locationNumbers = append(locationNumbers, getLocation(seed, maps))
	}

	return locationNumbers
}

func Solution() {
	locations := readInput("p1/problem_input.txt")
	fmt.Println("----")
	fmt.Println("locations", locations)
	minLocation := math.MaxInt32
	for _, location := range locations {
		if location < minLocation {
			minLocation = location
		}
	}
	fmt.Println("minLocation", minLocation)
}
