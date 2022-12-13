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

	data := strings.Split(string(file), "\n\n")

	var init_state [][]rune
	for _, s := range strings.Split(data[0], "\n") {
		init_state = append(init_state, []rune(s))
	}

	var stacks_p1, stacks_p2 [][]rune

	for p, r := range init_state[len(init_state)-1] {
		if r != ' ' {
			var stack []rune

			for x := len(init_state) - 2; x >= 0; x -= 1 {
				c := init_state[x][p]
				if c != ' ' {
					stack = append(stack, c)

				}
			}

			stacks_p1 = append(stacks_p1, stack)
			stacks_p2 = append(stacks_p2, stack)
		}

	}

	for _, o := range strings.Split(strings.TrimRight(data[1], "\n"), "\n") {
		var count, from, to int
		fmt.Sscanf(o, "move %d from %d to %d", &count, &from, &to)

		for x := 1; x <= count; x += 1 {
			el := stacks_p1[from-1][len(stacks_p1[from-1])-1]
			stacks_p1[from-1] = stacks_p1[from-1][:len(stacks_p1[from-1])-1]

			stacks_p1[to-1] = append(stacks_p1[to-1], el)
		}

		sl := stacks_p2[from-1][len(stacks_p2[from-1])-count:]
		stacks_p2[from-1] = stacks_p2[from-1][:len(stacks_p2[from-1])-count]
		stacks_p2[to-1] = append(stacks_p2[to-1], sl...)
	}

	var ans1, ans2 string

	for _, s := range stacks_p1 {
		ans1 += string(s[len(s)-1])
	}

	for _, s := range stacks_p2 {
		ans2 += string(s[len(s)-1])
	}

	fmt.Println(ans1)
	fmt.Println(ans2)

}
