package db

import (
	"context"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type DriverContextKey string

func Driver(ctx context.Context) context.Context {
	dbUri := os.Getenv("DBURI")
	dbadmin := os.Getenv("DBADMIN")
	dbpassw := os.Getenv("DBPASSW")
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth(dbadmin, dbpassw, ""))
	if err != nil {
		panic(err)
	}
	return context.WithValue(ctx, DriverContextKey("driver"), driver)
}
