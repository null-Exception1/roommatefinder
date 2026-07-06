package handlers

import (
	"fmt"
	"golang/caching"
	"golang/globals"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func Rooms(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	q := req.URL.Query()

	blockno := q.Get("block")

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/rooms",
		"block":    blockno,
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Info("requested /rooms")

	if os.Getenv("CACHING") == "true" {
		if time.Now().After(globals.CacheBlocksExpiry[blockno]) {
			caching.CacheRoomsUpdate(blockno)
			globals.CacheMisses++
		} else {
			logrus.Debug("CACHE HIT!") // moment when the cache hits lol
			globals.CacheHits++
		}
	} else {
		caching.CacheRoomsUpdate(blockno) // do a normal update instead of holding off (no different from cache miss)
	}

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/rooms",
		"rooms":    len(globals.CachedRoomsJSON[blockno]),
		"status":   http.StatusOK,
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Info("response sent")

	fmt.Fprintf(w, "%s", globals.CacheBlocksExpiry[blockno])

}
