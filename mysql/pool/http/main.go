package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nbari/violetear"
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

	router := violetear.New()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		fmt.Fprintf(w, "Will run this query: SELECT NOW() UNION SELECT SLEEP(3) LIMIT 1\n\n")
		flusher.Flush()

		var wg sync.WaitGroup
		out := make(chan string)
		for i := 0; i < 30; i++ {
			wg.Add(1)
			go func(w *sync.WaitGroup, j int) {
				var now string
				if j%3 == 0 {
					pool.QueryRow("SELECT \"after 3 seconds\" UNION SELECT SLEEP(3) LIMIT 1").Scan(&now)
				} else {
					pool.QueryRow("SELECT NOW()").Scan(&now)
				}
				out <- now
				w.Done()
			}(&wg, i)
		}

		go func() {
			for o := range out {
				fmt.Fprintf(w, "now: %s\n", o)
				flusher.Flush()
			}
		}()

		wg.Wait()
		time.Sleep(time.Second)
		close(out)

		s := pool.Stats()
		fmt.Fprintf(w, fmt.Sprintf("OpenConnections: %+v\n", s.OpenConnections))
		flusher.Flush()
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())

}
