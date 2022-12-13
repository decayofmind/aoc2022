package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name  string
	size  int
	files map[string]*File
	dirs  map[string]*Directory
}

func getSizes(dir *Directory) int {
	var size int

	for _, f := range dir.files {
		size += f.size
	}

	for n, d := range dir.dirs {
		if n != ".." {
			size += getSizes(d)
		}
	}

	dir.size = size

	return size
}

func getSizesList(dir *Directory) []int {
	var sizes []int

	for n, d := range dir.dirs {
		if n != ".." {
			sizes = append(sizes, d.size)
			sizes = append(sizes, getSizesList(d)...)
		}
	}

	return sizes
}

func main() {
	var ans1, ans2 int
	path, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(path, "input"))

	root := &Directory{name: "/", files: make(map[string]*File), dirs: make(map[string]*Directory)}

	scanner := bufio.NewScanner(file)

	cwd := root

	for scanner.Scan() {
		l := scanner.Text()
		p := strings.Split(l, " ")

		if p[0] == "$" && p[1] == "cd" {
			if p[2] != "/" {
				cwd = cwd.dirs[p[2]]
			}
		} else if p[0] != "$" {
			if p[0] == "dir" {
				cwd.dirs[p[1]] = &Directory{name: p[1], files: make(map[string]*File), dirs: make(map[string]*Directory)}
				cwd.dirs[p[1]].dirs[".."] = cwd
			} else {
				size, _ := strconv.Atoi(p[0])
				cwd.files[p[1]] = &File{name: p[1], size: size}
			}
		}
	}

	need := 70000000 - getSizes(root) - 30000000

	sizes := getSizesList(root)

	sort.Ints(sizes)

	for _, s := range sizes {
		if s <= 100000 {
			ans1 += s
		}

		if s >= -need {
			ans2 = s
			break
		}
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}
