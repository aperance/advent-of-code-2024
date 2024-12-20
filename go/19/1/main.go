package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

func check(s string, m map[string]struct{}, c map[string]bool, depth int) bool {
	if match, ok := c[s]; ok {
		return match
	}

	for i := int(math.Min(float64(len(s)), 10)); i > 0; i-- {
		if _, ok := m[s[:i]]; ok {
			if i == len(s) {
				return true
			}
			match := check(s[i:], m, c, depth+1)
			c[s] = match
			if match {
				return true
			}
		}
	}
	return false
}

func main() {
	t := utils.StartTimer()
	defer t.PrintDuration()

	var designs []string
	towelMap := make(map[string]struct{})
	cache := make(map[string]bool)

	scanner := utils.GetScanner()
	for scanner.Scan() {
		row := scanner.Text()
		if len(towelMap) == 0 {
			for _, t := range strings.Split(row, ", ") {
				towelMap[t] = struct{}{}
			}
		} else if len(row) > 0 {
			designs = append(designs, row)
		}
	}

	count := 0
	for _, design := range designs {
		if ok := check(design, towelMap, cache, 0); ok {
			fmt.Println("Pass:", design)
			count++
		} else {
			fmt.Println("Fail:", design)
		}
	}

	fmt.Println("Possible designs:", count)
}
