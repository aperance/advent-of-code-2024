package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"github.com/aperance/advent-of-code-2024/go/pkg/utils"
)

func appendMultiple(slice []string, count int, elems ...string) []string {
	for i := 0; i < count; i++ {
		slice = append(slice, elems...)
	}
	return slice
}

type file struct {
	empty     bool
	id        string
	blocks    int
	relocated []struct {
		id     string
		blocks int
	}
}

func (f *file) relocate(newFile *file) {
	if !newFile.empty || newFile.blocks < f.blocks {
		panic("Not enough free space")
	}

	newFile.blocks -= f.blocks
	newFile.relocated = append(newFile.relocated, struct {
		id     string
		blocks int
	}{
		id:     f.id,
		blocks: f.blocks,
	})

	f.empty = true
	f.id = ""
}

func (f *file) toArray() []string {
	var result []string
	if f.empty {
		for _, r := range f.relocated {
			result = appendMultiple(result, r.blocks, r.id)
		}
		result = appendMultiple(result, f.blocks, ".")
	} else {
		result = appendMultiple(result, f.blocks, f.id)
	}
	return result
}

func main() {
	scanner := utils.GetScanner()
	scanner.Split(bufio.ScanRunes)

	fileID := 0
	freeSpace := false
	var files []file
	var disk []string
	for scanner.Scan() {
		blocks, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		newFile := file{blocks: blocks}
		if freeSpace {
			newFile.empty = true
		} else {
			newFile.id = strconv.Itoa(fileID)
		}
		files = append(files, newFile)

		if freeSpace {
			fileID++
		}
		freeSpace = !freeSpace
	}

	for i := len(files) - 1; i >= 0; i-- {
		if files[i].empty || files[i].blocks == 0 {
			continue
		}

		for j := 0; j <= i; j++ {
			if files[j].empty && files[j].blocks >= files[i].blocks {
				files[i].relocate(&files[j])
				break
			}
		}
	}

	for _, file := range files {
		disk = append(disk, file.toArray()...)
	}

	checksum := 0
	for pos, id := range disk {
		if id != "." {
			num, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}

			checksum += pos * num
		}
	}

	fmt.Println("Checksum:", checksum)
}
