package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		ans1, ans2 int
	)

	file, _ := os.ReadFile("./input")

	for _, l := range strings.Split(string(file), "\n") {
		var (
			opponent string
			you      string
		)

		fmt.Sscanf(l, "%s %s", &opponent, &you)

		if opponent == "A" && you == "X" {
			ans1 += 4
			ans2 += 3
		} else if opponent == "A" && you == "Y" {
			ans1 += 8
			ans2 += 4
		} else if opponent == "A" && you == "Z" {
			ans1 += 3
			ans2 += 8
		} else if opponent == "B" && you == "X" {
			ans1 += 1
			ans2 += 1
		} else if opponent == "B" && you == "Y" {
			ans1 += 5
			ans2 += 5
		} else if opponent == "B" && you == "Z" {
			ans1 += 9
			ans2 += 9
		} else if opponent == "C" && you == "X" {
			ans1 += 7
			ans2 += 2
		} else if opponent == "C" && you == "Y" {
			ans1 += 2
			ans2 += 6
		} else if opponent == "C" && you == "Z" {
			ans1 += 6
			ans2 += 7
		}
	}

	fmt.Println(ans1, ans2)
}
