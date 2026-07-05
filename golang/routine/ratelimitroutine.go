package routine

import (
	utils "golang/globals"
	"time"

	"github.com/sirupsen/logrus"
)

func Routine(ticker *time.Ticker) {
	for range ticker.C {
		select {
		case utils.RatelimitChannel <- time.Now():
			logrus.WithFields(logrus.Fields{
				"package":  "routine",
				"function": "ratelimit",
			}).Trace("added a token to the ratelimit bucket")
		default:
			logrus.WithFields(logrus.Fields{
				"package":  "routine",
				"function": "ratelimit",
			}).Trace("ratelimit is full")
		}
	}
}
