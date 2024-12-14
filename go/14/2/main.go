package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aperance/advent-of-code-2024/go/pkg/stdin"
)

const width = 101
const height = 103

type robot struct {
	xVec int
	xPos int
	yVec int
	yPos int
}

func (r *robot) move(count int) {
	for i := 0; i < count; i++ {
		r.xPos += r.xVec
		r.yPos += r.yVec

		if r.xPos >= width {
			r.xPos -= width
		}
		if r.xPos < 0 {
			r.xPos = width + r.xPos
		}
		if r.yPos >= height {
			r.yPos -= height
		}
		if r.yPos < 0 {
			r.yPos = height + r.yPos
		}
	}
}

func (r *robot) getQuadrant() int {
	if r.xPos < width/2 && r.yPos < height/2 {
		return 0
	}
	if r.xPos > width/2 && r.yPos < height/2 {
		return 1
	}
	if r.xPos < width/2 && r.yPos > height/2 {
		return 2
	}
	if r.xPos > width/2 && r.yPos > height/2 {
		return 3
	}
	return -1
}

func strToIntArray(strings []string) []int {
	result := make([]int, len(strings))
	for i, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		result[i] = num
	}
	return result
}

func parseInput(input string) ([]int, []int) {
	row := strings.Split(input, " ")
	pos := strings.Split(strings.TrimPrefix(row[0], "p="), ",")
	vec := strings.Split(strings.TrimPrefix(row[1], "v="), ",")

	return strToIntArray(pos), strToIntArray(vec)
}

func filledGrid() [height][width]string {
	grid := [height][width]string{}

	for i, row := range grid {
		for j, _ := range row {
			grid[i][j] = " "
		}
	}

	return grid
}

func main() {

	scanner := stdin.GetScanner()
	var bots []robot
	for scanner.Scan() {
		pos, vec := parseInput(scanner.Text())
		bot := robot{xPos: pos[0], yPos: pos[1], xVec: vec[0], yVec: vec[1]}
		bots = append(bots, bot)
	}

	for i := 1; ; i++ {
		grid := filledGrid()

		quadCounts := [4]int{}
		safetyFactor := 1

		for j := 0; j < len(bots); j++ {
			bots[j].move(1)
			grid[bots[j].yPos][bots[j].xPos] = "0"
			q := bots[j].getQuadrant()
			if q >= 0 && q < 4 {
				quadCounts[q]++
			}
		}

		for _, count := range quadCounts {
			safetyFactor *= count
		}

		if safetyFactor < 100000000 {
			fmt.Println("Generation", i)
			for _, line := range grid {
				fmt.Println(line)
			}
			return
		}
	}
}
