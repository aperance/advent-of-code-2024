package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

type queue struct {
	data [][2]int
}

func (q *queue) pop() [2]int {
	if len(q.data) == 0 {
		log.Fatal("Queue empty!")
	}

	result := q.data[0]
	q.data = q.data[1:]
	return result
}

func (q *queue) push(el [2]int) {
	q.data = append(q.data, el)
}

type cell struct {
	blocked  bool
	visited  bool
	distance int
}

type race struct {
	cells [][]cell
	start [2]int
	end   [2]int
}

func (r *race) print() {
	for y := 0; y < len(r.cells); y++ {
		for x := 0; x < len(r.cells[0]); x++ {
			c := r.cells[y][x]
			if c.visited {
				fmt.Printf("\033[10%vm", strconv.Itoa(c.distance%6+1))
				fmt.Printf("\033[30m %v ", strconv.Itoa(c.distance%10))
			} else if c.blocked {
				fmt.Print("\033[40m   ")
			} else {
				fmt.Print("\033[107m   ")
			}
			fmt.Print("\033[0m")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (r *race) getNeighbors(loc [2]int) [][2]int {
	var result [][2]int
	vectors := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for _, vector := range vectors {
		x := loc[0] + vector[0]
		y := loc[1] + vector[1]

		if x < 0 || x >= len(r.cells[0]) || y < 0 || y >= len(r.cells) {
			continue
		}
		result = append(result, [2]int{x, y})
	}

	return result
}

func (r *race) findPath() {
	var queue queue

	queue.push(r.start)

	for {
		ancestor := queue.pop()
		ancestorCell := &r.cells[ancestor[1]][ancestor[0]]

		for _, loc := range r.getNeighbors(ancestor) {
			cell := &r.cells[loc[1]][loc[0]]

			if cell.blocked || cell.visited {
				continue
			}

			cell.visited = true
			cell.distance = ancestorCell.distance + 1

			if loc == r.end {
				return
			}

			queue.push(loc)
		}
	}
}

func (r *race) findCheats() {
	count := 0

	for y := 1; y < len(r.cells)-1; y++ {
		for x := 1; x < len(r.cells[0])-1; x++ {
			if !r.cells[y][x].blocked {
				continue
			}

			up := r.cells[y-1][x]
			down := r.cells[y+1][x]
			left := r.cells[y][x-1]
			right := r.cells[y][x+1]

			if up.visited && down.visited {
				diff := math.Abs(float64(up.distance)-float64(down.distance)) - 2
				if diff >= 100 {
					count++
				}
			}

			if left.visited && right.visited {
				diff := math.Abs(float64(left.distance)-float64(right.distance)) - 2
				if diff >= 100 {
					count++
				}
			}
		}
	}

	fmt.Println("Cheats of over 100 picoseconds:", count)
}

func main() {
	t := utils.StartTimer()
	defer t.PrintDuration()

	race := race{}

	scanner := utils.GetScanner()
	for scanner.Scan() {
		var row []cell
		for x, char := range scanner.Text() {
			y := len(race.cells)
			if char == 'S' {
				race.start = [2]int{x, y}
			} else if char == 'E' {
				race.end = [2]int{x, y}
			}
			row = append(row, cell{blocked: char == '#', visited: char == 'S'})
		}
		race.cells = append(race.cells, row)
	}

	race.print()
	race.findPath()
	race.print()
	race.findCheats()
}
