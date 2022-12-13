package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path, _ := os.Getwd()

	file, _ := os.ReadFile(filepath.Join(path, "input"))

	data := []rune(strings.TrimRight(string(file), "\n"))

	for pos := 0; pos <= len(data)-3; pos++ {
		counts := make(map[rune]int)
		for _, c := range data[pos : pos+4] {
			counts[c]++
		}
		if len(counts) == 4 {
			fmt.Println(pos + 4)
			break
		}
	}

	for pos := 0; pos <= len(data)-13; pos++ {
		counts := make(map[rune]int)
		for _, c := range data[pos : pos+14] {
			counts[c]++
		}
		if len(counts) == 14 {
			fmt.Println(pos + 14)
			break
		}
	}

}
