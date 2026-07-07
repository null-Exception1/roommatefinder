package routine

import (
	"golang/globals"

	"github.com/sirupsen/logrus"
)

func WorkersResults() {
	for result := range globals.CacheRoomsJobsResults {
		logrus.WithFields(logrus.Fields{
			"package":       "routine",
			"resultblockno": result.Blockno,
			"resultlen":     len(result.JSON),
		}).Debug("updating cache with result from worker")
		globals.CachedRoomsJSON.Store(result.Blockno, result.JSON)
		//<-globals.CacheRoomsJobsResults
	}
}
