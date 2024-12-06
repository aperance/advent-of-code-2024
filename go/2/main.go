package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func removeIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func checkReport(row []string) (bool, int) {
	var previousValue int
	var previousSign bool

	for i, str := range row {
		value, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}

		if i > 0 {
			rawDiff := float64(previousValue - value)
			diff := math.Abs(rawDiff)
			sign := math.Signbit(rawDiff)
			signChange := i > 1 && sign != previousSign

			if diff < 1 || diff > 3 || signChange {
				return false, i
			}

			previousSign = sign
		}

		previousValue = value
	}

	return true, 0
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReports := 0
	dampenedReports := 0

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		safe, errorIndex := checkReport(row)
		if safe {
			safeReports++
		} else {
			for i := 0; i <= errorIndex; i++ {
				safe, _ = checkReport(removeIndex(row, i))
				if safe {
					dampenedReports++
					break
				}
			}
		}
	}

	fmt.Println("Safe count:", safeReports)
	fmt.Println("Safe count with dampener:", safeReports+dampenedReports)
}
