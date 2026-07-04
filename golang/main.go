package main

import (
	"database/sql"
	"fmt"
	globals "golang/globals"
	handlers "golang/handlers"
	goroutines "golang/routine"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var ratelimit chan time.Time

func main() {
	globals.Ticker = time.NewTicker(100 * time.Millisecond)
	globals.RatelimitChannel = make(chan time.Time, 10)
	connStr := "postgres://devuser:devpass@localhost:5432/roommatefinder?sslmode=disable"
	go goroutines.Routine(globals.Ticker)
	go goroutines.StartSessionCleanup(globals.Globaldb)
	globals.Globaldb, _ = sql.Open("postgres", connStr)
	defer globals.Globaldb.Close()

	fmt.Println("Connected to Postgres!")

	/*
		utils.Insert("12345", "fuck you", "hello", "hello", "012", 16, db)
		db.Exec(`DELETE FROM people WHERE admn_hash='12345';`)
		utils.Query("SELECT * FROM people;", db)
	*/

	ratelimit = make(chan time.Time, 5)
	http.HandleFunc("/registration", handlers.Ratelimit(handlers.RegistrationHandler))
	http.HandleFunc("/rooms", handlers.Ratelimit(handlers.Rooms))
	http.HandleFunc("/blocks", handlers.Ratelimit(handlers.Blocks))
	http.HandleFunc("/login", handlers.Ratelimit(handlers.Login))
	http.HandleFunc("/verify", handlers.Ratelimit(handlers.Verify))

	http.ListenAndServe("localhost:8080", nil)
}
