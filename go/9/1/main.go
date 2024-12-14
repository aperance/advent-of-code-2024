package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"github.com/aperance/advent-of-code-2024/go/pkg/stdin"
)

func main() {
	scanner := stdin.GetScanner()
	scanner.Split(bufio.ScanRunes)

	fileID := 0
	freeSpace := false
	var disk []string
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < num; i++ {
			if freeSpace {
				disk = append(disk, ".")
			} else {
				disk = append(disk, strconv.Itoa(fileID))
			}
		}

		if freeSpace {
			fileID++
		}
		freeSpace = !freeSpace
	}

	left := 0
	right := len(disk) - 1
	for {
		if left > right {
			break
		}
		if disk[left] != "." {
			left++
			continue
		}
		if disk[right] == "." {
			right--
			continue
		}

		disk[left] = disk[right]
		disk[right] = "."
	}

	checksum := 0
	for pos, id := range disk {
		if id == "." {
			break
		}
		num, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}
		checksum += pos * num
	}

	fmt.Println("Checksum:", checksum)
}
