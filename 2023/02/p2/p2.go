package p2

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

func parseColor(color string) (int, Color) {
	vals := strings.Split(strings.Trim(color, " "), " ")
	num, _ := strconv.Atoi(vals[0])
	return num, Color(vals[1])
}

func parseGameID(text string) (int, int) {
	startOfReveals := strings.Index(text, ":")
	gameId, _ := strconv.Atoi(text[5:startOfReveals])
	reveals := text[startOfReveals+1:]

	sets := strings.Split(reveals, ";")

	numRequiredPerColor := map[Color]int{}
	for _, set := range sets {
		colors := strings.Split(set, ",")
		for _, color := range colors {
			num, color := parseColor(color)
			if numRequiredPerColor[color] == 0 || num > numRequiredPerColor[color] {
				numRequiredPerColor[color] = num
			}
		}
	}

	power := 1
	for _, numRequired := range numRequiredPerColor {
		power *= numRequired
	}
	return gameId, power
}

func readInput(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	powersPerGame := []int{}
	for scanner.Scan() {
		text := scanner.Text()
		gameId, power := parseGameID(text)
		fmt.Println(gameId, power)
		powersPerGame = append(powersPerGame, power)
	}
	return powersPerGame
}

func Solution() {
	powers := readInput("p2/problem_input.txt")
	sum := 0
	for _, power := range powers {
		sum += power
	}
	fmt.Println("----")
	fmt.Println(sum)
}
