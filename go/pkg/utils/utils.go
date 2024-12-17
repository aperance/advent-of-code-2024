package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type exeutionTimer struct {
	start time.Time
}

func (t *exeutionTimer) PrintDuration() {
	fmt.Println("Elapsed time:", time.Since(t.start))
}

func StartTimer() exeutionTimer {
	return exeutionTimer{start: time.Now()}
}

func GetScanner() *bufio.Scanner {
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		log.Fatal("Input data must be piped in via stdin")
	}

	return bufio.NewScanner(os.Stdin)
}

func SetCleanup(f func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		f()
		os.Exit(1)
	}()
}

func EncodeMapKey(i int, j int) string {
	return strconv.Itoa(i) + ":" + strconv.Itoa(j)
}

func DecodeMapKey(s string) (int, int) {
	var x, y int
	var err error

	vals := strings.Split(s, ":")

	x, err = strconv.Atoi(vals[0])
	if err != nil {
		log.Fatal(err)
	}

	y, err = strconv.Atoi(vals[1])
	if err != nil {
		log.Fatal(err)
	}

	return x, y
}
