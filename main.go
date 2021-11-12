package main

import (
	"fmt"
)

type imgMap struct {
	Inmap   [][]int
	Visited [][]int
	Outmap  [][]int
}

func makeImgMap(input [][]int) imgMap {
	ret := imgMap{}
	ret.Inmap = input
	ret.Outmap = make([][]int, len(input))
	for i := range ret.Outmap {
		ret.Outmap[i] = make([]int, len(input[i]))
	}
	ret.Visited = make([][]int, len(input))
	for i := range ret.Visited {
		ret.Visited[i] = make([]int, len(input[i]))
	}
	return ret
}
func (i *imgMap) processNode(x, y int) {
	fmt.Printf("currently at position %d, %d \n", x, y)
	i.Visited[x][y] = 1
	fmt.Printf("visited: %d", i.Visited)
	//fmt.Print("Press 'Enter' to continue...")
	//bufio.NewReader(os.Stdin).ReadBytes('\n')
	dirs := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	if i.Inmap[x][y] == 1 || x == 0 || y == 0 || x == len(i.Inmap[0])-1 || y == len(i.Inmap)-1 {
		i.Outmap[x][y] = i.Inmap[x][y]
		for j := 0; j < len(dirs); j++ {
			if x+dirs[j][0] >= 0 && y+dirs[j][1] >= 0 && x+dirs[j][0] < len(i.Inmap[0]) && y+dirs[j][1] < len(i.Inmap) {
				if i.Visited[x+dirs[j][0]][y+dirs[j][1]] == 0 {
					i.processNode(x+dirs[j][0], y+dirs[j][1])
				}
			}
		}
	}
}
func removeIslands(input [][]int) [][]int {
	mp := makeImgMap(input)
	mp.processNode(0, 0)
	return mp.Outmap
}
func main() {
	queue := [][]int{
		{0, 1, 1, 0, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{0, 0, 1, 0, 0},
		{1, 1, 0, 0, 1},
	}
	mp := makeImgMap(queue)
	fmt.Println(mp.Visited)
	fmt.Println(mp.Outmap)
	mp.processNode(0, 0)
	fmt.Println(mp.Inmap)
	fmt.Println(mp.Visited)
	fmt.Println(mp.Outmap)

}
