package main

import (
	"golang/caching"
	globals "golang/globals"
	handlers "golang/handlers"
	goroutines "golang/routine"
	"net/http"
	"time"

	initfuncs "golang/init"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	initfuncs.Logging()
	initfuncs.Database()

	defer globals.Globaldb.Close()

	globals.RatelimitChannel = make(chan time.Time, 20)

	go goroutines.WorkersResults()
	go goroutines.CacheBlocksCleanup()
	go goroutines.CacheRoomsCleanup()
	go goroutines.Routine(globals.Ticker)
	go goroutines.StartSessionCleanup(globals.Globaldb)

	for range globals.NumWorkers {
		globals.CacheRoomsJobsWaitGroup.Go(func() {
			for job := range globals.CacheRoomsJobs {
				caching.CacheRoomsUpdate(job.Blockno)
			}
		})
	}
	http.HandleFunc("/registration", handlers.Ratelimit(handlers.RegistrationHandler))
	http.HandleFunc("/rooms", handlers.Ratelimit(handlers.Rooms))
	http.HandleFunc("/blocks", handlers.Ratelimit(handlers.Blocks))
	http.HandleFunc("/login", handlers.Ratelimit(handlers.Login))
	http.HandleFunc("/verify", handlers.Ratelimit(handlers.Verify))
	http.HandleFunc("/logout", handlers.Ratelimit(handlers.Logout))

	http.ListenAndServe("localhost:8080", nil)
}
