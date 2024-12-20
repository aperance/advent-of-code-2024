package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

func check(s string, m map[string]struct{}, c map[string]int, depth int) int {
	if count, ok := c[s]; ok {
		return count
	}

	sum := 0

	for i := int(math.Min(float64(len(s)), 100)); i > 0; i-- {
		if _, ok := m[s[:i]]; ok {
			if i == len(s) {
				sum++
			}
			count := check(s[i:], m, c, depth+1)
			c[s[i:]] = count
			sum += count
		}
	}

	return sum
}

func main() {
	t := utils.StartTimer()
	defer t.PrintDuration()

	var designs []string
	towelMap := make(map[string]struct{})
	cache := make(map[string]int)

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

	sum := 0
	for _, design := range designs {
		count := check(design, towelMap, cache, 0)
		fmt.Println(design, count)
		sum += count
	}

	fmt.Println("Sum:", sum)
}
