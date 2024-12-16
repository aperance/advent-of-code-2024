package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

type warehouse struct {
	floorPlan [][]rune
	backup    [][]rune
	robotX    int
	robotY    int
}

func (w *warehouse) appendRow(str string) {
	var row []rune
	for _, rune := range str {
		switch rune {
		case '#':
			row = append(row, '#', '#')
		case 'O':
			row = append(row, '[', ']')
		case '.':
			row = append(row, '.', '.')
		case '@':
			w.robotX = len(row)
			w.robotY = len(w.floorPlan)
			row = append(row, '@', '.')
		}
	}
	w.floorPlan = append(w.floorPlan, row)
}

func (w *warehouse) saveBackup() {
	w.backup = make([][]rune, len(w.floorPlan))
	for i := range w.backup {
		w.backup[i] = make([]rune, len(w.floorPlan[i]))
		copy(w.backup[i], w.floorPlan[i])
	}
}

func (w *warehouse) restoreBackup() {
	for i := range w.backup {
		copy(w.floorPlan[i], w.backup[i])
	}
}

func (w *warehouse) moveRobot(xVec int, yVec int) {
	neighbor := w.floorPlan[w.robotY+yVec][w.robotX+xVec]

	if neighbor == '#' {
		return
	}

	if neighbor == '[' || neighbor == ']' {
		w.saveBackup()
		result := w.moveBox(w.robotX+xVec, w.robotY+yVec, xVec, yVec)
		if !result {
			w.restoreBackup()
			return
		}
	}

	w.floorPlan[w.robotY][w.robotX] = '.'
	w.robotX += xVec
	w.robotY += yVec
	w.floorPlan[w.robotY][w.robotX] = '@'

}

func (w *warehouse) moveBox(xPos int, yPos int, xVec int, yVec int) bool {
	var xPosLeft, xPosRight int
	if w.floorPlan[yPos][xPos] == '[' {
		xPosLeft = xPos
		xPosRight = xPos + 1
	} else if w.floorPlan[yPos][xPos] == ']' {
		xPosLeft = xPos - 1
		xPosRight = xPos
	}

	if xVec != 0 && yVec == 0 {
		neighbor := w.floorPlan[yPos][xPos+xVec*2]

		if neighbor == '#' {
			return false
		}

		if neighbor == '[' || neighbor == ']' {
			result := w.moveBox(xPos+xVec*2, yPos, xVec, 0)
			if !result {
				return false
			}
		}
	} else if xVec == 0 && yVec != 0 {
		neighborLeft := w.floorPlan[yPos+yVec][xPosLeft]
		neighborRight := w.floorPlan[yPos+yVec][xPosRight]

		if neighborLeft == '#' || neighborRight == '#' {
			return false
		}

		if neighborLeft == ']' && neighborRight == '[' { // Two neighboring boxes
			resultLeft := w.moveBox(xPosLeft, yPos+yVec, 0, yVec)
			if !resultLeft {
				return false
			}
			resultRight := w.moveBox(xPosRight, yPos+yVec, 0, yVec)
			if !resultRight {
				return false
			}
		} else if neighborLeft == '[' || neighborLeft == ']' { // Single neighboring box aligned or skewed left
			result := w.moveBox(xPosLeft, yPos+yVec, 0, yVec)
			if !result {
				return false
			}
		} else if neighborRight == '[' { // Sigle neighboring box skewed right only
			result := w.moveBox(xPosRight, yPos+yVec, 0, yVec)
			if !result {
				return false
			}
		}
	} else {
		panic("Invalid move instruction")
	}

	w.floorPlan[yPos][xPosLeft] = '.'
	w.floorPlan[yPos][xPosRight] = '.'
	w.floorPlan[yPos+yVec][xPosLeft+xVec] = '['
	w.floorPlan[yPos+yVec][xPosRight+xVec] = ']'

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
			case '[', ']':
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
			if c != '[' {
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
					time.Sleep(time.Millisecond * 50)
				}
			}
		}
	}

	fmt.Println("Sum of all GPS coordinates:", w.getSumOfGPS())
}
