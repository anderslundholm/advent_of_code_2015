package day01

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2015/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2015/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	lines, err := reader.ReadLines("day01/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	result := findFloor(lines, true)

	fmt.Printf("Day01 Part2 result: %v\n", result)
}
