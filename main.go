/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"Garnery/cmd"
	"Garnery/db"
	"log"
	"fmt"
	"context"
	"Garnery/config"
)

func main() {
	cmd.Execute()

	// Attempt to connect to the database
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		// os.Exit(1) // log.Fatalf already exits the program
	}
	defer func() {
		// Ensure the connection is closed when main exits
		fmt.Println("Closing database connection...")
		conn.Close(context.Background())
		fmt.Println("Database connection closed.")
	}()

	fmt.Println("Successfully connected to the database!")

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("Database ping failed: %v\n", err)
	}
	fmt.Println("Database ping successful!")

	var serverVersion string
	err = conn.QueryRow(context.Background(), "SELECT version();").Scan(&serverVersion)
	if err != nil {
		log.Fatalf("Failed to execute version query: %v\n", err)
	}
	fmt.Printf("PostgreSQL Server Version: %s\n", serverVersion)

	// --- Connection Test 3: Check database name
	var currentDB string
	err = conn.QueryRow(context.Background(), "SELECT current_database();").Scan(&currentDB)
	if err != nil {
		log.Fatalf("Failed to get current database name: %v\n", err)
	}
	fmt.Printf("Connected to database: %s\n", currentDB)

	fmt.Println("\nAll connection tests passed!")



	// From here adding application's logic that interacts with the database.

	// Testing conf stuff

	connInfo, err := config.Get_Connect_Info()
	if err != nil {
		log.Fatalf("Error getting connection info: %v", err)
	}

	fmt.Println("\nConnection details recieved in main:")
	fmt.Printf("User: %s\n", connInfo.User)
	fmt.Printf("Pass: %s\n", connInfo.Password)
	fmt.Printf("Host: %s\n", connInfo.Host)
	fmt.Printf("Port: %s\n", connInfo.Port)
	fmt.Printf("Data: %s\n", connInfo.Database)

}
