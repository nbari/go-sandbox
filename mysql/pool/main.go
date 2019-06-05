package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
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
	pool.SetMaxIdleConns(30)
	pool.SetMaxOpenConns(50)

	err = pool.Ping()
	if err != nil {
		log.Fatalf("mysql: could not connect to the database: %s", err.Error())
	}

	var wg sync.WaitGroup
	out := make(chan string)
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			var now string
			pool.QueryRow("SELECT NOW() UNION SELECT SLEEP(1) LIMIT 1").Scan(&now)
			out <- now
			w.Done()
		}(&wg)
	}

	go func() {
		for o := range out {
			fmt.Printf("out = %+v\n", o)
		}
	}()

	wg.Wait()
	close(out)

	s := pool.Stats()
	fmt.Printf("OpenConnections: %+v\n", s.OpenConnections)
}
