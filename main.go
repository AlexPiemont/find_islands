package main

import (
	"fmt"
)

type coordinate struct {
	X       int
	Y       int
	Visited bool
	Value   int
}

/*func processNodes(x int, y int, visited [][]int, queue [][]int) {
	result := [len(queue)][len(queue[0])]int
	visited[x][y] = 1
	if queue[x][y] == 1 {

	}

}*/

/*func removeIslands(input [][]int) [][]int{
	output := [][]int{
		input[0]
	}
	output[len(input)-1] = input[len(input)-1]
	for i := 1; i<len(input)
}*/

func main() {
	queue := [][]int{
		{0, 1, 1, 0, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{0, 0, 1, 0, 0},
		{1, 1, 0, 0, 1},
	}
	result := make([][]int, len(queue))
	for i := range result {
		result[i] = make([]int, len(queue[i]))
	}
	fmt.Println(result)
	fmt.Println(queue)
}
