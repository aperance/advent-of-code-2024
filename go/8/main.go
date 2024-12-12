package main

import (
	"fmt"

	"github.com/aperance/advent-of-code-2024/go/pkg/stdin"
)

func getInput() (map[rune][][2]int, int, int) {
	antennaMap := make(map[rune][][2]int)
	width := 0

	scanner := stdin.GetScanner()
	rowIndex := 0
	for scanner.Scan() {
		row := scanner.Text()
		for colIndex, char := range row {
			if string(char) == "." {
				continue
			}
			position := [2]int{rowIndex, colIndex}
			antennaMap[char] = append(antennaMap[char], position)
		}

		if width == 0 {
			width = len(row)
		}
		rowIndex++
	}

	return antennaMap, rowIndex, width
}

type antinodes struct {
	initialSet  map[string]any
	harmonicSet map[string]any
}

func (a *antinodes) findAntinodes(start [2]int, step [2]int, direction int, max [2]int) {
	for i := 0; ; i++ {
		new0 := start[0] + step[0]*direction*i
		new1 := start[1] + step[1]*direction*i
		fmt.Println(new0, new1)
		if new0 < 0 || new0 > max[0] || new1 < 0 || new1 > max[1] {
			break
		}
		id := string(new0) + ":" + string(new1)
		a.harmonicSet[id] = struct{}{}
		if i == 1 {
			a.initialSet[id] = struct{}{}
		}
	}
}

func main() {
	antennaMap, height, width := getInput()
	antinodes := antinodes{
		initialSet:  make(map[string]any),
		harmonicSet: make(map[string]any),
	}

	for _, positions := range antennaMap {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				step := [2]int{positions[i][0] - positions[j][0], positions[i][1] - positions[j][1]}
				max := [2]int{height - 1, width - 1}
				antinodes.findAntinodes(positions[i], step, 1, max)
				antinodes.findAntinodes(positions[j], step, -1, max)
			}
		}
	}

	fmt.Println("Antinode count:", len(antinodes.initialSet))
	fmt.Println("Antinode harmonic count:", len(antinodes.harmonicSet))
}
