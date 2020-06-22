package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Show() {
	db, err := sql.Open("mysql", "devreadonly:ug<H7+hsSmw1@tcp(10.30.60.122:3306)/towngas_user_1")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select id,uwl_seq  from t_user_withdrawal order by id")
	for rows.Next() {
		var i interface{}
		var s interface{}
		// var sl []string
		err = rows.Scan(&i, &s)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("found row containing %d%q", i, s)
	}
	rows.Close()
	db.Close()
}
