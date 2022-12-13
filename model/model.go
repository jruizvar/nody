package model

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type IPv4 struct {
	IPv4   string `json:"ip"`
	IsMeli bool   `json:"ismeli"`
}

type HasIP struct {
	Name string `json:"name"`
	IPv4 string `json:"ip"`
}

func CreateUser(ctx context.Context, user User) neo4j.ManagedTransactionWork {
	return func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(
			ctx,
			"CREATE (u:User { name:$name, email:$email }) RETURN u.name, u.email",
			map[string]any{
				"name":  user.Name,
				"email": user.Email,
			})
		if err != nil {
			return nil, err
		}
		record, err := records.Single(ctx)
		if err != nil {
			return nil, err
		}
		return &User{
			Name:  record.Values[0].(string),
			Email: record.Values[1].(string),
		}, nil
	}
}

func CreateIPv4(ctx context.Context, ipv4 IPv4) neo4j.ManagedTransactionWork {
	return func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(
			ctx,
			"CREATE (n:IPv4 { ip:$ip, ismeli:$ismeli }) RETURN n.ip, n.ismeli", map[string]any{
				"ip":     ipv4.IPv4,
				"ismeli": ipv4.IsMeli,
			})
		if err != nil {
			return nil, err
		}
		record, err := records.Single(ctx)
		if err != nil {
			return nil, err
		}
		return &IPv4{
			IPv4:   record.Values[0].(string),
			IsMeli: record.Values[1].(bool),
		}, nil
	}
}

func CreateHasIP(ctx context.Context, hasip HasIP) neo4j.ManagedTransactionWork {
	return func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(
			ctx,
			"MATCH (u:User) MATCH (n:IPv4) WHERE u.name=$name AND n.ip=$ip MERGE (u)-[:HAS_IP]->(n) RETURN u.name, n.ip",
			map[string]any{
				"name": hasip.Name,
				"ip":   hasip.IPv4,
			})
		if err != nil {
			return nil, err
		}
		record, err := records.Single(ctx)
		if err != nil {
			return nil, err
		}
		return &HasIP{
			Name: record.Values[0].(string),
			IPv4: record.Values[1].(string),
		}, nil
	}
}

func GetUsernames(ctx context.Context) neo4j.ManagedTransactionWork {
	return func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(
			ctx,
			"MATCH (u:User) RETURN u.name",
			map[string]any{})
		if err != nil {
			return nil, err
		}
		var arr []string
		for records.Next(ctx) {
			value, found := records.Record().Get("u.name")
			if found {
				arr = append(arr, value.(string))
			}
		}
		if err = records.Err(); err != nil {
			return nil, err
		}
		return arr, nil
	}
}

func GetIPsByUsername(ctx context.Context, name string) neo4j.ManagedTransactionWork {
	return func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(
			ctx,
			"MATCH (u:User)-->(n:IPv4) WHERE u.name=$name RETURN n.ip as ipv4",
			map[string]any{
				"name": name,
			})
		if err != nil {
			return nil, err
		}
		var arr []string
		for records.Next(ctx) {
			value, found := records.Record().Get("ipv4")
			if found {
				arr = append(arr, value.(string))
			}
		}
		if err = records.Err(); err != nil {
			return nil, err
		}
		return arr, nil
	}
}

func GetUsersByIP(ctx context.Context, ip string) neo4j.ManagedTransactionWork {
	return func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(
			ctx,
			"MATCH (u:User)-->(n:IPv4) WHERE n.ip=$ip RETURN u.name",
			map[string]any{
				"ip": ip,
			})
		if err != nil {
			return nil, err
		}
		var arr []string
		for records.Next(ctx) {
			value, found := records.Record().Get("u.name")
			if found {
				arr = append(arr, value.(string))
			}
		}
		if err = records.Err(); err != nil {
			return nil, err
		}
		return arr, nil
	}
}
