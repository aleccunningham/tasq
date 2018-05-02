package main

import (
	"log"

	"github.com/aleccunningham/tasq/tasqctl/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
