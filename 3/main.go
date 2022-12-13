package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getIndex(slice []string, element string) int {
	for i, e := range slice {
		if e == element {
			return i
		}
	}
	return -1

}

func main() {
	path, _ := os.Getwd()

	file, _ := os.ReadFile(filepath.Join(path, "input"))

	data := strings.Split(string(file), "\n")

	data = data[:len(data)-1]

	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var ans1, ans2 int

	for _, l := range data {
		splitted := strings.Split(l, "")
		seen := []string{}

		for _, c := range splitted[:len(splitted)/2] {
			if getIndex(splitted[len(splitted)/2:], c) > -1 && getIndex(seen, c) < 0 {
				ans1 += getIndex(letters, c) + 1
				seen = append(seen, c)
			}
		}
	}

	for p := 0; p < len(data)-1; p += 3 {
		splitted := strings.Split(data[p], "")
		seen := []string{}

		for _, c := range splitted {
			if getIndex(strings.Split(data[p+1], ""), c) > -1 && getIndex(strings.Split(data[p+2], ""), c) > -1 && getIndex(seen, c) < 0 {
				seen = append(seen, c)
				ans2 += getIndex(letters, c) + 1
			}
		}
	}

	fmt.Println(ans1, ans2)
}
