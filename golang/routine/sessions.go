package routine

import (
	"database/sql"
	"golang/globals"
	"time"

	"github.com/sirupsen/logrus"
)

func StartSessionCleanup(db *sql.DB) {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			logrus.WithFields(logrus.Fields{
				"package":  "routine",
				"function": "sessions",
			}).Info("session cleanup")
			_, err := globals.Globaldb.Exec(`DELETE FROM sessions WHERE expires_at < NOW()`)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"package":  "routine",
					"function": "sessions",
					"errror":   err,
				}).Error("couldn't do session cleanup")
			} else {
				logrus.WithFields(logrus.Fields{
					"package":  "routine",
					"function": "sessions",
				}).Info("session cleanup done with no errors")
			}
		}
	}()
}
