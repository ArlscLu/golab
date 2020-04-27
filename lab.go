package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB // Database connection pool.
// Ping the database to verify DSN provided by the user is valid and the
// server accessible. If the ping fails exit the program with an error.
func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

// Query the database for the information requested and prints the results.
// If the query fails exit the program with an error.
func Query(ctx context.Context, id int64) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var city_id int
	var achiev_city_id int
	var achiev_city_name string
	err := pool.QueryRowContext(ctx, "select city_id,achiev_city_id,achiev_city_name  where city_id =4").Scan(&city_id, &achiev_city_id, &achiev_city_name)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}
	log.Println("city_id=", city_id)
	log.Println("achiev_city_id=", achiev_city_id)
	log.Println("achiev_city_name=", achiev_city_name)
}
func main() {
	// db, err := sql.Open("mysql", "root:111111@127.0.0.1/dbname")
	// fmt.Println(db)
	os.Setenv("DSN", "root:111111@/lab")
	id := flag.Int64("id", 0, "person ID to find")
	dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")
	flag.Parse()
	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}
	// if *id == 0 {
	// 	log.Fatal("missing person ID")
	// }
	var err error
	// Opening a driver typically will not attempt to connect to the database.
	pool, err = sql.Open("mysql", *dsn)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()
	fmt.Println(err)

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		select {
		case <-appSignal:
			stop()
		}
	}()
	Ping(ctx)

	Query(ctx, *id)
}
