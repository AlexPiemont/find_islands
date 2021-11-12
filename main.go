package main

import (
	"fmt"
	"strconv"
)

type direction struct {
	Right bool
	Up    bool
	Left  bool
	Down  bool
}

type imgMap struct {
	Inmap   [][]int
	Visited [][]direction
	Outmap  [][]int
}

func (mp *imgMap) toStringCompare() string {
	resStr := "   Input           Output \n"

	for i := 0; i < len(mp.Inmap); i++ {
		inStr := ""
		visStr := ""
		outStr := ""
		for j := 0; j < len(mp.Inmap[i]); j++ {
			inStr = inStr + " " + strconv.Itoa(mp.Inmap[i][j])
			outStr = outStr + " " + strconv.Itoa(mp.Outmap[i][j])
		}
		resStr = resStr + inStr + "    " + visStr + "    " + outStr + "\n"
	}
	return resStr
}

func makeImgMap(input [][]int) imgMap {
	ret := imgMap{}
	ret.Inmap = input
	ret.Outmap = make([][]int, len(input))
	for i := range ret.Outmap {
		ret.Outmap[i] = make([]int, len(input[i]))
	}
	ret.Visited = make([][]direction, len(input))
	for i := range ret.Visited {
		ret.Visited[i] = make([]direction, len(input[i]))
	}
	return ret
}
func (i *imgMap) isEdge(x, y int) bool {
	if x == 0 || y == 0 || x == len(i.Inmap[0])-1 || y == len(i.Inmap)-1 {
		return true
	} else {
		return false
	}
}

func (i *imgMap) inBounds(x, y int) bool {
	if x >= 0 && y >= 0 && x < len(i.Inmap[0]) && y < len(i.Inmap) {
		return true
	} else {
		return false
	}
}

func (d *direction) add(inD int) {
	switch inD {
	case 0:
		d.Right = true
	case 1:
		d.Up = true
	case 2:
		d.Left = true
	case 3:
		d.Down = true
	default:
		break
	}
}

func (d *direction) checkVisited(inD int) bool {
	switch inD {
	case 0:
		return d.Right
	case 1:
		return d.Up
	case 2:
		return d.Left
	case 3:
		return d.Down
	default:
		return false
	}
}

func (i *imgMap) processNode(x, y int, fromOne bool, fromDir int) {
	//fromDir schema {up,down,left,right}
	//fmt.Printf("currently at position %d, %d \n", x, y)
	i.Visited[x][y].add(fromDir)
	//fmt.Printf("visited: %d", i.Visited)
	//fmt.Print("Press 'Enter' to continue...")
	//bufio.NewReader(os.Stdin).ReadBytes('\n')
	dirs := [][]int{
		{-1, 0}, //from right
		{0, 1},  //from up
		{1, 0},  //from left
		{0, -1}, //from down
	}
	isOne := i.Inmap[x][y] == 1

	if i.isEdge(x, y) || fromOne {
		i.Outmap[x][y] = i.Inmap[x][y]
		for j := 0; j < len(dirs); j++ {
			if i.inBounds(x+dirs[j][0], y+dirs[j][1]) && !i.Visited[x+dirs[j][0]][y+dirs[j][1]].checkVisited(j) {
				i.processNode(x+dirs[j][0], y+dirs[j][1], isOne, j)
			}
		}
	}

}
func removeIslands(input [][]int) [][]int {
	mp := makeImgMap(input)
	mp.processNode(0, 0, false, 4)
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
	fmt.Println(mp.toStringCompare())
	mp.processNode(0, 0, false, 4)
	fmt.Println(mp.toStringCompare())

}
