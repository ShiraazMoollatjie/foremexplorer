package state

import (
	"flag"
	"log"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/db"
	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
	"github.com/jinzhu/gorm"
)

var devAPIKey = flag.String("dev_api_key", "", "the api key to use with dev.to")

type State struct {
	ForemClient *gophorem.Client
	DB          *gorm.DB
}

func NewState() *State {
	s := &State{}
	var err error

	s.ForemClient = gophorem.NewDevtoClient(gophorem.WithAPIKey(*devAPIKey))
	s.DB, err = db.Connect()
	if err != nil {
		log.Fatalf("cannot connect to the database: %+v", err)
	}

	return s
}
