package main

import (
	"log"
	"os"

	"github.com/vanillaiice/snaparser_server"
)

func main() {
	if err := snaparser_server.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
