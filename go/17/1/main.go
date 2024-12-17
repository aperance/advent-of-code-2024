package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

type computer struct {
	regA, regB, regC int
	program          []int
	pointer          int
	output           []int
}

func (c *computer) getComboOperand() int {
	operand := c.program[c.pointer+1]

	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.regA
	case 5:
		return c.regB
	case 6:
		return c.regC
	default:
		panic("Invalid operand")
	}
}

func (c *computer) execute() {
	opCode := c.program[c.pointer]
	literalOperand := c.program[c.pointer+1]
	nextPointer := c.pointer + 2

	switch opCode {
	case 0:
		c.regA = int(float64(c.regA) / math.Pow(float64(2), float64(c.getComboOperand())))
	case 1:
		c.regB = c.regB ^ literalOperand
	case 2:
		c.regB = c.getComboOperand() % 8
	case 3:
		if c.regA != 0 {
			nextPointer = literalOperand
		}
	case 4:
		c.regB = c.regB ^ c.regC
	case 5:
		c.output = append(c.output, c.getComboOperand()%8)
	case 6:
		c.regB = int(float64(c.regA) / math.Pow(float64(2), float64(c.getComboOperand())))
	case 7:
		c.regC = int(float64(c.regA) / math.Pow(float64(2), float64(c.getComboOperand())))
	default:
		log.Fatal("Invalid op code")
	}

	c.pointer = nextPointer
}

func (c *computer) runProgram() string {
	for c.pointer < len(c.program) {
		c.execute()
	}

	var output []string
	for _, val := range c.output {
		output = append(output, strconv.Itoa(val))
	}
	return strings.Join(output, ",")
}

func main() {
	t := utils.StartTimer()
	defer t.PrintDuration()

	c := computer{}

	scanner := utils.GetScanner()
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ": ")

		switch row[0] {
		case "Register A":
			c.regA = utils.Atoi(row[1])
		case "Register B":
			c.regB = utils.Atoi(row[1])
		case "Register C":
			c.regC = utils.Atoi(row[1])
		case "":
			continue
		case "Program":
			for _, s := range strings.Split(row[1], ",") {
				i := utils.Atoi(s)
				c.program = append(c.program, i)
			}
		default:
			log.Fatal("Invalid input")
		}
	}

	result := c.runProgram()

	fmt.Println("Result:", result)

}
