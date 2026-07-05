package handlers

import (
	utils "golang/globals"
	"net/http"
	"time"
)

func Ratelimit(handlerfunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		select {

		case <-utils.RatelimitChannel:
			handlerfunc(w, req)
		default:

		}
	}
}

func Routine(ticker *time.Ticker) {
	for range ticker.C {
		select {
		case utils.RatelimitChannel <- time.Now():
		default:
		}
	}
}
