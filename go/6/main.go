package main

import (
	"fmt"
	"slices"

	"github.com/aperance/advent-of-code-2024/go/pkg/stdin"
)

type visitedMap struct {
	_map map[string][][2]int
}

// Sets visited position and current vector in map.
// Returns true if value is unchanged, false otherwise
func (v *visitedMap) set(position [2]int, data [2]int) bool {
	id := string(position[0]) + ":" + string(position[1])
	previousData := v._map[id]
	if slices.Contains(previousData, data) {
		return true
	}
	v._map[id] = append(v._map[id], data)
	return false
}

func (v *visitedMap) count() int {
	return len(v._map)
}

type lab struct {
	field         [][]bool
	guardPosition [2]int
	extraObstacle [2]int
}

func (l *lab) loadData() {
	scanner := stdin.GetScanner()
	rowIndex := 0
	for scanner.Scan() {
		bytes := scanner.Bytes()
		row := make([]bool, len(bytes))
		for colIndex, char := range bytes {
			if char == 35 {
				row[colIndex] = true
			} else {
				row[colIndex] = false
			}

			if char == 94 {
				l.guardPosition = [2]int{rowIndex, colIndex}
			}
		}

		l.field = append(l.field, row)
		rowIndex++
	}
}

// Processes the guard movements. Returns a count of distinct
// positions visited and if the guard is stuck in a loop.
func (l lab) runGuard() (int, bool) {
	vectors := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := visitedMap{_map: make(map[string][][2]int)}

	for {
		for _, vec := range vectors {
			for {
				// fmt.Println(l.guardPosition, vec)
				unchanged := visited.set(l.guardPosition, vec)
				// fmt.Println(unchanged)
				if unchanged {
					return visited.count(), true
				}

				next := [2]int{l.guardPosition[0] + vec[0], l.guardPosition[1] + vec[1]}

				oob := next[0] < 0 || next[0] >= len(l.field) || next[1] < 0 || next[1] >= len(l.field[0])
				if oob {
					return visited.count(), false
				}

				blocked := l.field[next[0]][next[1]] || next == l.extraObstacle
				if blocked {
					break
				}

				l.guardPosition = next
			}
		}
	}
}

func newLab() lab {
	l := lab{extraObstacle: [2]int{-1, -1}}
	l.loadData()
	return l
}

func main() {
	lab := newLab()
	count, _ := lab.runGuard()

	fmt.Println("Distinct positions:", count)

	stuckCount := 0
	for i, row := range lab.field {
		for j := range row {
			lab.extraObstacle = [2]int{i, j}
			_, stuck := lab.runGuard()
			if stuck {
				stuckCount++
			}
		}
	}

	fmt.Println("Extra obstacles causing guard to be stuck:", stuckCount)
}
