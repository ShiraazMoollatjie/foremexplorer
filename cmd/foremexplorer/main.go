package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/db"
	"github.com/ShiraazMoollatjie/foremexplorer/pkg/ops"
	"github.com/ShiraazMoollatjie/foremexplorer/pkg/server"
	"github.com/ShiraazMoollatjie/foremexplorer/pkg/state"
)

func main() {
	flag.Parse()
	_, err := db.Connect()
	if err != nil {
		log.Fatalf("Cannot create db. Error: %+v", err)
	}

	s := state.NewState()
	go ops.QueryForemArticlesForever(s)
	go server.ServeHttp(s)
	waitForShutdown()
}

func waitForShutdown() {
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	fmt.Println("Shutting down")
}
