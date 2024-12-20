package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

type cell struct {
	blocked  bool
	end      bool
	visited  bool
	distance int
}

type race struct {
	cells [][]cell
	queue [][2]int
}

func (r *race) print() {
	// fmt.Print("\033[H")
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

func (r *race) findPath() {
	vectors := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for {
		if len(r.queue) == 0 {
			log.Fatal("Queue empty!")
		}

		ancestor := r.queue[0]
		r.queue = r.queue[1:]

		ancestorCell := &r.cells[ancestor[1]][ancestor[0]]

		if ancestorCell.end {
			fmt.Println(ancestorCell)
			return
		}

		for _, vector := range vectors {
			x := ancestor[0] + vector[0]
			y := ancestor[1] + vector[1]

			if x < 0 || x >= len(r.cells[0]) || y < 0 || y >= len(r.cells) {
				continue
			}

			cell := &r.cells[y][x]
			if cell.blocked || cell.visited {
				continue
			}

			cell.visited = true
			cell.distance = ancestorCell.distance + 1

			r.queue = append(r.queue, [2]int{x, y})
		}
	}
}

func (r *race) findCheats() int {
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

	return count
}

func main() {
	t := utils.StartTimer()
	defer t.PrintDuration()

	race := race{}

	scanner := utils.GetScanner()
	for scanner.Scan() {
		var row []cell
		for i, char := range scanner.Text() {
			row = append(row, cell{
				blocked: char == '#',
				end:     char == 'E',
				visited: char == 'S',
			})
			if char == 'S' {
				race.queue = append(race.queue, [2]int{i, len(race.cells)})
			}
		}
		race.cells = append(race.cells, row)
	}

	race.print()
	race.findPath()
	race.print()
	result := race.findCheats()
	fmt.Println("Cheats of over 100 picoseconds:", result)
}
