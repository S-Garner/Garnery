package db

import (
	//"fmt"
	"context"
	//"log"
	"github.com/jackc/pgx/v5"
)

func connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://test:password@localhost:5432/testdb")
	if err != nil {
		return nil, err
	}
	return conn, nil
}
