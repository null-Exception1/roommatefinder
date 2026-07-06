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
		if time.Now().After(globals.CacheExpiry) {
			logrus.Info("cache rooms cleaning started")
			for block := range globals.CacheRooms {
				caching.CacheRoomsUpdate(block)
			}
			logrus.Info("cache rooms cleaning ended")
		}
	}
}
