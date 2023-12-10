package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/slices"
)

var symbolMap = map[byte]string{
	'F': "┏",
	'7': "┓",
	'J': "┛",
	'|': "┃",
	'L': "┗",
	'-': "━",
	'.': "╳",
	'S': "S",
}

func printMap(m [][]Coord) {
	for _, row := range m {
		for _, c := range row {
			print(symbolMap[c.symbol])
		}
		println()
	}
}

type Coord struct {
	y      int
	x      int
	symbol byte
}

type CoordWithDist struct {
	coord Coord
	dist  int
}

func parseMap(filePath string) [][]Coord {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := [][]Coord{}
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		row := []Coord{}
		for x, symbol := range scanner.Text() {
			row = append(row, Coord{y: y, x: x, symbol: byte(symbol)})
		}
		y++
		m = append(m, row)
	}
	return m
}

var ConnectedSymbols = map[byte]map[byte][]byte{
	'S': map[byte][]byte{
		'N': []byte{'F', '7', '|'},
		'E': []byte{'7', 'J', '-'},
		'S': []byte{'J', 'L', '|'},
		'W': []byte{'L', 'F', '-'},
	},
	'F': map[byte][]byte{
		'E': []byte{'-', '7', 'J'},
		'S': []byte{'|', 'J', 'L'},
	},
	'7': map[byte][]byte{
		'S': []byte{'|', 'J', 'L'},
		'W': []byte{'-', 'L', 'F'},
	},
	'J': map[byte][]byte{
		'N': []byte{'F', '7', '|'},
		'W': []byte{'-', 'L', 'F'},
	},
	'L': map[byte][]byte{
		'N': []byte{'F', '7', '|'},
		'E': []byte{'7', 'J', '-'},
	},
	'|': map[byte][]byte{
		'N': []byte{'F', '7', '|'},
		'S': []byte{'|', 'J', 'L'},
	},
	'-': map[byte][]byte{
		'E': []byte{'-', '7', 'J'},
		'W': []byte{'-', 'L', 'F'},
	},
}

func getNeighbors(c Coord, m [][]Coord) []Coord {
	neighbors := []Coord{}
	possibleNeighbors := ConnectedSymbols[c.symbol]
	if c.y != 0 {
		north := m[c.y-1][c.x]
		if slices.Contains(possibleNeighbors['N'], north.symbol) {
			neighbors = append(neighbors, north)
		}
	}

	if c.x != 0 {
		west := m[c.y][c.x-1]
		if slices.Contains(possibleNeighbors['W'], west.symbol) {
			neighbors = append(neighbors, west)
		}
	}

	if c.y != len(m)-1 {
		south := m[c.y+1][c.x]
		if slices.Contains(possibleNeighbors['S'], south.symbol) {
			neighbors = append(neighbors, south)
		}
	}

	if c.x != len(m[0])-1 {
		east := m[c.y][c.x+1]
		if slices.Contains(possibleNeighbors['E'], east.symbol) {
			neighbors = append(neighbors, east)
		}
	}

	return neighbors
}

func findMaxDistFromStart(m [][]Coord) int {
	var start Coord
	for _, row := range m {
		for _, c := range row {
			if c.symbol == 'S' {
				start = c
			}
		}
	}
	queue := []CoordWithDist{
		CoordWithDist{start, 0},
	}

	visited := map[Coord]bool{
		start: true,
	}
	maxDist := 0
	for len(queue) > 0 {
		c := queue[0]
		fmt.Println(c)
		queue = queue[1:]
		n := getNeighbors(c.coord, m)
		for _, neighbor := range n {
			if !visited[neighbor] {
				queue = append(queue, CoordWithDist{neighbor, c.dist + 1})
				visited[neighbor] = true
				if c.dist+1 > maxDist {
					maxDist = c.dist + 1
				}
			}
		}
	}
	return maxDist
}

func Solution() {
	pipeMap := parseMap("p1/problem_input.txt")
	printMap(pipeMap)
	maxDist := findMaxDistFromStart(pipeMap)
	fmt.Println(maxDist)
}
