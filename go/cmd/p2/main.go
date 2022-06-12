package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	res := []line{}
	for _, sLine := range strings.Split(getInput(), "\n") {
		lin, err := ParseLine(sLine)
		if err != nil {
			log.Fatalln(err)
		}
		if isHOrV(*lin) {
			res = append(res, *lin)
		}
	}

	fmt.Printf("%+v\n", res)
}

func getInput() string {
	return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
}

type point struct {
	x int
	y int
}

type line struct {
	p1 point
	p2 point
}

func isHOrV(line line) bool {
	return line.p1.x == line.p2.x || line.p1.y == line.p2.y
}

func parsePoint(p string) (*point, error) {
	parts := strings.Split(p, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return &point{
		x: x,
		y: y,
	}, nil
}

func ParseLine(l string) (*line, error) {
	li := strings.Split(l, " -> ")
	p1, err := parsePoint(li[0])
	if err != nil {
		return nil, err
	}
	p2, err := parsePoint(li[1])
	if err != nil {
		return nil, err
	}

	return &line{
		p1: *p1,
		p2: *p2,
	}, nil
}
