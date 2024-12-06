package stdin

import (
	"bufio"
	"log"
	"os"
)

func GetScanner() *bufio.Scanner {
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		log.Fatal("No data provided via stdin")
	}

	return bufio.NewScanner(os.Stdin)
}
