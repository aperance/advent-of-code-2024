package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

func (s *memorySpace) print() {
	fmt.Print("\033[H")
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			c := s.cells[x][y]
			if c.visited {
				fmt.Printf("\033[10%vm", strconv.Itoa(c.distance%6+1))
				fmt.Printf("\033[30m %v ", strconv.Itoa(c.distance%10))
			} else if c.corrupt {
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

func (s *memorySpace) run(animate bool) int {
	vectors := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for {
		if len(s.queue) == 0 {
			log.Fatal("Queue empty!")
		}

		ancestor := s.queue[0]
		s.queue = s.queue[1:]

		ancestorCell := &s.cells[ancestor[0]][ancestor[1]]

		if ancestor[0] == size-1 && ancestor[1] == size-1 {
			return ancestorCell.distance
		}

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

			s.queue = append(s.queue, [2]int{x, y})
		}

		if animate {
			s.print()
		}
	}
}

func main() {
	t := utils.StartTimer()
	defer t.PrintDuration()

	animate := false
	if len(os.Args) > 1 && os.Args[1] == "--animate" {
		animate = true
	}

	if animate {
		fmt.Print("\033[2J\033[?25l")
		f := func() { fmt.Print("\033[0m\033[?25h") }
		utils.SetCleanup(f)
		defer f()
	}

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

	result := space.run(animate)
	fmt.Println("Found Path, distance:", result)
}
