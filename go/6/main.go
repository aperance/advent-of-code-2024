package main

import (
	"fmt"
	"slices"
	"sync"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

type visitedMap struct {
	_map map[string][][2]int
}

// Sets visited position and current vector in map.
// Returns true if value is unchanged, false otherwise
func (v *visitedMap) set(position [2]int, data [2]int) bool {
	id := string(rune(position[0])) + ":" + string(rune(position[1]))
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
}

func (l *lab) loadData() {
	scanner := utils.GetScanner()
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
func (l lab) runGuard(extraObstacle [2]int) (int, bool) {
	vectors := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := visitedMap{_map: make(map[string][][2]int)}

	for {
		for _, vec := range vectors {
			for {
				unchanged := visited.set(l.guardPosition, vec)
				if unchanged {
					return visited.count(), true
				}

				next := [2]int{l.guardPosition[0] + vec[0], l.guardPosition[1] + vec[1]}

				oob := next[0] < 0 || next[0] >= len(l.field) || next[1] < 0 || next[1] >= len(l.field[0])
				if oob {
					return visited.count(), false
				}

				blocked := l.field[next[0]][next[1]] || next == extraObstacle
				if blocked {
					break
				}

				l.guardPosition = next
			}
		}
	}
}

func newLab() lab {
	l := lab{}
	l.loadData()
	return l
}

type stuckCount struct {
	mu    sync.Mutex
	count int
}

func (s *stuckCount) increment() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.count++
}

func main() {
	t := utils.StartTimer()
	defer t.PrintDuration()

	lab := newLab()
	count, _ := lab.runGuard([2]int{-1, -1})

	fmt.Println("Distinct positions:", count)

	var wg sync.WaitGroup
	stuckCount := stuckCount{}

	for i, row := range lab.field {
		for j := range row {
			wg.Add(1)

			go func() {
				defer wg.Done()

				_, stuck := lab.runGuard([2]int{i, j})
				if stuck {
					stuckCount.increment()
				}
			}()
		}
	}

	wg.Wait()

	fmt.Println("Extra obstacles causing guard to be stuck:", stuckCount.count)
}
