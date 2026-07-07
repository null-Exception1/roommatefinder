package handlers

import (
	"encoding/json"
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
		if val, ok := globals.CacheBlocksExpiry.Load(blockno); ok {
			expiry := val.(time.Time) // type assertion
			if time.Now().After(expiry) {
				logrus.Debug("CACHE MISS!")
				caching.AddCacheRoomsJob(blockno)
				globals.CacheMisses.Add(1)
			} else {
				logrus.Debug("CACHE HIT!") // moment when the cache hits lol
				globals.CacheHits.Add(1)
			}
		} else { // edge case where we don't even have the key in cache
			logrus.Debug("CACHE MISS! || FIRST TIME EVER LOAD")

			// this time, actually wait and do the process instead of skipping
			rooms := caching.FormRooms(blockno)
			bytes, _ := json.Marshal(rooms)
			globals.CachedRoomsJSON.Store(blockno, string(bytes))
			globals.CacheBlocksExpiry.Store(blockno, time.Now().Add(globals.CacheRoomsSeconds*time.Second))

			//caching.AddCacheRoomsJob(blockno)
			globals.CacheMisses.Add(1)
		}

	} else {
		caching.AddCacheRoomsJob(blockno) // do a normal update instead of holding off (no different from cache miss)
	}

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/rooms",
		"status":   http.StatusOK,
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Info("response sent")

	json, ok := globals.CachedRoomsJSON.Load(blockno)

	if !ok {
		json = "{}" // no cache built yet (worse case scenario, this only happens on startup ig)
	}

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/rooms",
		"status":   http.StatusOK,
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Debug("fetched cache")

	fmt.Fprintf(w, "%s", json)
}
