package db

import (
	"flag"
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
)

var (
	dbHost     = flag.String("dbHost", "0.0.0.0", "the db schema to use")
	dbUser     = flag.String("dbUser", "foremexplorer", "the db user")
	dbPassword = flag.String("dbPassword", "password", "the db password")
	dbPort     = flag.Int("dbPort", 3306, "the db schema to use")
	dbSchema   = flag.String("dbSchema", "foremexplorer", "the db schema to use")
)

// Connect returns a pooled connection to the database.
func Connect() (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", *dbUser, *dbPassword, *dbHost, *dbPort, *dbSchema)
	return gorm.Open("mysql", url)
}

// connectForTesting returns a pooled connection to the database for integration testing.
func connectForTesting(_ *testing.T) (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/test?charset=utf8&parseTime=True&loc=Local", *dbUser, *dbPassword, *dbHost, *dbPort)
	return gorm.Open("mysql", url)
}
