package config

import (
	"fmt"
	//"log" //
	"github.com/bigkevmcd/go-configparser"
)

type Conn_Info struct {
	User 		string
	Password	string
	Host		string
	Port		string
	Database	string
}

func Get_Connect_Info() (Conn_Info, error) { // Changed signature
	var connInfo Conn_Info

	// Using NewConf
	config, err := configparser.NewConfigParserFromFile("config/config.conf")
	if err != nil {
		return Conn_Info{}, fmt.Errorf("failed to load config file: %w", err) // Return error
	}

	connInfo.User, err = config.Get("database", "user")
	if err != nil {
		return Conn_Info{}, fmt.Errorf("failed to get 'user' from config: %w", err) // Return error
	}

	connInfo.Password, err = config.Get("database", "pass")
	if err != nil {
		return Conn_Info{}, fmt.Errorf("failed to get 'pass' from config: %w", err) // Return error
	}

	connInfo.Host, err = config.Get("database", "host")
	if err != nil {
		return Conn_Info{}, fmt.Errorf("failed to get 'host' from config: %w", err) // Return error
	}

	connInfo.Port, err = config.Get("database", "port")
	if err != nil {
		fmt.Println("Setting port to default: 5432") // This message is still fine
		connInfo.Port = "5432"
	}

	connInfo.Database, err = config.Get("database", "data")
	if err != nil {
		return Conn_Info{}, fmt.Errorf("failed to get 'data' from config: %w", err) // Return error
	}


	//fmt.Printf("Parsed Connection Info:\n")
	//fmt.Printf("User: %s\n", connInfo.User)
	//fmt.Printf("Password: %s\n", connInfo.Password)
	//fmt.Printf("Host: %s\n", connInfo.Host)
	//fmt.Printf("Port: %s\n", connInfo.Port)
	//fmt.Printf("Database: %s\n", connInfo.Database)

	return connInfo, nil // Return the struct and nil for no error
}
