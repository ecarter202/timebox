package main

import (
	"log"

	"github.com/ecarter202/timebox/cmd"
	"github.com/ecarter202/timebox/db"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
