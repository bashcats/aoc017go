package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	f, err := os.ReadFile("in1701")
	if err != nil {
		panic("could not read file")
	}
	digits := strings.TrimSpace(string(f))
	fmt.Println(sumMatchingIntervals(digits, 1))  // part one
	fmt.Println(sumMatchingIntervals(digits, len(digits) / 2))  // part two
}

func sumMatchingIntervals(seq string, dist int) int {
	if len(seq) == 0 {
		return 0
	}
	total := 0
	nseq := len(seq)
	for i := 0; i < len(seq); i++ {
		if (seq[i] == seq[(i + dist) % nseq]) {
			total += int(seq[i]) - 48
		}
	}
	return total
}

