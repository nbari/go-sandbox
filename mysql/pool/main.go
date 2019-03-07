package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))
	pool, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatalf("mysql: Error on initializing database connection: %s", err.Error())
	}

	// sql pool options
	pool.SetConnMaxLifetime(time.Minute)
	pool.SetMaxIdleConns(50)
	pool.SetMaxOpenConns(50)

	err = pool.Ping()
	if err != nil {
		log.Fatalf("mysql: could not connect to the database: %s", err.Error())
	}

	var now string
	for i := 0; i < 50; i++ {
		go func() {
			pool.QueryRow("SELECT NOW()").Scan(&now)
		}()
	}

	for {
		err = pool.QueryRow("SELECT current_timestamp(3)").Scan(&now)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("now = %+v\n", now)
		time.Sleep(3 * time.Second)
		s := pool.Stats()
		fmt.Printf("s = %+v\n", s)
	}
}
