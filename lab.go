package main

import (
	"bufio"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/arlsclu7/golab/block"
	"github.com/arlsclu7/golab/db"
	applog "github.com/arlsclu7/golab/log"
	"github.com/arlsclu7/golab/types"
	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB // Database connection pool.

func chans() {
	donechan := make(chan string)
	go func() {
		donechan <- "i am done"
	}()
	<-donechan
}
func main() {
	f, err := os.Lstat("go.mod")
	fmt.Println(f.Mode().Perm())
	newfile := "a.log"
	_, err = os.OpenFile(newfile, os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	gn, err := os.Lstat(newfile)
	fmt.Println(gn.Mode().Perm())
	os.Stdout.Write([]byte{1, 2, 3})
	applog := applog.Applog()
	applog.Println("aaa")
	c, err := os.Create("/usr/local/var/log/goslab/a.log")
	if err != nil {
		applog.Fatalln(err)
	}
	applog.Fatal(c)
	os.Exit(1)
	// chans()
	os.Exit(12)
	block.MyBlock()
	os.Exit(111)
	db.Show()
	os.Exit(1)
	// os.Stdout.Write([]byte{1, 2, 3})
	// fmt.Println(os.Stdout)
	// os.Exit(1)
	types.Show()
	os.Exit(1)
	os.Exit(1)
	block.Change()
	block.Change()
	fmt.Println(block.A)
	os.Exit(1)
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	os.Exit(0)
	id := flag.Int64("id", 77, "person ID to find")
	dsn := flag.String("dsn", "devreadonly:ug<H7+hsSmw1@tcp(10.30.60.122)/towngas_user_1", "connection data source name")
	flag.Parse()

	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}
	if *id == 0 {
		log.Fatal("missing person ID")
	}

	// Opening a driver typically will not attempt to connect to the database.
	pool, err := sql.Open("mysql", *dsn)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()

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

	var uwd_seq string
	// err := pool.QueryRowContext(ctx, "select  *  from t_user_withdrawal where id = :id;",77).Scan(&uwd_seq)
	rst, err := pool.Query("select  *  from t_user_withdrawal where id = 0")
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}
	log.Println("name=", uwd_seq)
	log.Println("record=", rst)
}
