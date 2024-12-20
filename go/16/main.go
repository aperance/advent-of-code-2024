package main

import (
	"fmt"
	"maps"
	"sync"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

var wg sync.WaitGroup

func getDirectionData(i int) (int, int, rune) {
	vecX := [4]int{1, 0, -1, 0}
	vecY := [4]int{0, 1, 0, -1}
	markers := [4]rune{'>', 'v', '<', '^'}

	return vecX[i], vecY[i], markers[i]
}

func changeDirection(direction int, turn int) int {
	if turn < -1 || turn > 1 {
		panic("Unexpected turn param")
	}

	result := direction + turn
	if result == -1 {
		result = 3
	} else if result == 4 {
		result = 0
	}

	return result
}

type minScore struct {
	data map[string]int
	mu   sync.Mutex
}

type Maze struct {
	_map      [][]rune
	position  [2]int
	direction int // E=0, S=1, W=2, N=3
	path      map[string]rune
	score     int
	minScore  *minScore
}

func (m *Maze) endReached() bool {
	return m._map[m.position[1]][m.position[0]] == 'E'
}

func (m *Maze) checkNeighbor(i int) (bool, bool) {
	d := changeDirection(m.direction, i)
	vecX, vecY, _ := getDirectionData(d)

	x := m.position[0] + vecX
	y := m.position[1] + vecY

	key := utils.EncodeMapKey(x, y)
	if _, ok := m.path[key]; ok {
		return true, false
	}

	switch m._map[y][x] {
	case 'S':
		return true, false
	case '#':
		return false, true
	default:
		return false, false
	}
}

func (m *Maze) turn(i int) {
	if i != 0 {
		m.direction = changeDirection(m.direction, i)
		m.score += 1000
	}
}

func (m *Maze) step() {
	vecX, vecY, marker := getDirectionData(m.direction)

	x := m.position[0] + vecX
	y := m.position[1] + vecY

	m.position = [2]int{x, y}

	m.path[utils.EncodeMapKey(x, y)] = marker

	m.score++
}

func (m *Maze) run(queue chan Maze, results chan Maze) {
	var moves []int

	if m.endReached() {
		results <- *m
		return
	}

	for i := -1; i <= 1; i++ {
		loop, wall := m.checkNeighbor(i)
		if loop {
			results <- *m
			return
		} else if !wall {
			moves = append(moves, i)
		}
	}

	if len(moves) == 0 {
		results <- *m
		return
	}

	if len(moves) == 1 {
		m.minScore.mu.Lock()
		key := utils.EncodeMapKey(m.position[0], m.position[1])
		minScore, ok := m.minScore.data[key]
		if !ok || m.score < minScore {
			m.minScore.data[key] = m.score
		}
		m.minScore.mu.Unlock()

		if ok && m.score > minScore {
			results <- *m
			return
		}
	}

	for i := 1; i < len(moves); i++ {
		newMaze := *m
		newMaze.path = maps.Clone(m.path)

		newMaze.turn(moves[i])
		newMaze.step()
		queue <- newMaze
	}

	m.turn(moves[0])
	m.step()
	m.run(queue, results)
}

func main() {
	t := utils.StartTimer()
	defer t.PrintDuration()

	minScore := minScore{data: map[string]int{}}
	maze := Maze{path: map[string]rune{}, minScore: &minScore}

	scanner := utils.GetScanner()
	for scanner.Scan() {
		row := scanner.Text()
		for i, r := range row {
			if r == 'S' {
				maze.position = [2]int{i, len(maze._map)}
				maze.path[utils.EncodeMapKey(i, len(maze._map))] = '>'
			}
		}
		maze._map = append(maze._map, []rune(scanner.Text()))
	}

	lowestScore := 0
	resultMap := make(map[int]map[string]struct{})
	queue := make(chan Maze)
	results := make(chan Maze)

	spawn := func(m *Maze) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.run(queue, results)
		}()
	}

	spawn(&maze)

	go func() {
		for m := range queue {
			spawn(&m)
		}
	}()

	go func() {
		for m := range results {
			wg.Add(1)
			if m.endReached() {
				fmt.Printf("Successful path found! (score: %v)\n", m.score)
				if lowestScore == 0 || lowestScore > m.score {
					lowestScore = m.score
				}
				if _, ok := resultMap[m.score]; !ok {
					resultMap[m.score] = make(map[string]struct{})
				}
				for key := range m.path {
					resultMap[m.score][key] = struct{}{}
				}
			}
			wg.Done()
		}
	}()

	wg.Wait()

	fmt.Println("Lowest score:", lowestScore)
	fmt.Println("Best path tile count:", len(resultMap[lowestScore]))
}
