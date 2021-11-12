package main

import (
	"reflect"
	"testing"
)

func TestFindIslands(t *testing.T) {
	input := [][]int{
		{0, 1, 1, 0, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{0, 0, 1, 0, 0},
		{1, 1, 0, 0, 1},
	}
	expected := [][]int{
		{0, 1, 1, 0, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
	}
	output := removeIslands(input)
	if !reflect.DeepEqual(output, expected) {
		t.Fatalf(`Expected: %v Got: %v`, expected, output)
	}
}
func TestCenterIsland(t *testing.T) {
	input := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	expected := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	output := removeIslands(input)
	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("\nExpected: %v \nGot: %v", expected, output)
	}
}

/*
func TestUp(t *testing.T) {
	queues := [][]int{
		{},        // G
		{},        // 1
		{5, 5, 5}, // 2
		{},        // 3
		{},        // 4
		{},        // 5
		{},        // 6
	}
	bldg := NewBuilding(queues, 5)
	bldg.moveElevator()
	if !reflect.DeepEqual(bldg.El.History, []int{0, 2, 5, 0}) {
		t.Fatalf(`Expected: {0, 2, 5, 0} Got: %v`, bldg.El.History)
	}
}
*/
