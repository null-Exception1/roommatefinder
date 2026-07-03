package handlers

import (
	"fmt"
	utils "golang/globals"
	"net/http"
	"time"
)

func Ratelimit(handlerfunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//fmt.Println("[PACKET] [RATELIMIT]", len(utils.RatelimitChannel))
		select {

		case <-utils.RatelimitChannel:
			handlerfunc(w, req)
		default:
			fmt.Println("[RATELIMIT EXCEEDED] Packet left out")
		}
	}
}

func Routine(ticker *time.Ticker) {
	for range ticker.C {
		//fmt.Println("[TICKER] ", len(utils.RatelimitChannel))
		select {
		case utils.RatelimitChannel <- time.Now():
		default:
		}
	}
}
