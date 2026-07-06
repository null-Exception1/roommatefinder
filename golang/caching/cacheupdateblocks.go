package caching

import (
	"golang/db"
	"golang/globals"
	"golang/structs"
	"time"

	"github.com/sirupsen/logrus"
)

func CacheBlocksUpdate() {

	Occupancies := make(map[string]map[string]int, 0)
	rows := db.Query("SELECT * FROM people", globals.Globaldb)

	logrus.WithFields(logrus.Fields{
		"package":  "caching",
		"endpoint": "/blocks",
	}).Debug("fetched rows from postgres db")

	for _, row := range rows {
		if _, ok := Occupancies[row.Blockno]; ok {
			Occupancies[row.Blockno][row.Roomno] = Occupancies[row.Blockno][row.Roomno] + 1
		} else {
			Occupancies[row.Blockno] = make(map[string]int, 0)
			Occupancies[row.Blockno][row.Roomno] = 1
		}
	}

	logrus.WithFields(logrus.Fields{
		"package":     "caching",
		"endpoint":    "/blocks",
		"occupancies": len(Occupancies),
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
		"package":  "caching",
		"endpoint": "/blocks",
		"blocks":   len(Block),
	}).Debug("made blocks map")

	logrus.WithFields(logrus.Fields{
		"package":  "caching",
		"endpoint": "/blocks",
		"blocks":   len(Block),
	}).Debug("updating the cache..")

	globals.CacheMutex.Lock()
	clear(globals.CacheBlocks) // clear cache
	globals.CacheBlocks = Block
	globals.CacheExpiry = time.Now().Add(30 * time.Second)
	globals.CacheMutex.Unlock()

	logrus.WithFields(logrus.Fields{
		"package": "caching",
		"cache":   len(Block),
	}).Debug("updated the cache.")

}
