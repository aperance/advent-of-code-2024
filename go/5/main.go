package main

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

func checkUpdates(pages []string, ruleMap map[string][]string) bool {
	for i := range pages {
		for j := i + 1; j < len(pages); j++ {
			f := slices.Contains(ruleMap[pages[i]], pages[j])
			if !f {
				return false
			}
		}
	}
	return true
}

func main() {
	ruleMap := make(map[string][]string)
	parsingRules := true
	sum := 0
	correctedSum := 0

	scanner := utils.GetScanner()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			rule := strings.Split(line, "|")
			ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
			continue
		}

		pages := strings.Split(line, ",")
		ok := checkUpdates(pages, ruleMap)

		if !ok {
			sort.Slice(pages, func(i, j int) bool {
				return slices.Contains(ruleMap[pages[i]], pages[j])
			})
		}

		centerValue, err := strconv.Atoi(pages[len(pages)/2])
		if err != nil {
			log.Fatal(err)
		}

		if ok {
			sum += centerValue
		} else {
			correctedSum += centerValue
		}
	}

	fmt.Println("Sum of passing center pages:", sum)
	fmt.Println("Sum of corrected center pages:", correctedSum)
}
