package day01

import "fmt"

func findFloor(lines []string, part2 bool) int {
	// fmt.Println(lines)
	count := 0
	for _, line := range lines {
		fmt.Println(len(line))
		for i, char := range line {
			switch char {
			case '(':
				count++
			case ')':
				count--
			}
			if part2 && count < 0 {
				return i + 1
			}
			// fmt.Printf("%c\n", char)
		}
	}

	return count
}
