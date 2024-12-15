package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

func main() {
	re := regexp.MustCompile(`(do|don't)\(\)|mul\((\d+),(\d+)\)`)
	scanner := utils.GetScanner()
	enabled := true
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		matches := re.FindAllStringSubmatch(text, -1)

		for _, match := range matches {
			fmt.Println(match)

			command := match[1]
			if command == "do" {
				enabled = true
			} else if command == "don't" {
				enabled = false
			} else if enabled {
				v1, err := strconv.Atoi(match[2])
				if err != nil {
					log.Fatal(err)
				}
				v2, err := strconv.Atoi(match[3])
				if err != nil {
					log.Fatal(err)
				}
				sum += v1 * v2
			}
		}
	}

	fmt.Println("Sum of multiplications:", sum)
}
