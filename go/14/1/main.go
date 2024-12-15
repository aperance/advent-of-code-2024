package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
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

func main() {
	t := utils.StartTimer()
	defer t.PrintDuration()

	quadCounts := [4]int{}
	safetyFactor := 1

	scanner := utils.GetScanner()
	for scanner.Scan() {
		pos, vec := parseInput(scanner.Text())
		bot := robot{xPos: pos[0], yPos: pos[1], xVec: vec[0], yVec: vec[1]}
		bot.move(100)
		q := bot.getQuadrant()
		if q >= 0 && q < 4 {
			quadCounts[q]++
		}
	}

	for _, count := range quadCounts {
		safetyFactor *= count
	}

	fmt.Println("Safety factor:", safetyFactor)
}
