package main

import (
	"fmt"

	"github.com/aperance/advent-of-code-2024/go/pkg/stdin"
)

func main() {
	scanner := stdin.GetScanner()

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	wordCount := 0
	crossCount := 0
	for x, line := range input {
		for y, char := range line {
			// Check for XMAS
			if string(char) == "X" {
				for xDirection := -1; xDirection <= 1; xDirection++ {
					for yDirection := -1; yDirection <= 1; yDirection++ {
						word := make([]byte, 4)
						for magnitude := 0; magnitude <= 3; magnitude++ {
							i := xDirection*magnitude + x
							j := yDirection*magnitude + y
							if i >= 0 && i < len(input) && j >= 0 && j < len(input[x]) {
								word[magnitude] = input[i][j]
							}
						}
						if string(word) == "XMAS" {
							wordCount++
						}
					}
				}
			}

			// Check for X-MAS
			if string(char) == "A" {
				slashCount := 0
				for xDirection := -1; xDirection <= 1; xDirection += 2 {
					for yDirection := -1; yDirection <= 1; yDirection += 2 {
						i := x + xDirection
						j := y + yDirection
						if i < 0 || i >= len(input) || j < 0 || j >= len(input[x]) {
							continue
						}
						if string(input[i][j]) != "M" {
							continue
						}
						i = x - xDirection
						j = y - yDirection
						if i < 0 || i >= len(input) || j < 0 || j >= len(input[x]) {
							continue
						}
						if string(input[i][j]) != "S" {
							continue
						}
						slashCount++
					}
				}
				if slashCount == 2 {
					crossCount++
				}
			}
		}
	}

	fmt.Println("XMAS count:", wordCount)
	fmt.Println("X-MAS count:", crossCount)
}
