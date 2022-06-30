package main

import (
	"fmt"
	"log"

	"github.com/17xande/projector/pkg/projector"
)

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("unable to get options %v\n", err)
	}

	config, err := projector.NewConfig(opts)
	if err != nil {
		log.Fatalf("unable to get config %v\n", err)
	}

	fmt.Printf("%+v\n", config)
}
