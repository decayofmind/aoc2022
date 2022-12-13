package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		ans1, ans2 int
	)

	file, err := os.ReadFile("./input")

	if err != nil {
		panic(err)
	}

	var sums []int

	data := strings.Split(string(file), "\n\n")

	for _, elf := range data {
		meals := strings.Split(elf, "\n")

		var sum int

		for _, m := range meals {
			i, _ := strconv.Atoi(m)
			sum += i
		}

		sums = append(sums, sum)
	}

	sort.Ints(sums)

	ans1 = sums[len(sums)-1]

	for _, el := range sums[len(sums)-3:] {
		ans2 += el
	}

	fmt.Println(ans1, ans2)
}
