package main

import "fmt"

type GoEnum = int

const (
	Foo GoEnum = iota // auto incrementing integer. this is how you do enums in go.
	Bar
	Baz
)

func main() {
	fmt.Printf("%d\n", Foo)
}
