package main

import (
	"flag"
	"log"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/db"
	"github.com/ShiraazMoollatjie/foremexplorer/pkg/ops"
	"github.com/ShiraazMoollatjie/foremexplorer/pkg/state"
)

func main() {
	flag.Parse()
	_, err := db.Connect()
	if err != nil {
		log.Fatalf("Cannot create db. Error: %+v", err)
	}

	s := state.NewState()
	ops.SeedDevData(s)
}
