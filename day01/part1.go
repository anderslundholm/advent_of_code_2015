package day01

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2015/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2015/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day01/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	result := findFloor(lines, false)

	fmt.Printf("Day01 Part1 result: %v\n", result)
}
