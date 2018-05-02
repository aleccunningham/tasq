package main

import (
	"log"

	"github.com/aleccunningham/tasq/cmd/tasqctl/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
