package handlers

import (
	"encoding/json"
	"fmt"
	"golang/db"
	"golang/globals"
	"golang/structs"
	"net/http"

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

	Occupancies := make(map[string]map[string]int, 0)
	rows := db.Query("SELECT * FROM people", globals.Globaldb)

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/blocks",
		"rows":     len(rows),
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Info("fetched rows from postgres db")

	for _, row := range rows {
		if _, ok := Occupancies[row.Blockno]; ok {
			Occupancies[row.Blockno][row.Roomno] = Occupancies[row.Blockno][row.Roomno] + 1
		} else {
			Occupancies[row.Blockno] = make(map[string]int, 0)
			Occupancies[row.Blockno][row.Roomno] = 1
		}
	}

	logrus.WithFields(logrus.Fields{
		"package":     "handlers",
		"endpoint":    "/blocks",
		"occupancies": len(Occupancies),
		"method":      req.Method,
		"remote":      req.RemoteAddr,
	}).Debug("made occupancies map")

	Block := make(map[string]*structs.Block, 0)

	for key := range Occupancies {
		Block[key] = &structs.Block{Partial: 0, Full: 0}
		for _, occupancy := range Occupancies[key] {
			if occupancy >= 2 { // remind me to check the occupancy max limit later
				Block[key].Full = Block[key].Full + 1
			} else {
				Block[key].Partial = Block[key].Partial + 1
			}
		}
	}

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/blocks",
		"blocks":   len(Block),
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Debug("made blocks map")

	str, err := json.Marshal(Block)

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
