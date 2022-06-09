package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(getInput(), "\n")
	res := point{x: 0, y: 0}

	for _, line := range lines {
		r := parseLine(line)
		res.x += r.x
		res.y += r.y
	}
	fmt.Printf("%d\t %d\n", res, res.x*res.y)
}

type point struct {
	x int
	y int
}

func parseLine(line string) point {
	parts := strings.Split(line, " ")
	dir := parts[0]
	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalln("this should never happen.")
	}

	if dir == "forward" {
		return point{
			x: amount,
			y: 0,
		}
	} else if dir == "up" {
		return point{
			x: 0,
			y: -amount,
		}
	} else {
		return point{
			x: 0,
			y: amount,
		}
	}
}

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}
