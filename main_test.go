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
