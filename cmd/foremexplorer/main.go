package main

import (
	"flag"
	"log"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/db"
)

func main() {
	flag.Parse()
	_, err := db.Connect()
	if err != nil {
		log.Fatalf("Cannot create db. Error: %+v", err)
	}
}
