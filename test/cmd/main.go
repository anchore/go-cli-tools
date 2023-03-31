package main

import (
	"log"

	"github.com/anchore/go-cli-tools/test/cmd/cli"
)

func main() {
	c, err := cli.New()
	if err != nil {
		log.Fatalf("error during command construction: %v", err)
	}

	if err := c.Execute(); err != nil {
		log.Fatalf("error during command execution: %v", err)
	}
}
