package main

import (
	"fmt"
	"log"

	"github.com/theprimeagen/projector/pkg/projector"
)

func main() {
	p, err := projector.GetOpts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", p)
}
