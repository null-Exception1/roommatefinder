package routine

import (
	"golang/caching"
	"golang/globals"
	"time"

	"github.com/sirupsen/logrus"
)

func CacheBlocksCleanup() {
	ticker := time.NewTicker(globals.CacheBlocksRoutine * time.Minute)
	for range ticker.C {
		if time.Now().After(globals.CacheExpiry) {
			logrus.Info("cache cleaning started")
			caching.CacheBlocksUpdate()
		}
	}
}
