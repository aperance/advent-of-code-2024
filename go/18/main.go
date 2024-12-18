package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

const size = 71
const byteLimit = 1024

type memoryCell struct {
	corrupt  bool
	visited  bool
	distance int
}

type memorySpace struct {
	cells [size][size]memoryCell
	queue [][2]int
}

// func (s *memorySpace) print() {
// 	for y := 0; y < size; y++ {
// 		var row []string
// 		for x := 0; x < size; x++ {
// 			c := s.cells[x][y]
// 			if x == size-1 && y == size-1 {
// 				row = append(row, "E")
// 			} else if c.corrupt {
// 				row = append(row, "#")
// 			} else if c.visited {
// 				row = append(row, strconv.Itoa(c.distance))
// 			} else {
// 				row = append(row, ".")
// 			}
// 		}
// 		fmt.Println(row)
// 	}
// }

func (s *memorySpace) run() int {
	vectors := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for {
		if len(s.queue) == 0 {
			log.Fatal("Queue empty!")
		}

		ancestor := s.queue[0]
		s.queue = s.queue[1:]

		ancestorCell := &s.cells[ancestor[0]][ancestor[1]]

		for _, v := range vectors {
			x := ancestor[0] + v[0]
			y := ancestor[1] + v[1]

			if x < 0 || x >= size || y < 0 || y >= size {
				continue
			}

			cell := &s.cells[x][y]
			if cell.corrupt || cell.visited {
				continue
			}

			cell.visited = true
			cell.distance = ancestorCell.distance + 1

			if x == size-1 && y == size-1 {
				return cell.distance
			}

			s.queue = append(s.queue, [2]int{x, y})
		}
	}
}

func main() {
	t := utils.StartTimer()
	defer t.PrintDuration()

	space := memorySpace{}
	space.cells[0][0] = memoryCell{visited: true, distance: 0}
	space.queue = append(space.queue, [2]int{0, 0})

	scanner := utils.GetScanner()
	count := 0
	for scanner.Scan() {
		loc := strings.Split(scanner.Text(), ",")
		x, y := utils.Atoi(loc[0]), utils.Atoi(loc[1])
		space.cells[x][y].corrupt = true
		count++

		if count >= byteLimit {
			break
		}
	}

	result := space.run()
	fmt.Println("Found Path, distance:", result)
	// space.print()
}
