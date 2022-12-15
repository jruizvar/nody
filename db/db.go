package db

import (
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Driver() neo4j.DriverWithContext {
	dbUri := os.Getenv("DBURI")
	dbadmin := os.Getenv("DBADMIN")
	dbpassw := os.Getenv("DBPASSW")
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth(dbadmin, dbpassw, ""))
	if err != nil {
		panic(err)
	}
	return driver
}
