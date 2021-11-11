package main
import (
	"testing"
	"reflect"
)

func (t *testing) findIslands {
	input := [][]int {
		{0,1,1,0,0},
		{1,0,1,1,0},
		{1,0,0,0,1},
		{0,0,1,0,0},
		{1,1,0,0,1}
	}
	expected := [][]int {
		{0,1,1,0,0},
		{1,0,1,1,0},
		{1,0,0,0,1},
		{0,0,0,0,0},
		{1,1,0,0,1}
	}
	output := removeIslands(input)
	if !reflect.DeepEqual(output, expected) {
		t.Fatalf(`Expected: %v Got: %v`, expected, output)
	}
}