package controler

import (
	"context"
	"nody/db"
	"nody/model"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateUser(ctx context.Context, user model.User) error {
	driver := ctx.Value(db.DriverContextKey("driver")).(neo4j.DriverWithContext)
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx, model.CreateUser(ctx, user))
	if err != nil {
		return err
	}
	return nil
}

func CreateIPv4(ctx context.Context, ipv4 model.IPv4) error {
	driver := ctx.Value(db.DriverContextKey("driver")).(neo4j.DriverWithContext)
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx, model.CreateIPv4(ctx, ipv4))
	if err != nil {
		return err
	}
	return nil
}

func CreateHasIP(ctx context.Context, hasip model.HasIP) error {
	driver := ctx.Value(db.DriverContextKey("driver")).(neo4j.DriverWithContext)
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx, model.CreateHasIP(ctx, hasip))
	if err != nil {
		return err
	}
	return nil
}

func GetUsernames(ctx context.Context) ([]string, error) {
	driver := ctx.Value(db.DriverContextKey("driver")).(neo4j.DriverWithContext)
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	result, err := session.ExecuteWrite(ctx, model.GetUsernames(ctx))
	if err != nil {
		return nil, err
	}
	return result.([]string), nil
}

func GetIPsByUsername(ctx context.Context, name string) ([]string, error) {
	driver := ctx.Value(db.DriverContextKey("driver")).(neo4j.DriverWithContext)
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	result, err := session.ExecuteWrite(ctx, model.GetIPsByUsername(ctx, name))
	if err != nil {
		return nil, err
	}
	return result.([]string), nil
}

func GetUsersByIP(ctx context.Context, ip string) ([]string, error) {
	driver := ctx.Value(db.DriverContextKey("driver")).(neo4j.DriverWithContext)
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	result, err := session.ExecuteWrite(ctx, model.GetUsersByIP(ctx, ip))
	if err != nil {
		return nil, err
	}
	return result.([]string), nil
}
