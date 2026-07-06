package routine

import (
	"golang/caching"
	"golang/globals"
	"time"

	"github.com/sirupsen/logrus"
)

func CacheCleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		if time.Now().After(globals.CacheExpiry) {
			logrus.Info("cache cleaning started")
			caching.CacheUpdate()
		}
	}
}
