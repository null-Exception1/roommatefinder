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

func Blocks(w http.ResponseWriter, req *http.Request) {

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/blocks",
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Info("requested /blocks by ", req.RemoteAddr)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/text")

	if os.Getenv("CACHING") == "true" {
		if time.Now().After(globals.CacheExpiry) {
			caching.CacheBlocksUpdate()
			globals.CacheMisses++
			logrus.Debug("CACHE MISS!")
		} else {
			logrus.Debug("CACHE HIT!") // moment when the cache hits lol
			globals.CacheHits++
		}
	} else {
		caching.CacheBlocksUpdate() // do a normal update instead of holding off (no different from cache miss)
	}
	globals.CacheMutex.RLock()
	str, err := json.Marshal(globals.CacheBlocks)
	globals.CacheMutex.RUnlock()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":   "handlers",
			"endpoint":  "/blocks",
			"jsonified": len(str),
			"method":    req.Method,
			"remote":    req.RemoteAddr,
		}).Error("error in JSON.Marshal")
	}

	logrus.WithFields(logrus.Fields{
		"package":   "handlers",
		"endpoint":  "/blocks",
		"jsonified": len(str),
		"status":    http.StatusOK,
		"method":    req.Method,
		"remote":    req.RemoteAddr,
	}).Info("response sent")

	fmt.Fprintf(w, "%s", string(str))
}
