package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

func checkEquation(operands []int, result int) bool {
	matching := false
	operationCount := len(operands) - 1

	// Iterate over range of values where max number of bits equals number of operations.
	// This represents all operation combinations (0 = addition and 1 = multiplication).
	for i := 0; i < 1<<operationCount; i++ {
		accumulator := operands[0]

		// Iterate over operand positions
		for j := 0; j < operationCount; j++ {
			// If bit value of i at position j is 0 use addition, otherwise multiplication
			if i&(1<<j) == 0 {
				accumulator = accumulator + operands[j+1]
			} else {
				accumulator = accumulator * operands[j+1]
			}
		}

		if accumulator == result {
			matching = true
		}
	}

	return matching
}

func main() {
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

		matching := checkEquation(operands, result)
		if matching {
			sum += result
		}
	}

	fmt.Println("Sum of matching equations:", sum)
}
