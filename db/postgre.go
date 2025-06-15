package db

import (
	//"fmt"
	"context"
	"Garnery/config"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	connInfo, err := config.Get_Connect_Info()
	if err != nil {
		log.Fatalf("Error getting connection info: %v", err)
	}

	conn, err := pgx.Connect(context.Background(), "postgres://" + connInfo.User + ":" + connInfo.Password + "@" + connInfo.Host + ":" + connInfo.Port + "/" + connInfo.Database)
	//conn, err := pgx.Connect(context.Background(), "postgres://test:password@localhost:5432/testdb")
	if err != nil {
		return nil, err
	}
	return conn, nil
}
