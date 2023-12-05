package p2

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

func getSeedRanges(seedText string) []Range {
	seedsStr := seedsRegex.FindAllString(seedText, -1)
	seedRanges := []Range{}

	i := 0
	for i < len(seedsStr) {
		start, _ := strconv.Atoi(seedsStr[i])
		rangeLen, _ := strconv.Atoi(seedsStr[i+1])

		seedRanges = append(seedRanges, Range{start, start + rangeLen - 1})
		i += 2
	}

	return seedRanges
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

func getOverlap(range1 Range, range2 Range) Range {
	if range1.start > range2.start {
		return getOverlap(range2, range1)
	}

	if range1.end < range2.start {
		return Range{-1, -1}
	}

	if range1.end < range2.end {
		return Range{range2.start, range1.end}
	}
	return Range{range2.start, range2.end}
}

func getNestRanges(r Range, mapping RangeMapping) []Range {
	nextRanges := []Range{}
	initialRange := Range{r.start, r.end}
	remaining := []Range{initialRange}

	// for each range in remaining, find the partial dstRange from the srcRange overlaps
	i := 0
	for len(remaining) > 0 {
		pop := remaining[0]
		remaining = remaining[1:]

		// process overlapping ranges
		foundOverlap := false
		for srcRange, dstRange := range mapping {
			overlap := getOverlap(pop, srcRange)
			if overlap.start != -1 && overlap.end != -1 {
				foundOverlap = true
				nextRanges = append(nextRanges, Range{dstRange.start + (overlap.start - srcRange.start), dstRange.start + (overlap.end - srcRange.start)})
				if pop.start < overlap.start {
					remaining = append(remaining, Range{pop.start, overlap.start - 1})
				}
				if pop.end > overlap.end {
					remaining = append(remaining, Range{overlap.end + 1, pop.end})
				}
			}
		}
		if !foundOverlap {
			nextRanges = append(nextRanges, pop)
		}
		i += 1
		if i == 2 {
			break
		}
	}

	return nextRanges
}

func getMinLocation(seedRange Range, mapping map[string]RangeMapping) int {
	key := "seed"
	currRanges := []Range{seedRange}
	for key != "location" {
		nextRanges := []Range{}
		for _, r := range currRanges {
			nextRanges = append(nextRanges, getNestRanges(r, mapping[key])...)
		}
		currRanges = nextRanges
		key = entityMapping[key]
	}

	minLocation := math.MaxInt32
	for _, currRange := range currRanges {
		if currRange.start < minLocation {
			minLocation = currRange.start
		}
	}
	return minLocation
}

func readInput(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	seedRanges := getSeedRanges(scanner.Text())
	scanner.Scan()
	maps := setupMaps()
	populateMapping(scanner, maps)
	locationNumbers := []int{}
	for _, seedRange := range seedRanges {
		locationNumbers = append(locationNumbers, getMinLocation(seedRange, maps))
	}

	return locationNumbers
}

func Solution() {
	locations := readInput("p2/problem_input.txt")
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
