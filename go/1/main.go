package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

func main() {
	scanner := utils.GetScanner()

	arr := make([][]int, 2)

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "   ")
		if len(row) != 2 {
			log.Fatal("Row parsing error")
		}

		for i, str := range row {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			arr[i] = append(arr[i], num)
		}
	}

	sort.Ints(arr[0])
	sort.Ints(arr[1])

	distance, similarity := 0, 0

	for i, v0 := range arr[0] {
		for j, v1 := range arr[1] {
			if i == j {
				if v0 < v1 {
					distance += v1 - v0
				} else {
					distance += v0 - v1
				}
			}

			if v0 == v1 {
				similarity += v1
			}
		}
	}

	fmt.Println("Total distance:", distance)
	fmt.Println("Similarity score:", similarity)
}
