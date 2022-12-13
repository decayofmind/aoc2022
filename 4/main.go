package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func isNotOverlap(range1, range2 string) bool {
	e1 := strings.Split(range1, "-")
	e2 := strings.Split(range2, "-")

	a1, _ := strconv.Atoi(e1[0])
	b1, _ := strconv.Atoi(e1[1])
	a2, _ := strconv.Atoi(e2[0])
	b2, _ := strconv.Atoi(e2[1])

	if (a2 > b1) || (a1 > b2) {
		return true
	}

	return false
}

func isFullOverlap(range1, range2 string) bool {
	e1 := strings.Split(range1, "-")
	e2 := strings.Split(range2, "-")

	a1, _ := strconv.Atoi(e1[0])
	b1, _ := strconv.Atoi(e1[1])
	a2, _ := strconv.Atoi(e2[0])
	b2, _ := strconv.Atoi(e2[1])

	if (a1 >= a2 && b1 <= b2) || (a2 >= a1 && b2 <= b1) {
		return true
	}

	return false
}

func main() {
	path, _ := os.Getwd()

	file, _ := os.ReadFile(filepath.Join(path, "input"))

	data := strings.Split(string(file), "\n")

	data = data[:len(data)-1]

	var ans1, ans2 int

	for _, l := range data {
		splitted := strings.Split(l, ",")
		if isFullOverlap(splitted[0], splitted[1]) {
			ans1 += 1
		}

		if isNotOverlap(splitted[0], splitted[1]) {
			ans2 += 1
		}
	}

	fmt.Println(ans1, len(data)-ans2)
}
