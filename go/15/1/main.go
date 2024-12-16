package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

type warehouse struct {
	floorPlan [][]rune
	robotX    int
	robotY    int
}

func (w *warehouse) appendRow(row string) {
	for i, r := range row {
		if r == '@' {
			w.robotX = i
			w.robotY = len(w.floorPlan)
		}
	}
	w.floorPlan = append(w.floorPlan, []rune(row))
}

func (w *warehouse) moveRobot(xVec int, yVec int) {
	neighbor := w.floorPlan[w.robotY+yVec][w.robotX+xVec]

	if neighbor == '#' {
		return
	}

	if neighbor == 'O' {
		result := w.moveBox(w.robotX+xVec, w.robotY+yVec, xVec, yVec)
		if !result {
			return
		}
	}

	w.floorPlan[w.robotY][w.robotX] = '.'
	w.robotX += xVec
	w.robotY += yVec
	w.floorPlan[w.robotY][w.robotX] = '@'
}

func (w *warehouse) moveBox(xPos int, yPos int, xVec int, yVec int) bool {
	neighbor := w.floorPlan[yPos+yVec][xPos+xVec]

	if neighbor == '#' {
		return false
	}

	if neighbor == 'O' {
		result := w.moveBox(xPos+xVec, yPos+yVec, xVec, yVec)
		if !result {
			return false
		}
	}

	w.floorPlan[yPos][xPos] = '.'
	w.floorPlan[yPos+yVec][xPos+xVec] = 'O'

	return true
}

func (w *warehouse) printFloorPlan() {
	fmt.Print("\033[H")
	for i := 0; i < len(w.floorPlan); i++ {
		for j := 0; j < len(w.floorPlan[0]); j++ {
			char := w.floorPlan[i][j]
			switch char {
			case '#':
				fmt.Print("\033[33m" + string(char))
			case 'O':
				fmt.Print("\033[32m" + string(char))
			case '@':
				fmt.Print("\033[35m" + string(char))
			case '.':
				fmt.Print(" ")
			}
			fmt.Print(" \033[0m")
		}
		fmt.Print("\n")
	}
}

func (w *warehouse) getSumOfGPS() int {
	sum := 0
	for y, row := range w.floorPlan {
		for x, c := range row {
			if c != 'O' {
				continue
			}
			sum += y*100 + x
		}
	}
	return sum
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
		f := func() { fmt.Print("\033[0m\033[?25h\n") }
		utils.SetCleanup(f)
		defer f()
	}

	w := warehouse{}

	scanner := utils.GetScanner()
	floorplanDone := false
	for scanner.Scan() {
		row := scanner.Text()

		if len(row) == 0 {
			floorplanDone = true
			continue
		}

		if !floorplanDone {
			w.appendRow(row)
		} else {
			for _, instruction := range row {
				switch instruction {
				case '^':
					w.moveRobot(0, -1)
				case 'v':
					w.moveRobot(0, 1)
				case '<':
					w.moveRobot(-1, 0)
				case '>':
					w.moveRobot(1, 0)
				}

				if animate {
					w.printFloorPlan()
					time.Sleep(time.Millisecond * 100)
				}
			}
		}
	}

	fmt.Println("Sum of all GPS coordinates:", w.getSumOfGPS())
}
