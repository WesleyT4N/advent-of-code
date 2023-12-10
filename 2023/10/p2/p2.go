package p2

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
	'I': "I",
	'C': "C",
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

	m := [][]Coord{
		[]Coord{},
	}
	scanner := bufio.NewScanner(file)
	y := 0
	scanner.Scan()
	for i := 0; i < len(scanner.Text())+2; i++ {
		m[0] = append(m[0], Coord{y: y, x: i, symbol: '.'})
	}
	y++
	row := []Coord{
		Coord{y: y, x: 0, symbol: '.'},
	}
	for x, symbol := range scanner.Text() {
		row = append(row, Coord{y: y, x: x + 1, symbol: byte(symbol)})
	}
	row = append(row, Coord{y: y, x: len(scanner.Text()) + 1, symbol: '.'})
	m = append(m, row)
	y++
	for scanner.Scan() {
		row := []Coord{
			Coord{y: y, x: 0, symbol: '.'},
		}
		for x, symbol := range scanner.Text() {
			row = append(row, Coord{y: y, x: x + 1, symbol: byte(symbol)})
		}
		row = append(row, Coord{y: y, x: len(scanner.Text()) + 1, symbol: '.'})
		y++
		m = append(m, row)
	}
	row = []Coord{}
	for x := range m[0] {
		row = append(row, Coord{y: y, x: x, symbol: '.'})
	}
	return append(m, row)
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

func determineStartSymbol(start Coord, visited map[Coord]bool, m [][]Coord) byte {
	fmt.Println(start)
	if visited[m[start.y-1][start.x]] && visited[m[start.y][start.x-1]] {
		return 'J'
	}
	if visited[m[start.y-1][start.x]] && visited[m[start.y][start.x+1]] {
		return 'L'
	}
	if visited[m[start.y+1][start.x]] && visited[m[start.y][start.x+1]] {
		return 'F'
	}
	if visited[m[start.y+1][start.x]] && visited[m[start.y][start.x-1]] {
		return '7'
	}
	if visited[m[start.y-1][start.x]] && visited[m[start.y+1][start.x]] {
		return '|'
	}
	if visited[m[start.y][start.x-1]] && visited[m[start.y][start.x+1]] {
		return '-'
	}
	return 'S'

}

func findLoop(m [][]Coord) map[Coord]bool {
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

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		n := getNeighbors(c.coord, m)
		for _, neighbor := range n {
			if !visited[neighbor] {
				queue = append(queue, CoordWithDist{neighbor, c.dist + 1})
				visited[neighbor] = true
			}
		}
	}
	return visited
}

func numInsideLoopAtY(c Coord, m [][]Coord) int {
	y := c.y
	numInside := 0
	isInside := false
	var prevSymbol byte
	for x := c.x; x < len(m[y]); x++ {
		switch m[y][x].symbol {
		case '.':
			if isInside {
				numInside++
			}
			prevSymbol = '.'
			continue
		case '|':
			isInside = !isInside
			continue
		default:
			if prevSymbol == '.' {
				isInside = !isInside
				prevSymbol = m[y][x].symbol
			} else {
				if (prevSymbol == 'L' && m[y][x].symbol == 'J') || (prevSymbol == 'F' && m[y][x].symbol == '7') {
					isInside = !isInside
				}
				prevSymbol = '.'
			}
		}
	}
	return numInside
}

func findNunInsideLoop(m [][]Coord, loop map[Coord]bool) int {
	fmt.Println("finding num inside loop")
	var start Coord
	for _, row := range m {
		for _, c := range row {
			if !loop[c] {
				c.symbol = '.'
			}
			if c.symbol == 'S' {
				start = c
			}
		}
	}
	m[start.y][start.x].symbol = determineStartSymbol(start, loop, m)

	n := 0
	// use raycasting algorithm to find all points inside loop
	for y := range m {
		n += numInsideLoopAtY(m[y][0], m)
	}

	return n
}

func Solution() {
	pipeMap := parseMap("p2/test_input5.txt")
	fmt.Println("----")
	loop := findLoop(pipeMap)
	println(findNunInsideLoop(pipeMap, loop))
}
