package p2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Node struct {
	value string
	left  *Node
	right *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%s = (%s, %s)", n.value, n.left.value, n.right.value)
}

var nodeRegex = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

func parseNode(nodeText string, nodeMap map[string]*Node) *Node {
	matches := nodeRegex.FindStringSubmatch(nodeText)
	nodeKey := matches[1]
	leftNodeKey := matches[2]
	rightNodeKey := matches[3]

	node := nodeMap[nodeKey]
	if node == nil {
		node = &Node{
			value: nodeKey,
		}
		nodeMap[nodeKey] = node
	}

	left := nodeMap[leftNodeKey]
	if left == nil {
		left = &Node{
			value: leftNodeKey,
		}
		nodeMap[leftNodeKey] = left
	}
	node.left = nodeMap[leftNodeKey]

	right := nodeMap[rightNodeKey]
	if right == nil {
		right = &Node{
			value: rightNodeKey,
		}
		nodeMap[rightNodeKey] = right
	}
	node.right = nodeMap[rightNodeKey]
	return node
}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a int, integers ...int) int {
	result := a
	for _, b := range integers {
		result = result * b / gcd(result, b)
	}
	return result
}

func countSteps(instructions string, startNode *Node) int {
	currNode := startNode
	steps := 0
	i := 0
	n := len(currNode.value)

	for currNode.value[n-1] != 'Z' {
		if i >= len(instructions) {
			i = 0
		}

		instruction := instructions[i]
		switch instruction {
		case 'L':
			currNode = currNode.left
		case 'R':
			currNode = currNode.right
		}
		steps++
		i++
	}
	return steps
}

func readInput(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := scanner.Text()
	fmt.Println(instructions)
	scanner.Scan()
	nodeMap := make(map[string]*Node)
	for scanner.Scan() {
		parseNode(scanner.Text(), nodeMap)
	}

	stepCount := []int{}
	startNodes := []*Node{}
	for nodeKey := range nodeMap {
		if nodeKey[len(nodeKey)-1] == 'A' {
			startNodes = append(startNodes, nodeMap[nodeKey])
		}
	}
	for _, node := range startNodes {
		steps := make(chan int)
		go func() {
			steps <- countSteps(instructions, node)
		}()
		stepCount = append(stepCount, <-steps)
	}
	return lcm(stepCount[0], stepCount[1:]...)
}

func Solution() {
	steps := readInput("p2/problem_input.txt")
	fmt.Println("steps", steps)
}
