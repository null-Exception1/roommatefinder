package main

import (
	"fmt"
	"net/http"
	"time"
)

var ratelimit chan time.Time

var ticker = time.NewTicker(100 * time.Millisecond)

func Ratelimit(handlerfunc http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		select {
		case <-ratelimit:
			handlerfunc(w, req)
		default:
			fmt.Println("[RATELIMIT EXCEEDED] Packet left out")
		}
	}

}

func Handler1(w http.ResponseWriter, req *http.Request) {

}

func Routine() {
	for range ticker.C {
		select {
		case ratelimit <- time.Now():
		default:
		}
	}
}
func main() {
	ratelimit = make(chan time.Time, 5)
	http.HandleFunc("/api", Ratelimit(Handler1))

	http.ListenAndServe("localhost:8080", nil)
}
