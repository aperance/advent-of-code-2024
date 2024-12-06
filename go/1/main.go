package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../../inputs/1/full.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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

	sum := 0

	for i, v0 := range arr[0] {
		v1 := arr[1][i]
		if v0 < v1 {
			sum += v1 - v0
		} else {
			sum += v0 - v1
		}
	}

	fmt.Println("Result:", sum)
}
