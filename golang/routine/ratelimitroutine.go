package routine

import (
	utils "golang/globals"
	"time"
)

func Routine(ticker *time.Ticker) {
	for range ticker.C {
		select {
		case utils.RatelimitChannel <- time.Now():
		default:
		}
	}
}
