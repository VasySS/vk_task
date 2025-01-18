package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"vktask/dijkstra"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	firstLine := readArrInt(in)
	if len(firstLine) != 2 {
		log.Fatalf("invalid input data, expected 2 elements (n, m), got %d", len(firstLine))
	}

	n, m := firstLine[0], firstLine[1]

	maze := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		nums := readArrInt(in)
		if len(nums) != m {
			log.Fatalf("invalid input data, expected %d elements, got %d", m, len(nums))
		}

		maze = append(maze, nums)
	}

	lastLine := readArrInt(in)
	if len(lastLine) != 4 {
		log.Fatalf("invalid input data, expected 4 elements (start and end coordinates), got %d", len(lastLine))
	}

	start := dijkstra.Point{X: lastLine[0], Y: lastLine[1]}
	end := dijkstra.Point{X: lastLine[2], Y: lastLine[3]}

	distance, path := dijkstra.Run(maze, start, end)
	if distance == -1 {
		log.Fatalf("no path found")
	}

	for _, p := range path {
		fmt.Printf("%d %d\n", p.X, p.Y)
	}

	fmt.Println(".")
}

func readArrInt(in *bufio.Reader) []int {
	nums := readArrString(in)
	arr := make([]int, len(nums))

	for i, n := range nums {
		val, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("can't parse input data: %v", err)
		}

		arr[i] = val
	}

	return arr
}

func readArrString(in *bufio.Reader) []string {
	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatalf("can't read input data: %v", err)
	}

	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")

	strs := strings.Split(line, " ")

	return strs
}
