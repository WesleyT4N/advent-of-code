package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Color string

const (
	RED   Color = "red"
	GREEN Color = "green"
	BLUE  Color = "blue"
)

var maxOfEachColor = map[Color]int{
	RED:   12,
	GREEN: 13,
	BLUE:  14,
}

func parseColor(color string) (int, Color) {
	vals := strings.Split(strings.Trim(color, " "), " ")
	num, _ := strconv.Atoi(vals[0])
	return num, Color(vals[1])
}

func parseGameID(text string) (int, bool) {
	startOfReveals := strings.Index(text, ":")
	gameId, _ := strconv.Atoi(text[5:startOfReveals])
	reveals := text[startOfReveals+1:]

	sets := strings.Split(reveals, ";")

	for _, set := range sets {
		colors := strings.Split(set, ",")
		for _, color := range colors {
			num, color := parseColor(color)
			if num > maxOfEachColor[color] {
				return gameId, false
			}
		}
	}
	return gameId, true
}

func readInput(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validGames := []int{}
	for scanner.Scan() {
		text := scanner.Text()
		gameId, isValid := parseGameID(text)
		fmt.Println(gameId, isValid)
		if isValid {
			validGames = append(validGames, gameId)
		}
	}
	return validGames
}

func Solution() {
	validGames := readInput("p1/problem_input.txt")
	sum := 0
	for _, gameID := range validGames {
		sum += gameID
	}
	fmt.Println("----")
	fmt.Println(sum)
}
