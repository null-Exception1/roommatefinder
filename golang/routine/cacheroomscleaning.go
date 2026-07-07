package routine

import (
	"golang/caching"
	"golang/globals"
	"time"

	"github.com/sirupsen/logrus"
)

func CacheRoomsCleanup() {
	ticker := time.NewTicker(globals.CacheRoomsRoutine * time.Minute)
	for range ticker.C {
		logrus.Info("cache rooms cleaning started")
		globals.CachedRoomsJSON.Range(func(key, value any) bool {
			if block, ok := key.(string); ok {
				caching.CacheRoomsUpdate(block)
			} else {
				logrus.Warn("found a non-string key inside CachedRoomsJSON map")
			}
			return true
		})
		logrus.Info("cache rooms cleaning ended")

	}
}
