package main

import (
	"fmt"
	"strings"
)

func getInput() string {
	return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
}

func main() {
	things := strings.Split(getInput(), "\n")
	colLen := len(things[0])
	treeCount := 0
	for i, thing := range things {
		if string(thing[i*3%colLen]) == "#" {
			treeCount++
		}
	}

	fmt.Printf("Trees: %d\n", treeCount)
}
