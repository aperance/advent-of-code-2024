package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

func checkEquation(operands []int, operations []int, expectedResult int) bool {
	if len(operations) == len(operands)-1 {
		accumulator := operands[0]

		for i, op := range operations {
			if op == 0 {
				accumulator = accumulator + operands[i+1]
			} else if op == 1 {
				accumulator = accumulator * operands[i+1]
			} else {
				concat, err := strconv.Atoi(strconv.Itoa(accumulator) + strconv.Itoa(operands[i+1]))
				if err != nil {
					log.Fatal(err)
				}
				accumulator = concat
			}
		}

		return accumulator == expectedResult
	} else {
		for i := 0; i < 3; i++ {
			newOperations := append(operations, i)
			result := checkEquation(operands, newOperations, expectedResult)
			if result {
				return true
			}
		}

		return false
	}
}

func main() {
	t := utils.StartTimer()
	defer t.PrintDuration()

	sum := 0

	scanner := utils.GetScanner()
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ": ")

		result, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatal(err)
		}

		operands := []int{}
		for _, str := range strings.Split(row[1], " ") {
			operand, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			operands = append(operands, operand)
		}

		matching := checkEquation(operands, []int{}, result)
		if matching {
			sum += result
		}
	}

	fmt.Println("Sum of matching equations:", sum)
}
