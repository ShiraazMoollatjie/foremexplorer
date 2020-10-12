package main

import (
	"flag"

	"log"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/db"
)

func main() {
	flag.Parse()
	dbc, err := db.Connect()
	dbc.LogMode(true)
	dbc.Debug()
	if err != nil {
		log.Fatalf("Cannot connect to the database! Error: %+v", err)
	}

	// TODO(shiraaz): This is only a short term solution. Gorm does not create foreign keys.
	// Rather use sql.db.Exec() statements.
	dbc.DropTableIfExists(db.Article{})
	if err := dbc.AutoMigrate(db.Article{}).Error; err != nil {
		log.Fatalf("Errors occurred during migration! Error: %+v", err)
	}
}
